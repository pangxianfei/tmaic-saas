package helper

import (
	"errors"
	"reflect"
	"strings"

	"github.com/jinzhu/copier"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"tmaic/vendors/framework/model"
)

type Helper struct {
	model.BaseModel
}

func (h *Helper) SetTX(db *gorm.DB) {
	// helper cannot setTX only can setDB in transaction_helpers
}

func isNull(value reflect.Value, nullData reflect.Value) bool {
	if value.Type().Kind() == reflect.Ptr {
		if !nullData.IsNil() { // nullData has null.xxx set
			return true
		}
	}
	return false
}

func isSameField(a reflect.StructField, b reflect.StructField) bool {
	return a.Type == b.Type && a.Name == b.Name
}

func fillStruct(data interface{}, fill interface{}, nullData interface{}, mustFill bool) (interface{}, error) {
	dataType := reflect.TypeOf(data).Elem()
	dataValue := reflect.ValueOf(data).Elem()
	fillType := reflect.TypeOf(fill)
	fillValue := reflect.ValueOf(fill)

	var nullValue reflect.Value

	if nullData != nil {
		nullValue = reflect.ValueOf(nullData).Elem()
	}

	// new struct
	newDataType := reflect.TypeOf(fill)
	newDataValue := reflect.New(reflect.TypeOf(fill)).Elem()
	for i := 0; i < newDataType.NumField(); i++ {
		if !newDataValue.Field(i).IsValid() || !newDataValue.Field(i).CanSet() {
			return nil, errors.New("model value cannot be filled")
		}

		if nullData != nil {
			if isNull(fillValue.Field(i), nullValue.Field(i)) {
				continue
			}
		}

		// fill original data
		isFilled := false
		for j := 0; j < dataType.NumField(); j++ {
			if isSameField(newDataType.Field(i), dataType.Field(j)) {
				newDataValue.Field(i).Set(dataValue.Field(j))
				isFilled = true
				break
			}
		}
		if !mustFill && isFilled {
			continue
		}

		// fill input data
		for j := 0; j < fillType.NumField(); j++ {
			if fillValue.Field(j).Kind() == reflect.Ptr {
				// if not set, then do not fill
				if fillValue.Field(j).IsNil() {
					continue
				}

			}

			if isSameField(newDataType.Field(i), fillType.Field(j)) {
				newDataValue.Field(i).Set(fillValue.Field(j))
				break
			}

		}
	}

	return newDataValue.Addr().Interface(), nil
}

// out must be a struct pointer
func (h *Helper) Create(outPtr interface{}) error {

	defaultData := outPtr.(model.Modeller)
	inData, err := fillStruct(outPtr, defaultData.Default(), nil, false)
	if err != nil {
		return err
	}

	// validate data
	if err := validator.New().Struct(inData); err != nil {
		//// this check is only needed when your code could produce
		//// an invalid value for validation such as interface with nil
		//// value most including myself do not usually have code like this.
		//if _, ok := err.(*validator.InvalidValidationError); ok {
		//	fmt.Println(err)
		//	return
		//}
		//
		//for _, err := range err.(validator.ValidationErrors) {
		//
		//	fmt.Println(err.Namespace())
		//	fmt.Println(err.Field())
		//	fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
		//	fmt.Println(err.StructField())     // by passing alt name to ReportError like below
		//	fmt.Println(err.Tag())
		//	fmt.Println(err.ActualTag())
		//	fmt.Println(err.Kind())
		//	fmt.Println(err.Type())
		//	fmt.Println(err.Value())
		//	fmt.Println(err.Param())
		//	fmt.Println()
		//}
		//
		//// from here you can create your own error messages in whatever language you wish
		return err
	}

	// create record
	if err := h.DB().Create(inData).Error; err != nil {
		return err
	}

	if err := copier.Copy(outPtr, inData); err != nil {
		return err
	}

	return nil
}

// outPtr must be a struct pointer, modify and nullDataPtr must be a same type
func (h *Helper) Save(outPtr interface{}, modify interface{}, nullDataPtr ...interface{}) error {

	if err := h.DB().Debug().Model(outPtr).UpdateColumns(modify).Error; err != nil {
		return err
	}

	return nil

}

func (h *Helper) SaveByID(id interface{}, outPtr interface{}, modify interface{}) error {

	return h.Save(outPtr, modify)
}

// outPtr must be a struct pointer
// outPtr必须是结构指针
func (h *Helper) First(outPtr interface{}, withTrashed bool) error {
	_db := h.DB()
	if withTrashed {
		_db = _db.Unscoped()
	}
	if err := _db.Where(outPtr).First(outPtr).Error; err != nil {
		return err
	}
	return nil

}

func (h *Helper) FirstForUpdate(outPtr interface{}, withTrashed bool) error {
	_db := h.DB().Set("gorm:query_option", "FOR UPDATE")
	if withTrashed {
		_db = _db.Unscoped()
	}
	if err := _db.Where(outPtr).First(outPtr).Error; err != nil {
		return err
	}
	return nil
}

func (h *Helper) Delete(in interface{}, force bool) error {
	_db := h.DB()
	if force {
		_db = _db.Unscoped()
	}
	if err := _db.Delete(in).Error; err != nil {
		return err
	}
	return nil
}

func deletedAtKeyName(in interface{}) (string, error) {
	dataType, ok := reflect.TypeOf(in).Elem().FieldByName("DeletedAt")
	if !ok {
		return "", errors.New("cannot get DeletedAt key name type")
	}

	gormTag := dataType.Tag.Get("gorm")
	for _, gormTagType := range strings.Split(gormTag, ";") {
		tmp := strings.Split(gormTagType, ":")
		key := tmp[0]
		value := tmp[1]
		if key == "column" {
			return value, nil
		}
	}
	return "", errors.New("cannot get DeletedAt key name")
}

func (h *Helper) Restore(in interface{}) error {
	deletedAtKeyName, err := deletedAtKeyName(in)
	if err != nil {
		return err
	}
	if err := h.DB().Unscoped().Model(in).Update(deletedAtKeyName, gorm.Expr("NULL")).Error; err != nil {
		return err
	}
	return nil
}

func (h *Helper) Query(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) *gorm.DB {
	_db := h.DB()

	/*
		if len(sortArr) > 0{
			for _, value := range sortArr {
				if value.Key != ""{
					_db = _db.Debug().Order(value.Key + " " + value.Direction.String())
				}
			}
		}

		if limit != 0 {
			_db = _db.Debug().Limit(limit)
		}

		if withTrashed {
			_db = _db.Unscoped()
		}
	*/
	return _db
}

func (h *Helper) Q(filterArr []model.Filter, sortArr []model.Sort, limit int, withTrashed bool) *gorm.DB {
	_db := mapFilter(h.DB(), filterArr)
	if len(sortArr) > 0 {
		for _, value := range sortArr {
			_db = _db.Debug().Order(value.Key + " " + value.Direction.String())
		}
	}

	if limit != 0 {
		_db = _db.Debug().Limit(limit)
	}

	if withTrashed {
		_db = _db.Unscoped()
	}

	return _db
}

func mapFilter(_db *gorm.DB, filterArr []model.Filter) *gorm.DB {
	/*
		if len(filterArr) > 0 {

			for _, filter := range filterArr {

				switch filter.Op {

				case "is_null":
					_db = _db.Debug().Where(filter.Key + " is null")
					break
				case "is_not_null":
					_db = _db.Debug().Where(filter.Key + " is not null")
					break

				// xxx in array
				case "in":
					f, ok := filter.Val.([]string)
					if !ok {
						// error cannot parse array conditions for in
						panic(errors.New("cannot parse array conditions for IN"))
					}
					_db = _db.Debug().Where(filter.Key + " in (?)", f)
					break

				// xxx not in array
				case "not_in":
					f, ok := filter.Val.([]string)
					if !ok {
						// error cannot parse array conditions for in
						panic(errors.New("cannot parse array conditions for NOT IN"))
					}
					_db = _db.Where(filter.Key+" not in (?)", f)
					break

				// xxx between array
				case "between":
					f, ok := filter.Val.([]string)
					if !ok || len(f) != 2 {
						panic(errors.New("cannot parse array conditions for BETWEEN"))
					}
					_db = _db.Debug().Where(filter.Key+" between ? and ?", f[0], f[1])
					break

				// xxx like><!= yyy
				default:
					_db = _db.Debug().Where(filter.Key+" "+filter.Op+" ?", filter.Val)
				}
			}
		}
	*/
	return _db
}

/*
*

	pangxianfei by ad count int64
*/
func (h *Helper) Count(in model.Modeller, filterArr []model.Filter, withTrashed bool) (count int64, err error) {
	err = h.Q(filterArr, []model.Sort{}, 0, withTrashed).Model(in).Count(&count).Error
	return count, err
}

func (h *Helper) Exist(in model.Modeller, withTrashed bool) (exist bool) {
	err := h.First(in, withTrashed)
	if err == nil {
		return true
	}
	return false
}

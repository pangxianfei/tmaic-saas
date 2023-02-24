package model

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"gorm.io/gorm"

	"github.com/iancoleman/strcase"

	"tmaic/vendors/framework/database"
)

type OrderModeller interface {
	DB() *gorm.DB
	SetTX(db *gorm.DB)
}

type OrderModel struct {
	db *gorm.DB
}

func (bm *OrderModel) DB() *gorm.DB {
	if bm.db == nil {
		//_, bm.db = database.Initialize("order")
		bm.db = database.DB()
	}
	return bm.db
}

func (bm *OrderModel) SetTX(db *gorm.DB) {
	bm.SetDB(db)
}
func (bm *OrderModel) SetDB(db *gorm.DB) {
	bm.db = db
}
func (bm *OrderModel) SetTableName(tableName string) string {
	return fmt.Sprintf("%s%s", database.Prefix(), tableName)
}

func (bm *OrderModel) BeforeSave(scope *gorm.DB) (err error) {
	defer func() {
		if _err := recover(); _err != nil {
			if __err, ok := _err.(error); ok {
				err = __err
				return
			}
			err = errors.New(fmt.Sprint(_err))
			return
		}
	}()
	callMutator(scope, false)
	return nil
}
func (bm *OrderModel) BeforeCreate(scope *gorm.DB) (err error) {
	defer func() {
		if _err := recover(); _err != nil {
			if __err, ok := _err.(error); ok {
				err = __err
				return
			}
			err = errors.New(fmt.Sprint(_err))
			return
		}
	}()
	callMutator(scope, false)
	return nil
}

func (bm *OrderModel) BeforeUpdate(scope *gorm.DB) (err error) {
	defer func() {
		if _err := recover(); _err != nil {
			if __err, ok := _err.(error); ok {
				err = __err
				return
			}
			err = errors.New(fmt.Sprint(_err))
			return
		}
	}()
	callMutatorOrder(scope, false)
	return nil
}

func (bm *OrderModel) AfterFind(scope *gorm.DB) (err error) {
	defer func() {
		if _err := recover(); _err != nil {
			if __err, ok := _err.(error); ok {
				err = __err
				return
			}
			err = errors.New(fmt.Sprint(_err))
			return
		}
	}()
	callMutatorOrder(scope, true)
	return nil
}

func callMutatorOrder(scope *gorm.DB, isGetter bool) {

	var reflectValue reflect.Value
	//只从非指针获取地址
	if reflectValue.CanAddr() && scope.Statement.ReflectValue.Kind() != reflect.Ptr {
		reflectValue = scope.Statement.ReflectValue
	}

	for i := 0; i < len(scope.Statement.Schema.Fields); i++ {
		switch scope.Statement.ReflectValue.Kind() {
		case reflect.Struct:
			structReflectOrder(&scope.Statement.ReflectValue, isGetter)
		case reflect.Slice:
			for i := 0; i < reflectValue.Elem().Len(); i++ {
				rv := reflectValue.Elem().Index(i)
				structReflectOrder(&rv, isGetter)
			}
		default:
			panic("cannot use mutator in type:" + reflectValue.Type().Elem().Kind().String())
		}
	}
}

func structReflectOrder(reflectValue *reflect.Value, isGetter bool) {
	wg := &sync.WaitGroup{}
	if reflectValue.CanAddr() && reflectValue.Kind() != reflect.Ptr {
		tmp := reflectValue.Addr()
		reflectValue = &tmp
	}
	for i := 0; i < reflectValue.Type().Elem().NumField(); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, index int) {
			fieldName := reflectValue.Type().Elem().Field(index).Name
			fieldValue := reflectValue.Elem().Field(index).Interface()
			if isGetter {
				getterOrder(reflectValue, fieldName, fieldValue)
			} else {
				setterOrder(*reflectValue, fieldName, fieldValue)
			}
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
}
func setterOrder(reflectValue reflect.Value, fieldName string, fieldValue interface{}) {
	const setterMethodTemplate = "Set%sAttribute"
	methodName := fmt.Sprintf(setterMethodTemplate, strcase.ToCamel(fieldName))
	if methodValue := reflectValue.MethodByName(methodName); methodValue.IsValid() {
		methodValue.Interface().(func(value interface{}))(fieldValue)
	}
}
func getterOrder(reflectValue *reflect.Value, fieldName string, fieldValue interface{}) {
	const getterMethodTemplate = "Get%sAttribute"
	methodName := fmt.Sprintf(getterMethodTemplate, strcase.ToCamel(fieldName))
	if methodValue := reflectValue.MethodByName(methodName); methodValue.IsValid() {
		newGetData := methodValue.Interface().(func(value interface{}) interface{})(fieldValue)
		reflectValue.Elem().FieldByName(fieldName).Set(reflect.ValueOf(newGetData))
	}
}

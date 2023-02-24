package model

import (
	"errors"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"math"
	"reflect"
	"strconv"
)

type Model gorm.DB

type Pagination struct {
	currentPageItemCount int
	currentPageNum       int
	totalPageNum         int
	totalItemCount       int64
	itemArr              interface{}
	perPage              int
}

func (p *Pagination) Count() int {
	return p.currentPageItemCount
}
func (p *Pagination) CurrentPage() int {
	return p.currentPageNum
}
func (p *Pagination) LastPage() int {
	return p.totalPageNum
}
func (p *Pagination) FirstItem() interface{} {

	if len(p.itemArr.([]interface{})) > 0 {
		return p.itemArr.([]interface{})[0]
	}

	return nil
}
func (p *Pagination) LastItem() interface{} {

	if len(p.itemArr.([]interface{})) > 0 {
		return p.itemArr.([]interface{})[len(p.itemArr.([]interface{}))-1]
	}

	return nil
}
func (p *Pagination) ItemArr() interface{} {
	return p.itemArr
}
func (p *Pagination) Total() int64 {
	return p.totalItemCount
}
func (p *Pagination) PerPage() int {
	return p.perPage
}

func (bm *Model) Paginate(model interface{}, c iris.Context, perPage int) (pagination Pagination, err error) {

	p := c.URLParamIntDefault("page", 1)

	if err := validator.New().Var(p, "numeric,gt=0"); err != nil {

		return pagination, err
	}

	_p, err := strconv.ParseUint(strconv.Itoa(p), 10, 32)
	if err != nil {
		return pagination, err
	}
	page := int(_p)

	pagination.perPage = perPage

	pagination.currentPageNum = page

	count := gorm.DB(*bm)

	totalItemCount := &pagination.totalItemCount

	if err = count.Model(model).Count(totalItemCount).Error; err != nil {

		return pagination, err
	}

	pagination.totalPageNum = int(math.Ceil(float64(pagination.totalItemCount) / float64(perPage)))

	data := gorm.DB(*bm)
	if err = data.Offset(int(perPage * (page - 1))).Limit(int(perPage)).Find(model).Error; err == nil {
		return pagination, err
	}
	pagination.itemArr = model

	if reflect.ValueOf(model).Elem().Type().Kind() != reflect.Slice {
		return pagination, errors.New("result is not a slice")
	}
	s := reflect.ValueOf(model).Elem()
	pagination.currentPageItemCount = int(s.Len())
	return pagination, nil
}

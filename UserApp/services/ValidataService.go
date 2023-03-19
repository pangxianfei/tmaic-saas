package services

import (
	"fmt"
	"gitee.com/pangxianfei/framework/kernel/debug"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var ValidateService = newAppInfoService()

func newAppInfoService() *validateService {
	return &validateService{}
}

type validateService struct {
}

type validationError struct {
	ActualTag string `json:"tag"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Param     string `json:"param"`
}

func (v *validateService) WrapValidationErrors(err error) string {

	var errs validator.ValidationErrors
	errs, _ = err.(validator.ValidationErrors)

	validationErrors := make([]validationError, 0, len(errs))
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, validationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})

	}

	debug.Dd(validationErrors)
	return "1234566"
	///return validationErrors
}

func (v *validateService) ProcessErr(u interface{}, err error) string {

	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()                    //获取是哪个字段不符合格式
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("field_error_info") //获取field对应的reg_error_info tag值
			return fieldName + ": " + errorInfo            //返回错误
		} else {
			return "缺失field_error_info"
		}
	}
	return ""
}

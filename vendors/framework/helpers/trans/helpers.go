package trans

import (
	"tmaic/vendors/framework/resources/lang"

	"gopkg.in/go-playground/validator.v9"

	"tmaic/vendors/framework/resources/lang/helper"
)

func ValidationTranslate(v *validator.Validate, langName string, e validator.ValidationErrors) lang.ValidationError {
	return helper.ValidationTranslate(v, langName, e)
}
func CustomTranslate(messageID string, data map[string]interface{}, langName string) string {
	return helper.CustomTranslate(messageID, data, langName)
}

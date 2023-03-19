package requests

import (
	"github.com/kataras/iris/v12"
)

func ReadJSON(cxt iris.Context, data interface{}) error {
	if err := cxt.ReadJSON(data); err != nil {
		return err
	}
	return nil
}

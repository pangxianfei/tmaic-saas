package log

import (
	"tmaic/vendors/framework/errors"
	"tmaic/vendors/framework/helpers/toto"
	"tmaic/vendors/framework/logs"
)

func Error(err error, v ...toto.V) error {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	errors.ErrPrintln(err, fields)
	return err
}

func Info(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.INFO, msg, fields)
}
func Warn(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.WARN, msg, fields)
}
func Fatal(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.FATAL, msg, fields)
}
func Debug(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.DEBUG, msg, fields)
}
func Panic(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.PANIC, msg, fields)
}
func Trace(msg interface{}, v ...toto.V) {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.TRACE, msg, fields)
}
func ErrorStr(err error, v ...toto.V) string {
	var fields toto.V
	if len(v) > 0 {
		fields = v[0]
	}
	return errors.ErrPrint(err, 2, fields)
}

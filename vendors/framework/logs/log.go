package logs

import (
	"github.com/sirupsen/logrus"
	"os"
	"tmaic/vendors/framework/helpers/tmaic"

	"tmaic/vendors/framework/config"
)

var log *logrus.Logger
var logLevel Level

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.Out = os.Stdout
}

func Initialize() {
	levelStr := config.GetString("app.log_level")
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		panic(err)
	}

	logLevel = level
	log.SetLevel(logLevel)
}

type Field = tmaic.V

func Println(level Level, msg interface{}, fields Field) {

	if fields == nil {
		log.Log(level, msg)
	} else {
		log.WithFields(logrus.Fields(fields)).Log(level, msg)
	}

	if level <= logLevel {
		_fields := make(map[string]interface{})
		if fields != nil {
			_fields = fields
		}

		switch level {
		case PANIC:
			//sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
		case FATAL:
			//sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
		case ERROR:
			//sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
		case WARN:
			_fields["level"] = "WARN"
			//sentry.CaptureMsg(fmt.Sprintf("%v", msg), _fields)
		}
	}
}

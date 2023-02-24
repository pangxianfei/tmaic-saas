package helpers

import (
	"tmaic/vendors/framework/helpers/trans"
)

func L(messageID string, dataNlocale ...interface{}) string {

	data := make(map[string]interface{})
	switch len(dataNlocale) {
	case 1:
		data = dataNlocale[0].(map[string]interface{})
		break
	default:
	}

	return trans.CustomTranslate(messageID, data)
}

package m

import (
	"tmaic/vendors/framework/database"
	"tmaic/vendors/framework/model/helper"
)

var h helper.Helper

type Helper = helper.Helper

func Initialize() {
	h = helper.Helper{}
	h.SetDB(database.DB())
}

func H() *helper.Helper {
	return &h
}

func GetDB() *helper.Helper {
	return &h
}

func Transaction(tf func(TransactionHelper *helper.Helper), attempts uint) {
	helper.Transaction(tf, attempts)
}

package Orm

import (
	"tmaic/vendors/framework/database"
	"tmaic/vendors/framework/model/helper"
)

var h helper.Helper
var DB = GetDB()

type Helper = helper.Helper

func Initialize() {
	h = helper.Helper{}
	h.SetDB(database.DB())
}

func GetDB() *helper.Helper {
	return &h
}

func Transaction(tf func(TransactionHelper *helper.Helper), attempts uint) {
	helper.Transaction(tf, attempts)
}

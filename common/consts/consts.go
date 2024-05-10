package consts

import "gorm.io/gorm"

var ERROR_NOT_FOUND = gorm.ErrRecordNotFound

const (
	GENDER_MAN = iota + 1
	GENDER_WOMAN

	Inter string = "http://192.168.152.129:11434"
)

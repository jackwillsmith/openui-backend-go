package utils

import "github.com/google/uuid"

var ollurl string

func SetOllUrl(url string) {
	ollurl = url
}

func GetOllUrl() string {
	return ollurl
}

// uuid库创建随机 uuid
func GengerateUUID() string {
	return uuid.New().String()
}

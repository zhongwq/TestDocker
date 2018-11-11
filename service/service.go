package service

import (
	"github.com/zhongwq/TestDocker/Model"
	"github.com/zhongwq/TestDocker/database"
	"fmt"
)

func GetUserInfo(info map[string][]string) (Model.User, string) {
	if info[`username`] == nil || info[`password`] == nil {
		fmt.Println("Login: Error Parameter", info)
		return Model.User{"", "", "", ""}, "Login: Error Parameter"
	}
	return database.GetUserInfo(info[`username`][0], info[`password`][0])
}

func UserRegister(info map[string][]string) (bool, string) {
	if info[`username`] == nil || info[`password`] == nil || info[`email`] == nil || info[`phone`] == nil {
		return false, "UserRegister: Error Parameter"
	}

	return database.InsertUser(info[`username`][0], info[`password`][0], info[`email`][0], info[`phone`][0])
}
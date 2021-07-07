package main

import (
	"fmt"
	"gingin/model"
)

func main() {
	model.InitDB()
	userModel := new(model.UserModel)
	user,err := userModel.FindMany("he",1,10)
	fmt.Println(user)
	fmt.Println(err)
}
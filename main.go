package main

import (
	"fmt"
	"gin-restful-api-mysql/Config"
	"gin-restful-api-mysql/Models"
	"gin-restful-api-mysql/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	fmt.Println("Golang gin restful api mysql")
	Config.DB, err = gorm.Open("mysql", Config.DBURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	r.Run()

}

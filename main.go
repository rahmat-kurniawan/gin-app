package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rahmat-kurniawan/gin-app/Config"
	"github.com/rahmat-kurniawan/gin-app/Models"
	"github.com/rahmat-kurniawan/gin-app/Routers"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/mux_app?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routers.SetupRouter()
	r.Run()
}

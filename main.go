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

	Config.DB, err = gorm.Open("postgres", "user=admin password=Admin@123 dbname=go_app port=5432 sslmode=disable TimeZone=Asia/Jakarta")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{}, &Models.Song{})

	r := Routers.SetupRouter()
	r.Run()
}

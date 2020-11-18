package app

import (
	"github.com/gin-gonic/gin"	"github.com/rahmat-kurniawan/gin-app/controllers/song"
)



func Init() {
	router := gin.Default()

	router.GET("/songs", song.GetAll)
	router.GET("/song/:id", song.Get)
	router.POST("/song", song.Create)
	router.PUT("/song/:id", song.Update)
	router.DELETE("/song/:id", song.Delete)

	router.Run()
}

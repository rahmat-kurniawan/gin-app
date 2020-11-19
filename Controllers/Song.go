package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmat-kurniawan/gin-app/Helpers"
	"github.com/rahmat-kurniawan/gin-app/Models"
)

func ListSong(c *gin.Context) {
	var song []Models.Song
	err := Models.GetAllSong(&song)
	if err != nil {
		Helpers.RespondJSON(c, 404, song)
	} else {
		Helpers.RespondJSON(c, 200, song)
	}
}

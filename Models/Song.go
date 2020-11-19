package Models

import (
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/rahmat-kurniawan/gin-app/Config"
)

func GetAllSong(s *[]Song) (err error) {
	if err = Config.DB.Find(s).Error; err != nil {
		return err
	}
	log.Println(err)
	return nil
}

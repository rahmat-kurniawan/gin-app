package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmat-kurniawan/gin-app/Helpers"
	"github.com/rahmat-kurniawan/gin-app/Models"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		Helpers.RespondJSON(c, 404, book)
	} else {
		Helpers.RespondJSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.AddNewBook(&book)
	if err != nil {
		Helpers.RespondJSON(c, 404, book)
	} else {
		Helpers.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, book)
	} else {
		Helpers.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, book)
	} else {
		Helpers.RespondJSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		Helpers.RespondJSON(c, 404, book)
	} else {
		Helpers.RespondJSON(c, 200, book)
	}
}

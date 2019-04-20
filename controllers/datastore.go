package controllers

import (
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"github.com/iPolyomino/awacha/models"
	"google.golang.org/appengine"
)

func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)

		address := c.Param("address")

		text, err := models.GetNote(ctx, address)
		if err == datastore.ErrNoSuchEntity {
			c.String(http.StatusNotFound, "Not Found\n")
			return
		} else if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error\n")
			return
		}
		c.String(http.StatusOK, text+"\n")
	}
}

func PutNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)

		address := c.Param("address")

		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		text := string(buf[0:n])

		note := models.Note{
			Address: address,
			Text:    text,
		}
		_, err := models.PutNote(ctx, note)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error\n")
			return
		}
		c.String(http.StatusCreated, "Created\n")
	}
}

func DeleteNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)

		address := c.Param("address")

		err := models.DeleteNote(ctx, address)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error\n")
			return
		}
		c.String(http.StatusOK, "Deleted\n")
	}
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iPolyomino/awacha/models"
	"google.golang.org/appengine"
)

func StoreingData() gin.HandlerFunc {
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
		_, err := models.StoreingData(ctx, note)
		if err != nil {
			c.String(http.StatusInternalServerError, "The value was not saved")
			return
		}
		c.String(http.StatusOK, "saved")
	}
}

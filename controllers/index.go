package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.Header.Get("User-Agent"), "curl") {
			c.File("views/awacha.txt")
		} else {
			c.HTML(http.StatusOK, "index.html", nil)
		}
	}
}

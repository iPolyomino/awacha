package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iPolyomino/awacha/controllers"
)

func Run() {
	r := gin.New()

	r.LoadHTMLGlob("views/*.html")
	group := r.Group("/")
	{
		group.GET("/", indexHandler())
		group.GET("/:address", controllers.FetchData())
		group.POST("/:address", controllers.StoreingData())
	}

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

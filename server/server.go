package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iPolyomino/awacha/controllers"
)

func Run() {
	r := gin.New()

	r.LoadHTMLGlob("views/*.html")
	group := r.Group("/")
	{
		group.GET("/", controllers.IndexHandler())
		group.GET("/:address", controllers.GetNote())
		group.POST("/:address", controllers.PutNote())
		group.DELETE("/:address", controllers.DeleteNote())
	}

	http.Handle("/", http.TimeoutHandler(r, 10*time.Second, "timeout"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

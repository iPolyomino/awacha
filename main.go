package main

import (
	"github.com/iPolyomino/awacha/server"
	"google.golang.org/appengine"
)

func main() {
	server.Run()
	appengine.Main()
}

package main

import (
	"post-article/connection"
	"post-article/handlers"
)

func main() {
	connection.Connect()
	handlers.HandleReq()
}

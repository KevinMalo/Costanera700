package main

import (
	"github.com/kevinmalo/Costanera700/internal/logs"
)


func main()  {

	_ = logs.InitLogger()

	//Run server
	mux := Routes()
	server := NewServer(mux)
	server.Run()

}

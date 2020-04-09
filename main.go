package main

import (
	"./db"
	"./service"
	)

func main()  {
	done := make(chan bool, 1)

	db.Init()
	service.GRPCInit()

	<-done
}

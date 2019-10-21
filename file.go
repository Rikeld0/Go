package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"./log"

	"./service"
)

func main() {
	service.Init()
	log.Init()
	defer service.Close()
	service.Run()

	service.InitDataInst(&service.ConfigData{Filename: "list.json"})
	inst := service.GetInsData()
	inst.Data()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C for quit.")
	<-c

}

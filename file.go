package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"./service"
)

func main() {
	service.Init()
	defer service.Close()
	availabilityResources()
	//getResources()
	service.Run()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C for quit.")
	<-c
}

//Emulation request from the outside
func availabilityResources() {
	data, err := ioutil.ReadFile("list.json")
	if err != nil {
		service.LogG.Error("file not found")
	}
	fmt.Println(string(data))

	var ss map[string]interface{}
	json.Unmarshal(data, &ss)
	q := ss["list"].([]interface{})
	for _, v := range q {
		url := "https://" + v.(string)
		var r string
		resp, err := http.Get(url)
		if err != nil {
			r = "404 FATAL"
		} else {
			r = resp.Status
		}
		service.Db.Put(url, r)
	}

}

/*func getResources() {
	data, _ := ioutil.ReadFile("list.json")
	fmt.Println(string(data))

	var ss map[string]interface{}
	json.Unmarshal(data, &ss)
	q := ss["list"].([]interface{})
	for _, v := range q {
		url := "https://" + v.(string)
		service.Get(url)
	}
}*/

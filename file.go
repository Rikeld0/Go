package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	idb "./service"
)

func main() {
	idb.Init()
	defer idb.Close()
	availabilityResources()
	getResources()
}

//Emulation request from the outside
func availabilityResources() {
	data, _ := ioutil.ReadFile("list.json")
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
		idb.Put(url, r)
	}

}

func getResources() {
	data, _ := ioutil.ReadFile("list.json")
	fmt.Println(string(data))

	var ss map[string]interface{}
	json.Unmarshal(data, &ss)
	q := ss["list"].([]interface{})
	for _, v := range q {
		url := "https://" + v.(string)
		idb.Get(url)
	}
}

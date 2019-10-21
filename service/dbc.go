package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../log"
)

type ConfigData struct {
	Filename string
	Urlname  string
}

var inst Idata = nil

type file struct {
	Idata
	file string
}

type web struct {
	Idata
	url string
}

func (this *file) Data() {
	data, err := ioutil.ReadFile(this.file)
	if err != nil {
		log.Fatal("json file not found")
	}
	var ss map[string]interface{}
	json.Unmarshal(data, &ss)
	q := ss["list"].([]interface{})
	qq := Db.GetAll()
	for _, v := range q {
		var find bool = false
		url := "https://" + v.(string)
		for _, w := range qq {
			if w == url {
				find = true
				break
			}
		}
		if find == false {
			var r string
			resp, err := http.Get(url)
			if err != nil {
				r = "404 FATAL"
			} else {
				r = resp.Status
			}
			Db.Put(url, r)
		}
	}
}

func (this *web) Data() {
	resp, err := http.Get(this.url)
	if err != nil {
		log.Fatal("url not found")
	}
	data, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatal("err")
	}
	var ss map[string]interface{}
	json.Unmarshal(data, &ss)
	q := ss["list"].([]interface{})
	qq := Db.GetAll()
	for _, v := range q {
		url := "https://" + v.(string)
		fmt.Println(qq)
		var r string
		resp, err := http.Get(url)
		if err != nil {
			r = "404 FATAL"
		} else {
			r = resp.Status
		}
		Db.Put(url, r)
	}
}

//For access in object
func GetInsData() Idata {
	if inst == nil {
		log.Fatal("don't init")
	}
	return inst
}

func InitDataInst(data *ConfigData) {
	if data.Urlname != "" && data.Filename != "" {
		log.Fatal("use only Filename or Urlname")
	}
	if data.Urlname == "" {
		if data.Filename == "" {
			log.Fatal("Filename not found")
		}
		inst = &file{file: data.Filename}
	} else {
		inst = &web{url: data.Urlname}
	}
}

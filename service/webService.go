package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	mux "github.com/gorilla/mux"
)

type httpService struct {
	server *http.Server
}

type jsonStruct struct {
	Status string
}

func (srv *httpService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	buf := req.URL.String()
	paramPars, _ := url.Parse(buf)
	paramURL := paramPars.Query().Get("url")
	paramURLr := strings.Replace(paramURL, "'", "", len(paramURL))
	if paramURLr != paramURL || len(paramURL) == 0 {
		return
	}
	status := &jsonStruct{Status: Get(paramURLr)}
	JSON, _ := json.Marshal(status)
	fmt.Fprint(resp, string(JSON))
}

func (srv *httpService) Start() {
	if srv.server != nil {
		panic("Server already started.")
	}

	router := mux.NewRouter()

	router.HandleFunc("/get", srv.ServeHTTP).Methods("GET")

	srv.server = &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: router,
	}

	go func() {
		if err := srv.server.ListenAndServe(); err != nil {
			panic("Failed to listen.")
		}
	}()
}

//Run ...
func Run() {
	server := httpService{}
	fmt.Println("Server is listening...")
	server.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	fmt.Println("Press Ctrl+C for quit.")
	<-c
}

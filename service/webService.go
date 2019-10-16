package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
	paramPars, err := url.Parse(buf)
	if err != nil {
		LogG.Error("error")
	}
	paramURL := paramPars.Query().Get("url")
	paramURLr := strings.Replace(paramURL, "'", "", len(paramURL))
	if paramURLr != paramURL || len(paramURL) == 0 {
		return
	}
	status := &jsonStruct{Status: Db.Get(paramURLr)}
	JSON, err1 := json.Marshal(status)
	if err1 != nil {
		LogG.Error("error")
	}
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
}

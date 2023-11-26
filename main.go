package main

import (
	"fmt"
	"log"
	"net/http"
	"test-vscode-go-module/model"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello From VSCode!"))
}

func main() {
	log.Println("Read config")
	var cfg model.Config
	cleanenv.ReadConfig("config.yml", &cfg)

	router := httprouter.New()
	router.GET("/", IndexHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}

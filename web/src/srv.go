package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var config *Config
var ih *ImageHandler
var th = new(TextHandler)

func predict(w http.ResponseWriter, r *http.Request) {
	var tags []string
	var tessText string
	_ = r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println("No Image")
	} else {
		defer file.Close()
		tags, tessText = ih.handle(file, handler)
	}
	log.Printf("Iai resualt:%v", tags)
	log.Printf("Tesseract result:%v", tessText)
	text := r.PostFormValue("text")
	isMostPop, _ := strconv.ParseBool(r.PostFormValue("isMostPop"))
	city := r.PostFormValue("location")
	if len(text) != 0 || len(tessText) != 0 {
		text += tessText
		tags = append(tags, th.handle(text)...)
	}
	result := mineTag(tags, isMostPop, strings.ToLower(city))
	body, _ := json.Marshal(result)
	_, _ = w.Write(body)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	config = NewConfig()
	log.Println(config)
	ih = NewImageHandler(config.UrlYolo, config.UrlTess)
	*th = TextHandler{UrlNer: config.UrlNER, UrlTr: config.UrlTR}
	r := mux.NewRouter()
	r.Handle("/predict", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(predict))).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Println("Start listening on port 80")
	_ = http.ListenAndServe(":80", r)
}

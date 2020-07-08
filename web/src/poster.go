package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func post(url string, body bytes.Buffer, contentType string, pipe chan []byte) {
	defer close(pipe)
	var res []byte
	resp, err := http.Post(url, contentType, &body)
	if err != nil {
		log.Println(err)
		pipe<-res
		return
	}
	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	pipe<-res
	return
}

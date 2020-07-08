package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type TextHandler struct {
	UrlNer string
	UrlTr  string
}

type requestNer struct {
	X []string `json:"x"`
}

func (th TextHandler) handle(text string) (result []string) {
	var rp responseNer
	bodyNer, _ := json.Marshal(requestNer{X: []string{text}})
	buf := bytes.NewBuffer(bodyNer)
	ner := make(chan []byte)
	tr := make(chan []byte)
	go post(th.UrlNer, *buf, "application/json", ner)
	go post(th.UrlTr, *buf, "application/json", tr)
	if err := json.Unmarshal(<-tr, &result); err != nil {
		log.Println("TextRank was broken")
	}
	log.Printf("TextRank:%v", result)
	if err := json.Unmarshal(<-ner, &rp); err != nil {
		log.Println("Ner was broken")
	}
	log.Print(rp)
	if len(rp) != 0 {
		result = append(result, rp.unionNerTag()...)
	}
	return
}

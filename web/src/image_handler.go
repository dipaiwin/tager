package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type ImageHandler struct {
	UrlYolo    string
	UrlTess    string
	Extensions map[string]bool
}

func NewImageHandler(UrlYolo, UrlTess string) *ImageHandler {
	if UrlYolo == "" || UrlTess == "" {
		log.Println("Are you sure that should use an empty Url?")
	}
	return &ImageHandler{
		UrlYolo:    UrlYolo,
		UrlTess:    UrlTess,
		Extensions: map[string]bool{"jpeg": true, "jpg": true, "png": true},
	}
}

func (ih *ImageHandler) IsAvailable(filename string) (isImage bool) {
	splitName := strings.Split(filename, ".")
	return ih.Extensions[splitName[len(splitName)-1]]
}

func (ih *ImageHandler) handle(file multipart.File, fh *multipart.FileHeader) (objects []string, text string) {
	if !ih.IsAvailable(fh.Filename) {
		log.Println("File format is not available")
		return
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", filepath.Base(fh.Filename))
	_, _ = io.Copy(part, file)
	_ = writer.Close()
	ct := writer.FormDataContentType()
	do := make(chan []byte)
	tess := make(chan []byte)
	go post(ih.UrlYolo, *body, ct, do)
	go post(ih.UrlTess, *body, ct, tess)
	if err := json.Unmarshal(<-do, &objects); err != nil {
		log.Println("Detection object was broken")
	}
	if err := json.Unmarshal(<-tess, &text); err != nil {
		log.Println("Tesseract was broken")
	}
	return
}

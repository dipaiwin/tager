package main

import (
	"fmt"
	"os"
	"reflect"
)

// Config - структура для хранения параметров из env-ов
type Config struct {
	UrlYolo string
	UrlTess string
	UrlTR   string
	UrlNER  string
}

// Считывание заданных env-ов
func NewConfig() *Config {
	return &Config{
		UrlYolo: getEnv("URL_YOLO", ""),
		UrlTess: getEnv("URL_TESS", ""),
		UrlTR:   getEnv("URL_TR", ""),
		UrlNER:  getEnv("URL_NER", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func (conf *Config) String() string {
	s := reflect.ValueOf(conf).Elem()
	typeOfT := s.Type()
	out := "Config\n"
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		out += fmt.Sprintf("%s = %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	return out
}

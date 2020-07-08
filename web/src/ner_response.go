package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"unicode"
)

type responseNer [][][]string

func (rp responseNer) String() string {
	var buf bytes.Buffer
	buf.WriteString("Ner result:\n")
	if len(rp) == 1 {
		buf.WriteString(fmt.Sprintf("Out text:%v\nOut tags:%v\n", rp[0][0], rp[0][1]))
	}
	return buf.String()
}

func isPunct(line string) bool {
	arr := []rune(strings.ToUpper(line))
	if len(arr) > 1 {
		return false
	}
	return unicode.IsSpace(arr[0]) || unicode.IsPunct(arr[0])
}

func (rp responseNer) unionNerTag() (result []string) {
	tokText := rp[0][0]
	tagTok := rp[0][1]
	var cache []string
	for index, tag := range tagTok {
		if tag == "O" || isPunct(tokText[index]) {
			continue
		}
		suffix := strings.Split(tag, "-")[0]
		subText := strings.ToLower(tokText[index])
		if suffix == "I" {
			if len(subText) != 1 {
				cache = append(cache, subText)
			}
		} else if suffix == "B" {
			if len(cache) != 0 {
				result = append(result, strings.Join(cache, ""))
				cache = nil
			}
			cache = append(cache, subText)
		}
	}
	if len(cache) != 0 {
		result = append(result, strings.Join(cache, ""))
	}
	log.Printf("Ner result: %v\n", result)
	return
}

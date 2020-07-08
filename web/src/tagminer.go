package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"sync"
)

const maxTags = 30

type tagInfo struct {
	Count       float64
	RelatedTags []string
	Tag         string
}

type tagArray []tagInfo

func (ta tagInfo) String() string {
	return fmt.Sprintf("Tag: %v\nCount: %v\nReleted tags: %v\n", ta.Tag, ta.Count, ta.RelatedTags)
}

func extractInfo(obj map[string]interface{}) (ti tagInfo) {
	var relTags []string
	var ok bool
	for _, key := range []string{"graphql", "hashtag"} {
		obj, ok = obj[key].(map[string]interface{})
		if !ok {
			return
		}
	}
	hm := obj["edge_hashtag_to_media"]
	ti.Count = hm.(map[string]interface{})["count"].(float64)
	rt := obj["edge_hashtag_to_related_tags"]
	edges := rt.(map[string]interface{})["edges"].([]interface{})
	for _, item := range edges {
		mapItem := item.(map[string]interface{})
		node := mapItem["node"].(map[string]interface{})
		relTags = append(relTags, node["name"].(string))
	}
	ti.RelatedTags = relTags
	return
}

func getTagPopularity(tag string, group *sync.WaitGroup, tagPos int, tags []tagInfo) {
	defer group.Done()
	client := &http.Client{}
	url := fmt.Sprintf("https://www.instagram.com/explore/tags/%v/", tag)
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("__a", "1")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	instResult := make(map[string]interface{})
	_ = json.NewDecoder(resp.Body).Decode(&instResult)
	resp.Body.Close()
	res := extractInfo(instResult)
	res.Tag = tag
	tags[tagPos] = res
}

func selectTags(tis []tagInfo) (tags []string) {
	tags = make([]string, 0)
	for i := 0; i < len(tis) && len(tags) < maxTags; i++ {
		value := tis[i]
		tags = append(tags, value.Tag)
		delta := maxTags - len(tags)
		if len(value.RelatedTags) > delta {
			value.RelatedTags = value.RelatedTags[:delta]
		}
		tags = append(tags, value.RelatedTags...)
	}
	return
}

func mineTag(keyWords []string, isMostPop bool, city string) (result []string) {
	log.Printf("Origin keywords: %v\n", keyWords)
	var wg sync.WaitGroup
	tags := make([]tagInfo, len(keyWords))
	for index, tag := range keyWords {
		wg.Add(1)
		go getTagPopularity(tag, &wg, index, tags)
	}
	tiCity := make([]tagInfo, 1)
	if city != "null" {
		wg.Add(1)
		go getTagPopularity(city, &wg, 0, tiCity)
	}
	wg.Wait()
	if isMostPop {
		log.Println("Use most popular")
		sort.SliceStable(tags, func(i, j int) bool { return tags[i].Count > tags[j].Count })
	} else {
		log.Println("Use avg popular")
		var sum float64
		for _, item := range tags {
			sum += item.Count
		}
		avg := sum / float64(len(tags))
		log.Printf("AVG: %v\n", avg)
		sort.SliceStable(tags, func(i, j int) bool { return math.Abs(tags[i].Count-avg) < math.Abs(tags[j].Count-avg) })
	}
	if city != "null" {
		tags = append(tiCity, tags...)
	}
	log.Printf("Top tags: %v\n", tags)
	result = selectTags(tags)
	log.Println(len(result), result)
	return result
}

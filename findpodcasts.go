package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

func Search(query string) ([]string){
	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.podcastindex.org/api/1.0/search/byterm?q=%s&pretty", query), nil)
	if err != nil {
		print("failed")
		os.Exit(1)
	}

	var apiKey string = "3PFHXD4S7QHRL5LTWGDB"
	var apiSecret string = "6GyBFHFz75etzpVLA2LR^KEFHMDDJZ$zPSpmZf75"
	var apiTime string = strconv.FormatInt(time.Now().Unix(), 10)

	var data2hash string = apiKey + apiSecret + apiTime
	h := sha1.New()
	h.Write([]byte(data2hash))
	hash := h.Sum(nil)
	var hashString string = fmt.Sprintf("%x", hash)

	request.Header.Add("User-Agent", "CastCLI/1.0")
	request.Header.Add("X-Auth-Date", apiTime)
	request.Header.Add("X-Auth-Key", apiKey)
	request.Header.Add("Authorization", hashString)
	print(apiTime)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		print("failed")
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print("failed")
		os.Exit(1)
	}

	podcastsList := gjson.Get(string(data), "feeds")
	var podcasts = make([]string, len(podcastsList.Array()))
	for i := range podcastsList.Array(){
		podcasts[i] = podcastsList.Array()[i].Get("title").String()
	}
	return podcasts
}

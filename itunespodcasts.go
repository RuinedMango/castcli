package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

func main(){
	request, err := http.NewRequest("GET", "https://itunes.apple.com/search?term=scp&entity=podcast", nil)
	if err != nil {
		print("failed")
		os.Exit(1)
	}

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

	podcasts := gjson.Get(string(data), "results")

	for i := range podcasts.Array(){
		print(podcasts.Array()[i].Get("collectionName").String())
	}
}

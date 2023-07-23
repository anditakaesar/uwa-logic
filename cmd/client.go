package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ResponseBody struct {
	Number string
	Result string
}

func main() {
	responseBody := &ResponseBody{}
	for i := 0; i < 1000; i++ {
		url := fmt.Sprintf("http://localhost:8888?number=%d", i)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(responseBody)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[main] request num: %d, resp: %s", i, responseBody.Result)
		time.Sleep(1 * time.Second)
	}
}

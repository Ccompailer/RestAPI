package awesomeProject

import (
	"awesomeProject/Data"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {
	http.HandleFunc("/fastest-mirror", func(writer http.ResponseWriter, request *http.Request) {
		response := FindFastestUrl(Data.MirrorsList)
		responseJSON, _ := json.Marshal(response)

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseJSON)
	})
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %sn", port)
	log.Fatal(server.ListenAndServe())
}

func FindFastestUrl(urls []string) Response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return Response{<-urlChan, <-latencyChan}
}

package main

import (
	"fmt"
	httprouter2 "github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type Response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {
	//http.HandleFunc("/fastest-mirror", func(writer http.ResponseWriter, request *http.Request) {
	//	response := FindFastestUrl(Data.MirrorsList)
	//	responseJSON, _ := json.Marshal(response)
	//
	//	writer.Header().Set("Content-Type", "application/json")
	//	writer.Write(responseJSON)
	//})
	//port := ":8000"
	//server := &http.Server{
	//	Addr:           port,
	//	ReadTimeout:    15 * time.Second,
	//	WriteTimeout:   15 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//fmt.Printf("Starting server on port %s", port)
	//log.Fatal(server.ListenAndServe())

	router := httprouter2.New()

	router.ServeFiles("/static/*filepath", http.Dir("C:\\Users\\1\\Desktop"))

	log.Fatal(http.ListenAndServe(":8000", router))
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

func GetFileContent(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter2.Params) {
	fmt.Fprintf(writer, GetCommandOutput("/bin/cat", params.ByName("name")))
}

func GetCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func GoVersion(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	response := GetCommandOutput("D:\\GoSDK\\go1.21.5\\bin\\go", "version")
	io.WriteString(w, response)
	return
}

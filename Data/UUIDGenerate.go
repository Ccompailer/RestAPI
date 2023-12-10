package Data

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

type UUID struct {
}

func (p *UUID) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		GiveRandomUUID(writer, request)
		return
	}
	http.NotFound(writer, request)
	return
}

func GiveRandomUUID(writer http.ResponseWriter, request *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(writer, fmt.Sprintf("%x", b))
}

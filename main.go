package main

import (
	"awesomeProject/Data"
	"encoding/json"
	httprouter2 "github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	UserId int `json:"user_id"`
}

var storage Data.Storage = Data.NewMemoryStorage()

func main() {
	router := httprouter2.New()

	router.ServeFiles("/static/*filepath", http.Dir("C:\\Users\\1\\Desktop"))
	router.GET("/api/v1/getEmployee", getEmployeeHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getEmployeeHandler(w http.ResponseWriter, r *http.Request, param httprouter2.Params) {
	val, e := strconv.Atoi(param.ByName("id"))
	if e != nil {
		panic(e.Error())
	}

	res, err := storage.Get(val)

	if err != nil {
		panic(err.Error())
	}

	jsonString, _ := json.Marshal(Response{UserId: res.Id})
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

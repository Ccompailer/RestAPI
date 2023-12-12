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
	router.GET("/api/v1/getEmployee/:id", getEmployeeHandler)
	router.DELETE("/api/v1/deleteEmployee/:id", deleteEmployeeHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getEmployeeHandler(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	val, e := strconv.Atoi(params.ByName("id"))
	if e != nil {
		panic(e.Error())
	}

	resp, err := storage.Get(val)

	if err != nil {
		panic(err.Error())
	}

	jsonString, _ := json.Marshal(Response{UserId: resp.Id})
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	val, er := strconv.Atoi(params.ByName("id"))

	if er != nil {
		panic(er.Error())
	}

	err := storage.Delete(val)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

package main

import (
	"awesomeProject/Data"
	"encoding/json"
	"fmt"
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
	originalHandler := http.HandlerFunc(handlerForMiddleware)

	router.ServeFiles("/api/v1/static/*filepath", http.Dir("C:\\Users\\1\\Desktop"))
	router.GET("/api/v1/employee/get/:id", getEmployeeHandler)
	router.DELETE("/api/v1/employee/delete/:id", deleteEmployeeHandler)
	router.POST("/api/v1/employee/create", createEmployeeHandler)
	router.PATCH("/api/v1/employee/update/:id", updateEmployeeHandler)
	router.GET("/api/v1/employee/get/:id/test", httpRouterMiddleware(getEmployeeHandler))
	http.Handle("/", customMiddleware(originalHandler))

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getEmployeeHandler(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	val, e := strconv.Atoi(params.ByName("id"))
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	resp, err := storage.Get(val)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	jsonString, _ := json.Marshal(Response{UserId: resp.Id})

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	val, er := strconv.Atoi(params.ByName("id"))

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := storage.Delete(val)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

func createEmployeeHandler(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	sex := r.Form.Get("sex")
	salary, _ := strconv.Atoi(r.Form.Get("salary"))

	storage.Insert(name, age, sex, salary)
}

func updateEmployeeHandler(w http.ResponseWriter, r *http.Request, params httprouter2.Params) {
	id, er := strconv.Atoi(params.ByName("id"))
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	sex := r.Form.Get("sex")
	salary, _ := strconv.Atoi(r.Form.Get("salary"))

	if er == nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	emp := storage.Update(id, Data.Employee{
		Age:    age,
		Name:   name,
		Sex:    sex,
		Salary: salary,
	})

	jsonString, _ := json.Marshal(emp)
	w.Write(jsonString)
}

func handlerForMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Execute mainHandler ...")
	w.Write([]byte("OK"))
}

func customMiddleware(originHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Work middleware before handler")
		originHandler.ServeHTTP(w, r)
		fmt.Print("Work middleware after handler")
	})
}

func httpRouterMiddleware(handle httprouter2.Handle) httprouter2.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter2.Params) {
		handle(writer, request, params)
	}
}

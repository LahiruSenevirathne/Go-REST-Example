package main

import (
	"RESTExample/Data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
)

func Index (w http.ResponseWriter ,r *http.Request){
	vars := mux.Vars(r)
	s := vars["number"]
	number,_ := strconv.ParseInt(s,10,64)
	if  number ==0 || number >10 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404-NOT FOUND"))
	}else {
		fmt.Fprintf(w,html.EscapeString(Data.Numerals[number]))
	}

}
func Index2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}


func TodoIndex (w http.ResponseWriter , r *http.Request){
	//fmt.Fprint(w,"Todo Index")
	todos := Todos{
		{Name:"Meetup",Completed:false},{Name:"Exam",Completed:false},
	}
	json.NewEncoder(w).Encode(todos)

}


func TodoShow (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w,todoID)
}
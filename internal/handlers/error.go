package handlers

import (
	"log"
	"net/http"
	"text/template"
)

type errors struct {
	ErrorCode int
	ErrorMsg  string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, errCode int, msg string) {
	t, err := template.ParseFiles("ui/error.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Errors := errors{
		ErrorCode: errCode,
		ErrorMsg:  msg,
	}
	t.Execute(w, Errors)
}

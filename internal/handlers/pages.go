package handlers

import (
	"groupie-tracker/internal"
	"net/http"
	"strconv"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("ui/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	data, err := internal.GetData(internal.ArtistUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("ui/artistpage.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if r.URL.Path != "/artist" {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	idString := r.URL.Query().Get("ID")
	if len(idString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if idString[0] == '0' {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	MaxId := internal.IdLimit()
	if id < 0 || id > MaxId {
		w.WriteHeader(http.StatusNotFound)
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	dataid, err := internal.ParseApi(idString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = ts.Execute(w, dataid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

package test

import (
	"groupie-tracker/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	// test method
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handlers.HomePage(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	// test method
	req = httptest.NewRequest("POST", "/", nil)
	w = httptest.NewRecorder()
	handlers.HomePage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	// test url
	req = httptest.NewRequest("GET", "/salam", nil)
	w = httptest.NewRecorder()
	handlers.HomePage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist?ID=1", nil)
	w = httptest.NewRecorder()
	handlers.HomePage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("error code, %v", resp.StatusCode)
	}
}

func TestArtistPage(t *testing.T) {
	// test method
	req := httptest.NewRequest("GET", "/artist?ID=1", nil)
	w := httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("POST", "/artist?ID=1", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	// test url
	req = httptest.NewRequest("GET", "/salam", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artists", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	// test id
	req = httptest.NewRequest("GET", "/artist?ID=0", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist?ID=52", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist?ID=53", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist?ID=asdsada", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist?ID=-1asda", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error code, %v", resp.StatusCode)
	}
	req = httptest.NewRequest("GET", "/artist?ID=1ooo0", nil)
	w = httptest.NewRecorder()
	handlers.ArtistPage(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("error code, %v", resp.StatusCode)
	}
}

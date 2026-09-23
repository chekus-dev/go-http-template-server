package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testHandler(t *testing.T) http.Handler {
	t.Helper()
	tmpl, err := template.ParseFS(templateFiles, "templates/home.html")
	if err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		homeHandler(tmpl, w, r)
	})
	return mux
}

func TestHomeHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	testHandler(t).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/home", nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if !strings.Contains(recorder.Body.String(), "Welcome") {
		t.Fatal("response does not contain the welcome message")
	}
}

func TestRootRedirectsToHome(t *testing.T) {
	recorder := httptest.NewRecorder()
	testHandler(t).ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))

	if recorder.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusFound)
	}
	if location := recorder.Header().Get("Location"); location != "/home" {
		t.Fatalf("location = %q, want %q", location, "/home")
	}
}

func TestHomeRejectsNonGet(t *testing.T) {
	recorder := httptest.NewRecorder()
	testHandler(t).ServeHTTP(recorder, httptest.NewRequest(http.MethodPost, "/home", nil))

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusMethodNotAllowed)
	}
}

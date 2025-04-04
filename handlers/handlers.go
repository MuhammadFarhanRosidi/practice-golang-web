package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to the Home Page!"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Product Page!"))
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Article Page!"))
}

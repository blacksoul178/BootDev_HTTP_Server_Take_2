package handlers

import "net/http"

const filepathRoot = "app"

func App(mux *http.ServeMux) {
	// Register all “app” related routes here
	fileServer := http.FileServer(http.Dir(filepathRoot))

	mux.Handle("/app/", fileServer)

	mux.HandleFunc("/app/test", func(w http.ResponseWriter, r *http.Request) { //custom handler to beautify urls
		http.ServeFile(w, r, "app/pages/test.html")
	})
}

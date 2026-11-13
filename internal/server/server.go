package server

import "net/http"

func NewRouter(webDir string) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(webDir))
	mux.Handle("/", fileServer)

	return mux
}

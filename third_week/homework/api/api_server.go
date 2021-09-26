package api

import "net/http"

func NewServerApiInit() (*http.ServeMux) {
	server := http.NewServeMux()
	server.HandleFunc("/help", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("help: ..."))
	})
	return server
}

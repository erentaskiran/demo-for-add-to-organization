package main

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/organization", HandleAddUser)

	return mux
}

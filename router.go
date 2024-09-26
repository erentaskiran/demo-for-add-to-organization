package main

import "net/http"

func (app *application) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/organization", app.HandleAddUser)

	return mux
}

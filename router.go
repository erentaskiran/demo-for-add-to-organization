package main

import "net/http"

func (app *application) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	h := newHandler(app.Db, &app.Cfg)
	mux.HandleFunc("/user", h.HandleUser)
	mux.HandleFunc("/organization", h.HandleOrganization)
	mux.HandleFunc("/organization_user", h.HandleOrganizationUser)
	return mux
}

package handler

import (
	"database/sql"
	"goweb/service"
	"net/http"
)

func NewHandler(db *sql.DB) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		rsp := Response{Message: "pong"}
		RespondJSON(w, &rsp, http.StatusOK)
	})

	cs := service.NewCdsDataService()
	gcdlc := NewGetCdsDataListController(db, cs)
	scdc := NewSaveCdsDataController(db, cs)
	r.HandleFunc("/cds/get", gcdlc.ServeHTTP)
	r.HandleFunc("/cds/post", scdc.ServeHTTP)
	return r
}

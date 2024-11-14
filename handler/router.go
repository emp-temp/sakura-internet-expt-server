package handler

import (
	"database/sql"
	"net/http"
	"sakura-internet-expt/repository"
	"sakura-internet-expt/service"
)

func NewHandler(db *sql.DB) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		rsp := Response{Message: "pong"}
		RespondJSON(w, &rsp, http.StatusOK)
	})

	// repository
	cr := repository.NewCdsDataRepository(db)

	// services
	cs := service.NewCdsDataService(cr)

	// controllers
	scdc := NewSaveCdsDataController(db, cs)
	ifuc := NewIsFrequentUrinationController(db, cs)

	r.HandleFunc("/cds", scdc.ServeHTTP)
	r.HandleFunc("/frequent_urination", ifuc.ServeHTTP)
	return r
}

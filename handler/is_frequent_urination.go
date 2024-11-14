package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"sakura-internet-expt/service"
)

type IsFrequentUrinationController struct {
	DB      *sql.DB
	Service service.ICdsDataService
}

func NewIsFrequentUrinationController(db *sql.DB, service service.ICdsDataService) *IsFrequentUrinationController {
	return &IsFrequentUrinationController{
		DB:      db,
		Service: service,
	}
}

func (ifuc *IsFrequentUrinationController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ifu, err := ifuc.Service.IsFrequentUrination()
	if err != nil {
		rsp := &Response{
			Message: err.Error(),
		}
		RespondJSON(w, &rsp, http.StatusInternalServerError)
		return
	}
	rsp := &Response{
		Message: fmt.Sprintf("%t", ifu),
	}
	RespondJSON(w, &rsp, http.StatusOK)
}

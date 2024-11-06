package handler

import (
	"database/sql"
	"goweb/service"
	"net/http"
	"strconv"
)

type GetCdsDataListController struct {
	DB      *sql.DB
	Service service.ICdsDataService
	// TODO: validator
}

func NewGetCdsDataListController(db *sql.DB, service service.ICdsDataService) *GetCdsDataListController {
	return &GetCdsDataListController{
		DB:      db,
		Service: service,
	}
}

func (gcdlc *GetCdsDataListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	var limit, offset int
	if limitStr == "" {
		limit = 10
	} else {
		limit, _ = strconv.Atoi(limitStr)
	}
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(offsetStr)
	}
	cdl, err := gcdlc.Service.GetCdsDataList(gcdlc.DB, limit, offset)
	if err != nil {
		rsp := &Response{
			Message: err.Error(),
		}
		RespondJSON(w, rsp, http.StatusInternalServerError)
		return
	}
	RespondJSON(w, &cdl, http.StatusOK)
}

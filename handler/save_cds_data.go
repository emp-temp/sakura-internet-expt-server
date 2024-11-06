package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sakura-internet-expt/entity"
	"sakura-internet-expt/service"
	"time"
)

type SaveCdsDataController struct {
	DB      *sql.DB
	Service service.ICdsDataService
}

func NewSaveCdsDataController(db *sql.DB, service service.ICdsDataService) *SaveCdsDataController {
	return &SaveCdsDataController{
		DB:      db,
		Service: service,
	}
}

func (scdc *SaveCdsDataController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		rsp := &Response{
			Message: err.Error(),
		}
		RespondJSON(w, &rsp, http.StatusBadRequest)
		return
	}
	cd := &entity.CdsData{
		StartTime: reqBody.StartTime,
		EndTime:   reqBody.EndTime,
	}
	if err := scdc.Service.SaveCdsData(scdc.DB, cd); err != nil {
		rsp := &Response{
			Message: err.Error(),
		}
		RespondJSON(w, &rsp, http.StatusInternalServerError)
		return
	}
	rsp := &Response{
		Message: "success",
	}
	RespondJSON(w, &rsp, http.StatusOK)
}

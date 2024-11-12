package service

import (
	"sakura-internet-expt/entity"
	"sakura-internet-expt/repository"
	"time"
)

var (
	table = "cds_data"
)

type ICdsDataService interface {
	IsFrequentUrination() (bool, error)
}

type CdsDataService struct {
	CdsDataRepository repository.ICdsDataRepository
}

func NewCdsDataService(repo repository.ICdsDataRepository) ICdsDataService {
	return &CdsDataService{
		CdsDataRepository: repo,
	}
}

func (cds *CdsDataService) IsFrequentUrination() (bool, error) {
	now := time.Now().AddDate(0, 0, 1)
	// 一日分
	rd, err := cds.CdsDataRepository.GetDailyCdsDataList(now)
	if err != nil {
		return false, err
	}
	// 一週間分
	// rw, err := cds.CdsDataRepository.GetWeeklyCdsDataList(now)
	// if err != nil {
	// 	return false, err
	// }
	// return isDaytimeFrequentUrination(rd) || isNighttimeFrequentUrination(rw), nil
	return isDaytimeFrequentUrination(rd), nil
}

func isDaytimeFrequentUrination(cdsData []entity.CdsData) bool {
	daytime := make([]entity.CdsData, 0)
	for _, v := range cdsData {
		if v.StartTime.Hour() > 8 || v.StartTime.Hour() < 20 {
			daytime = append(daytime, v)
		}
	}
	return len(daytime) >= 10
}

// func isNighttimeFrequentUrination(weeklyCdsData []entity.CdsData) bool {
// 	var wd map[string][]entity.CdsData
// 	for _, c := range weeklyCdsData {
// 		wd[c.StartTime.Weekday().String()] = append(wd[c.StartTime.Weekday().String()], c)
// 	}
// 	for _, cdl := range wd {
// 		count := 0
// 		for _, cd := range cdl {
// 			if cd.StartTime.Hour()
// 		}
// 	}
// }

package service

import (
	"sakura-internet-expt/entity"
	"sakura-internet-expt/repository"
	"time"
)

type ICdsDataService interface {
	SaveCdsData(cdsData entity.CdsData) error
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

func (cds *CdsDataService) SaveCdsData(cdsData entity.CdsData) error {
	return cds.CdsDataRepository.SaveCdsData(cdsData)
}

func (cds *CdsDataService) IsFrequentUrination() (bool, error) {
	yd := time.Now().AddDate(0, 0, -1)
	rd, err := cds.CdsDataRepository.GetDailyCdsDataList(yd)
	if err != nil {
		return false, err
	}
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

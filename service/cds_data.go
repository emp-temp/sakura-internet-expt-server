package service

import (
	"database/sql"
	"errors"
	"fmt"
	"sakura-internet-expt/entity"
)

var (
	table = "cds_data"
)

type ICdsDataService interface {
	GetCdsDataList(db *sql.DB, limit, offset int) ([]entity.CdsData, error)
	SaveCdsData(db *sql.DB, cdsData *entity.CdsData) error
}

type CdsDataService struct{}

func NewCdsDataService() ICdsDataService {
	return &CdsDataService{}
}

func (cds *CdsDataService) GetCdsDataList(db *sql.DB, limit, offset int) ([]entity.CdsData, error) {
	result := make([]entity.CdsData, 0)

	// `?` プレースホルダーを使い、`limit` と `offset` をパラメータとして渡す
	rows, err := db.Query("SELECT id, start_time, end_time FROM cds_data LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.CdsData{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c entity.CdsData
		if err := rows.Scan(&c.ID, &c.StartTime, &c.EndTime); err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func (cds *CdsDataService) SaveCdsData(db *sql.DB, c *entity.CdsData) error {
	query := fmt.Sprintf("INSERT INTO %s (start_time, end_time) VALUES (?, ?)", table)
	_, err := db.Exec(query, c.StartTime, c.EndTime)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sakura-internet-expt/entity"
	"time"
)

type ICdsDataRepository interface {
	GetCdsDataList(limit, offset int) ([]entity.CdsData, error)
	GetDailyCdsDataList(date time.Time) ([]entity.CdsData, error)
	GetWeeklyCdsDataList(date time.Time) ([]entity.CdsData, error)
	SaveCdsData(cdsData entity.CdsData) error
}

var (
	table = "cds_data"
)

type CdsDataRepository struct {
	DB *sql.DB
}

func NewCdsDataRepository(db *sql.DB) ICdsDataRepository {
	return &CdsDataRepository{
		DB: db,
	}
}

func (cdr *CdsDataRepository) GetCdsDataList(limit, offset int) ([]entity.CdsData, error) {
	result := make([]entity.CdsData, 0)

	rows, err := cdr.DB.Query("SELECT id, start_time, end_time FROM cds_data LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.CdsData{}, nil
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		var c entity.CdsData
		if err := rows.Scan(&c.ID, &c.StartTime, &c.EndTime); err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func (cdr *CdsDataRepository) GetDailyCdsDataList(date time.Time) ([]entity.CdsData, error) {
	result := make([]entity.CdsData, 0)

	rows, err := cdr.DB.Query(fmt.Sprintf("SELECT * FROM %s WHERE start_time = CAST(? AS DATE)", table), date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.CdsData{}, nil
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		var c entity.CdsData
		if err := rows.Scan(&c.ID, &c.StartTime, &c.EndTime); err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func (cdr *CdsDataRepository) GetWeeklyCdsDataList(date time.Time) ([]entity.CdsData, error) {
	result := make([]entity.CdsData, 0)
	oneWeekAgo := date.AddDate(0, 0, -7)

	rows, err := cdr.DB.Query(fmt.Sprintf("SELECT * FROM %v WHERE start_time >= ? ORDER BY start_time DESC", table), oneWeekAgo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.CdsData{}, nil
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}(rows)

	for rows.Next() {
		var c entity.CdsData
		if err := rows.Scan(&c.ID, &c.StartTime, &c.EndTime); err != nil {
			return nil, err
		}
		result = append(result, c)
	}

	return result, nil
}

func (cdr *CdsDataRepository) SaveCdsData(cdsData entity.CdsData) error {
	query := fmt.Sprintf("INSERT INTO %s (start_time, end_time) VALUES (?, ?)", table)
	_, err := cdr.DB.Exec(query, cdsData.StartTime, cdsData.EndTime)
	if err != nil {
		return err
	}
	return nil
}

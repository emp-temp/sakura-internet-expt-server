package entity

import "time"

type CdsData struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
}

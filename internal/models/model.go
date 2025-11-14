package models

import "time"

type Subscription struct {
	ID           int        `json:"id"`
	UserID       string     `json:"user_id"`
	ServiceName  string     `json:"service_name"`
	MonthlyPrice float64    `json:"monthly_price"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
}

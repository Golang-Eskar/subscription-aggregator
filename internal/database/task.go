package database

import (
	"time"

	"github.com/Golang-Eskar/subscription-aggregator/internal/models"
	"github.com/google/uuid"
)

func Baysub(w models.Subscription) (int, error) {
	if w.UserID == "" {
		w.UserID = uuid.New().String()
	}

	var id int
	err := DB.QueryRow(
		`INSERT INTO subscriptions (service_name, monthly_price, user_id, start_date)
         VALUES ($1, $2, $3, $4)
         RETURNING id`,
		w.ServiceName,
		w.MonthlyPrice,
		w.UserID,
		w.StartDate,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
func NewTask(service_name string, price float64, user_id string, start_date time.Time) *models.Subscription {
	if user_id == "" {
		user_id = uuid.New().String()
	}
	return &models.Subscription{
		ServiceName:  service_name,
		MonthlyPrice: price,
		UserID:       user_id,
		StartDate:    start_date,
	}
}

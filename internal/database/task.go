package database

import "github.com/Golang-Eskar/subscription-aggregator/internal/models"

func Baysub(w models.Task) (int, error) {
	result, err := db.Exec("INSERT INTO subscriptions (service_name, monthly_price, user_id, start_date) VALUES ($1, $2, $3, $4)",
		w.Service_name, w.Monthly_price, w.User_id, w.Start_date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func NewTask(service_name string, price int, user_id string, start_date string) *models.Task {
	return &models.Task{
		Service_name:  service_name,
		Monthly_price: price,
		User_id:       user_id,
		Start_date:    start_date,
	}
}

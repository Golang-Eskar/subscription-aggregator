package database

import (
	"fmt"

	"github.com/Golang-Eskar/subscription-aggregator/internal/models"
)

func Create(sub models.Subscription) (int, error) {
	var id int
	err := DB.QueryRow(`
		INSERT INTO subscriptions (user_id, service_name, monthly_price, start_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		sub.UserID, sub.ServiceName, sub.MonthlyPrice, sub.StartDate,
	).Scan(&id)
	return id, err
}

func GetAll() ([]models.Subscription, error) {
	rows, err := DB.Query(`SELECT id, user_id, service_name, monthly_price, start_date, end_date FROM subscriptions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []models.Subscription

	for rows.Next() {
		var s models.Subscription
		rows.Scan(&s.ID, &s.UserID, &s.ServiceName, &s.MonthlyPrice, &s.StartDate, &s.EndDate)
		subs = append(subs, s)
	}
	return subs, nil
}

func GetByID(id int) (*models.Subscription, error) {
	var s models.Subscription
	err := DB.QueryRow(`
		SELECT id, user_id, service_name, monthly_price, start_date, end_date
		FROM subscriptions WHERE id = $1`, id).
		Scan(&s.ID, &s.UserID, &s.ServiceName, &s.MonthlyPrice, &s.StartDate, &s.EndDate)

	return &s, err
}

func Update(id int, s models.Subscription) error {
	_, err := DB.Exec(`
		UPDATE subscriptions SET service_name=$1, monthly_price=$2, end_date=$3
		WHERE id=$4`,
		s.ServiceName, s.MonthlyPrice, s.EndDate, id)
	return err
}

func Delete(id int) error {
	_, err := DB.Exec(`DELETE FROM subscriptions WHERE id=$1`, id)
	return err
}

func Filter(userID, service string) ([]models.Subscription, error) {
	query := `
	SELECT id, user_id, service_name, monthly_price, start_date, end_date
	FROM subscriptions WHERE 1=1`

	args := []any{}
	i := 1

	if userID != "" {
		query += " AND user_id = $" + fmt.Sprint(i)
		args = append(args, userID)
		i++
	}

	if service != "" {
		query += " AND service_name = $" + fmt.Sprint(i)
		args = append(args, service)
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Subscription

	for rows.Next() {
		var s models.Subscription
		rows.Scan(&s.ID, &s.UserID, &s.ServiceName, &s.MonthlyPrice, &s.StartDate, &s.EndDate)
		result = append(result, s)
	}

	return result, nil
}

func Sum(from, to string) (float64, error) {
	var sum float64
	err := DB.QueryRow(`
		SELECT COALESCE(SUM(monthly_price), 0)
		FROM subscriptions
		WHERE start_date >= $1 AND start_date <= $2`,
		from, to,
	).Scan(&sum)

	return sum, err
}

package database

type Task struct {
	service_name  string
	monthly_price int
	user_id       string
	start_date    string
}

func Baysub(w Task) (int, error) {
	result, err := db.Exec("INSERT INTO subscriptions (service_name, monthly_price, user_id, start_date) VALUES ($1, $2, $3, $4)",
		w.service_name, w.monthly_price, w.user_id, w.start_date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func NewTask(service_name string, price int, user_id string, start_date string) *Task {
	return &Task{
		service_name:  service_name,
		monthly_price: price,
		user_id:       user_id,
		start_date:    start_date,
	}
}

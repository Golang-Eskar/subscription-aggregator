package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const creatdb = `CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    service_name VARCHAR(100) NOT NULL,
    monthly_price DECIMAL(10,2) NOT NULL CHECK (monthly_price > 0),
    start_date DATE NOT NULL,
    end_date DATE NULL,
    created_at TIMESTAMP DEFAULT NOW()
);`

var db *sql.DB

func Init() error {
	log.Println("Ждем запуска PostgreSQL...")
	time.Sleep(5 * time.Second)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "subscriptions_db"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к БД: %w", err)
	}

	if err := dbConn.Ping(); err != nil {
		return fmt.Errorf("ошибка соединения с БД: %w", err)
	}

	_, err = dbConn.Exec(creatdb)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %w", err)
	}

	db = dbConn
	return nil
}

/*func DB() *sql.DB {
	return db
}*/

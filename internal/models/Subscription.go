package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Subscription представляет запись о подписке пользователя.
// Соответствует таблице `subscriptions` в PostgreSQL.
type Subscription struct {
	ID           uuid.UUID       `json:"id" db:"id"`                       // Уникальный идентификатор (UUID)
	UserID       uuid.UUID       `json:"user_id" db:"user_id"`             // UUID пользователя
	ServiceName  string          `json:"service_name" db:"service_name"`   // Название сервиса
	MonthlyPrice decimal.Decimal `json:"monthly_price" db:"monthly_price"` // Стоимость подписки (decimal)
	StartDate    time.Time       `json:"start_date" db:"start_date"`       // Дата начала подписки
	EndDate      *time.Time      `json:"end_date,omitempty" db:"end_date"` // Дата окончания (опционально)
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`       // Дата создания записи
}

package scheme

import "time"

type OrderID int64

type Order struct {
	ID            OrderID   `db:"id"`
	Price         int64     `db:"price"`
	UserID        int64     `db:"user_id"`
	Status        string    `db:"status"`
	Items         []Item    `db:"items"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

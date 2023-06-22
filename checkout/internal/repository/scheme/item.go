package scheme

import "time"

type Item struct {
	UserID        int64     `db:"user_id"`
	SKU           int64     `db:"sku"`
	Count         int32     `db:"count"`
	Price         int64     `db:"price"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

package scheme

import "time"

type Stock struct {
	WarehouseID   int64     `db:"warehouse_id"`
	Count         int32     `db:"count"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

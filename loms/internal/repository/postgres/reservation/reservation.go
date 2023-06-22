package reservation

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"route256/libs/postgresdb"
	"route256/libs/tx"
	"route256/loms/internal/model"
	"route256/loms/internal/repository/postgres/types"
	"route256/loms/internal/repository/scheme"
)

const (
	tableReservation  = "reservation"
	columnWarehouseID = "warehouse_id"
	columnOrderID     = "order_id"
	columnSKU         = "sku"
	columnCount       = "count"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Reservation struct {
	provider tx.DBProvider
}

func NewReservationRepository(db *postgresdb.Postgres) *Reservation {
	return &Reservation{provider: tx.New(db)}
}

func (repo *Reservation) GetReservation(ctx context.Context, orderID int64) (model.Reservation, error) {
	db := repo.provider.GetDB(ctx)

	orderIDScheme := scheme.OrderIDToScheme(orderID)
	query, args, err := psql.Select(columnSKU, columnWarehouseID, columnCount).
		From(tableReservation).
		Where(sq.Eq{columnOrderID: orderIDScheme}).
		ToSql()
	if err != nil {
		return model.Reservation{}, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return model.Reservation{}, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}

	var reservation = scheme.Reservation{
		OrderID: orderID,
	}
	for rows.Next() {
		var reservationStock scheme.StockReservation
		err := rows.Scan(&reservationStock.SKU, &reservationStock.WarehouseID, &reservationStock.Count)
		if err != nil {
			return model.Reservation{}, err
		}
		reservation.Stocks = append(reservation.Stocks, reservationStock)
	}
	return scheme.ReservationFromScheme(reservation), nil
}

func (repo *Reservation) CreateReservation(ctx context.Context, reservation model.Reservation) error {
	db := repo.provider.GetDB(ctx)

	builder := psql.Insert(tableReservation).
		Columns(columnOrderID, columnSKU, columnWarehouseID, columnCount)

	orderIDScheme := scheme.OrderIDToScheme(reservation.OrderID)
	for _, stock := range reservation.Stocks {
		stockScheme := scheme.StockReservationToScheme(stock)
		builder = builder.Values(
			orderIDScheme,
			stockScheme.SKU,
			stockScheme.WarehouseID,
			stockScheme.Count,
		)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	if _, err := db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

func (repo *Reservation) DeleteReservation(ctx context.Context, orderID int64) error {
	db := repo.provider.GetDB(ctx)

	schemeOrderID := scheme.OrderIDToScheme(orderID)
	query, args, err := psql.Delete(tableReservation).
		Where(sq.Eq{columnOrderID: schemeOrderID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	if _, err := db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

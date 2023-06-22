package order

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"route256/libs/postgresdb"
	"route256/libs/tx"
	"route256/loms/internal/model"
	"route256/loms/internal/repository/postgres/types"
	"route256/loms/internal/repository/scheme"
	"time"
)

const (
	tableOrders        = "orders"
	columnID           = "id"
	columnUserID       = "user_id"
	columnStatus       = "status"
	columnCreatedAt    = "created_at"
	columnLastUpdateAt = "last_updated_at"

	tableOrderItems = "order_items"
	columnOrderID   = "order_id"
	columnSKU       = "sku"
	columnCount     = "count"
	columnPrice     = "price"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Order struct {
	provider tx.DBProvider
}

func NewOrderRepository(db *postgresdb.Postgres) *Order {
	return &Order{provider: tx.New(db)}
}

func (repo *Order) CreateOrder(ctx context.Context, order model.Order) (int64, error) {
	db := repo.provider.GetDB(ctx)

	orderScheme := scheme.OrderToScheme(order)
	query, args, err := psql.Insert(tableOrders).
		Columns(columnUserID, columnStatus, columnLastUpdateAt, columnCreatedAt).
		Values(orderScheme.UserID, orderScheme.Status, sq.Expr("CURRENT_TIMESTAMP"), sq.Expr("CURRENT_TIMESTAMP")).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	var orderID scheme.OrderID
	row := db.QueryRow(ctx, query, args...)
	if err := row.Scan(&orderID); err != nil {
		return 0, fmt.Errorf("%w: %w", types.ErrFailedToScanRow, err)
	}
	return scheme.OrderIDFromScheme(orderID), nil
}

func (repo *Order) GetOrder(ctx context.Context, orderID int64) (model.Order, error) {
	db := repo.provider.GetDB(ctx)

	orderIDScheme := scheme.OrderIDToScheme(orderID)
	query, args, err := psql.
		Select(columnOrderID, columnSKU, columnCount, columnPrice, columnUserID, columnStatus).
		From(tableOrders).
		InnerJoin(fmt.Sprintf("%s ON %s=%s", tableOrderItems, tableOrders+"."+columnID, tableOrderItems+"."+columnOrderID)).
		Columns().
		Where(sq.Eq{columnID: orderIDScheme}).
		ToSql()
	if err != nil {
		return model.Order{}, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return model.Order{}, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}

	var result model.Order
	for rows.Next() {
		var itemSch scheme.Item
		var orderSch scheme.Order

		err := rows.Scan(&itemSch.OrderID, &itemSch.SKU, &itemSch.Count, &itemSch.Price, &orderSch.UserID, &orderSch.Status)
		if err != nil {
			return model.Order{}, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
		}
		if result.User == 0 {
			order := scheme.OrderFromScheme(orderSch)
			result.User = order.User
			result.Status = order.Status
		}
		result.Items = append(result.Items, scheme.ItemFromScheme(itemSch))
	}
	return result, nil
}

func (repo *Order) GetExpiredOrderIDs(ctx context.Context, comparisonTime time.Time, status string) ([]int64, error) {
	db := repo.provider.GetDB(ctx)

	query, args, err := psql.
		Select(columnID).
		From(tableOrders).
		Where(sq.Lt{columnLastUpdateAt: comparisonTime}).
		Where(sq.Eq{columnStatus: status}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}

	var orderIDs []scheme.OrderID
	for rows.Next() {
		var orderID scheme.OrderID

		err := rows.Scan(&orderID)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
		}
		orderIDs = append(orderIDs, orderID)
	}
	return scheme.OrderIDsFromScheme(orderIDs), nil
}

func (repo *Order) SetOrderStatus(ctx context.Context, orderID int64, status string) error {
	db := repo.provider.GetDB(ctx)

	orderIDScheme := scheme.OrderIDToScheme(orderID)
	query, args, err := psql.Update(tableOrders).
		Set(columnStatus, status).
		Where(sq.NotEq{columnStatus: status}).
		Where(sq.Eq{columnID: orderIDScheme}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(columnID, query, args)

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

func (repo *Order) CreateOrderItems(ctx context.Context, order model.Order) error {
	db := repo.provider.GetDB(ctx)

	orderScheme := scheme.OrderToScheme(order)
	builder := psql.Insert(tableOrderItems).
		Columns(columnOrderID, columnSKU, columnCount, columnPrice)

	for _, item := range order.Items {
		builder = builder.Values(orderScheme.ID, item.SKU, item.Count, item.Price)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

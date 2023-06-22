package purchase

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/postgres/types"
	"route256/checkout/internal/repository/scheme"
	"route256/libs/postgresdb"
	"route256/libs/tx"
)

const (
	tablePurchase   = "purchase"
	columnOrderID   = "user_id"
	columnUserID    = "sku"
	columnStatus    = "status"
	columnCreatedAt = "created_at"

	tablePurchaseItems = "purchase_items"
	columnPurchaseID   = "purchase_id"
	columnSKU          = "sku"
	columnCount        = "count"
	columnPrice        = "price"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Purchase struct {
	provider tx.DBProvider
}

func NewPurchaseRepository(db *postgresdb.Postgres) *Purchase {
	return &Purchase{provider: tx.New(db)}
}

func (repo *Purchase) Create(ctx context.Context, userID, orderID int64, status int16) (purchaseID int64, err error) {
	db := repo.provider.GetDB(ctx)
	query, args, err := psql.Insert(tablePurchase).
		Columns(columnUserID, columnOrderID, columnStatus, columnCreatedAt).
		Values(userID, orderID, status, sq.Expr("CURRENT_TIMESTAMP")).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	row := db.QueryRow(ctx, query, args...)
	if row == nil {
		return 0, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	if err := row.Scan(&purchaseID); err != nil {
		return 0, fmt.Errorf("%w: %w", types.ErrFailedToScanRow, err)
	}

	return purchaseID, nil
}

func (repo *Purchase) UpdateStatus(ctx context.Context, userID, orderID int64, status int16) error {
	db := repo.provider.GetDB(ctx)
	query, args, err := psql.Update(tablePurchase).
		Set(columnStatus, status).
		Where(sq.Eq{
			columnUserID:  userID,
			columnOrderID: orderID,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}

	return nil
}

func (repo *Purchase) AddToPurchaseItems(ctx context.Context, purchaseID int64, items []model.Item) error {
	db := repo.provider.GetDB(ctx)
	builder := psql.Insert(tablePurchaseItems).
		Columns(columnPurchaseID, columnSKU, columnCount, columnPrice)

	schemeItem := scheme.ItemsToScheme(items)
	for _, item := range schemeItem {
		builder = builder.Values(purchaseID, item.SKU, item.Count, item.Price)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query)

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w : %w", types.ErrFailedExecuteQuery)
	}
	return nil
}

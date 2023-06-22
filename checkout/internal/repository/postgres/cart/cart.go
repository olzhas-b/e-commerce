package cart

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/postgres/types"
	"route256/checkout/internal/repository/scheme"
	"route256/libs/postgresdb"
	"route256/libs/tx"
	"time"
)

type Status int16

const (
	statusUndefined Status = iota - 1
)

const (
	tableCart           = "cart"
	columnUserID        = "user_id"
	columnSKU           = "sku"
	columnCount         = "count"
	columnStatus        = "Status"
	columnLastUpdatedAt = "last_updated_at"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Cart struct {
	provider tx.DBProvider
}

func NewCartRepository(db *postgresdb.Postgres) *Cart {
	return &Cart{provider: tx.New(db)}
}

func (repo *Cart) GetItemCount(ctx context.Context, userID int64, sku uint32) (uint16, error) {
	db := repo.provider.GetDB(ctx)

	query, args, err := psql.Select(columnCount).
		From(tableCart).
		Where(sq.Eq{
			columnUserID: userID,
			columnSKU:    scheme.SkuToScheme(sku),
		}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	var item scheme.Item
	row := db.QueryRow(ctx, query, args...)
	if err = row.Scan(&item.Count); err != nil {
		if err == pgx.ErrNoRows {
			return 0, types.ErrItemNotFound
		}
		return 0, fmt.Errorf("%w: %w", types.ErrFailedToScanRow, err)
	}

	return scheme.ItemFromScheme(item).Count, nil
}

func (repo *Cart) GetCartItems(ctx context.Context, userID int64) ([]model.Item, error) {
	db := repo.provider.GetDB(ctx)

	query, args, err := psql.Select(columnUserID, columnSKU, columnCount, columnLastUpdatedAt).
		From(tableCart).
		Where(sq.Eq{
			columnUserID: userID,
		}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	fmt.Println(query)
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	var results []scheme.Item
	for rows.Next() {
		var result scheme.Item
		err = rows.Scan(&result.UserID, &result.SKU, &result.Count, &result.LastUpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", types.ErrFailedToScanRow, err)
		}
		results = append(results, result)
	}
	return scheme.ItemsFromScheme(results), nil
}

func (repo *Cart) AddToCart(ctx context.Context, userID int64, item model.Item) error {
	db := repo.provider.GetDB(ctx)

	schemeItem := scheme.ItemToScheme(item)
	fmt.Println(schemeItem, userID, schemeItem.SKU, schemeItem.Count, time.Now())
	query := `
INSERT INTO cart("user_id", "sku", "count", "last_updated_at") VALUES 
    ($1, $2, $3, $4)
ON CONFLICT ("user_id", "sku") DO UPDATE 
	SET count=cart.count+excluded.count`

	_, err := db.Exec(ctx, query, userID, schemeItem.SKU, schemeItem.Count, time.Now())
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

func (repo *Cart) UpdateItemFromCart(ctx context.Context, userID int64, item model.Item) error {
	db := repo.provider.GetDB(ctx)

	schemeItem := scheme.ItemToScheme(item)
	query, args, err := psql.Update(tableCart).
		Set(columnCount, schemeItem.Count).
		Set(columnLastUpdatedAt, time.Now()).
		Where(sq.Eq{columnUserID: userID}).
		Where(sq.Eq{columnSKU: schemeItem.SKU}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

func (repo *Cart) DeleteFromCart(ctx context.Context, userID int64, sku uint32) error {
	db := repo.provider.GetDB(ctx)
	query, args, err := psql.Delete(tableCart).
		Where(
			sq.Eq{
				columnUserID: userID,
				columnSKU:    scheme.SkuToScheme(sku),
			},
		).ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println(query, args)

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

func (repo *Cart) DeleteAllFromCart(ctx context.Context, userID int64) error {
	db := repo.provider.GetDB(ctx)

	query, args, err := psql.Delete(tableCart).
		Where(sq.Eq{columnUserID: userID}).
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

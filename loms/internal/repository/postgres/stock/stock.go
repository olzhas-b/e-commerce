package stock

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
	tableStocks         = "stocks"
	columnWarehouseID   = "warehouse_id"
	columnSKU           = "sku"
	columnCount         = "count"
	columnLastUpdatedAt = "last_updated_at"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Stock struct {
	provider tx.DBProvider
}

func NewStocksRepository(db *postgresdb.Postgres) *Stock {
	return &Stock{provider: tx.New(db)}
}

func (repo *Stock) GetStocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	db := repo.provider.GetDB(ctx)

	skuScheme := scheme.SkuToScheme(sku)
	query, args, err := psql.Select(columnWarehouseID, columnCount).
		From(tableStocks).
		Where(sq.Eq{columnSKU: skuScheme}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}

	var result []model.Stock
	for rows.Next() {
		var r scheme.Stock
		if err := rows.Scan(&r.WarehouseID, &r.Count); err != nil {
			return nil, fmt.Errorf("%w: %w", types.ErrFailedToScanRow, err)
		}
		result = append(result, scheme.StockFromScheme(r))
	}
	return result, nil
}

func (repo *Stock) ModifyStocksCount(ctx context.Context, stock model.Stock, sku uint32, operation string) error {
	db := repo.provider.GetDB(ctx)

	skuScheme := scheme.SkuToScheme(sku)
	stockScheme := scheme.StockToScheme(stock)
	query, args, err := psql.Update(tableStocks).
		Set(columnCount, sq.Expr(fmt.Sprintf("%s %s %d", columnCount, operation, stockScheme.Count))).
		Set(columnLastUpdatedAt, stockScheme.LastUpdatedAt).
		Where(sq.Eq{
			columnSKU:         skuScheme,
			columnWarehouseID: stockScheme.WarehouseID,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedBuildQuery, err)
	}
	fmt.Println("ModifyStocksCount", query, args)

	_, err = db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrFailedExecuteQuery, err)
	}
	return nil
}

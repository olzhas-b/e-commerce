-- +goose Up
-- +goose StatementBegin
insert into stocks(sku, warehouse_id, count, last_updated_at)
values (1076963, 1000, 1000000, '2019-01-01 00:00:00'),
       (1625903, 1000, 1000000, '2019-01-01 00:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from stocks
where sku in (1076963, 1625903);
-- +goose StatementEnd

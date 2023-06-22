-- +goose Up
-- +goose StatementBegin
create table if not exists orders (
    id bigserial primary key,
    user_id bigint not null,
    status varchar default 'new',
    created_at timestamp not null,
    last_updated_at timestamp not null
);
create index orders_user_id_hx on orders using HASH (user_id);

create table if not exists order_items (
    order_id bigint not null,
    sku bigint not null,
    count int not null,
    price bigint not null,
    constraint count check (count >= 0)
);
create index order_items_order_id_hx on order_items using HASH(order_id);

create table stocks (
    sku bigint not null,
    warehouse_id bigint not null,
    count int,
    last_updated_at timestamp not null,
    constraint count check (count >= 0)
);
create index stocks_sku_hx on stocks using HASH(sku);


create table reservation (
    order_id bigint not null,
    warehouse_id bigint,
    sku int not null,
    count int not null
);
create index reservation_order_id_hx on reservation USING hash(order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists orders;
drop table if exists order_items;
drop table if exists reservation;
drop table if exists stocks;
-- +goose StatementEnd

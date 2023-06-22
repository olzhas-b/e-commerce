-- +goose Up
-- +goose StatementBegin
create table if not exists cart (
    user_id bigint not null,
    sku bigint not null,
    count int not null,
    status int default 0,
    last_updated_at timestamp not null,
    constraint count check (count >= 0),
    constraint uniqueness unique(user_id, sku)
);
create index cart_user_id_hx on cart using HASH (user_id);

create table if not exists purchase (
    id bigserial primary key,
    order_id bigint not null,
    user_id bigint not null,
    status int default 0,
    created_at timestamp not null
);

create index purchase_user_id_hx on purchase using HASH (user_id);

create table purchase_items (
    purchase_id bigint not null,
    sku bigint not null,
    count int not null,
    price bigint not null,
    constraint count check (count >= 0)
);
create index  purchase_items_purchase_id_hx on purchase_items using HASH (purchase_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cart;
drop table if exists purchase;
drop table if exists purchase_items;
-- +goose StatementEnd

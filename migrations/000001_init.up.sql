create
    extension if not exists "uuid-ossp";

create table orders
(
    id                uuid default uuid_generate_v4() not null unique primary key,
    symbol            varchar(15)                     not null,
    exchanger_id      varchar(31)                     not null,
    exchange_order_id varchar(255)                    null,
    created_at        timestamp                       not null
);
create unique index idx_order_symbol on orders(symbol, exchanger_id);

create table signals
(
    symbol     varchar(15)                     not null unique primary key,
    value      varchar(255)                    not null,
    created_at timestamp                       not null
);
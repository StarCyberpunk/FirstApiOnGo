-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA bank
CREATE TABLE bank.bank_account(
    id_ba uuid PRIMARY KEY,
    pass_serial integer not null ,
    pass_number integer not null ,
    -- decimal
    cash_total double precision
);
CREATE TABLE bank.currency(
    id serial PRIMARY KEY,
    name text,
    one_to_rub double precision
);
CREATE TABLE bank.users(
    id_user uuid PRIMARY KEY,
    login text not null ,
    password bytea not null ,
    id_role integer,
    email text,
    id_ba uuid,
    FOREIGN KEY (id_ba) REFERENCES bank.bank_account(id_ba)
);
CREATE TABLE bank.cards(
    id_card uuid PRIMARY KEY,
    id_currency integer,
    type_card_id integer,
    cash double precision,
    number_card bigint,
    valid_date date,
    --hash
    cvv smallint,
    block boolean,
    id_ba uuid,
    FOREIGN KEY (id_currency) REFERENCES bank.currency(id),
    FOREIGN KEY (id_ba) REFERENCES bank.bank_account(id_ba)
);
CREATE TABLE bank.operations_card(
    id_op uuid PRIMARY KEY,
    date_op date,
    id_card_from uuid,
    id_card_to uuid,
    total double precision,
    id_cur integer,
    description text,
    FOREIGN KEY (id_card_from) REFERENCES bank.cards(id_card),
    FOREIGN KEY (id_card_to) REFERENCES bank.cards(id_card),
    FOREIGN KEY (id_cur) REFERENCES bank.currency(id)
);
CREATE TABLE bank.operations_bank_account(
    id_op uuid PRIMARY KEY,
    date_op date,
    id_ba_from uuid,
    id_ba_to uuid,
    total double precision,
    id_cur integer,
    description text,
    FOREIGN KEY (id_ba_from) REFERENCES bank.bank_account(id_ba),
    FOREIGN KEY (id_ba_to) REFERENCES bank.bank_account(id_ba),
    FOREIGN KEY (id_cur) REFERENCES bank.currency(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd

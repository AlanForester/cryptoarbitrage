BEGIN;

CREATE TABLE assets (
	id uuid NOT NULL PRIMARY KEY,
	symbol text NOT NULL,
	name text NOT NULL,
	is_fiat boolean NOT NULL
);


CREATE TABLE pairs (
	id uuid NOT NULL PRIMARY KEY,
	symbol text NOT NULL,
	base_id uuid REFERENCES assets(id),
	quote_id uuid REFERENCES assets(id)
);


CREATE TABLE exchanges (
	id uuid NOT NULL PRIMARY KEY,
	symbol text NOT NULL,
	name text NOT NULL,
	is_active boolean NOT NULL
);


CREATE TABLE differences (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	pair_id uuid REFERENCES pairs(id),
	base_id uuid REFERENCES exchanges(id),
	quote_id uuid REFERENCES exchanges(id),
	delta real NOT NULL
);


CREATE TABLE exchange_assets (
	id uuid NOT NULL PRIMARY KEY,
	asset_id uuid REFERENCES assets(id),
	exchange_id uuid REFERENCES exchanges(id),
	transaction_fee real NOT NULL
);


CREATE TABLE markets (
	id uuid NOT NULL PRIMARY KEY,
	pair_id uuid REFERENCES pairs(id),
	exchange_id uuid REFERENCES exchanges(id),
	is_active boolean NOT NULL
);


CREATE TABLE prices (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	pair_id uuid REFERENCES pairs(id),
	exchange_id uuid REFERENCES exchanges(id),
	market_id uuid REFERENCES markets(id),
	price real NOT NULL,
	pair_symbol text NOT NULL,
	exchange_symbol text NOT NULL
);


CREATE TABLE users (
	id uuid NOT NULL PRIMARY KEY,
	email text NOT NULL,
	password text NOT NULL,
	last_login timestamptz NOT NULL,
	subscribe_to timestamptz NOT NULL,
	role text NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL
);


CREATE TABLE orders (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	user_id uuid REFERENCES users(id),
	exchange_id uuid REFERENCES exchanges(id),
	pair_id uuid REFERENCES pairs(id),
	market_id uuid REFERENCES markets(id),
	order_type text NOT NULL,
	open_price real NOT NULL,
	close_price real NOT NULL,
	ordered_volume real NOT NULL,
	swapped_volume real NOT NULL,
	is_closed boolean NOT NULL,
	stop_loss real NOT NULL,
	take_profit real NOT NULL,
	buy_fee real NOT NULL,
	sell_fee real NOT NULL,
	delta real NOT NULL
);


CREATE TABLE trades (
	id uuid NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	user_id uuid REFERENCES users(id),
	exchange_id uuid REFERENCES exchanges(id),
	pair_id uuid REFERENCES pairs(id),
	market_id uuid REFERENCES markets(id),
	order_id uuid REFERENCES orders(id),
	type text NOT NULL,
	volume real NOT NULL,
	price real NOT NULL
);


CREATE TABLE user_balances (
	id uuid NOT NULL PRIMARY KEY,
	user_id uuid REFERENCES users(id),
	exchange_id uuid REFERENCES exchanges(id),
	asset_id uuid REFERENCES assets(id),
	volume real NOT NULL
);


COMMIT;

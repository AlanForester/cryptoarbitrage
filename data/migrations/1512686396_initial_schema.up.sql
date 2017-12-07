BEGIN;

ALTER TABLE differences ADD COLUMN pair_symbol text NOT NULL;

ALTER TABLE differences ADD COLUMN exchange_symbol text NOT NULL;

ALTER TABLE prices DROP COLUMN pair_symbol;

ALTER TABLE prices DROP COLUMN exchange_symbol;

ALTER TABLE prices ADD COLUMN pair_symbols text[] NOT NULL;

ALTER TABLE prices ADD COLUMN exchange_symbols text[] NOT NULL;

COMMIT;

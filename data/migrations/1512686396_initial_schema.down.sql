BEGIN;

ALTER TABLE differences DROP COLUMN pair_symbol;

ALTER TABLE differences DROP COLUMN exchange_symbol;

ALTER TABLE prices ADD COLUMN pair_symbol text NOT NULL;

ALTER TABLE prices ADD COLUMN exchange_symbol text NOT NULL;

ALTER TABLE prices DROP COLUMN pair_symbols;

ALTER TABLE prices DROP COLUMN exchange_symbols;

COMMIT;

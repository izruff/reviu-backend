BEGIN;

ALTER TABLE users ADD COLUMN rating integer default 0;

COMMIT;
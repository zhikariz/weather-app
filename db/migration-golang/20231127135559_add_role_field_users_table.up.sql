BEGIN;

ALTER TABLE
    "public"."users"
ADD
    COLUMN role VARCHAR(255);

COMMIT;
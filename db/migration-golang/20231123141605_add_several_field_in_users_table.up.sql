BEGIN;

ALTER TABLE "public"."users"
ADD COLUMN email VARCHAR(255),
ADD COLUMN password VARCHAR(255);

COMMIT;
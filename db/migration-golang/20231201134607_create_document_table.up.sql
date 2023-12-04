BEGIN;

CREATE TABLE IF NOT EXISTS "public"."documents" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "user_id" INT NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6)
);

COMMIT;
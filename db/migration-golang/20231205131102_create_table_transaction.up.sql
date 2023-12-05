BEGIN;

CREATE TABLE IF NOT EXISTS "public"."transactions" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "order_id" varchar(255) NOT NULL,
    "user_id" INT NOT NULL,
    "amount" INT NULL,
    "status" varchar(255) NOT NULL
);

COMMIT;
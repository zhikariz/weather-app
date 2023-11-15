CREATE TABLE
  "public"."users" (
    "id" SERIAL NOT NULL,
    "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6)
  );

ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
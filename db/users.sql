/*
 Navicat Premium Data Transfer

 Source Server         : DOCKER-POSTGRES-LOCAL
 Source Server Type    : PostgreSQL
 Source Server Version : 130011 (130011)
 Source Host           : localhost:5432
 Source Catalog        : weather-app
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130011 (130011)
 File Encoding         : 65001

 Date: 15/11/2023 21:08:00
*/


-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."users" OWNER TO "postgres";

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

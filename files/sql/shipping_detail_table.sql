-- +goose Up
CREATE TABLE shipping_detail (
	"id" SERIAL,
	"shipping_detail_id" int8 NOT NULL,
	"status" int2 NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);


-- +goose down
DROP TABLE shipping_detail;
-- +goose Up
CREATE TABLE shop_order_detail (
	"id" SERIAL,
	"order_id" int8 NOT NULL,
	"product_id" int8 NOT NULL,
	"amount" int4 NOT NULL,
	"price" int8 NOT NULL,
	 "total" int8 NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);

-- +goose down
DROP TABLE shop_order_detail;
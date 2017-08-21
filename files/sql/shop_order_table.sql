-- +goose Up
CREATE TABLE shop_order (
	"id" SERIAL,
	"order_detail_id" int8 NOT NULL,
	"shipping_id" int8 NOT NULL,
	"voucher_id" int8 NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);


-- +goose down
DROP TABLE shop_order;
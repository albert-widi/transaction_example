-- +goose Up
CREATE TABLE shop_order (
	"id" SERIAL,
	"user_id" int8 NOT NULL,
	"shipping_id" int8 NULL,
	"voucher_id" int8 NULL,
	"payment_confirmed" boolean NULL,
	"total" int8 NULL,
	"status" int2 NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);

-- +goose down
DROP TABLE shop_order;
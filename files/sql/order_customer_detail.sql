-- +goose Up
CREATE TABLE order_customer_detail (
	"id" SERIAL,
	"order_id" int8 NOT NULL,
    "name" varchar(100) NOT NULL,
    "phone_number" varchar(100) NOT NULL,
    "email_address" varchar(100) NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);

-- +goose down
DROP TABLE order_shipping;
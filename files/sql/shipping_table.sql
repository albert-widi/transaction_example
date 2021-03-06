-- +goose Up
CREATE TABLE shipping (
	"id" SERIAL,
    "shipper_id" int8 NOT NULL,
    "price" int8 NOT NULL,
    "from" varchar(100) NOT NULL,
    "to" varchar(100) NOT NULL,
	"status" int2 NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);

-- +goose down
DROP TABLE shipping;
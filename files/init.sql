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

CREATE TABLE order_customer_detail (
	"id" SERIAL,
	"order_id" int8 NOT NULL,
    "name" varchar(100) NOT NULL,
    "phone_number" varchar(100) NOT NULL,
    "email" varchar(100) NOT NULL,
    "address" varchar(100) NOT NULL,
	"created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  	"updated_at" timestamp NULL, 
	PRIMARY KEY ("id")
);
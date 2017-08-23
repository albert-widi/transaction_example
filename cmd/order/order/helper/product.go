package helper

import (
	"context"

	"github.com/albert-widi/transaction_example/apicalls"
	"github.com/albert-widi/transaction_example/log"
)

// type productHelper struct{}

// Product struct
type Product struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  int64  `json:"price,omitempty"`
	Stock  int    `json:"stock,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

func FindProductByID(ctx context.Context, productID int64) (Product, error) {
	prod := Product{}
	err := apicalls.Product.GetProductByID(ctx, productID, &prod)
	if err == nil {
		log.Debug("[FindProductByID] Product: %+v", prod)
	}
	return prod, err
}

func DecreaseProductStock(ctx context.Context, productID, amount int64) error {
	return apicalls.Product.DecreaseProductStock(ctx, productID, amount)
}

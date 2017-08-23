package product

import "sync"
import "github.com/albert-widi/transaction_example/errors"

// Product struct
type Product struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  int64  `json:"price,omitempty"`
	Stock  int    `json:"stock,omitempty"`
	Weight int    `json:"weight,omitempty"`
	mutex  sync.Mutex
}

// Hardcoded product list
var productList = map[int64]*Product{
	1: &Product{
		ID:     1,
		Name:   "Mesin cuci",
		Price:  100,
		Stock:  10,
		Weight: 10,
	},
	2: &Product{
		ID:     2,
		Name:   "Pohon natal",
		Price:  120,
		Stock:  3,
		Weight: 12,
	},
	3: &Product{
		ID:     2,
		Name:   "Penjepit baju",
		Price:  3,
		Stock:  200,
		Weight: 1,
	},
}

// GetProductList return list of product
func GetProductList() ([]*Product, error) {
	prodLength := len(productList)
	p := make([]*Product, prodLength)
	counter := 0
	for _, prod := range productList {
		p[counter] = prod
		counter++
	}
	return p, nil
}

func GetProduct(productID int64) (*Product, error) {
	p, ok := productList[productID]
	if !ok {
		return nil, errors.New(errors.ProductNotExists)
	}
	return p, nil
}

func DecreaseStock(productID int64, amount int) error {
	prod := productList[productID]
	prod.mutex.Lock()
	defer prod.mutex.Unlock()

	if prod.Stock-amount > 0 {
		prod.Stock -= amount
	} else {
		errors.New(errors.ProductStockEmpty)
	}
	return nil
}

func IncreaseStock(productID int64, amount int) error {
	prod := productList[productID]
	prod.mutex.Lock()
	defer prod.mutex.Unlock()
	prod.Stock += amount
	return nil
}

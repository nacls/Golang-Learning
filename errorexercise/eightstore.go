package errorexercise

import (
	"fmt"
	"sort"
	"strings"
)

type Store struct {
	products []Product
}

type Product struct {
	name  string
	price float64
	count int
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	for _, product := range s.products {
		if strings.ToLower(product.name) == strings.ToLower(name) {
			return fmt.Errorf("%s already exists", name)
		}
	}

	if price <= 0 {
		return fmt.Errorf("price should be positive")
	}

	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	s.products = append(s.products, Product{name: strings.ToLower(name), price: price, count: count})
	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	for _, product := range s.products {
		if strings.ToLower(product.name) == strings.ToLower(name) {
			return product.count, nil
		}
	}
	return 0, fmt.Errorf("invalid product name")
}

func (s *Store) GetProductPrice(name string) (float64, error) {
	for _, product := range s.products {
		if strings.ToLower(product.name) == strings.ToLower(name) {
			return product.price, nil
		}
	}
	return 0, fmt.Errorf("invalid product name")
}

func (s *Store) Order(name string, count int) error {
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	productExists := false
	for _, product := range s.products {
		if product.name == name {
			productExists = true
		}
	}
	if !productExists {
		return fmt.Errorf("invalid product name")
	}

	for i := range s.products {
		if s.products[i].name == name {
			if s.products[i].count == 0 {
				return fmt.Errorf("there is no %s in the store", name)
			}
			if s.products[i].count < count {
				return fmt.Errorf("not enough %s in the store. there are %d left", name, s.products[i].count)
			}
			s.products[i].count -= count
		}
	}
	return nil
}

func (s *Store) ProductsList() ([]string, error) {
	if len(s.products) == 0 {
		return nil, fmt.Errorf("store is empty")
	}
	atLeastOne := false
	productNames := make([]string, 0)
	for _, product := range s.products {
		if product.count > 0 {
			productNames = append(productNames, product.name)
			atLeastOne = true
		}
	}
	if atLeastOne {
		sort.Strings(productNames)
		return productNames, nil
	}
	return nil, fmt.Errorf("store is empty")
}

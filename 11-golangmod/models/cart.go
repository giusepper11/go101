package models

import (
	"errors"
	"time"
)

type ShoppingCart struct {
	Market    string
	Items     []Items
	CreatedAt time.Time
}

type Items struct {
	Name string
}

func StartShoppingCart(market string, items_list []string) (*ShoppingCart, error) {

	if market == "" {
		return nil, errors.New("market is mandatory")
	}

	if len(items_list) == 0 {
		return nil, errors.New("items are mandatory")
	}

	var items []Items
	for _, item := range items_list {
		items = append(items, Items{Name: item})
	}

	cart := ShoppingCart{
		Market:    market,
		Items:     items,
		CreatedAt: time.Now(),
	}

	return &cart, nil
}

package app

import "github.com/mhdianrush/go-ecommerce-api/app/entities"

type Entity struct {
	Entity any
}

func RegisterEntities() []Entity {
	return []Entity{
		{
			Entity: entities.User{},
		},
		{
			Entity: entities.Address{},
		},
		{
			Entity: entities.Product{},
		},
		{
			Entity: entities.ProductImage{},
		},
		{
			Entity: entities.Section{},
		},
		{
			Entity: entities.Category{},
		},
		{
			Entity: entities.Order{},
		},
		{
			Entity: entities.OrderItem{},
		},
		{
			Entity: entities.OrderCustomer{},
		},
		{
			Entity: entities.Payment{},
		},
		{
			Entity: entities.Shipment{},
		},
	}
}

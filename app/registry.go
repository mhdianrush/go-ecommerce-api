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
	}
}
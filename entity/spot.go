package entity

import "github.com/shengng325/clean-architecture/v2/model"

type Spot struct {
	ID          int
	DisplayName string
	Position    *model.Position
	Category    *model.PlaceCategory
}

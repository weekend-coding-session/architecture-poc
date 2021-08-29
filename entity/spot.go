package entity

import "architecture-poc/model"

type Spot struct {
	ID          int
	DisplayName string
	Position    *model.Position
	Category    *model.PlaceCategory
}

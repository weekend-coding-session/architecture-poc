package dto

import "architecture-poc/model"

type Spot struct {
	DisplayName string
	Position    *model.Position
	Category    *model.PlaceCategory
}

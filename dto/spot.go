package dto

import "github.com/shengng325/clean-architecture/v2/model"

type Spot struct {
	DisplayName string
	Position    *model.Position
	Category    *model.PlaceCategory
}

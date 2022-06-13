package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	ItemName     string             `json:"itemname,omitempty" validate:"required"`
	Location     string             `json:"location,omitempty" validate:"required"`
	ItemPrice    int                `json:"itemprice,omitempty" validate:"required"`
	ItemBrand    string             `json:"itembrand,omitempty" validate:"required"`
	ItemCategory string             `json:"itemcategory,omitempty" validate:"required"`
}

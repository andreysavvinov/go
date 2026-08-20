package main

import "time"

type Order struct {
	ID int `db:"id" json:"id" validate:"required"`
	UserID int `db:"user_id" json:"user_id" validate:"required"`
	Product string `db:"product" json:"product" validate:"required,min=1"`
	Price  float32 `db:"price" json:"price" validate:"gte=0.0"`
	CreatedAt time.Time `db:"created_at" json:"created_at" validate:"required"`
}
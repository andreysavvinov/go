package main

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func TestValidateOrder(t *testing.T) {
	validate = validator.New()

	order := Order{
		ID:        1,
		UserID:    10,
		Product:   "MacBook",
		Price:     160000.0,
		CreatedAt: time.Now(),
	}

	if err := validate.Struct(order); err != nil {
		t.Fatalf("valid order failed validation: %v", err)
	}
}

func TestValidateOrderInvalid(t *testing.T) {
	validate = validator.New()

	tests := []struct {
		name  string
		order Order
	}{
		{
			name: "missing ID",
			order: Order{
				UserID:    10,
				Product:   "MacBook",
				Price:     160000.0,
				CreatedAt: time.Now(),
			},
		},
		{
			name: "missing UserID",
			order: Order{
				ID:        1,
				Product:   "MacBook",
				Price:     160000.0,
				CreatedAt: time.Now(),
			},
		},
		{
			name: "empty Product",
			order: Order{
				ID:        1,
				UserID:    10,
				Product:   "",
				Price:     160000.0,
				CreatedAt: time.Now(),
			},
		},
		{
			name: "negative Price",
			order: Order{
				ID:        1,
				UserID:    10,
				Product:   "MacBook",
				Price:     -100,
				CreatedAt: time.Now(),
			},
		},
		{
			name: "missing CreatedAt",
			order: Order{
				ID:      1,
				UserID:  10,
				Product: "MacBook",
				Price:   160000.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.Struct(tt.order); err == nil {
				t.Errorf("expected validation error for %s", tt.name)
			}
		})
	}
}
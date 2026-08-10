package main

import "testing"

func TestValidateEmail(t *testing.T) {

	tests := []struct {
		email string
		valid bool
	}{
		{"test@mail.ru", true},
		{"abc@gmail.com", true},
		{"bad@", false},
		{"@", false},
		{"", false},
	}

	for _, tt := range tests {

		err := ValidateEmail(tt.email)

		if tt.valid && err != nil {
			t.Errorf("%s should be valid", tt.email)
		}

		if !tt.valid && err == nil {
			t.Errorf("%s should be invalid", tt.email)
		}
	}
}

func TestValidateAge(t *testing.T) {

	tests := []struct {
		age   int
		valid bool
	}{
		{20, true},
		{0, true},
		{150, true},
		{-1, false},
		{200, false},
	}

	for _, tt := range tests {

		err := ValidateAge(tt.age)

		if tt.valid && err != nil {
			t.Fail()
		}

		if !tt.valid && err == nil {
			t.Fail()
		}
	}
}

func TestValidateName(t *testing.T) {

	tests := []struct {
		name  string
		valid bool
	}{
		{"Ivan", true},
		{"Анна", true},
		{"", false},
		{"   ", false},
	}

	for _, tt := range tests {

		err := ValidateName(tt.name)

		if tt.valid && err != nil {
			t.Fail()
		}

		if !tt.valid && err == nil {
			t.Fail()
		}
	}
}

func TestValidateUser(t *testing.T) {

	user := User{
		Name:     "Ivan",
		Email:    "ivan@mail.ru",
		Age:      25,
		IsActive: true,
	}

	if err := ValidateUser(user); err != nil {
		t.Fatal(err)
	}
}

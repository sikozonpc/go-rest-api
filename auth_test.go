package main

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	userID := int64(49)

	token, err := CreateJWT(secret, userID)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("expected token to not be empty")
	}
}

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if len(hash) == 0 {
		t.Errorf("expected hash to not be empty")
	}

	if string(hash) == "password" {
		t.Errorf("expected hash to not be equal to password")
	}
}


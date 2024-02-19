package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestValidateUserPayload(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name string
		args args
		want error
	}{

		{
			name: "should return error if email is empty",
			args: args{
				user: &User{
					FirstName: "John",
					LastName:  "Doe",
				},
			},
			want: errEmailRequired,
		},
		{
			name: "should return error if first name is empty",
			args: args{
				user: &User{
					Email:    "joe@mail.com",
					LastName: "Doe",
				},
			},
			want: errFirstNameRequired,
		},
		{
			name: "should return error if last name is empty",
			args: args{
				user: &User{
					Email:     "joe@mail.com",
					FirstName: "John",
				},
			},
			want: errLastNameRequired,
		},
		{
			name: "should return error if the password is empty",
			args: args{
				user: &User{
					Email:     "joe@mail.com",
					FirstName: "John",
				},
			},
			want: errLastNameRequired,
		},
		{
			name: "should return nil if all fields are present",
			args: args{
				user: &User{
					Email:     "joe@mail.com",
					FirstName: "John",
					LastName:  "Doe",
					Password:  "password",
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateUserPayload(tt.args.user); got != tt.want {
				t.Errorf("validateUserPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	// Create a new project
	ms := &MockStore{}
	service := NewUserService(ms)

	t.Run("should validate if the email is not empty", func(t *testing.T) {
		payload := &RegisterPayload{
			Email:     "",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "password",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users/register", service.handleUserRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var response ErrorResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Error != errEmailRequired.Error() {
			t.Errorf("expected error message %s, got %s", response.Error, errEmailRequired.Error())
		}
	})

	t.Run("should create a user", func(t *testing.T) {
		payload := &RegisterPayload{
			Email:     "joe@mail.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "password",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users/register", service.handleUserRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

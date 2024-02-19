package main

import "time"

type CreateProjectPayload struct {
	Name        string `json:"name"`
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
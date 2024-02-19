package main

import "database/sql"

type Storage struct {
	db *sql.DB
}

type Store interface {
	CreateProject(p *Project) error
	GetProject(id string) (*Project, error)
	DeleteProject(id string) error
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateProject(p *Project) error {
	_, err := s.db.Exec("INSERT INTO projects (name) VALUES (?)", p.Name)
	return err
}

func (s *Storage) GetProject(id string) (*Project, error) {
	var p Project
	err := s.db.QueryRow("SELECT id, name, createdAt FROM projects WHERE id = ?", id).Scan(&p.ID, &p.Name, &p.CreatedAt)
	return &p, err
}

func (s *Storage) DeleteProject(id string) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

package main

import "database/sql"

type Storage struct {
	db *sql.DB
}

type Store interface {
	// Users
	CreateUser(u *User) (*User, error)
	GetUserByID(id string) (*User, error)
	// Projects
	CreateProject(p *Project) error
	GetProject(id string) (*Project, error)
	DeleteProject(id string) error
	// Tasks
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
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

func (s *Storage) CreateUser(u *User) (*User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)", u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (s *Storage) GetUserByID(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedToID)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)
	return &t, err
}

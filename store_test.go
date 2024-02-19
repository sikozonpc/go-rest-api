package main

// Mocks

type MockStore struct{}

func (s *MockStore) CreateProject(p *Project) error {
	return nil
}

func (s *MockStore) GetProject(id string) (*Project, error) {
	return &Project{Name: "Super cool project"}, nil
}

func (s *MockStore) DeleteProject(id string) error {
	return nil
}

func (s *MockStore) CreateUser(u *User) (*User, error) {
	return &User{}, nil
}

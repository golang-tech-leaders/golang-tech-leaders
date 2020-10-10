package database

// RuntimeStorage is a runtime storage manager
type RuntimeStorage struct {
	users map[string]string
}

// NewRuntimeStorage creates new runtme storage
func NewRuntimeStorage() *RuntimeStorage {
	return &RuntimeStorage{
		users: make(map[string]string),
	}
}

// GetUser returns a user
func (s *RuntimeStorage) GetUser(id string) string {
	return s.users[id]
}

// SetUser adds a useer to the storage
func (s *RuntimeStorage) SetUser(id, name string) {
	s.users[id] = name
}

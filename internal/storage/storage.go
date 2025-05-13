package storage

import (
	"errors"
	"sync"

	"github.com/yourusername/todo-app/internal/models"
)

// Common errors
var (
	ErrNotFound = errors.New("todo not found")
)

// TodoStorage defines the interface for todo storage implementations
type TodoStorage interface {
	Create(todo *models.Todo) error
	GetByID(id int) (*models.Todo, error)
	GetAll() ([]*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id int) error
}

// MemoryStorage implements TodoStorage using in-memory storage
type MemoryStorage struct {
	todos     map[int]*models.Todo
	nextID    int
	todoMutex sync.RWMutex
}

// NewMemoryStorage creates a new in-memory storage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos:  make(map[int]*models.Todo),
		nextID: 1,
	}
}

// Create adds a new todo to storage
func (s *MemoryStorage) Create(todo *models.Todo) error {
	s.todoMutex.Lock()
	defer s.todoMutex.Unlock()

	todo.ID = s.nextID
	s.todos[todo.ID] = todo
	s.nextID++
	return nil
}

// GetByID retrieves a todo by its ID
func (s *MemoryStorage) GetByID(id int) (*models.Todo, error) {
	s.todoMutex.RLock()
	defer s.todoMutex.RUnlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, ErrNotFound
	}
	return todo, nil
}

// GetAll retrieves all todos
func (s *MemoryStorage) GetAll() ([]*models.Todo, error) {
	s.todoMutex.RLock()
	defer s.todoMutex.RUnlock()

	todos := make([]*models.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Update updates an existing todo
func (s *MemoryStorage) Update(todo *models.Todo) error {
	s.todoMutex.Lock()
	defer s.todoMutex.Unlock()

	if _, exists := s.todos[todo.ID]; !exists {
		return ErrNotFound
	}
	s.todos[todo.ID] = todo
	return nil
}

// Delete removes a todo by its ID
func (s *MemoryStorage) Delete(id int) error {
	s.todoMutex.Lock()
	defer s.todoMutex.Unlock()

	if _, exists := s.todos[id]; !exists {
		return ErrNotFound
	}
	delete(s.todos, id)
	return nil
}

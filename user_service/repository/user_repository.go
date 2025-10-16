package repository

import (
	"sync"
	"time"

	"user-service/model"
)

type UserRepository interface {
	List(page, size int) []model.User
	Create(name string) (model.User, error)
	GetByID(id int) (model.User, error)
}

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users []model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []model.User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}},
	}
}

func (r *InMemoryUserRepository) List(page, size int) []model.User {
	start := (page - 1) * size
	if start >= len(r.users) {
		return []model.User{}
	}

	end := start + size
	if end > len(r.users) {
		end = len(r.users)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]model.User, end-start)
	copy(out, r.users[start:end])
	return out
}

func (r *InMemoryUserRepository) Create(name string) (model.User, error) {
	user := model.User{ID: r.nextId(), Name: name, CreatedAt: time.Now().UnixMicro(), UpdatedAt: time.Now().UnixMicro()}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.users = append(r.users, user)
	return user, nil
}

func (r *InMemoryUserRepository) GetByID(id int) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, ErrNotFound
}

func (r *InMemoryUserRepository) nextId() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	maxID := 0
	for _, u := range r.users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	return maxID + 1
}

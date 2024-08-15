package server

import (
	"sync"
)

type Storage struct {
	mu       sync.Mutex
	handlers map[*Handler]bool
}

func (s *Storage) addUser(h *Handler, f func(*Handler)) {
	s.mu.Lock()
	f(h)
	s.handlers[h] = true
	defer s.mu.Unlock()
}

// add delete user

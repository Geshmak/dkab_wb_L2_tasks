package memory

import (
	"sync"
	"wb_2_11/internal/model"
)

type MemoryDB struct {
	sync.RWMutex
	data   map[int]model.Event
	lastID int
}

func New() *MemoryDB {
	return &MemoryDB{
		RWMutex: sync.RWMutex{},
		data:    make(map[int]model.Event, 0),
	}
}

func (m *MemoryDB) nextID() int {
	m.lastID++
	return m.lastID
}

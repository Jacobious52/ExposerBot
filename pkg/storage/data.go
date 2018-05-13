package storage

import (
	"strings"
	"sync"
)

// ReadStore interface allows get only access to the table without add functions
type ReadStore interface {
	GetTable(string) (FreqTable, bool)
	GetUsers() []string
}

// DataStore defines the storage structure that owns a model
type DataStore struct {
	sync.RWMutex

	model model
	path  string
	dirty bool
}

// NewDataStore creates a new datastore with an empty model
func NewDataStore(path string) *DataStore {
	return &DataStore{
		model: make(map[string]FreqTable),
		path:  path,
	}
}

// AddUser adds a user to the model
// If user already exists, nothing will happen
func (s *DataStore) AddUser(user string) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.model[user]; ok {
		return
	}

	s.model[user] = make(FreqTable)
	s.dirty = true
}

// AddWord adds a word to the frequency table for a user
// if user does not exist, it will be created
func (s *DataStore) AddWord(user string, word string) {
	s.Lock()
	defer s.Unlock()

	word = strings.ToLower(word)

	s.dirty = true
	if table, ok := s.model[user]; ok {
		table[word]++
		return
	}

	// create the user with a table with that word
	s.model[user] = FreqTable{word: 1}
}

// GetTable returns a frequency table for a user
func (s *DataStore) GetTable(user string) (FreqTable, bool) {
	s.RLock()
	defer s.RUnlock()

	if f, ok := s.model[user]; ok {
		// make a copy so we can release the lock
		cpy := make(FreqTable, len(f))
		for k, v := range f {
			cpy[k] = v
		}
		return cpy, ok
	}

	return nil, false
}

// GetUsers returns a array of users for the table
func (s *DataStore) GetUsers() []string {
	s.RLock()
	defer s.RUnlock()

	users := make([]string, 0, len(s.model))
	for user := range s.model {
		users = append(users, user)
	}
	return users
}

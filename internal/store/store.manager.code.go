package store

import (
	"github.com/lishimeng/x/container"
	"github.com/patrickmn/go-cache"
	"time"
)

func GetManager() Manager {
	var m Manager
	_ = container.Get(&m)
	return m
}

type Manager interface {
	GetDefaultStore() CodeStore
}

type SimpleStoreManager struct {
	defaultStore CodeStore
}

func NewStoreManager() (s Manager) {
	ssm := &SimpleStoreManager{}
	ssm.defaultStore = NewRamStore()
	s = ssm
	return
}

func (s *SimpleStoreManager) GetDefaultStore() CodeStore {
	return s.defaultStore
}

type Option func()

type CodeStore interface {
	Save(code string, payload string, options ...Option)
	Load(code string) (string, bool) // payload, found
}

type MapRamCodeStore struct {
	c *cache.Cache
}

func (m *MapRamCodeStore) Save(code string, payload string, options ...Option) {
	m.c.Set(code, payload, cache.DefaultExpiration)
}

func (m *MapRamCodeStore) Load(code string) (string, bool) {
	v, found := m.c.Get(code)
	if !found {
		return "", false
	}
	m.c.Delete(code)
	return v.(string), true
}

func NewRamStore() (cs CodeStore) {
	s := &MapRamCodeStore{c: cache.New(5*time.Minute, 10*time.Minute)}
	cs = s
	return
}

package syncmap

import "sync"

// SyncMap is a map with a builtin mutex guard that lock and unlocks itself, as required
type SyncMap[K comparable, V any] struct {
	lock sync.Mutex
	m    map[K]V
}

func NewSyncMap[K comparable, V any](cap int) SyncMap[K, V] {
	return SyncMap[K, V]{
		lock: sync.Mutex{},
		m:    make(map[K]V, cap),
	}
}

// Get same semantics as a regular map's get
func (s *SyncMap[K, V]) Get(key K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.m[key]
	return v, ok
}

// Insert add the value for key, returning the old value and true if any was set before
func (s *SyncMap[K, V]) Insert(key K, val V) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	prev, ok := s.m[key]

	s.m[key] = val
	return prev, ok
}

func (s *SyncMap[K, V]) Delete(key K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	prev, ok := s.m[key]
	delete(s.m, key)
	return prev, ok
}

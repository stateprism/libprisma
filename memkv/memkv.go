package memkv

import (
	"strings"
	"sync"
)

type Opts struct {
	CaseInsensitive bool
}

type MemKV struct {
	sync.RWMutex
	sep       string
	caseSense bool
	m         map[string]any
}

func NewMemKV(sep string, opts *Opts) *MemKV {
	s := &MemKV{
		RWMutex:   sync.RWMutex{},
		sep:       sep,
		caseSense: true,
		m:         make(map[string]any),
	}

	if opts == nil {
		return s
	}

	if opts.CaseInsensitive {
		s.caseSense = false
	}

	return s
}

func (m *MemKV) Get(key string) (any, bool) {
	m.RLock()
	defer m.RUnlock()
	if !m.caseSense {
		key = strings.ToLower(key)
	}
	view := m.m
	keys := strings.Split(key, ".")
	if strings.Contains(key, ".") {
		for i, k := range keys {
			if i == len(keys)-1 {
				break
			}
			viewTemp, ok := view[k].(map[string]any)
			if !ok {
				return nil, false
			}
			view = viewTemp
		}
	}
	key = keys[len(keys)-1]
	val, ok := view[key]
	if !ok {
		return nil, ok
	}
	return val, true
}

func (m *MemKV) Set(key string, val any) bool {
	m.Lock()
	defer m.Unlock()
	if !m.caseSense {
		key = strings.ToLower(key)
	}
	view := m.m
	keys := strings.Split(key, ".")
	if strings.Contains(key, ".") {
		for i, k := range keys {
			if i == len(keys)-1 {
				break
			}
			if v, ok := view[k]; !ok {
				view[k] = map[string]any{}
				view = view[k].(map[string]any)
			} else {
				viewTemp, ok := v.(map[string]any)
				if !ok {
					return false
				}
				view = viewTemp
			}
		}
	}
	key = keys[len(keys)-1]
	view[key] = val
	return true
}

func (m *MemKV) Contains(key string) bool {
	_, ok := m.Get(key)
	return ok
}

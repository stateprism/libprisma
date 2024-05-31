package memkv

import (
	"strings"
	"sync"
)

type MemKV struct {
	sync.RWMutex
	sep string
	m   map[string]any
}

func NewMemKV(sep string) *MemKV {
	return &MemKV{
		sync.RWMutex{},
		sep,
		make(map[string]any),
	}
}

func (m *MemKV) Get(k string) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	view := m.m
	keys := strings.Split(k, ".")
	if strings.Contains(k, ".") {
		for i, k := range keys {
			if i == len(keys)-1 {
				break
			}
			viewTemp, ok := view[k].(map[string]interface{})
			if !ok {
				return nil, false
			}
			view = viewTemp
		}
	}
	k = keys[len(keys)-1]
	val, ok := view[k]
	if !ok {
		return nil, ok
	}
	return val, true
}

func (m *MemKV) Set(k string, val interface{}) bool {
	m.Lock()
	defer m.Unlock()
	view := m.m
	keys := strings.Split(k, ".")
	if strings.Contains(k, ".") {
		for i, k := range keys {
			if i == len(keys)-1 {
				break
			}
			if v, ok := view[k]; !ok {
				view[k] = map[string]interface{}{}
				view = view[k].(map[string]interface{})
			} else {
				viewTemp, ok := v.(map[string]interface{})
				if !ok {
					return false
				}
				view = viewTemp
			}
		}
	}
	k = keys[len(keys)-1]
	view[k] = val
	return true
}

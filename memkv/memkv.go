package memkv

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"
)

type Opts struct {
	CaseInsensitive bool
}

type EventType int

const (
	E_KEY_CREATED  = iota
	E_KEY_UPDATED  = iota
	E_KEY_ACCESSED = iota
)

type Event struct {
	Key        string
	Type       EventType
	When       time.Time
	Success    bool
	FailReason string
	OldVal     any
	NewVal     any
}

type WatchHook func(e Event)
type Trigger func(self *MemKV, e Event)

type eHandler struct {
	hook         WatchHook
	trigger      Trigger
	eventsFilter []EventType
}

type MemKV struct {
	l         sync.RWMutex
	sep       string
	caseSense bool
	m         map[string]any
	watchers  map[string][]eHandler
}

func NewMemKV(sep string, opts *Opts) *MemKV {
	s := &MemKV{
		l:         sync.RWMutex{},
		sep:       sep,
		caseSense: true,
		m:         make(map[string]any),
		watchers:  make(map[string][]eHandler),
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
	m.l.RLock()
	defer m.l.RUnlock()
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
	e := Event{
		Key:     key,
		Type:    E_KEY_ACCESSED,
		When:    time.Now(),
		Success: true,
	}
	m.dispatchWatchers(e)
	return val, true
}

func (m *MemKV) Set(key string, val any) bool {
	m.l.Lock()
	defer m.l.Unlock()
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
	if v, ok := view[key]; !ok {
		e := Event{
			Key:     key,
			Type:    E_KEY_CREATED,
			NewVal:  val,
			When:    time.Now(),
			Success: true,
		}
		m.dispatchWatchers(e)
	} else {
		e := Event{
			Key:     key,
			Type:    E_KEY_UPDATED,
			NewVal:  val,
			OldVal:  v,
			When:    time.Now(),
			Success: true,
		}
		m.dispatchWatchers(e)
	}
	view[key] = val
	return true
}

func (m *MemKV) dispatchWatchers(e Event) {
	var wg sync.WaitGroup
	for _, w := range m.watchers[e.Key] {
		if slices.Contains(w.eventsFilter, e.Type) && w.hook != nil {
			wg.Add(1)
			go func() {
				defer wg.Done()
				w.hook(e)
			}()
		}
	}
	wg.Wait()
	fmt.Println("done")
}

func (m *MemKV) AddWatcherHook(key string, hook WatchHook, eFilter []EventType) {
	m.l.Lock()
	defer m.l.Unlock()
	if !m.caseSense {
		key = strings.ToLower(key)
	}

	handler := eHandler{
		hook:         hook,
		trigger:      nil,
		eventsFilter: eFilter,
	}

	if _, ok := m.watchers[key]; !ok {
		m.watchers[key] = []eHandler{handler}
	} else {
		m.watchers[key] = append(m.watchers[key], handler)
	}
}

func (m *MemKV) Contains(key string) bool {
	_, ok := m.Get(key)
	return ok
}

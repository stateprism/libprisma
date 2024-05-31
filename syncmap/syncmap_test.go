package syncmap_test

import (
	"github.com/stateprism/libprisma/syncmap"
	"testing"
)

func TestSyncMap_Insert(t *testing.T) {
	tests := []struct {
		Name      string
		Key       string
		Value     int
		OldValue  int
		HasOldVal bool
	}{
		{
			Name:  "insert no old",
			Key:   "test1",
			Value: 2,
		},
		{
			Name:      "insert with old",
			Key:       "test_old_val",
			Value:     2,
			OldValue:  5,
			HasOldVal: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sm := syncmap.NewSyncMap[string, int](0)
			if tt.HasOldVal {
				sm.Insert(tt.Key, tt.OldValue)
			}

			old, had := sm.Insert(tt.Key, tt.Value)
			if tt.HasOldVal && !had {
				t.Error("test didn't expect old value, but it was present")
			} else if tt.HasOldVal && old != tt.OldValue {
				t.Error("the old value didn't match what was expected")
			} else if !tt.HasOldVal && had {
				t.Error("test expected old value, but it was not present")
			}

			nv, _ := sm.Get(tt.Key)
			if nv != tt.Value {
				t.Error("set value could not be retrieved")
			}
		})
	}
}

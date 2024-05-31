package memkv_test

import (
	"github.com/stateprism/libprisma/memkv"
	"testing"
)

type expectedValues struct {
	k    string
	v    int
	fail bool
}

func TestMemKV_Get(t *testing.T) {
	tests := []struct {
		Name   string
		Expect []expectedValues
	}{
		{
			Name: "Find simple",
			Expect: []expectedValues{
				{k: "testKey", v: 123},
			},
		},
		{
			Name: "Find path",
			Expect: []expectedValues{
				{k: "Some.test.path", v: 42},
			},
		},
		{
			Name: "Get multi path",
			Expect: []expectedValues{
				{k: "Some.test.path", v: 42},
				{k: "Some.test.leaf", v: 100},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			kvs := memkv.NewMemKV(".")
			for _, e := range tt.Expect {
				kvs.Set(e.k, e.v)
				if v, ok := kvs.Get(e.k); ok {
					if v != e.v {
						t.Error("Failed to retrieve expected value")
					}
				}
			}
		})
	}
}

func TestMemKV_Set(t *testing.T) {
	tests := []struct {
		Name   string
		Expect []expectedValues
	}{
		{
			Name: "Set simple",
			Expect: []expectedValues{
				{k: "testKey", v: 123},
			},
		},
		{
			Name: "Set path",
			Expect: []expectedValues{
				{k: "Some.test.path", v: 42},
			},
		},
		{
			Name: "Set multi path",
			Expect: []expectedValues{
				{k: "Some.test.path", v: 42},
				{k: "Some.test.leaf", v: 100},
				{k: "Some.test.leaf.notAllowed", v: 100, fail: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			kvs := memkv.NewMemKV(".")
			for _, e := range tt.Expect {
				if ok := kvs.Set(e.k, e.v); !ok && !e.fail {
					t.Error("Failed to set key unexpectedly")
				} else if ok && e.fail {
					t.Errorf("Succeess when failure was expected")
				}
				v, _ := kvs.Get(e.k)
				if v != e.v && !e.fail {
					t.Error("Failed to retrieve expected value")
				}
			}
		})
	}
}

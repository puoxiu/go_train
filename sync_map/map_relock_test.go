package syncmap

import (
	"sync"
	"testing"
)

type pair struct {
	key string
	val string
}

func TestMapSyncMap(t *testing.T) {
	mp := NewSyncMap[string, string]()
	var wg sync.WaitGroup
	tests := []pair{
		{"1", "aaaa"},
		{"2", "bbbb"},
		{"3", "hello word"},
		{"4", "golang"},
		{"5", "sync map"},
		{"6", "sync map"},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				mp.Set(test.key, test.val)
				val, ok := mp.Get(test.key)
				if !ok || val != test.val {
					t.Errorf("key:%s, want:%s, got:%s", test.key, test.val, val)
				}
			}()
		})
	}
}
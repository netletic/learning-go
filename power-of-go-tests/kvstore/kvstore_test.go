package kvstore_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/netletic/kvstore"
)

func TestKVStore_HasNoDataRace(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	wg.Add(1)
	store := kvstore.NewStore()
	go func() {
		for i := range 1000 {
			store.Set("foo", strconv.Itoa(i))
		}
		wg.Done()
	}()
	for range 1000 {
		store.Get("foo")
		runtime.Gosched()
	}
	wg.Wait()
}

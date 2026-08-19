package db

import (
	"sync"
	"testing"
)

func TestGetDB_SingleInit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDB()
		}()
	}
	wg.Wait()

	if got := InitCount(); got != 1 {
		t.Fatalf("expected connect() to run exactly once, ran %d times", got)
	}
}

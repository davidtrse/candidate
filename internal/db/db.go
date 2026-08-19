package db

import (
	"sync"
	"sync/atomic"
	"time"
)

// Conn is a stand-in for a real database connection/pool.
type Conn struct {
	ID int
}

var (
	instance  *Conn
	initCount int32
	once sync.Once
)

// GetDB returns the shared Conn, creating it on first use.
//
// TODO(candidate): under concurrent access this has a race condition —
// multiple goroutines can all see instance == nil and each create their own
// Conn. Fix it so connect() runs exactly once no matter how many goroutines
// call GetDB() concurrently. `go test -race ./internal/db/...` should pass
// once this is fixed.
func GetDB() *Conn {
	once.Do(func() {
		if instance == nil {
			instance = connect()
		}
	})

	return instance
}

func connect() *Conn {
	atomic.AddInt32(&initCount, 1)
	time.Sleep(time.Millisecond) // simulate connection setup latency
	return &Conn{ID: 1}
}

// InitCount reports how many times connect() actually ran. Used by tests.
func InitCount() int32 {
	return atomic.LoadInt32(&initCount)
}

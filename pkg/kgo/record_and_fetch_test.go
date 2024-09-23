package kgo

import (
	"log"
	"testing"
)

func TestPooling(t *testing.T) {
	pool := newRecordPool()

	r := pool.get()
	r.Value = []byte("foo bar baz")
	r.pool = pool

	processedValue := r.Value
	r.Close()

	r2 := pool.get()

	if "foo bar baz" != string(processedValue) {
		log.Fatal("not foo bar baz")
	}

	if r2 == nil {
		log.Fatal("r is nil")
	}

}

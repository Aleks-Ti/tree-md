package main

import "sync"

var (
	mu = sync.Mutex{}
	i  = 0
)

func safeInc() {
	mu.Lock()
	i++
	mu.Unlock()
}

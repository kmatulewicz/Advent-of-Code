package main

import "sync"

// safe int
type block struct {
	mu sync.Mutex
	i  int
}

// Add adds 1 to int
func (b *block) Add() {
	b.mu.Lock()
	b.i++
	b.mu.Unlock()
}

// Sub subtracts 1 from int
func (b *block) Sub() {
	b.mu.Lock()
	b.i--
	b.mu.Unlock()
}

// Check checks if int is equal to 0
func (b *block) Check() bool {
	b.mu.Lock()
	if b.i == 0 {
		return true
	}
	b.mu.Unlock()
	return false
}

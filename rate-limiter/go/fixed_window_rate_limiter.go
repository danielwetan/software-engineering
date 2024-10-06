package main

import (
	"sync"
	"time"
)

var _ RateLimiter = (*FixedWindowRateLimiter)(nil)

type FixedWindowRateLimiter struct {
	maxRequests      int
	windowSize       time.Duration
	requestCounts    map[string]int
	windowStartTimes map[string]time.Time
	mu               sync.Mutex
}

func (r *FixedWindowRateLimiter) AllowRequest(clientID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	currentTime := time.Now()
	if _, exists := r.windowStartTimes[clientID]; !exists {
		r.windowStartTimes[clientID] = currentTime
		r.requestCounts[clientID] = 0
	}

	windowStartTime := r.windowStartTimes[clientID]
	if currentTime.Sub(windowStartTime) >= r.windowSize {
		r.windowStartTimes[clientID] = currentTime
		r.requestCounts[clientID] = 0
	}

	if r.requestCounts[clientID] < r.maxRequests {
		r.requestCounts[clientID]++
		return true
	}

	return false
}

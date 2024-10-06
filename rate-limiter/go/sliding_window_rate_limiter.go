package main

import (
	"sync"
	"time"
)

var _ RateLimiter = (*SlidingWindowRateLimiter)(nil)

type SlidingWindowRateLimiter struct {
	maxRequests       int
	windowSize        time.Duration
	requestTimestamps map[string][]time.Time
	mu                sync.Mutex
}

func (r *SlidingWindowRateLimiter) AllowRequest(clientID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	currentTime := time.Now()
	if _, exists := r.requestTimestamps[clientID]; !exists {
		r.requestTimestamps[clientID] = []time.Time{}
	}

	timestamps := r.requestTimestamps[clientID]

	// Remove timestamps that are outside the window
	validTimestamp := []time.Time{}
	for _, ts := range timestamps {
		if currentTime.Sub(ts) <= r.windowSize {
			validTimestamp = append(validTimestamp, ts)
		}
	}

	// Update the timestamps with only valid ones
	r.requestTimestamps[clientID] = validTimestamp

	if len(validTimestamp) < r.maxRequests {
		r.requestTimestamps[clientID] = append(r.requestTimestamps[clientID], currentTime)
		return true
	}

	return false
}

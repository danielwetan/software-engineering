package main

import (
	"errors"
	"time"
)

type RateLimiterFactory struct{}

func (r *RateLimiterFactory) CreateRateLimiter(rateLimiterType string, maxRequests int, windowSize time.Duration) (RateLimiter, error) {
	switch rateLimiterType {
	case "fixed":
		return &FixedWindowRateLimiter{
			maxRequests:      maxRequests,
			windowSize:       windowSize,
			requestCounts:    make(map[string]int),
			windowStartTimes: make(map[string]time.Time),
		}, nil

	case "sliding":
		return &SlidingWindowRateLimiter{
			maxRequests:       maxRequests,
			windowSize:        windowSize,
			requestTimestamps: make(map[string][]time.Time),
		}, nil

	default:
		return nil, errors.New("unknown rate limiter type")
	}
}

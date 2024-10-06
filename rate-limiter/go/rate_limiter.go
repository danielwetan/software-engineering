package main

type RateLimiter interface {
	AllowRequest(clientID string) bool
}

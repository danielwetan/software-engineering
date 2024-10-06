package main

import (
	"fmt"
	"time"
)

func main() {
	factory := &RateLimiterFactory{}
	fixedWindowRateLimiter, err := factory.CreateRateLimiter("fixed", 2, 2*time.Second)
	if err != nil {
		fmt.Printf("Error creating sliding window rate limiter: %v\n", err)
		return
	}
	slidingWindowRateLimiter, err := factory.CreateRateLimiter("sliding", 2, 1*time.Minute)
	if err != nil {
		fmt.Printf("Error creating sliding window rate limiter: %v\n", err)
		return
	}
	fmt.Println("Fixed Window Rate Limiter:")
	for i := 0; i < 12; i++ {
		allowed := fixedWindowRateLimiter.AllowRequest("client1")
		fmt.Println(allowed)
	}

	fmt.Println("Sliding Window Rate Limiter:")
	for i := 0; i < 12; i++ {
		allowed := slidingWindowRateLimiter.AllowRequest("client2")
		fmt.Println(allowed)
	}
}

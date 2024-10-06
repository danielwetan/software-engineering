package main

import "errors"

var _ LoadBalancingStrategy = (*RoundRobinStrategy)(nil)

type RoundRobinStrategy struct {
	currentIndex int
}

func (r *RoundRobinStrategy) GetServer(servers []Server, request Request) (Server, error) {
	totalServers := len(servers)
	if totalServers == 0 {
		return Server{}, errors.New("no servers available")
	}

	server := servers[r.currentIndex]
	r.currentIndex = (r.currentIndex + 1) % totalServers

	return server, nil
}

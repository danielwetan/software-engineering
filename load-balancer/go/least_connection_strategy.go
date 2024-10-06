package main

import (
	"errors"
	"math"
)

var _ LoadBalancingStrategy = (*LeastConnectionStrategy)(nil)

type LeastConnectionStrategy struct{}

func (r *LeastConnectionStrategy) GetServer(servers []Server, request Request) (Server, error) {
	minConnections := math.MaxInt
	var selectedServer Server

	for _, server := range servers {
		if server.IsHealthy {
			connections := r.getConnections(server)

			if connections < minConnections {
				minConnections = connections
				selectedServer = server
			}
		}
	}

	if selectedServer.ServerID == "" {
		return Server{}, errors.New("no healthy servers available")
	}

	return selectedServer, nil
}

// Perform logic to get current connections for the server
func (r *LeastConnectionStrategy) getConnections(_ Server) int {
	return 0
}

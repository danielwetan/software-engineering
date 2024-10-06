package main

import "sync"

type LoadBalancer struct {
	servers  []Server
	strategy LoadBalancingStrategy
}

var instance *LoadBalancer
var once sync.Once

// GetInstance ensures a singleton instance of LoadBalancer.
func GetInstance() *LoadBalancer {
	once.Do(func() {
		instance = &LoadBalancer{
			servers: make([]Server, 0),
		}
	})
	return instance
}

func (lb *LoadBalancer) AddServer(server Server) {
	lb.servers = append(lb.servers, server)
}

func (lb *LoadBalancer) RemoveServer(server Server) {
	for i, s := range lb.servers {
		if s == server {
			lb.servers = append(lb.servers[:i], lb.servers[i+1:]...)
			break
		}
	}
}

func (lb *LoadBalancer) GetServer(request Request) (Server, error) {
	return lb.strategy.GetServer(lb.servers, request)
}

func (lb *LoadBalancer) SetLoadBalancingStrategy(strategy LoadBalancingStrategy) {
	lb.strategy = strategy
}

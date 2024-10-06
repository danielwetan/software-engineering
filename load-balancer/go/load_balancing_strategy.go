package main

type LoadBalancingStrategy interface {
	GetServer(servers []Server, request Request) (Server, error)
}

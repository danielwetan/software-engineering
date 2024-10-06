package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create servers
	server1 := Server{ServerID: "1", IsHealthy: true}
	server2 := Server{ServerID: "2", IsHealthy: true}

	// Create load balancer
	loadBalancer := GetInstance()
	loadBalancer.AddServer(server1)
	loadBalancer.AddServer(server2)

	// Set load balancing strategy
	roundRobinStrategy := RoundRobinStrategy{}
	loadBalancer.SetLoadBalancingStrategy(&roundRobinStrategy)

	// Create requests
	request1 := Request{Method: http.MethodPost, Path: "/login"}
	request2 := Request{Method: http.MethodGet, Path: "/profile"}
	request3 := Request{Method: http.MethodGet, Path: "/dashboard"}

	// Get server for each requests
	selectedServer1, _ := loadBalancer.GetServer(request1)
	fmt.Printf("server: %s | %s %s \n", selectedServer1.ServerID, request1.Method, request1.Path)

	selectedServer2, _ := loadBalancer.GetServer(request2)
	fmt.Printf("server: %s | %s %s \n", selectedServer2.ServerID, request1.Method, request2.Path)

	selectedServer3, _ := loadBalancer.GetServer(request3)
	fmt.Printf("server: %s | %s %s \n", selectedServer3.ServerID, request1.Method, request3.Path)
}

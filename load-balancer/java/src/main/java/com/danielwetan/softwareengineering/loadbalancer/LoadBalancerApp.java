package com.danielwetan.softwareengineering.loadbalancer;

public class LoadBalancerApp {
  public static void main(String[] args) {
    // Create servers
    Server server1 = new Server("server-a");
    Server server2 = new Server("server-b");

    // Create load balancer
    LoadBalancer loadBalancer = LoadBalancer.getInstance();
    loadBalancer.addServer(server1);
    loadBalancer.addServer(server2);

    // Set load balancing strategy
    LoadBalancingStrategy roundRobinStrategy = new RoundRobinStrategy();
    loadBalancer.setLoadBalancingStrategy(roundRobinStrategy);

    // Create requests 
    Request request1 = new Request("POST", "/login");
    Request request2 = new Request("GET", "/profile");
    Request request3 = new Request("GET", "/dashboard");

    // Get server for request 1
    Server selectedServer1 = loadBalancer.getServer(request1);
    System.out.println("Selected server for request1: " + selectedServer1.getServerId());

    Server selectedServer2 = loadBalancer.getServer(request2);
    System.out.println("Selected server for request2: " + selectedServer2.getServerId());

    Server selectedServer3 = loadBalancer.getServer(request3);
    System.out.println("Selected server for request3: " + selectedServer3.getServerId());
  }  
}

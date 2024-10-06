package com.danielwetan.softwareengineering.loadbalancer;

import java.util.ArrayList;
import java.util.List;

public class LoadBalancer {
  private static LoadBalancer instance;
  private List<Server> servers;
  private LoadBalancingStrategy strategy;

  private LoadBalancer() {
    this.servers = new ArrayList<>();
  }

  public static LoadBalancer getInstance() {
    if (instance == null) {
      instance = new LoadBalancer();
    }

    return instance;
  }

  public void addServer(Server server) {
    servers.add(server);
  } 

  public void removeServer(Server server) {
    servers.remove(server);
  }

  public Server getServer(Request request) {
    return strategy.getServer(servers, request);
  }

  public void setLoadBalancingStrategy(LoadBalancingStrategy strategy) {
    this.strategy = strategy;
  }
}
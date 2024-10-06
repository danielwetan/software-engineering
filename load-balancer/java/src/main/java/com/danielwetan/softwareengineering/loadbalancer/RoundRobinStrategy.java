package com.danielwetan.softwareengineering.loadbalancer;

import java.util.List;

public class RoundRobinStrategy implements LoadBalancingStrategy {
  private int currentIndex;

  public RoundRobinStrategy() {
      this.currentIndex = 0;
  }

  @Override
  public Server getServer(List<Server> servers, Request request) {
    int totalServers = servers.size();
    if (totalServers == 0) {
      throw new IllegalStateException("No servers available");
    }

    Server server = servers.get(currentIndex);
    currentIndex = (currentIndex + 1) % totalServers;
    return server;
  }
}

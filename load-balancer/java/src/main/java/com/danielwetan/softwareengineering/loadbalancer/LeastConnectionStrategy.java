package com.danielwetan.softwareengineering.loadbalancer;

import java.util.List;

public class LeastConnectionStrategy implements LoadBalancingStrategy {
  @Override
  public Server getServer(List<Server> servers, Request request) {
    int minConnections = Integer.MAX_VALUE;
    Server selectedServer = null;

    for (Server server : servers) {
      if (server.isHealthy()) {
        int connections = getConnections(server);
        if (connections < minConnections) {
          minConnections = connections;
          selectedServer = server;
        }
      }
    }
    return selectedServer;
  }

  private int getConnections(Server server) {
    // Perform logic to get current connections for the server
    return 0; // Placeholder for connections count
  }
}

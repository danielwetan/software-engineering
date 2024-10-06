package com.danielwetan.softwareengineering.loadbalancer;

class Server {
  private String serverId;
  private boolean isHealthy;

  public Server(String serverId) {
    this.serverId = serverId;
    this.isHealthy = true;
  }

  public String getServerId() {
    return this.serverId;
  }

  public boolean isHealthy() {
    return this.isHealthy;
  }

  public void setHealth(boolean healthy) {
    this.isHealthy = healthy;
  }
}


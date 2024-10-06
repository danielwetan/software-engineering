package com.danielwetan.softwareengineering.loadbalancer;

// Mock user request
public class Request {
  private String method;
  private String path;

  public Request(String method, String path) {
    this.method = method;
    this.path = path;
  }

  public String getMethod() {
    return this.method;
  }

  public String getPath() {
    return this.path;
  }
}

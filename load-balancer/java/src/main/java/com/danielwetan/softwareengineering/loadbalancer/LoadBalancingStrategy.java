package com.danielwetan.softwareengineering.loadbalancer;

import java.util.List;

interface LoadBalancingStrategy {
  Server getServer(List<Server> servers, Request request);
}
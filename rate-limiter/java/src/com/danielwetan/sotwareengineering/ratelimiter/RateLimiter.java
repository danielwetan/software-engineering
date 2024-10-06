package com.danielwetan.sotwareengineering.ratelimiter;

public interface RateLimiter {
  boolean allowRequest(String clientId);
}

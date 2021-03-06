# go-micro-auth
[![Build Status](https://travis-ci.org/antklim/go-micro-auth.svg?branch=master)](https://travis-ci.org/antklim/go-micro-auth)

Authentication microservice in Go lang

# Service configuration
Authentication microservice is using `Consul` as a service discovery. `Consul` is also a default configuration source.

Command line options to define configuration:
- `-config [consul|file]` or `-config=[consul|file]`
- `-config_path path/to/service.cfg` or `-config_path=path/to/service.cfg` (used only in case of file config)

## Configuration on Consul
There are the following keys used:
- auth/config/jwssecret - signing key value for JWT(JWS) signing
- auth/config/jwtttl - default TTL of the token

## File configuration
There are the following keys used:
- jwssecret - signing key value for JWT(JWS) signing
- jwtttl - default TTL of the token

Please check `service.cfg` for config file example

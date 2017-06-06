# go-micro-auth
Authentication microservice in Go lang

## Consul
Authentication microservice is using `Consul` as a service discovery and KV store.
There are the following keys used:
- auth/config/jwssecret - signing key value for JWT(JWS) signing
- auth/config/jwtttl - default TTL of the token

# Fibonacci

Fibonacci is a service that calculates a numerical interval from x to y inclusive from the Fibonacci series. The service does not recalculate the numbers, but takes them from the cache if they are present there.

## Delivery
Service communicates using two protocols: gRPC and HTTP. Service considers all HTTP2 traffic is gRPC requests.

#### HTTP
```
GET /fibonacci?x=1&y=2
```
`x` and `y` are non negative integers greater than 1. `x` is less than or equal `y`

#### gRPC
Proto file can be found [here](https://github.com/Eretic431/fibonacci/blob/master/internal/fibonacci/proto/fibonacci.proto).
You can also compile gRPC Client to test out gRPC work:
```sh
make build-client
```
Then run:
```sh
./grpc_client -x=1 -y=5
```

## Deployment
To start the service simply run:
```sh
make up
```
Application listens on `:8080` and Redis listens on `:6379` by default so these ports must not be taken by another processes.
In order to change ports simply modify `docker-compose.yml` file.

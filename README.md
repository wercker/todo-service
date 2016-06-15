# todo-service

Todo Service is a middleware microservice that is in charge of storing
todos.

It uses [grpc](http://grpc.io) as the underlying communications
protocol.

Todos get stored either in memory (the default) or alternative storage
mechanisms (yet to come).

By default it listens on TCP port 3000

Todo-service lives on Kubernetes

## usage:

```
./todo-service start [--service-port XXXX] start
```



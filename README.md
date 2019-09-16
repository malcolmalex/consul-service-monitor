### Consul Service Monitor

A simple program to monitor the health check endpoint for services in consul.

Usage: 

```
go build consul-service-monitor
./consul-service-monitor -server=<hostname> -port=<consul API port> <list of services, separated by spaces>
```

## Planned updates

- continually loop
- notifications
- Example using the GO consul library

# GoBlockHub
A modular Go API server for interacting with multiple blockchain platforms.  
Currently supports **Binance** and **OKX**, with easy scalability for future platforms.


## Features
- Modular router with platform-specific handlers
- DI (Dependency Injection) pattern for Service → Handler → Response
- Standardized API responses: success code = 1, error codes starting from 10001
- Zap logger integration with levels (info, debug, error) and date-based log files
- Automatic logging for all API responses
- Graceful server shutdown, including in-progress requests
- Scalable handler registration via router registry
- Example routes for testing: `/ping` and `/slow`
- Ready for integration with real blockchain APIs
- Independent unit tests with mock servers for Binance and OKX


## Installation
Clone the repository and install dependencies
``` bash
git clone https://github.com/CharlesWhiteSun/goblockhub.git
cd goblockhub
go mod tidy
```


## Running the Server
Start the server on http://localhost:8080
``` bash
go run .\main.go
```


## API Examples

### Ping test
Response: pong
``` bash
curl http://localhost:8080/ping
```


### Slow test
Response: finished slow API
``` bash
curl http://localhost:8080/slow
```


### Binance status
Response: {"status":"Binance OK"}
``` bash
curl http://localhost:8080/api/binance/status
```


### OKX status
Response: {"status":"OKX OK"}
```bash
curl http://localhost:8080/api/okx/status
```


## How to Extend
1. Implement a new Service implementing `IPlatformService`
2. Implement a new Handler implementing `IPlatformHandler`
3. Register the handler in `router/registry.go`
4. Router automatically registers all handlers on server start


## CI / Testing
- Cross-platform GitHub Actions CI workflow for Windows
- Builds, runs tests, and performs `golangci-lint`
- Unit tests for handlers using mock servers, including graceful shutdown scenarios


## License
This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

# GoBlockHub

A modular Go API server for interacting with multiple blockchain platforms.  
Currently supports **Binance** and **OKX**, with easy scalability for future platforms.

---

## Features

- Modular router with platform-specific handlers
- DI (Dependency Injection) for Service → Handler
- Graceful shutdown of server
- Scalable handler registration via router registry
- Example routes for testing: `/ping` and `/slow`
- Ready for integration with real blockchain APIs

---

## Project Structure

```
goblockhub/
├─ main.go                 # Server entry point
├─ go.mod / go.sum
├─ internal/
│  ├─ server/              # Gin server wrapper with graceful shutdown
│  │  └─ server.go
│  ├─ router/              # Router setup and handler registry
│  │  ├─ router.go
│  │  └─ registry.go
│  ├─ handler/             # Platform handlers
│  │  ├─ handler.go        # Handler interface
│  │  ├─ binance.go
│  │  └─ okx.go
│  └─ service/             # Platform services
│     ├─ service.go        # Service interface
│     ├─ binance.go
│     └─ okx.go
```

---

## Installation

```bash
git clone https://github.com/CharlesWhiteSun/goblockhub.git
cd goblockhub
go mod tidy
```

---

## Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

---

## API Examples

### Ping test

```bash
curl http://localhost:8080/ping
# Response: pong
```

### Slow test

```bash
curl http://localhost:8080/slow
# Response: finished slow API
```

### Binance status

```bash
curl http://localhost:8080/api/binance/status
# Response: {"status":"Binance OK"}
```

### OKX status

```bash
curl http://localhost:8080/api/okx/status
# Response: {"status":"OKX OK"}
```

---

## How to Extend

1. Add a new Service implementing `IPlatformService`
2. Add a new Handler implementing `IPlatformHandler`
3. Register the handler in `router/registry.go`
4. Done! Router automatically registers all handlers on server start

---

## License
This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

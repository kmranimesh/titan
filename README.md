# Titan 

**Titan** is a high-performance, distributed task queue written in Go. 

It is designed to be a lightweight but durable alternative to systems like Redis/Celery, built from scratch for educational purposes and real-world utility.

## Features (Planned)
- **Distributed**: Support for multiple producers and workers.
- **Durable**: Custom Write-Ahead Log (WAL) for data persistence.
- **Fast**: gRPC-based communication protocol.
- **Reliable**: At-least-once delivery guarantees with ACK/NACK support.

## Project Structure
- `cmd/`: Entry points for executables (the Broker).
- `pkg/`: Public libraries (Client SDKs).
- `internal/`: Private application logic (Queue, Storage).

## Getting Started
```bash
go mod tidy
```

## License
MIT (See [LICENSE](LICENSE))

# Flussonic Go SDK

A Go SDK for interacting with Flussonic Media Server, Flussonic Watcher, Flussonic Central and Flussonic Vision APIs.

## Installation

```bash
go get github.com/flussonic/go-flussonic
```

## Overview

This SDK provides a simple and idiomatic Go interface for managing Flussonic products. The library includes public API clients for:

- **Flussonic Media Server** - Video streaming server management
- **Flussonic Central** - Centralized management of multiple Flussonic instances
- **Flussonic Vision Inference** - Video analytics and episode processing
- **Flussonic Vision Identification** - Object and person identification
- **Flussonic Watcher Admin** - Administrative API for Watcher platform
- **Flussonic Watcher Client** - Client API for Watcher platform

Examples for the SDK usage is located in Examples directory.

## Features

### Flussonic Media Server (`flussonic`)

- Stream management (create, update, delete, list)
- DVR archive management and querying
- Session monitoring and management
- Package management
- Server configuration and health checks
- SSL/TLS certificate management (Let's Encrypt)
- Real-time statistics and metrics
- Pagination support with iterators for large datasets

### Flussonic Central (`central`)

- Agent management (list, get, configure)
- Stream configuration across multiple instances
- Central configuration management
- Pagination support for agents and streams

### Flussonic Vision Inference (`vision-inference`)

- Episode management and listing
- Stream configuration for analytics
- Counter records and metrics
- Worker statistics
- Episode filtering and search
- Iterator support for large episode lists

### Flussonic Vision Identification (`vision-identification`)

- Person identification management
- Object recognition capabilities
- Vision stream configuration

### Flussonic Watcher Admin (`watcher-admin`)

- Administrative stream management
- Agent management and configuration
- Organization and user management
- Stream zone configuration

### Flussonic Watcher Client (`watcher-client`)

- Client-side stream access
- Attendance tracking (persons and vehicles)
- Agent activation token management
- Organization-based stream access

## Key Features

- **Type-safe API** - Comprehensive Go types generated from API schemas
- **Pagination support** - Built-in iterators for handling large datasets
- **Error handling** - Detailed error messages with status codes
- **Context support** - Full support for Go contexts (timeouts, cancellation)
- **Flexible configuration** - URL-based or struct-based configuration
- **Production ready** - Retry logic, proper error handling, comprehensive tests

## License

See [LICENSE](LICENSE) file for details.

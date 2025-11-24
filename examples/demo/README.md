# Demo Example

This example demonstrates how to use Flussonic Go SDK to interact with Central API.

## What it does

The demo application performs three main tasks:

1. **Creates streamers** in Central
2. **Reads `streams.csv`** and creates streams in Central
3. **Streams episodes** with type `qr_code` and outputs their payload

## Customization

To add custom logic for processing episodes, modify the `episodesCallback` function in `main.go`. This function is called for each episode with type `qr_code` and receives the episode data. You can add your own processing logic there, such as:

- Storing episodes to a database
- Sending notifications
- Processing payload data
- Integrating with external services

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Flussonic license key

### Run

```bash
cd examples/demo
LICENSE_KEY='your_license_key' make up
```

### Stop

```bash
make down
```

## View Logs

```bash
docker-compose logs -f demo-app
```

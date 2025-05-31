# Cerify - Smart Contract Analysis & Verification Platform

Cerify is a cutting-edge blockchain platform for automated analysis, verification, and certification of Ethereum smart contracts. The platform provides comprehensive security analysis, verification of contract properties, and certification of contract behavior.

## Features

- Smart contract static analysis
- Security vulnerability detection
- Contract property verification
- Automated certification
- RESTful API for integration
- Real-time analysis status tracking

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- MongoDB 6.0 or later
- Redis 7.0 or later

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/cerify/cerify-infra.git
cd cerify-infra
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Update the environment variables in `.env` with your configuration.

4. Run the application using Docker Compose:
```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`.

## API Documentation

### Endpoints

#### Contract Analysis
- `POST /api/v1/contracts/analyze` - Submit a contract for analysis
- `GET /api/v1/contracts/:address` - Get contract information
- `GET /api/v1/contracts/:address/analysis` - Get analysis results

#### Contract Verification
- `POST /api/v1/contracts/verify` - Submit a contract for verification
- `GET /api/v1/contracts/:address/verification` - Get verification status

#### Analysis
- `POST /api/v1/analysis/submit` - Submit new analysis
- `GET /api/v1/analysis/status/:id` - Get analysis status
- `GET /api/v1/analysis/report/:id` - Get analysis report

#### Verification
- `POST /api/v1/verification/submit` - Submit new verification
- `GET /api/v1/verification/status/:id` - Get verification status
- `GET /api/v1/verification/certificate/:id` - Get verification certificate

## Development

### Project Structure

```
cerify-infra/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── app/
│   ├── config/
│   ├── handlers/
│   └── services/
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
```

### Running Tests

```bash
go test ./...
```

### Building Locally

```bash
go build -o cerify-server ./cmd/server
```

## Contributing

Please read our contributing guidelines before submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

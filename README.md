# ArbForge gRPC Service

Hybrid Node.js/Go gRPC service for ArbForge deal validation, confidence scoring, and Monica AI integration.

## 🚀 Quick Start

### Using Docker Compose
```bash
docker-compose up
```

### Manual Setup

**Go gRPC Server:**
```bash
go mod download
buf generate
go run ./server
```

**Node.js Bridge:**
```bash
npm install
npm start
```

## 📡 API Endpoints

### POST /api/test-grpc

Test gRPC calls from Node.js.

**Actions:**
- `validate` — Validate a deal
- `monica` — Get Monica AI advice
- `confidence` — Calculate confidence score
- `generate` — Generate listing copy

**Example:**
```bash
curl -X POST http://localhost:3000/api/test-grpc \
  -H "Content-Type: application/json" \
  -d '{
    "action": "validate",
    "productName": "2026 Hot Wheels Drift Ender",
    "buyPrice": 1.99
  }'
```

## 📂 Project Structure

```
.
├── proto/
│   └── arbforge.proto          # gRPC service definition
├── server/
│   ├── main.go                 # gRPC server entry
│   └── grpc_server.go          # Service implementation
├── services/
│   └── grpcClient.js           # Node.js gRPC client
├── routes/
│   └── test-grpc.js            # Express test route
├── app.js                      # Express server
├── Dockerfile                  # Go service image
├── docker-compose.yml          # Multi-service orchestration
├── buf.yaml                    # Buf lint config
├── buf.gen.yaml                # Buf code gen config
├── go.mod                      # Go dependencies
└── package.json                # Node dependencies
```

## 🔧 Development

**Generate gRPC stubs:**
```bash
buf generate
```

**Format Go code:**
```bash
gofmt -s -w ./server
```

**Run Node.js in dev mode:**
```bash
npm run dev
```

## 📋 Service Methods

### ValidateDeal
Validates a deal based on product name and buy price.
- **Params:** `product_name` (string), `buy_price` (double)
- **Returns:** `valid` (bool), `message` (string), `estimated_roi` (double), `status` (string)

### AskMonica
GetsMonica AI advice on flip opportunities.
- **Params:** `prompt` (string)
- **Returns:** `answer` (string), `tone` (string)

### GetConfidenceScore
Calculates confidence score based on margins.
- **Params:** `buy_price` (double), `sell_price` (double)
- **Returns:** `score` (double), `label` (string), `summary` (string)

### GenerateListing
Generates listing copy for a deal.
- **Params:** `product_name` (string), `buy_price` (double), `sell_price` (double)
- **Returns:** `listing_title` (string), `listing_body` (string), `tags` (string)

## 🐳 Docker

**Build and run gRPC server:**
```bash
docker build -t arbforge-grpc .
docker run -p 50051:50051 arbforge-grpc
```

**Run full stack:**
```bash
docker-compose up --build
```

## 📝 License

MIT

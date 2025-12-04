# Automation Catalog Backend

### Endpoints
GET  /health
GET  /automations
POST /demo/{id}

### Run local
go run cmd/server/main.go

### Deploy on Railway
1. Create new service → Deploy from GitHub
2. Add Environment Variables
3. Expose PORT=8080

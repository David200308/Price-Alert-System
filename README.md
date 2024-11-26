# Go API Backend

## 1 Tech Stack

- Programming Language: Go Language
- Web Framework: Gin
- Database: PostgreSQL
- Cache Database: Redis
- Message Queue: RabbitMQ
- Container: Docker
- Document: Swagger
- Scheduler: Linux Cron
- Email Service: Mailgun (https://www.mailgun.com)
- Payment Service: Stripe (https://stripe.com)
- Stock Market Data Service: Alpha Vantage (https://www.alphavantage.co)
- Crypto Market Data Service: API-Ninjas (https://api-ninjas.com)
- Error Tracking: Sentry (https://sentry.io)

## 2 Program Build & Usage

```bash
## Init database (Make sure the postgreSQL is prepared)
cd Backend
go run ./migrate/migrate.go

## Generate Swagget Docs
go install github.com/swaggo/swag/cmd/swag@latest
$HOME/go/bin/swag init

## Running Backend
cd Backend
go build
./Backend
```

## 3 ES256 JWT Signing Key Pair Generate

Generate the ECDSA key pairs with prime256v1 curve

```
openssl ecparam -name prime256v1 -genkey -noout -out private_key.pem
openssl ec -in private_key.pem -pubout -out public_key.pem
```

## 4 AES Key for Encryption Generate

```
head /dev/urandom | sha256sum
```

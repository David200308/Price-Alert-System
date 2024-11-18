# Go API Backend

## 1 Tech Stack

- Programming Language: GoLang
- Web Framework: Gin
- Database: PostgreSQL
- Cache Database: Redis
- Message Queue: RabbitMQ
- Container: Docker
- Email Service: Mailgun (https://www.mailgun.com/)
- Document: Swagger
- Error Tracking: Sentry (https://sentry.io/)

## 2 Program Build & Usage

```bash
## Running Backend
cd Backend
go build
./Backend

## Generate Swagget Docs
go install github.com/swaggo/swag/cmd/swag@latest
$HOME/go/bin/swag init
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

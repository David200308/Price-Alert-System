FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Backend .

FROM alpine:latest  

RUN adduser -D backend

WORKDIR /root/

COPY --from=builder /app/Backend .

RUN chown backend:backend Backend

USER backend

EXPOSE 3010

CMD ["./Backend"]

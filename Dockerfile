FROM golang:1.24.5-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY db/migrations ./db/migrations

RUN go install github.com/google/wire/cmd/wire@latest
RUN go generate ./...

RUN go build -o main ./cmd

FROM alpine:3.21.3

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser

COPY --from=builder /app/main .
COPY --from=builder /app/db/migrations ./db/migrations

RUN chown -R appuser:appgroup /home/appuser

USER appuser

EXPOSE 8080

CMD ["./main"]

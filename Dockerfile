FROM golang:1.25-alpine AS builder

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY db/migrations ./db/migrations

EXPOSE 3000

CMD [ "/app/main" ]
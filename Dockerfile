FROM golang:1.25.4-alpine AS builder

WORKDIR /app

ENV GOTOOLCHAIN=auto

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=1 go build -o pedalea ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache sqlite-libs

COPY --from=builder /app/pedalea .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app

USER appuser
EXPOSE 8080
CMD ["/app/pedalea"]

# -------- Build Stage --------
    FROM golang:1.24-alpine AS builder

    WORKDIR /app
    
    RUN apk add --no-cache git
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    # Binary name change -> cryptoapiservice
    RUN go build -ldflags="-s -w" -o test-aws-go ./cmd/main.go

    CMD ["./test-aws-go", "serve"]
    
    
    # -------- Run Stage --------
    # FROM alpine:latest
    
    # RUN apk --no-cache add ca-certificates curl \
    #     && addgroup -S appgroup && adduser -S appuser -G appgroup
    
    # WORKDIR /app
    
    # # Copy binary
    # COPY --from=builder /app/test-aws-go .
    
    # # Copy ABI folder
    # # COPY --from=builder /app/internal/provider/evm ./abi
    
    # # (optional) copy migrations/config if required later
    # # COPY --from=builder /app/migrations ./migrations
    # # COPY --from=builder /app/config ./config
    
    # RUN chown -R appuser:appgroup /app
    
    # USER appuser
    
    # EXPOSE 7322
    
    # # Run solveronboarding service
    # CMD ["./test-aws-go", "serve"]
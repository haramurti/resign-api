# --- STAGE 1: Build Stage ---
FROM golang:1.25-alpine AS builder

# Install git & ca-certificates to connect to supabase
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy dependency files 
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# translating to binary
# CGO_ENABLED=0 so the binary can run Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# --- STAGE 2: Final Image ---
FROM alpine:latest

# Install timezone data & certs
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# get result from first tage
COPY --from=builder /app/main .

# Copy public folder
COPY --from=builder /app/public ./public

# Expose port 
EXPOSE 8686

# runn app
CMD ["./main"]
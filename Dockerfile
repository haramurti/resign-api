# --- STAGE 1: Build Stage ---
FROM golang:1.25-alpine AS builder

# Install git & ca-certificates buat koneksi HTTPS ke Supabase
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy dependency files duluan biar build-nya cepet (cached)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build aplikasi jadi binary "main"
# CGO_ENABLED=0 penting biar binary-nya statis & jalan di Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# --- STAGE 2: Final Image ---
FROM alpine:latest

# Install timezone data & certs
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Ambil hasil build dari Stage 1
COPY --from=builder /app/main .

# Copy folder public (Frontend lo wajib ikut!)
COPY --from=builder /app/public ./public

# Expose port sesuai settingan Go lo
EXPOSE 8686

# Jalankan aplikasi
CMD ["./main"]
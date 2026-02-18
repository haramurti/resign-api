# 🚀 Resign & Leave Management API - POROS 2026

Aplikasi manajemen cuti dan pengunduran diri karyawan berbasis Cloud-Native.

## 🏛️ Arsitektur Aplikasi
Aplikasi ini menggunakan **Micro-services Architecture** sederhana yang dikemas dalam kontainer:
- **Backend (Go-Fiber):** Menangani logic bisnis, auth, dan pengolahan data.
- **Frontend (Nginx):** Menyediakan antarmuka pengguna secara statis dan efisien.
- **Database (Supabase/Postgres):** Sebagai storage utama yang diakses secara cloud.
- **Orchestration (Kubernetes):** Mengatur siklus hidup container (Auto-healing & Service Discovery).

## 🛠️ Alur CI/CD
Pipeline otomatis menggunakan **GitHub Actions** dengan alur sebagai berikut:
1. **Push Trigger:** Pipeline aktif saat kode masuk ke branch `master`.
2. **Build Stage:** Kompilasi source code untuk verifikasi integritas kode.
3. **Registry Push:** Image Docker di-push ke Docker Hub dengan tag `latest`.
4. **Local Deployment:** Self-hosted runner pada MacBook M4 memicu `kubectl rollout restart` untuk memperbarui pod di cluster lokal (Orbstack) secara otomatis.

## 📂 Struktur Repository
```text
.
├── cmd/
│   └── main.go             # Entry point aplikasi (Inisialisasi server & database)
├── internal/
│   ├── database/           # Konfigurasi koneksi ke Supabase PostgreSQL
│   ├── domain/             # Enterprise business rules (Struct/Entity data)
│   ├── handler/            # Interface layer (HTTP/API handlers & Middleware)
│   ├── repository/         # Data access layer (Query database via GORM)
│   └── usecase/            # Application business logic (Alur proses utama)
├── public/
│   └── index.html          # Frontend sederhana untuk antarmuka pengguna
├── Dockerfile              # Konfigurasi container untuk Backend Go
├── Dockerfile.frontend     # Konfigurasi container untuk Frontend Nginx
├── k8s-configs.yaml        # Manifest Kubernetes: ConfigMaps & Secrets
├── k8s-deployment.yaml     # Manifest Kubernetes: Deployment & Service Backend
└── k8s-frontend.yaml       # Manifest Kubernetes: Deployment & Service Frontend
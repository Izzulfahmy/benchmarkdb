# 🚀 Benchmark: Golang vs Node.js vs Laravel dengan PostgreSQL

Repositori ini berisi percobaan **benchmark performa backend** menggunakan **PostgreSQL** dengan tiga framework berbeda:

* ⚡ **Golang (Gin)**
* ⚖️ **Node.js (Express)**
* 🐢 **Laravel (PHP)**

Pengujian dilakukan menggunakan **wrk** untuk mengukur *Requests per Second (Req/Sec)*, *Latency*, dan *Throughput*.

---

## 📦 Arsitektur

```
PostgreSQL (benchmarkdb)
   │
   ├── backend-golang  (Gin, port 8080)
   ├── backend-node    (Express, port 8081)
   └── backend-laravel (Laravel, port 8000)
```

---

## ⚡ Quick Start (1x Copy-Paste)

👉 Cukup salin & tempel blok di bawah ini di terminal (MacOS/Linux) untuk menjalankan semua backend sekaligus:

```bash
# === Clone Repo ===
git clone https://github.com/Izzulfahmy/benchmarkdb.git
cd benchmarkdb

# === Setup Database PostgreSQL ===
psql -U postgres <<EOF
CREATE DATABASE benchmarkdb;
\c benchmarkdb;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100)
);
INSERT INTO users (name, email)
SELECT 'User' || i, 'user' || i || '@mail.com'
FROM generate_series(1,1000) AS s(i);
EOF

# === Jalankan Golang Backend ===
cd backend-golang
go mod tidy
(go run main.go &)   # running in background (port 8080)
cd ..

# === Jalankan Node.js Backend ===
cd backend-node
npm install
(node index.js &)    # running in background (port 8081)
cd ..

# === Jalankan Laravel Backend ===
cd backend-laravel
composer install
php artisan migrate:fresh --seed
(php artisan serve --host=127.0.0.1 --port=8000 &)   # running (port 8000)
cd ..

# === Install wrk (MacOS) ===
brew install wrk

# === Jalankan Benchmark (30 detik) ===
wrk -t12 -c400 -d30s http://127.0.0.1:8080/users   # Golang
wrk -t12 -c400 -d30s http://127.0.0.1:8081/users   # Node.js
wrk -t12 -c400 -d30s http://127.0.0.1:8000/users   # Laravel
```

---

## 📖 Penjelasan Langkah

1. **Database Setup**
   Membuat database `benchmarkdb` dan mengisi tabel `users` dengan 1000 dummy data menggunakan `generate_series`.

2. **Backend Golang (Gin)**

   * File utama: `main.go`
   * Endpoint: `GET /users`

3. **Backend Node.js (Express)**

   * File utama: `index.js`
   * Endpoint: `GET /users`

4. **Backend Laravel (PHP)**

   * Framework Laravel standar
   * Endpoint: `GET /users` via `routes/api.php`

5. **Benchmark dengan wrk**

   * `-t12` → 12 thread
   * `-c400` → 400 koneksi simultan
   * `-d30s` → 30 detik uji beban

---

## 📊 Hasil Benchmark

| Backend     | Req/Sec    | Avg Latency | Requests Total | Transfer/sec |
| ----------- | ---------- | ----------- | -------------- | ------------ |
| **Golang**  | \~2.200    | 100–110 ms  | 130k+          | \~1.3 MB/s   |
| **Node.js** | \~700–1200 | 250–360 ms  | 40k+           | \~800 KB/s   |
| **Laravel** | \~27       | >1s         | 1.6k           | \~29 KB/s    |

---

## 📈 Visualisasi

```
Requests per Second (lebih tinggi lebih baik)

Golang   ██████████████████████████████████████  ~2200
Node.js  ██████████                            ~1000
Laravel  █                                     ~27
```

---

## 🎯 Kesimpulan

* **Golang (Gin)** → 🚀 Paling cepat, efisien, stabil.
* **Node.js (Express)** → ⚖️ Performa menengah, mudah digunakan.
* **Laravel (PHP)** → 🐢 Lambat jika pakai dev server (`artisan serve`),
  disarankan **Nginx + PHP-FPM** untuk production.

👉 Dari hasil ini, **Golang** lebih unggul untuk aplikasi dengan kebutuhan **high-concurrency** seperti **CBT, real-time API, dan sistem skala besar**.

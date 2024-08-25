# Backend API untuk Culinary Review

API ini menyediakan fitur CRUD untuk resep, ulasan, dan tag; autentikasi pengguna dengan JWT; penyimpanan gambar menggunakan Cloudinary; dan manajemen tag serta ulasan. Backend dibangun dengan prinsip Clean Architecture, menggunakan teknologi Go, Gin Framework, PostgreSQL, GORM, JWT, dan Cloudinary untuk mendukung skalabilitas dan maintainability. Dokumentasi API tersedia di endpoint Swagger setelah server dijalankan.

## Fitur Utama

- **CRUD untuk Resep, Ulasan, dan Tag**: Mengelola data resep, ulasan, dan tag.
- **Autentikasi Pengguna dengan JWT**: Mengamankan akses dengan JSON Web Token.
- **Penyimpanan Gambar Menggunakan Cloudinary**: Mengelola gambar resep dengan optimasi otomatis.
- **Manajemen Tag dan Ulasan**: Menandai resep dengan kategori dan mengelola ulasan pengguna.
- **Clean Architecture**: Struktur kode yang memudahkan pemeliharaan dan skalabilitas.

## Teknologi yang Digunakan

- **Go**: Bahasa pemrograman untuk backend.
- **Gin Framework**: Framework web untuk Go.
- **PostgreSQL**: Database relasional.
- **GORM**: ORM untuk Go.
- **JWT**: Untuk autentikasi pengguna.
- **Cloudinary**: Untuk penyimpanan dan pengelolaan gambar.

## Dokumentasi API

Dokumentasi API dapat diakses melalui Swagger setelah server dijalankan di endpoint `/swagger`.

## Instalasi

1. Clone repository ini:

   ```bash
   git clone <url-repository>
   ```

2. Masuk ke direktori proyek:

   ```bash
   cd <nama-direktori>
   ```

3. Instal dependensi:

   ```bash
   go mod tidy
   ```

4. Jalankan server:

   ```bash
   go run main.go
   ```

5. Akses dokumentasi API di `/swagger`.

## Kontribusi

Jika Anda ingin berkontribusi, silakan buat pull request atau buka isu dengan deskripsi masalah yang ditemukan.

## Lisensi

Tersedia di file [LICENSE](LICENSE).

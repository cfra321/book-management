# Book Management API

Repository ini berisi API untuk mengelola data buku dan kategori yang dibangun menggunakan **Golang** dan **Gin Gonic**.

## Fitur

- **Buku**
  - Menambahkan buku baru
  - Memperbarui buku berdasarkan ID
  - Menghapus buku berdasarkan ID
  - Mendapatkan daftar semua buku
  - Mendapatkan detail buku berdasarkan ID
- **Kategori**
  - Menambahkan kategori baru
  - Mendapatkan daftar semua kategori
  - Mendapatkan detail kategori berdasarkan ID


## Struktur Proyek

```bash
.
├── main.go          # Entry point aplikasi
├── controllers/     # Handler untuk endpoints
│   └── book.go      # Logika untuk resource buku
├── models/          # Struktur data dan model database
│   └── book.go      # Model untuk data buku
├── repositories/    # Interaksi dengan database
│   └── book_repo.go # Implementasi repository untuk buku
└── services/        # Logika bisnis aplikasi
    └── book_service.go # Layanan untuk memproses logika buku

```

Aplikasi akan berjalan di http://localhost:8080.

```bash
- GET (/api/books)
    Mendapatkan Daftar Semua Buku.
- GET (/api/books/:id)
    Mendapatkan Detail Buku Berdasarkan ID
- POST (/api/books)
    Menambahkan Buku Baru
- PUT (/api/books/:id)
    Memperbarui Buku Berdasarkan ID
- DELETE (/api/books/:id)
    Menghapus Buku Berdasarkan ID

- GET (/api/categories)
    Mendapatkan Daftar Semua Categories.
- GET (/api/categories/:id)
    Mendapatkan Detail Categories Berdasarkan ID.
- POST (/api/categories)
    Menambahkan Categories Baru.
- PUT (/api/categories/:id)
    Memperbarui Categories Berdasarkan ID.
- DELETE (/api/categories/:id)
    Menghapus Categories Berdasarkan ID.
```
## Teknologi yang Digunakan

- **Go**: Bahasa pemrograman untuk logika backend.
- **Gin Gonic**: Web framework untuk membangun API.
- **MySQL**: Basis data untuk menyimpan data buku dan kategori.

## Persiapan Sebelum Menjalankan

1. Pastikan **Go** sudah terinstall. Anda bisa mendownloadnya dari [sini](https://golang.org/dl/).
2. Buat dan konfigurasi database MySQL yang dibutuhkan untuk API ini.
3. Sesuaikan konfigurasi database di file `config.yaml` atau di dalam kode jika ada.

## Menjalankan Aplikasi

Untuk menjalankan aplikasi ini secara lokal:

1. Clone repository:
   ```bash
   git clone https://github.com/cfra321/sb-go-batch-59-kukuh.git
   cd sb-go-batch-59-kukuh

Jalankan aplikasi menggunakan perintah berikut:

```bash
go run main.go
```

## License

Dokumentasi ini memberikan panduan lengkap untuk menjalankan, dan menggunakan API. Anda dapat menyesuaikan informasi sesuai kebutuhan proyek Anda.



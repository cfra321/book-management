# Book Management API

Repository ini berisi API untuk mengelola data buku dan kategori yang dibangun menggunakan **Golang** dan **Gin Gonic**.

## Fitur and API Endpoints

- **Buku**
  - **`POST /api/books`**
    - Menambahkan buku baru
  - **`PUT /api/books/:id`**  
    - Memperbarui buku berdasarkan ID
  - **`DELETE /api/books/:id`** 
    - Menghapus buku berdasarkan ID
  - **`GET /api/books`**  
    - Mendapatkan daftar semua buku
  - **`GET /api/books/:id`**  
    - Mendapatkan detail buku berdasarkan ID

- **Kategori**
  - **`POST /api/categories`**  
   - Menambahkan kategori baru
  - **`GET /api/categories`**
    - Mendapatkan daftar semua kategori
  - **`GET /api/categories/:id`**  
    - Mendapatkan detail kategori berdasarkan ID
  - **`PUT /api/categories/:id`** 
    - Memperbarui detail kategori berdasarkan ID
  - **`DELETE /api/categories/:id`**  
    - Menghapus kategori dari sistem berdasarkan ID

## Struktur Proyek

```bash
.
├── go.mod             # Informasi dependensi dan versi modul Go
├── go.sum             # Checksum dari dependensi yang digunakan
├── config/            # Konfigurasi aplikasi
│   └── .env           # File environment untuk menyimpan variabel konfigurasi
├── controllers/       # Handler untuk endpoints
│   ├── book.go        # Logika untuk resource buku (CRUD Buku)
│   ├── category.go    # Logika untuk resource kategori buku (CRUD Kategori)
│   ├── login.go       # Logika untuk otentikasi pengguna (Login)
│   └── user.go        # Logika untuk resource pengguna (CRUD Pengguna)
├── database/          # Konfigurasi database dan migrasi
│   ├── database.go    # Koneksi dan pengaturan database
│   └── sql_migrations/  # Direktori untuk file migrasi SQL
│       ├── 1_initiate_book.sql     # SQL untuk membuat tabel buku
│       ├── 2_initiate_category.sql # SQL untuk membuat tabel kategori
│       └── 3_initiate_user.sql     # SQL untuk membuat tabel pengguna
├── helpers/           # Fungsi utilitas umum
│   └── jwt.go         # Fungsi helper untuk JWT (JSON Web Token)
├── middleware/        # Middleware untuk aplikasi
│   └── middleware.go  # Implementasi middleware (misalnya, otentikasi, logging)
├── repository/        # Interaksi langsung dengan database
│   ├── book.go        # Implementasi repository untuk resource buku
│   ├── category.go    # Implementasi repository untuk resource kategori
│   └── user.go        # Implementasi repository untuk resource pengguna
├── seeder/            # Data seeder untuk populasi awal database
│   └── user_seeder.go # Seeder untuk pengguna awal di database
└── structs/           # Definisi struktur data (model)
    └── struct.go      # Definisi struct untuk buku, kategori, pengguna, dll.

```

Aplikasi akan berjalan di http://localhost:8080.


  
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

Dokumentasi ini MIT



# Book Management API
<ul>
<li>Nama : Kukuh Wicaksono</li>
<li>Quiz 3 - Pekan 3 (Coding) Bootcamp Golang Sanbercode Batch 59</li>
</ul>


### ERD
Sistem Service ini memiliki 3 tabel :
<ul>
<li>Tabel books</li>
<li>Tabel categories</li>
<li>Tabel users</li>
</ul>
<img src="documentation/Database_design.png">

## AKSES TOKEN

## Endpoint Login 
Endpoint Login digunakan untuk autentikasi pengguna dan mendapatkan token JWT.

Method | Path | Keterangan | Auth | Body Request  
------------- | ------------- | ------------- | -------------  | -------------  
***POST*** | *`{base_url}/login`* | Mengautentikasi pengguna dan mengembalikan token JWT jika kredensial valid. | No  | `{ "username": "string" ,  "password": "string" }'

**Contoh CURL**
```json
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"admin",
    "password": "password"
}'
```
**Contoh Request Body**
> ⚠️ **Warning**: Username dan password yang valid adalah **admin** dan **password**.
  ```json
  {
    "username":"{input your usernamme}",
    "password": "{input your password}"
  }
  ```
  **Contoh Response Body**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzI1ODg5MzgxfQ.BcSTHFfh-ew_8UsOqWmbpzVw1QcQ7gUFE5sBi395Gp8"
  }
  ```

## Fitur and API Endpoints

### Buku

Method | Path | Keterangan | Auth | Body Request  
------------- | ------------- | ------------- | ------------- | -------------  
***GET*** | *`/api/books`* | Mendapatkan daftar semua buku | No | -  
***POST*** | *`/api/books`* | Menambahkan buku baru | Yes | `{ "title": "string", "description": "string", "image_url": "string", "release_year": 2023, "price": 250000, "total_page": 450, "category_id": 1 }`  
***GET*** | *`/api/books/:id`* | Mendapatkan detail buku berdasarkan ID | No | -  
***PUT*** | *`/api/books/:id`* | Memperbarui buku berdasarkan ID | Yes | `{ "title": "string", "description": "string", "image_url": "string", "release_year": 2023, "price": 250000, "total_page": 450, "category_id": 1 }`  
***DELETE*** | *`/api/books/:id`* | Menghapus buku berdasarkan ID | Yes | -  
  
  
### Kategori

Method | Path | Keterangan | Auth | Body Request  
------------- | ------------- | ------------- | ------------- | -------------  
***POST*** | *`/api/categories`* | Menambahkan kategori baru | Yes | `{ "name": "string", "description": "string" }`  
***GET*** | *`/api/categories`* | Mendapatkan daftar semua kategori | No | -  
***GET*** | *`/api/categories/:id`* | Mendapatkan detail kategori berdasarkan ID | No | -  
***PUT*** | *`/api/categories/:id`* | Memperbarui detail kategori berdasarkan ID | Yes | `{ "name": "string", "description": "string" }`  
***DELETE*** | *`/api/categories/:id`* | Menghapus kategori dari sistem berdasarkan ID | Yes | -  

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



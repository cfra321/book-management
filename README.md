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
***POST*** | *`{base_url}/login`* | Mengautentikasi pengguna dan mengembalikan token JWT jika kredensial valid. | No  | `{ "username": "string" , "password": "string" }`  


> ⚠️ **Warning**: Anda dapat mengunakan Username dan password yang valid yaitu  **admin** dan **password**

**Contoh CURL**
```json
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"your username",
    "password": "your password"
}'
```


## Fitur and API Endpoints

### Buku

Method | Path | Keterangan | Auth | Body Request  
------------- | ------------- | ------------- | ------------- | -------------  
***GET*** | *`{base_url}/books`* | Mendapatkan daftar semua buku | Yes | -  
***POST*** | *`{base_url}/books`* | Menambahkan buku baru | Yes | `{ "title": "string", "description": "string", "image_url": "string", "release_year": "int, "price": 280000, "total_page": 90, "category_id": 3 }`  
***GET*** | *`{base_url}/books/:id`* | Mendapatkan detail buku berdasarkan ID | Yes | -  
***PUT*** | *`{base_url}/books/:id`* | Memperbarui buku berdasarkan ID | Yes | `{ "title": "string", "description": "string", "image_url": "string", "release_year": "int, "price": 280000, "total_page": 90, "category_id": 3 }`  
***DELETE*** | *`{base_url}/books/:id`* | Menghapus buku berdasarkan ID | Yes | -  
  
  
### Kategori

Method | Path | Keterangan | Auth | Body Request  
------------- | ------------- | ------------- | ------------- | -------------  
***POST*** | *`{base_url}/categories`* | Menambahkan kategori baru | Yes | `{ "id": 7, "name": "string", "created_by": "string", "modified_by": "string" }`  
***GET*** | *`{base_url}/categories`* | Mendapatkan daftar semua kategori | Yes | -  
***GET*** | *`{base_url}/categories/:id`* | Mendapatkan detail kategori berdasarkan ID | Yes | -  
***PUT*** | *`{base_url}/categories/:id`* | Memperbarui detail kategori berdasarkan ID | Yes | `{ "name": "string", "created_by": "string", "modified_by": "string" }`  
***DELETE*** | *`{base_url}/categories/:id`* | Menghapus kategori dari sistem berdasarkan ID | Yes | -  
***GET*** | *`{base_url}/category/:id/books`* | Mendapatkan detail buku berdasarkan kategori | Yes | -  
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



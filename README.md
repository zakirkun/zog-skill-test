# Zero One Group Backend Test

Test Assesment Backend Dev

## Daftar Isi

- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Menggunakan Docker](#menggunakan-docker)
- [Testing](#testing)
- [Lisensi](#lisensi)

## Prasyarat

Pastikan Anda telah menginstal perangkat berikut:

- [Golang](https://golang.org/doc/install) versi 1.20 atau lebih baru
- [Docker](https://www.docker.com/get-started)

## Instalasi

1. Clone repositori ini

    ```sh
    git clone https://github.com/zakirkun/zot-skill-test backend-dev
    cd backend-dev
    ```

2. Instal dependensi Go

    ```sh
    go mod tidy
    ```

3. Buatlah directory assets

## Menjalankan Aplikasi

Untuk menjalankan aplikasi secara lokal tanpa Docker:

```sh
go run main.go -c config.toml
```

## Menggunakan Docker

### Langkah 1: Build Docker Image

```sh
docker build -t backend-dev .
```

### Langkah 2: Menjalankan Docker Container

Untuk menjalankan container:

```sh
docker run -p 9080:9080 backend-dev
```

### Menggunakan Docker Compose

Jika Anda menggunakan `docker-compose`, cukup jalankan:

```sh
docker-compose up --build -d
```

Docker Compose akan membaca `docker-compose.yml` dan mengatur semua layanan yang diperlukan.

## Testing

Untuk menjalankan unit tests:

```sh
go test ./...
```

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
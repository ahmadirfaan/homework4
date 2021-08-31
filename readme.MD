Merupakan Tugas Homework 4 dengan membuat suatu REST API CRUD dengan menggunakan :

- Go-Fiber
- GORM
- MySQL

Untuk dapat menjalankan source code ini, harap mengisi masing-masing nilai dari DB_HOST, dll di file .env.sample

Menjalankan command Dockerfile :

```dockerfile
docker build -t <nama image>.
```

Menjalankan container dengan image yang telah dibuat:

```dockerfile
docker container create --name golang-movie -e DB_HOST="host.docker.internal" -e DB_PORT=<port_db> -e DB_USER=<nama_user> -e DB_PASS=<db_passwordnya> -e DB_DATABASE=<nama_schema_database> -p 8000:8000 image-golang-movie
```

Disarankan DB_HOST memakai host.docker.internal dikarenakan MySQL berjalan pada localhost dan agar komunikasi antara Docker dan localhost menggunakan host seperti itu 

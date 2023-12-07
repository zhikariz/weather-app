# Gunakan image resmi Golang sebagai base image
FROM golang:latest

# Atur working directory di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum terlebih dahulu untuk mendownload dependensi
COPY go.mod .
COPY go.sum .

# Download dan instal dependensi
RUN go mod download

# Salin seluruh file dari direktori aplikasi ke dalam container
COPY . .

# Kompilasi aplikasi Golang
RUN go build -o main cmd/server/main.go

# Expose port yang digunakan oleh aplikasi
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["./main"]
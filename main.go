package main

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func createDatabaseIfNotExists() {
	// Koneksi sementara ke MySQL tanpa menyebutkan database tertentu
	dsn := "root:@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal terhubung ke MySQL:", err)
	}
	defer db.Close()

	// Perintah SQL untuk membuat database jika belum ada
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS angkatan_1_2025")
	if err != nil {
		log.Fatal("Gagal membuat database:", err)
	}
	fmt.Println("Database sudah ada atau berhasil dibuat.")
}

func ConnectDatabase() {
	// Pastikan database sudah ada
	createDatabaseIfNotExists()

	// Koneksi ke database
	dsn := "root:@tcp(127.0.0.1:3306)/angkatan_1_2025?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	DB = db
	fmt.Println("Berhasil terhubung ke database!")
}

func main() {
	ConnectDatabase()
}

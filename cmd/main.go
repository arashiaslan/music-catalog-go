package main

import (
	"log"

	"github.com/arashiaslan/music-catalog-go/configs"
	"github.com/arashiaslan/music-catalog-go/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal membaca konfigurasi", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal koneksi ke database, err: %v", err)
	}

	r := gin.Default()

	r.Run(cfg.Service.Port)
}
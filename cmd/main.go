package main

import (
	"log"

	"github.com/arashiaslan/music-catalog-go/internal/configs"
	membershipHandler "github.com/arashiaslan/music-catalog-go/internal/handler/memberships"
	"github.com/arashiaslan/music-catalog-go/internal/models/memberships"
	membershipRepo "github.com/arashiaslan/music-catalog-go/internal/repository/memberships"
	membershipSvc "github.com/arashiaslan/music-catalog-go/internal/services/memberships"
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

	db.AutoMigrate(&memberships.User{})
	// db.AutoMigrate(&memberships.SignupRequest{})

	r := gin.Default()

	membershipRepo := membershipRepo.NewRepository(db)
	membershipSvc := membershipSvc.NewService(cfg, membershipRepo)
	membershipHandler := membershipHandler.NewHandler(r, membershipSvc)

	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}

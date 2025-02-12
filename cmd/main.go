package main

import (
	"log"

	"xprasetio/go-account-recovery.git/internal/configs"
	membershipsHandler "xprasetio/go-account-recovery.git/internal/handler/memberships"
	"xprasetio/go-account-recovery.git/internal/helpers"
	"xprasetio/go-account-recovery.git/internal/models/memberships"
	membershipsRepo "xprasetio/go-account-recovery.git/internal/repository/memberships"
	membershipsSvc "xprasetio/go-account-recovery.git/internal/service/memberships"
	"xprasetio/go-account-recovery.git/pkg/internalsql"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)
	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/", // for local configs file path
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v", err)
	}

	db.AutoMigrate(&memberships.User{})

	// load log
	helpers.SetupLogger()
	
	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)
	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)
			
	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
package main

import (
	"log"
	"net/http"
	"web-desa/config"
	"web-desa/handler"
	"web-desa/repository"
	"web-desa/service"

	"github.com/gin-gonic/gin"
)

type (
	server struct {
		httpServer *gin.Engine
		cfg			config.Config
	}

	Server interface {
		Run()
	}
)

func InitServer(cfg config.Config) Server {
	r := gin.Default()
	
	// r.Use(cors.New(cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// }))

	return &server{
		httpServer: r,
		cfg:		cfg,
	}
}

func(s *server) Run() {

	s.httpServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "test success",
		})
	})

	// s.httpServer.GET("/seeder", func(ctx *gin.Context) {
	// 	helper.SeederRefresh(s.cfg)
		
	// 	helper.ResponseSuccessJson(ctx, "seeder success", "")
	// })

	umkmRepo := repository.NewUmkmRepository(s.cfg)
	umkmService := service.NewumkmService(umkmRepo)
	umkmHandler := handler.NewUmkmHandler(umkmService)
	umkmGroup := s.httpServer.Group("/umkm")
	umkmHandler.Mount(umkmGroup)

	wisataRepo := repository.NewWisataRepository(s.cfg)
	wisataService := service.NewWisataService(wisataRepo)
	wisataHandler := handler.NewWisataHandler(wisataService)
	wisataGroup := s.httpServer.Group("/wisata")
	wisataHandler.Mount(wisataGroup)

	if err := s.httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
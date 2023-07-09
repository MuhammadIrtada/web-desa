package main

import (
	"log"
	"net/http"
	"web-desa/config"

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

	// matkulRepo := repository.NewMatkulRepository(s.cfg)
	// matkulService := service.NewMatkulService(matkulRepo)
	// matkulHandler := handler.NewMatkulHandler(matkulService)
	// matkulGroup := s.httpServer.Group("/matkul")
	// matkulGroup.Use(middleware.ValidateToken())
	// matkulHandler.Mount(matkulGroup)

	if err := s.httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"log"
	"os"

	"user-service/handler"
	"user-service/repository"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	r.GET("/ping", h.Ping)
	r.GET("/users", h.ListUsers)
	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUserByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7001"
	}
	addr := ":" + port
	log.Printf("user-service listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

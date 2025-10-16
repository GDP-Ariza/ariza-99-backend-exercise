package main

import (
	"log"
	"os"
	"public_service/adapter"
	"public_service/handler"
	"public_service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userAdapter := adapter.NewUserAdapter()
	listingAdapter := adapter.NewListingAdapter()
	svc := service.NewPublicService(userAdapter, listingAdapter)
	h := handler.NewUserHandler(svc)

	r.GET("/public-api/ping", h.Ping)
	r.GET("/public-api/listings", h.Listings)
	r.POST("/public-api/listings", h.CreateListing)
	r.POST("/public-api/users", h.CreateUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7002"
	}
	addr := ":" + port
	log.Printf("public-service listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

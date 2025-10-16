package handler

import (
	"net/http"
	"public_service/model"
	"public_service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.PublicService
}

func NewUserHandler(svc service.PublicService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *UserHandler) Listings(c *gin.Context) {
	// Parse pagination params with defaults
	pageNumStr := c.Query("page_num")
	pageSizeStr := c.Query("page_size")
	userIdStr := c.Query("user_id")

	pageNum := 1
	pageSize := 10
	if pageNumStr != "" {
		if v, err := strconv.Atoi(pageNumStr); err == nil && v > 0 {
			pageNum = v
		}
	}
	if pageSizeStr != "" {
		if v, err := strconv.Atoi(pageSizeStr); err == nil && v > 0 {
			pageSize = v
		}
	}
	// Optional safety cap
	if pageSize > 100 {
		pageSize = 100
	}

	userId := 0
	if userIdStr != "" {
		if v, err := strconv.Atoi(userIdStr); err == nil && v > 0 {
			userId = v
		}
	}

	listings, err := h.svc.Listings(pageNum, pageSize, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.ListingListResponse{Result: true, Listings: listings})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	user, err := h.svc.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.UserResponse{Result: true, User: user})
}

func (h *UserHandler) CreateListing(c *gin.Context) {
	var req model.ListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	listing, err := h.svc.CreateListing(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.ListingResponse{Result: true, Listing: listing})
}

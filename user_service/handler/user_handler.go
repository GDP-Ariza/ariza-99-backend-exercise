package handler

import (
	"net/http"
	"strconv"

	"user-service/model"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	// Parse pagination params with defaults
	pageNumStr := c.Query("page_num")
	pageSizeStr := c.Query("page_size")

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

	response := h.svc.List(pageNum, pageSize)
	c.JSON(http.StatusOK, model.UserListResponse{Result: true, Users: response})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	user, err := h.svc.Create(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.UserResponse{Result: true, User: user})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.svc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, model.UserResponse{Result: true, User: user})
}

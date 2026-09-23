package handler

import (
	"day01/internal/model"
	"day01/internal/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{ store *store.UserStore }

func NewUserHandler(s *store.UserStore) *UserHandler { return &UserHandler{store: s} }
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "connected"})
}
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.store.List(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.store.Create(c.Request.Context(), req.Username, req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

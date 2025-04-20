package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/users")
	{
		group.GET("", h.GetAllUsers)
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
    users, err := h.UserService.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
        return
    }

    c.JSON(http.StatusOK, users)
}
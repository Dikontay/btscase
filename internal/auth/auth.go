package auth

import "github.com/gin-gonic/gin"

func SetRoutes(c *gin.RouterGroup, h HandlersInterface) {
	c.POST("/register", h.Register)
	c.POST("/login", h.Login)
}

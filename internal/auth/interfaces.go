package auth

import (
	"context"

	"github.com/Dikontay/btscase/internal/models"
	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Authorize() gin.HandlerFunc
}

type ServiceInterface interface {
	Register(context.Context, *models.User) (string, error)
	Login(context.Context, *models.User) (string, error)
}

type RepoInterface interface {
	Insert(context.Context, *models.User) (int, error)
	GetByEmail(context.Context, string) (*models.User, error)
	GetByID(context.Context, int) (*models.User, error)
	GetAll(context.Context) ([]models.User, error)
}

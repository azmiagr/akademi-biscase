package middleware

import (
	"akademi-business-case/internal/service"
	"akademi-business-case/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	AuthenticateUser(c *gin.Context)
	OnlyAdmin(c *gin.Context)
	OnlyOperational(c *gin.Context)
	Timeout() gin.HandlerFunc
	Cors() gin.HandlerFunc
}

type middleware struct {
	service *service.Service
	jwtAuth jwt.Interface
}

func Init(service *service.Service, jwtAuth jwt.Interface) Interface {
	return &middleware{
		service: service,
		jwtAuth: jwtAuth,
	}
}

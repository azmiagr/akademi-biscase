package rest

import (
	"akademi-business-case/internal/service"
	"akademi-business-case/pkg/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	router := r.router.Group("/api/v1")

	auth := router.Group("/auth")
	auth.POST("/register", r.Register)
	auth.PATCH("/register", r.VerifyUser)
	auth.PATCH("/register/resend", r.ResendOtp)
	auth.POST("/login", r.Login)

	user := router.Group("/users")
	user.Use(r.middleware.AuthenticateUser)
	user.GET("/profile", r.GetUserProfile)

	admin := router.Group("/admin")
	admin.Use(r.middleware.AuthenticateUser)
	admin.GET("/mentors", r.GetMentors)
	admin.POST("/classes", r.CreateClass)

	class := router.Group("/classes")
	class.GET("", r.GetAllClasses)
	class.GET("/search", r.GetClassByName)
	class.GET("/types/:classTypeID", r.GetClassesByType)
	class.GET("/:classID", r.GetClassDetail)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}

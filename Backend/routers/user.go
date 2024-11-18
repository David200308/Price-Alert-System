package routers

import (
	"github.com/David200308/go-api/Backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetUser)
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/token", controllers.TokenVerification)
	r.POST("/email/verification", controllers.EmailVerification)
	r.POST("/logout", controllers.LogoutUser)
	// r.PATCH("/", controllers.UpdateUser)
	// r.DELETE("/", controllers.DeleteUser)
}

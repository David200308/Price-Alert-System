package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/user")
	UserRoutes(userGroup)

	PaymentGroup := r.Group("/payment")
	PaymentRoutes(PaymentGroup)

	notificationGroup := r.Group("/notification")
	NotificationRoutes(notificationGroup)

	return r
}

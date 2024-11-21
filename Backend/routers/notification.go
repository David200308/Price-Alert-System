package routers

import (
	"github.com/David200308/go-api/Backend/controllers"
	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.RouterGroup) {
	r.GET("/user/created", controllers.GetUserCreatedNotification)
	r.GET("/user/verified", controllers.GetUserVerifiedNotification)
	r.GET("/payment/created", controllers.GetPaymentCreatedNotification)
	r.GET("/payment/successful", controllers.GetPaymentSuccessfulNotification)
	r.GET("/payment/cancelled", controllers.GetPaymentCancelledNotification)
}

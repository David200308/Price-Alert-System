package routers

import (
	"github.com/David200308/go-api/Backend/controllers"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup) {
	r.POST("/init", controllers.InitPayment)
	r.PATCH("/update", controllers.UpdatePaymentStatus)
	r.GET("/", controllers.GetPayment)
}

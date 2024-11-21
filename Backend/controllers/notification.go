package controllers

// import (
// 	"github.com/David200308/go-api/Backend/initializers"
// 	"github.com/David200308/go-api/Backend/tools"
// 	"github.com/gin-gonic/gin"
// )

// // @Router /notification/payment/created [Post]
// func GetPaymentCreatedNotification(c *gin.Context) {
// 	userUUID, _, err := tools.NormalRequestVerifyToken(c)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"status":  "error",
// 			"message": "Unauthorized",
// 			"error":   err,
// 		})
// 		return
// 	}

// 	res, err := initializers.MQInstance.Consume("payments.event.created", "Payment.Created", userUUID)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"status":  "error",
// 			"message": "Failed to consume message",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status": "success",
// 		"result": res,
// 	})
// }

// // @Router /notification/payment/successful [Post]
// func GetPaymentSuccessfulNotification(c *gin.Context) {
// 	userUUID, _, err := tools.NormalRequestVerifyToken(c)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"status":  "error",
// 			"message": "Unauthorized",
// 			"error":   err,
// 		})
// 		return
// 	}

// 	res, err := initializers.MQInstance.Consume("payments.event.success", "Payment.Success", userUUID)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"status":  "error",
// 			"message": "Failed to consume message",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status": "success",
// 		"result": res,
// 	})
// }

// // @Router /notification/payment/cancelled [Post]
// func GetPaymentCancelledNotification(c *gin.Context) {
// 	userUUID, _, err := tools.NormalRequestVerifyToken(c)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"status":  "error",
// 			"message": "Unauthorized",
// 			"error":   err,
// 		})
// 		return
// 	}

// 	res, err := initializers.MQInstance.Consume("payments.event.cancelled", "Payment.Cancelled", userUUID)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"status":  "error",
// 			"message": "Failed to consume message",
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status": "success",
// 		"result": res,
// 	})
// }

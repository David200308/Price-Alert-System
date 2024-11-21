package controllers

import (
	"time"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/tools"
	"github.com/gin-gonic/gin"
)

const userCreatedQueueName = "notification:user:created"
const userVerifiedQueueName = "notification:user:verified"
const paymentCreatedQueueName = "notification:payment:created"
const paymentSuccessQueueName = "notification:payment:success"
const paymentCancelledQueueName = "notification:payment:cancelled"

// @Router /notification/user/created [Post]
func GetUserCreatedNotification(c *gin.Context) {
	_, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	res, err := initializers.MQConsume(userCreatedQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": res,
	})
}

// @Router /notification/user/verified [Post]
func GetUserVerifiedNotification(c *gin.Context) {
	_, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	res, err := initializers.MQConsume(userVerifiedQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": res,
	})
}

// @Router /notification/payment/created [Post]
func GetPaymentCreatedNotification(c *gin.Context) {
	_, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	res, err := initializers.MQConsume(paymentCreatedQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": res,
	})
}

// @Router /notification/payment/successful [Post]
func GetPaymentSuccessfulNotification(c *gin.Context) {
	_, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	res, err := initializers.MQConsume(paymentSuccessQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": res,
	})
}

// @Router /notification/payment/cancelled [Post]
func GetPaymentCancelledNotification(c *gin.Context) {
	_, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	res, err := initializers.MQConsume(paymentCancelledQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": res,
	})
}

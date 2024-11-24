package controllers

import (
	"time"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/mq"
	"github.com/David200308/go-api/Backend/tools"
	"github.com/gin-gonic/gin"
)

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

	res, err := initializers.MQConsume(mq.UserCreatedQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	var resJson []interface{}
	for _, v := range res {
		resJson = append(resJson, tools.StringToJSON(v))
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": resJson,
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

	res, err := initializers.MQConsume(mq.UserVerifiedQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	var resJson []interface{}
	for _, v := range res {
		resJson = append(resJson, tools.StringToJSON(v))
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": resJson,
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

	res, err := initializers.MQConsume(mq.PaymentCreatedQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	var resJson []interface{}
	for _, v := range res {
		resJson = append(resJson, tools.StringToJSON(v))
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": resJson,
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

	res, err := initializers.MQConsume(mq.PaymentSuccessQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	var resJson []interface{}
	for _, v := range res {
		resJson = append(resJson, tools.StringToJSON(v))
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": resJson,
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

	res, err := initializers.MQConsume(mq.PaymentCancelledQueueName, 5*time.Second)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Failed to consume message",
			"error":   err,
		})
		return
	}

	var resJson []interface{}
	for _, v := range res {
		resJson = append(resJson, tools.StringToJSON(v))
	}

	c.JSON(200, gin.H{
		"status": "success",
		"result": resJson,
	})
}

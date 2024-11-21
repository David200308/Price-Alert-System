package controllers

import (
	"strconv"

	"github.com/David200308/go-api/Backend/models"
	"github.com/David200308/go-api/Backend/mq"
	"github.com/David200308/go-api/Backend/services"
	"github.com/David200308/go-api/Backend/tools"
	"github.com/gin-gonic/gin"
)

// @Router /payment/init [Post]
// @Param amount formData int true "Amount"
// @Param order_uuid formData string true "Order UUID"
// @Param method formData string true "Method"
// @Param currency formData string true "Currency"
// @Success 200
// @Failure 400
// @Failure 500
func InitPayment(c *gin.Context) {
	userUUID, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	payment := models.Payment{
		Amount: func() int64 {
			amount, err := strconv.ParseInt(c.PostForm("amount"), 10, 64)
			if err != nil {
				return 0
			}
			return amount
		}(),
		OrderUUID:   c.PostForm("order_uuid"),
		Method:      c.PostForm("method"),
		Currency:    c.PostForm("currency"),
		UserUUID:    userUUID,
		PaymentUUID: tools.GenerateUUID(),
		Status:      "pending",
	}

	if payment.Amount == 0 || payment.OrderUUID == "" || payment.Method == "" || payment.Currency == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Amount, Order UUID, Method and Currency are required",
		})
		return
	}

	paymentSessionLink, err := services.CreatePayment(&payment)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to create payment",
		})
		return
	}

	mq.PaymentCreated(userUUID, payment.PaymentUUID)

	c.JSON(200, gin.H{
		"status":       "success",
		"url":          paymentSessionLink,
		"payment_uuid": payment.PaymentUUID,
	})
}

// @Router /payment [get]
// @Param order_uuid query string false "Order UUID"
// @Param payment_uuid query string false "Payment UUID"
// @Success 200
// @Failure 500
func GetPayment(c *gin.Context) {
	userUUID, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	orderUUID := c.Query("order_uuid")
	paymentUUID := c.Query("payment_uuid")

	switch {
	case orderUUID != "" && userUUID != "" && paymentUUID == "":
		payment, err := services.GetPaymentByOrderUUID(orderUUID, userUUID)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to get payment",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "success",
			"payment": payment,
		})

	case userUUID != "" && orderUUID == "" && paymentUUID == "":
		payments, err := services.GetAllPaymentByUserUUID(userUUID)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to get payments",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":   "success",
			"payments": payments,
		})

	case paymentUUID != "" && orderUUID == "" && userUUID != "":
		payment, err := services.GetPaymentByPaymentUUID(paymentUUID, userUUID)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to get payment",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "success",
			"payment": payment,
		})

	default:
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Order UUID, User UUID or Payment UUID is required",
		})
	}
}

// @Router /payment/update [Patch]
// @Param payment_uuid formData string true "Payment UUID"
// @Param status formData string true "Status"
// @Param reference_id formData string false "Reference ID"
// @Success 200
// @Failure 400
// @Failure 500
func UpdatePaymentStatus(c *gin.Context) {
	userUUID, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	paymentUUID := c.PostForm("payment_uuid")
	status := c.PostForm("status")
	referenceID := c.PostForm("reference_id")

	if paymentUUID == "" || status == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Payment UUID and Status are required",
		})
		return
	}

	if err := services.UpdatePaymentStatus(userUUID, paymentUUID, status, referenceID); err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to update payment status",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
	})
}

// @Router /payment/callback/stripe/{payment_uuid} [Get]
// @Param payment_uuid path string true "Payment UUID"
// @Param checkout_session_id path string false "Checkout Session ID"
// @Success 200
// @Failure 400
// @Failure 500
func StripePaymentCallback(c *gin.Context) {
	userUUID, _, err := tools.NormalRequestVerifyToken(c)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Unauthorized",
			"error":   err,
		})
		return
	}

	paymentUUID := c.Param("payment_uuid")
	checkoutSessionID := c.Param("checkout_session_id")

	if paymentUUID == "" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Payment UUID is required",
		})
		return
	}

	payment, err := services.GetPaymentByPaymentUUID(paymentUUID, userUUID)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to get payment",
		})
		return
	}

	if payment.Method != "stripe" {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid payment method",
		})
		return
	}

	if checkoutSessionID == "" {
		if err := services.UpdatePaymentStatus(userUUID, paymentUUID, "cancelled", ""); err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Failed to update payment status",
			})
			return
		}

		mq.PaymentCancelled(userUUID, paymentUUID)

		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Payment cancelled",
		})
	}

	if err := services.UpdatePaymentStatus(userUUID, paymentUUID, "success", checkoutSessionID); err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to update payment status",
		})
		return
	}

	mq.PaymentSuccessful(userUUID, paymentUUID)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Payment success",
	})
}

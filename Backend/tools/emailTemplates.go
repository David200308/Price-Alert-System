package tools

import "fmt"

func ActivationEmail(activationLink string) string {
	return fmt.Sprintf("<p>Please activate your account by clicking <a href=\"%s\">here</a>.</p>", activationLink)
}

func PaymentSuccessEmail(paymentUUID string, amount int64, currency string) string {
	return fmt.Sprintf("<p>Your payment was successful. Total amount paid: %s %d. Payment Reference: %s</p>", currency, amount/100, paymentUUID)
}

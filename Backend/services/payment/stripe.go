package payment

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

func CreateStripePaymentSession(amount int64, currency, paymentUUID string) (*stripe.CheckoutSession, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(fmt.Sprintf("%s/payment/callback/stripe/%s/{CHECKOUT_SESSION_ID}", os.Getenv("BACKEND_URL"), paymentUUID)),
		CancelURL:  stripe.String(fmt.Sprintf("%s/payment/callback/stripe/%s", os.Getenv("BACKEND_URL"), paymentUUID)),
		InvoiceCreation: &stripe.CheckoutSessionInvoiceCreationParams{
			Enabled: stripe.Bool(true),
		},
		AllowPromotionCodes: stripe.Bool(true),
		PaymentMethodTypes:  stripe.StringSlice([]string{"card"}),
		Locale:              stripe.String("auto"),
		Mode:                stripe.String("payment"),
		PhoneNumberCollection: &stripe.CheckoutSessionPhoneNumberCollectionParams{
			Enabled: stripe.Bool(true),
		},
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Quantity: stripe.Int64(1),
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:    stripe.String(currency),
					UnitAmount:  stripe.Int64(amount),
					TaxBehavior: stripe.String("inclusive"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Payment"),
					},
				},
			},
		},
	}

	s, err := session.New(params)
	if err != nil {
		return nil, err
	}

	return s, nil
}

package mapper

import (
	"crypto/sha256"
	"encoding/hex"
	"example/zota_assignment/config"
	"example/zota_assignment/data"
	"example/zota_assignment/model"
	"fmt"
)

func OrderToZotaEntity(order model.Order) data.ZotaOrderRequest {
	return data.ZotaOrderRequest{
		Merchant:      *data.GetMerchant(),
		Customer:      *data.GetCustomer(),
		OrderAmount:   order.OrderAmount,
		OrderCurrency: order.OrderCurrency,
		RedirectUrl:   "https://www.example-merchant.com/payment-return/",
		CallbackUrl:   "https://www.example-merchant.com/payment-callback/",
		CheckoutUrl:   "https://www.example-merchant.com/account/deposit/?uid=3c39583c",
		CustomParam:   "{\"UserId\": \"3c39583c\"}",
		Language:      "English",
		Signature:     generateSignature(order),
	}
}

func generateSignature(order model.Order) string {
	input := config.EndpointID +
		data.GetMerchant().MerchantOrderID +
		fmt.Sprintf("%f", order.OrderAmount) +
		data.GetCustomer().CustomerEmail +
		config.MerchantSecretKey

	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

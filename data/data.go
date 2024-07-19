package data

type Merchant struct {
	MerchantOrderID   string `json:"merchantOrderID"`
	MerchantOrderDesc string `json:"merchantOrderDesc"`
}

type Customer struct {
	CustomerEmail             string `json:"customerEmail"`
	CustomerFirstName         string `json:"customerFirstName"`
	CustomerLastName          string `json:"customerLastName"`
	CustomerAddress           string `json:"customerAddress"`
	CustomerCountryCode       string `json:"customerCountryCode"`
	CustomerCity              string `json:"customerCity"`
	CustomerState             string `json:"customerState"`
	CustomerZipCode           string `json:"customerZipCode"`
	CustomerPhone             string `json:"customerPhone"`
	CustomerIP                string `json:"customerIP"`
	CustomerPersonalID        string `json:"customerPersonalID"`
	CustomerBankCode          string `json:"customerBankCode"`
	CustomerBankAccountNumber string `json:"customerBankAccountNumber"`
}

func GetMerchant() *Merchant {
	return &Merchant{
		MerchantOrderID:   "test1",
		MerchantOrderDesc: "Some description",
	}
}

func GetCustomer() *Customer {
	return &Customer{
		CustomerEmail:             "customer@email-address.com",
		CustomerFirstName:         "John",
		CustomerLastName:          "Smith",
		CustomerAddress:           "Test street",
		CustomerCountryCode:       "BG",
		CustomerCity:              "Sofia",
		CustomerState:             "Sofia",
		CustomerZipCode:           "02",
		CustomerPhone:             "+1 999 999 9998§",
		CustomerIP:                "1.1.1.1.",
		CustomerPersonalID:        "1234",
		CustomerBankCode:          "STSA123",
		CustomerBankAccountNumber: "123456789",
	}
}

type ZotaOrderRequest struct {
	Merchant      Merchant
	OrderAmount   float64  `json:"orderAmount"`
	OrderCurrency string   `json:"orderCurrency"`
	Customer      Customer `json:"merchantOrderID"`
	RedirectUrl   string   `json:"redirectUrl"`
	CallbackUrl   string   `json:"callbackUrl"`
	CheckoutUrl   string   `json:"checkoutUrl"`
	CustomParam   string   `json:"customParam"`
	Language      string   `json:"language"`
	Signature     string   `json:"signature"`
}

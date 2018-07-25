package Request

var CreatePaymentEndpoint string = "/payment/auth"

type CreatePaymentRequest struct {

	Locale          string                        `json:"locale"`
	ConversationId  string                        `json:"conversationId"`
	Price           string                        `json:"price"`
	PaidPrice       string                        `json:"paidPrice"`
	Installment     int                           `json:"installment"`
	PaymentChannel  string                        `json:"paymentChannel"`
	BasketId        string                        `json:"basketId"`
	PaymentGroup    string                        `json:"paymentGroup"`
	PaymentCard     CreatePayment_PaymentCard     `json:"paymentCard"`
	Buyer           CreatePayment_Buyer           `json:"buyer"`
	ShippingAddress CreatePayment_ShippingAddress `json:"shippingAddress"`
	BillingAddress  CreatePayment_BillingAddress  `json:"billingAddress"`
	BasketItems     []CreatePayment_BasketItems   `json:"basketItems"`
	Currency        string                        `json:"currency"`
}

type CreatePayment_PaymentCard struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	Cvc            string `json:"cvc"`
	RegisterCard   int    `json:"registerCard"`
}

type CreatePayment_Buyer struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identityNumber"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsmNumber"`
	RegistrationDate    string `json:"registrationDate"`
	LastLoginDate       string `json:"lastLoginDate"`
	RegistrationAddress string `json:"registrationAddress"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ZipCode             string `json:"zipCode"`
	Ip                  string `json:"ip"`
}

type CreatePayment_ShippingAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type CreatePayment_BillingAddress struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type CreatePayment_BasketItems struct {
	Id        string `json:"id"`
	Price     string `json:"price"`
	Name      string `json:"name"`
	Category1 string `json:"category1"`
	Category2 string `json:"category2"`
	ItemType  string `json:"itemType"`
}

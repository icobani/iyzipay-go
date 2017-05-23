/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:49
* Developer : ibrahimcobani
*
*******/
package Model

type PaymentResource struct {
	Price                        string `json:"price"`
	PaidPrice                    string `json:"paid_price"`
	Installment                  int `json:"installment"`
	Currency                     string `json:"currency"`
	PaymentId                    string `json:"payment_id"`
	PaymentStatus                string `json:"payment_status"`
	FraudStatus                  int `json:"fraud_status"`
	MerchantCommissionRate       string `json:"merchant_commission_rate"`
	MerchantCommissionRateAmount string `json:"merchant_commission_rate_amount"`
	IyziCommissionRateAmount     string `json:"iyzi_commission_rate_amount"`
	IyziCommissionFee            string `json:"iyzi_commission_fee"`
	CardType                     string `json:"card_type"`
	CardAssociation              string `json:"card_association"`
	CardFamily                   string `json:"card_family"`
	CardToken                    string `json:"card_token"`
	CardUserKey                  string `json:"card_user_key"`
	BinNumber                    string `json:"bin_number"`
	BasketId                     string `json:"basket_id"`
	PaymentItems                 []PaymentItem `json:"payment_items"`
	ConnectorName                string `json:"connector_name"`
	AuthCode                     string `json:"auth_code"`
	Phase                        string `json:"phase"`
}

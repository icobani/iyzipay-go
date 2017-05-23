/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:54
* Developer : ibrahimcobani
*
*******/
package Model

type PayoutCompletedTransaction struct {
	PaymentTransactionId string `json:"payment_transaction_id"`
	PayoutAmount         string `json:"payout_amount"`
	PayoutType           string `json:"payout_type"`
	SubMerchantKey       string `json:"sub_merchant_key"`
	Currency             string `json:"currency"`
}

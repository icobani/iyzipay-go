/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:46
* Developer : ibrahimcobani
*
*******/
package Model

type PaymentItem struct {
	ItemId                        string `json:"item_id"`
	PaymentTransactionId          string `json:"payment_transaction_id"`
	TransactionStatus             int `json:"transaction_status"`
	Price                         string `json:"price"`
	PaidPrice                     string `json:"paid_price"`
	MerchantCommissionRate        string `json:"merchant_commission_rate"`
	MerchantCommissionRateAmount  string `json:"merchant_commission_rate_amount"`
	IyziCommissionRateAmount      string `json:"iyzi_commission_rate_amount"`
	IyziCommissionFee             string `json:"iyzi_commission_fee"`
	BlockageRate                  string `json:"blockage_rate"`
	BlockageRateAmountMerchant    string `json:"blockage_rate_amount_merchant"`
	BlockageRateAmountSubMerchant string `json:"blockage_rate_amount_sub_merchant"`
	BlockageResolvedDate          string `json:"blockage_resolved_date"`
	SubMerchantKey                string `json:"sub_merchant_key"`
	SubMerchantPrice              string `json:"sub_merchant_price"`
	SubMerchantPayoutRate         string `json:"sub_merchant_payout_rate"`
	SubMerchantPayoutAmount       string `json:"sub_merchant_payout_amount"`
	MerchantPayoutAmount          string `json:"merchant_payout_amount"`
	ConvertedPayout               ConvertedPayout `json:"converted_payout"`
}

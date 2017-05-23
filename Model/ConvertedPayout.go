/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:27
* Developer : ibrahimcobani
*
*******/
package Model

type ConvertedPayout struct {
	PaidPrice                     string `json:"paid_price"`
	IyziCommissionRateAmount      string `json:"iyzi_commission_rate_amount"`
	IyziCommissionFee             string `json:"iyzi_commission_fee"`
	BlockageRateAmountMerchant    string `json:"blockage_rate_amount_merchant"`
	BlockageRateAmountSubMerchant string `json:"blockage_rate_amount_sub_merchant"`
	SubMerchantPayoutAmount       string `json:"sub_merchant_payout_amount"`
	MerchantPayoutAmount          string `json:"merchant_payout_amount"`
	IyziConversionRate            string `json:"iyzi_conversion_rate"`
	IyziConversionRateAmount      string `json:"iyzi_conversion_rate_amount"`
	Currency                      string `json:"currency"`
}

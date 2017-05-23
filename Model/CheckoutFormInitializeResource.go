/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:26
* Developer : ibrahimcobani
*
*******/
package Model

type CheckoutFormInitializeResource struct {
	Token               string `json:"token"`
	CheckoutFormContent string `json:"checkout_form_content"`
	TokenExpireTime     int32 `json:"token_expire_time"`
	PaymentPageUrl      string `json:"payment_page_url"`
}

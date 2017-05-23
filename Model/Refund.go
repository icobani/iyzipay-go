/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 21:00
* Developer : ibrahimcobani
*
*******/
package Model

type Refund struct {
	PaymentId            string `json:"payment_id"`
	PaymentTransactionId string `json:"payment_transaction_id"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	ConnectorName        string `json:"connector_name"`
}

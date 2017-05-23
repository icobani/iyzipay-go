/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:18
* Developer : ibrahimcobani
*
*******/
package Model

type Cancel struct {
	PaymentId     string `json:"payment_id"`
	Price         string `json:"price"`
	Currency      string `json:"currency"`
	ConnectorName string `json:"connector_name"`
}

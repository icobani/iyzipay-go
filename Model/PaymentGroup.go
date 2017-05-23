/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:45
* Developer : ibrahimcobani
*
*******/
package Model

type PaymentGroup int

const (
	PRODUCT PaymentGroup = 1 + iota
	LISTING
	SUBSCRIPTION
)

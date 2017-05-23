/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:42
* Developer : ibrahimcobani
*
*******/
package Model

type PaymentChannel int

const (
	MOBILE PaymentChannel = 1 + iota
	WEB
	MOBILE_WEB
	MOBILE_IOS
	MOBILE_ANDROID
	MOBILE_WINDOWS
	MOBILE_TABLET
	MOBILE_PHONE
)

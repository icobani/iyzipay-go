/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 21:05
* Developer : ibrahimcobani
*
*******/
package Model

type SubMerchantType int

const (
	PERSONAL SubMerchantType = 1 + iota
	PRIVATE_COMPANY
	LIMITED_OR_JOINT_STOCK_COMPANY
)

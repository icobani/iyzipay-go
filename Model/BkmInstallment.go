/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:14
* Developer : ibrahimcobani
*
*******/
package Model

type BkmInstallment struct {
	BankId int32 `json:"bank_id"`
	InstallmentPrices[] BkmInstallmentPrice `json:"installment_prices"`
}

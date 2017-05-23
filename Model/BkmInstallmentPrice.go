/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:13
* Developer : ibrahimcobani
*
*******/
package Model

type BkmInstallmentPrice struct {
	InstallmentNumber int `json:"installment_number"`
	TotalPrice        string `json:"total_price"`
}

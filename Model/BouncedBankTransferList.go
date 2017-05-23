/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:15
* Developer : ibrahimcobani
*
*******/
package Model

type BouncedBankTransferList struct {
	BankTransfers [] BankTransfer `json:"bank_transfers"`
}

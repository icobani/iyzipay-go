/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:32
* Developer : ibrahimcobani
*
*******/
package Installment

type Detail struct {
	BinNumber         string
	Price             string
	CardType          string
	CardAssociation   string
	CardFamilyName    string
	Force3Ds          int
	BankCode          int32
	BankName          string
	ForceCvc          int
	InstallmentPrices Price
}

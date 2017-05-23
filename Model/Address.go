/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 19:29
* Developer : ibrahimcobani
*
*******/
package Model

type Address struct {
	Description string `json:"description"`
	ZipCode     string `json:"zip_code"`
	ContactName string `json:"contact_name"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:16
* Developer : ibrahimcobani
*
*******/
package Model

type Buyer struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identity_number"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsm_number"`
	RegistrationDate    string `json:"registration_date"`
	LastLoginDate       string `json:"last_login_date"`
	RegistrationAddress string `json:"registration_address"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ZipCode             string `json:"zip_code"`
	Ip                  string `json:"ip"`
}

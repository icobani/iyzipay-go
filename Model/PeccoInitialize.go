/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:57
* Developer : ibrahimcobani
*
*******/
package Model

type PeccoInitialize struct {
	HtmlContent     string `json:"html_content"`
	RedirectUrl     string `json:"redirect_url"`
	Token           string `json:"token"`
	TokenExpireTime int32 `json:"token_expire_time"`
}

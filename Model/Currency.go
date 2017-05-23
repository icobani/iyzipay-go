/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 21:03
* Developer : ibrahimcobani
*
*******/
package Model

import "fmt"

type Currency struct {
	value string `json:"value"`
}

func (c Currency) String() string {
	return fmt.Sprintf("%s", c.value)
}

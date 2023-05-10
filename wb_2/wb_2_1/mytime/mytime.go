// 0.ru.pool.ntp.org
package mytime

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func DkabNow() {
	resp, err := ntp.Time("0.ru.pool.ntp.org")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Local())
}

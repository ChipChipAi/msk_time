package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("2.pool.ntp.org")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(time.String())
	}
}


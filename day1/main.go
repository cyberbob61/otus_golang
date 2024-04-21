package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	// ntpTime, err := ntp.Time("pool.ntp.org")
	// time.apple.com will give the same result as pool.ntp.org
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Println(err)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)

	fmt.Printf("Current time: %v\n", ntpTimeFormatted)
}

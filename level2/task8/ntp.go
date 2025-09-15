package task8

import (
	"fmt"
	"os"

	ntp "github.com/beevik/ntp"
)

func getTime(address string) int {
	t, err := ntp.Time(address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error ocurred: %s", err.Error())
		return 1
	}

	fmt.Println(t)

	return 0
}

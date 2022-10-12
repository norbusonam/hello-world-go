package main

import (
	"fmt"
)

type IPAddress [4]byte

func (ip IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddress{
		"localhost": {127, 0, 0, 1},
		"google":    {142, 250, 69, 206},
	}
	for host, ip := range hosts {
		fmt.Printf("%s: %v\n", host, ip)
	}
}

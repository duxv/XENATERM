package main

import (
	"fmt"

	"github.com/duxv/XENATERM/client"
)

func main() {
	if err := client.Run(); err != nil {
		fmt.Println(err)
	}
}

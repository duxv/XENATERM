package main

import (
	"fmt"

	"github.com/duxv/XENATERM/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}

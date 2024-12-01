package main

import (
	"fmt"

	"github.com/vikingpg3d/just-dance-2015-server/protocols"
)

func main() {
	fmt.Println("Starting Just Dance 2015 Custom Server...")

	protocols.HandleLogin()

	select {}
}

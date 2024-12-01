package main

import (
    "fmt"
    "github.com/vikingpg3d/Just-Dance-2015-Wii-U-online-features-restoration/auth"
    "github.com/vikingpg3d/Just-Dance-2015-Wii-U-online-features-restoration/protocols"
)

func main() {
    fmt.Println("Starting Just Dance 2015 Custom Server...")

}


	protocols.HandleLogin()

	select {}
}

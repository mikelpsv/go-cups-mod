package main

import (
	"fmt"
	cups "github.com/mikelpsv/go-cups-mod"
)

func main() {
	printers := cups.NewConnection()
	n, err := printers.EnumDestinations()
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	fmt.Printf("found %d", n)

	//for _, dest := range printers.Dests {
	//	fmt.Printf("%v (%v)\n", dest.Name, dest.Status())
	//}
}

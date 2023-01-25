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
	fmt.Printf("found %d\n", n)

	for _, dest := range printers.Dests {
		fmt.Printf("%v %s \n", dest.Name, dest.Status())
		if dest.IsDefault {
			jobId := 0
			jobId, err = dest.PrintFile("filename.pdf", "test file name")
			if err != nil {
				fmt.Println(err)

			}
			fmt.Printf("job %d", jobId)
		}
	}
	fmt.Printf("printed successfully")
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	switch m := strings.Title(os.Args[1]); m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month\n", m)
	}
}

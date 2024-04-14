package main

import (
	"fmt"
	"github.com/Dikontay/btscase/forte"
)

func main() {
	newParser, err := forte.ForteBankParser()

	if err != nil {
		fmt.Println(err)
	}
	err = newParser.ParseForte()
	if err != nil {
		fmt.Println(err)
	}

}
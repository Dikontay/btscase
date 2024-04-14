package main

import (
	"fmt"
	"github.com/Dikontay/btscase/internal/parsing/halyk"
)

func main() {

	err := halyk.ParseHalyk()
	if err != nil {
		fmt.Println(err)
	}

}

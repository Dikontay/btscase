package main

import (
	"fmt"
	"github.com/Dikontay/btscase/internal/parsing/halyk"
)

func main() {
	newParser, err := halyk.NewHalykParser()

	if err != nil {
		fmt.Println(err)
	}
	err = newParser.ParseHalyk()
	if err != nil {
		fmt.Println(err)
	}

}

package main

import (
	"fmt"
	"github.com/Dikontay/btscase/internal/parsing/jusan"
)

func main() {
	err := jusan.ParseJusan()
	if err != nil {
		fmt.Println(err)
	}
}

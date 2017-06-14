/*
sumutils

@Author Andres Natanael Soria <andresnatanael@gmail.com>
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andresnatanael/go-training1/utils"
)

var input string

func main() {

	for i := 1; i < len(os.Args); i++ {
		input = os.Args[i]
	}

	v, err := strconv.Atoi(input)
	if err == nil {
		fmt.Println(utils.SumValues(v))
	} else {
		fmt.Println("Integer Argument Required!")
	}

}

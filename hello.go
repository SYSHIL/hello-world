package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(hello("ilhan"))
}

func hello(name string) string {
	return strings.ToUpper(name)
}

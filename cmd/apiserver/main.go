package main

import (
	"apiserver/internal/app/apiserver"
	"fmt"
)

func main() {
	s := apiserver.New()
	fmt.Print(s)
}

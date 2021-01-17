package main

import "fmt"

func main() {
	var m map[string]string = map[string]string{"a": "aa", "b": "bb"}
	m["a"] = "aaa"
	fmt.Println(m)

}

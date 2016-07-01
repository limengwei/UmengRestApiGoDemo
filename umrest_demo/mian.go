package main

import (
	"fmt"
	"limengwei/umrest"
)

const (
	aes_key = "273d7e70c2d115e62e0e45656ff82b39"
)

func main() {

	fmt.Println("--um_rest_demo--")

	data := `{"user_info":{"name":"lmwww","gender":1},"source_uid":"123491239324228","source":"qq"}`

	umrest.BuildUrl(aes_key, data)
}

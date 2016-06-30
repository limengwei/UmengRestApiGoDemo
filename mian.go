package main

import (
	"fmt"
	"umrest"
)

const (
	ak = "273d7e70c2d115e62e0e45656ff82b39"
)

func main() {

	fmt.Println("--um_rest_demo--")

	data := "{'user_info':{'name':'test1','icon_url':'http: //umeng.com/1.jpg'},'source_uid':'4124','source':'self_account'}"
	fmt.Println(umrest.BuildUrl(ak, data))
}

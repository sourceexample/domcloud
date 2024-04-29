package main

import (
	"domcloud/modHttp"
	"domcloud/modUtility"
	"fmt"
)

func main() {
	err := modUtility.Utility_initialize()
	if err != nil {
		fmt.Println("utility initialize failed: ", err)
		return
	}
	err = modHttp.Http_Initialize()
	if err != nil {
		fmt.Println("http initialize failed: ", err)
		return
	}

	err = modHttp.Http_Start()
	if err != nil {
		fmt.Println("http start failed: ", err)
	}
}

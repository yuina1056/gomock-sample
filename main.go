package main

import (
	"gomock-sample/getaccess"
	"gomock-sample/httpaccess"
)

func main() {
	ga := getaccess.NewGetAccess(httpaccess.NewTestAPI("localhost:8080"))
	ga.GetAccessFunc()
}

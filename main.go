package main

import "gomock-sample/getaccess"
import "gomock-sample/httpaccess"

func main() {
	ga := getaccess.NewGetAccess(httpaccess.NewTestAPI("localhost:8080"))
	ga.GetAccessFunc()
}

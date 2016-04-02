package main

import "fmt"
import "github.com/ufengzhu/gocurl/curl"

func myWriteFunc(data []byte, userdata interface{}) int {
	// fmt.Printf("Get data: type = %T, len = %d\n", data, len(data))
	// fmt.Printf("Data data: %v\n", string(data))
	return len(data)
}

func main() {
	// var url string = "http://www.douban.com"
	var url string = "http://www.google.com"

	easy := curl.NewEasy()
	defer easy.EasyCleanup()

	easy.EasySetopt(curl.OPT_URL, url)
	easy.EasySetopt(curl.OPT_WRITEFUNCTION, myWriteFunc)
	easy.EasySetopt(curl.OPT_VERBOSE, 1)
	// easy.EasySetopt(curl.OPT_HEADER, 1)
	// easy.EasySetopt(curl.OPT_NOPROGRESS, 0)
	// easy.EasySetopt(curl.OPT_NOSIGNAL, 1)
	// easy.EasySetopt(curl.OPT_WILDCARDMATCH, 1)
	// Only allow HTTP, TFTP and SFTP.
	easy.EasySetopt(curl.OPT_PROTOCOLS, curl.PROTO_HTTP|curl.PROTO_TFTP|curl.PROTO_SFTP)
	easy.EasySetopt(curl.OPT_HTTPHEADER, []string{"Shoesize: 10", "Accept:"})
	// err := easy.EasySetopt(245, 10)
	// if err != nil {
	// 	fmt.Printf("EasySetopt failed: %v\n", err)
	// }

	err := easy.EasyPerform()
	fmt.Printf("EasyPerform return: %v\n", err)
}

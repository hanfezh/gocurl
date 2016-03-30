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
	var ret int = 0

	var easy curl.CURL

	ret = easy.EasyInit()
	fmt.Printf("EasyInit return %d\n", ret)
	defer easy.EasyCleanup()

	easy.EasySetopt(curl.OPT_URL, url)
	easy.EasySetopt(curl.OPT_WRITEFUNCTION, myWriteFunc)
	easy.EasySetopt(curl.OPT_VERBOSE, 1)
	// curl.EasySetopt(OPT_HEADER, 1)
	// curl.EasySetopt(OPT_NOPROGRESS, 0)
	// curl.EasySetopt(OPT_NOSIGNAL, 1)
	// curl.EasySetopt(OPT_WILDCARDMATCH, 1)
	// Only allow HTTP, TFTP and SFTP.
	easy.EasySetopt(curl.OPT_PROTOCOLS, curl.PROTO_HTTP|curl.PROTO_TFTP|curl.PROTO_SFTP)
	easy.EasySetopt(curl.OPT_HTTPHEADER, []string{"Shoesize: 10", "Accept:"})

	ret = easy.EasyPerform()
	fmt.Printf("EasyPerform return %d\n", ret)
}

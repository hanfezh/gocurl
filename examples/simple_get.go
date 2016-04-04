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
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, url)
	easy.Setopt(curl.OPT_WRITEFUNCTION, myWriteFunc)
	easy.Setopt(curl.OPT_VERBOSE, 1)
	// easy.Setopt(curl.OPT_HEADER, 1)
	// easy.Setopt(curl.OPT_NOPROGRESS, 0)
	// easy.Setopt(curl.OPT_NOSIGNAL, 1)
	// easy.Setopt(curl.OPT_WILDCARDMATCH, 1)
	// Only allow HTTP, TFTP and SFTP.
	easy.Setopt(curl.OPT_PROTOCOLS, curl.PROTO_HTTP|curl.PROTO_TFTP|curl.PROTO_SFTP)
	easy.Setopt(curl.OPT_HTTPHEADER, []string{"Shoesize: 10", "Accept:"})
	// err := easy.Setopt(245, 10)
	// if err != nil {
	// 	fmt.Printf("Setopt failed: %v\n", err)
	// }

	err := easy.Perform()
	fmt.Printf("Perform return: %v\n", err)
}

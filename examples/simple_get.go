package main

import "fmt"
import "github.com/ufengzh/gocurl"

func myWriteFunc(data []byte, userdata interface{}) int {
	// fmt.Printf("Get data: type = %T, len = %d\n", data, len(data))
	// fmt.Printf("Data data: %v\n", string(data))
	return len(data)
}

func main() {
	// var url string = "http://www.douban.com"
	var url string = "http://www.google.com"

	easy := gocurl.NewEasy()
	defer easy.Cleanup()

	easy.Setopt(gocurl.OPT_URL, url)
	easy.Setopt(gocurl.OPT_WRITEFUNCTION, myWriteFunc)
	easy.Setopt(gocurl.OPT_VERBOSE, 1)
	// easy.Setopt(gocurl.OPT_HEADER, 1)
	// easy.Setopt(gocurl.OPT_NOPROGRESS, 0)
	// easy.Setopt(gocurl.OPT_NOSIGNAL, 1)
	// easy.Setopt(gocurl.OPT_WILDCARDMATCH, 1)
	// Only allow HTTP, TFTP and SFTP.
	easy.Setopt(gocurl.OPT_PROTOCOLS, gocurl.PROTO_HTTP|gocurl.PROTO_TFTP|gocurl.PROTO_SFTP)
	easy.Setopt(gocurl.OPT_HTTPHEADER, []string{"Shoesize: 10", "Accept:"})
	// err := easy.Setopt(245, 10)
	// if err != nil {
	// 	fmt.Printf("Setopt failed: %v\n", err)
	// }

	err := easy.Perform()
	fmt.Printf("Perform return: %v\n", err)
}

package main

import "fmt"
import "gocurl/curl"

func myWriteFunction(data []byte, userdata interface{}) int {
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

	easy.EasySetopt(curl.CURLOPT_URL, url)
	easy.EasySetopt(curl.CURLOPT_WRITEFUNCTION, myWriteFunction)
	easy.EasySetopt(curl.CURLOPT_VERBOSE, 1)
	// curl.EasySetopt(CURLOPT_HEADER, 1)
	// curl.EasySetopt(CURLOPT_NOPROGRESS, 0)
	// curl.EasySetopt(CURLOPT_NOSIGNAL, 1)
	// curl.EasySetopt(CURLOPT_WILDCARDMATCH, 1)
	// Only allow HTTP, TFTP and SFTP.
	easy.EasySetopt(curl.CURLOPT_PROTOCOLS, curl.CURLPROTO_HTTP|curl.CURLPROTO_TFTP|curl.CURLPROTO_SFTP)
	easy.EasySetopt(curl.CURLOPT_HTTPHEADER, []string{"Shoesize: 10", "Accept:"})

	ret = easy.EasyPerform()
	fmt.Printf("EasyPerform return %d\n", ret)
}

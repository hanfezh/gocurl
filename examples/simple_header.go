package main

import "fmt"
import "github.com/ufengzhu/gocurl/curl"

func myHeaderFunc(data []byte, userdata interface{}) int {
	fmt.Printf("Recv headers: %v", string(data))
	return len(data)
}

func myWriteFunc(data []byte, userdata interface{}) int {
	// Ignore output to stdout
	return len(data)
}

func main() {
	var url string = "http://www.google.com"

	easy := curl.NewEasy()
	defer easy.EasyCleanup()

	easy.EasySetopt(curl.OPT_URL, url)
	// easy.EasySetopt(curl.OPT_VERBOSE, 1)
	easy.EasySetopt(curl.OPT_HEADERFUNCTION, myHeaderFunc)
	easy.EasySetopt(curl.OPT_WRITEFUNCTION, myWriteFunc)

	err := easy.EasyPerform()
	fmt.Printf("EasyPerform return %v\n", err)

	eurl, err := easy.EasyGetinfo(curl.INFO_EFFECTIVE_URL)
	fmt.Printf("eurl: %v, err: %v\n", eurl, err)

	elist, err := easy.EasyGetinfo(curl.INFO_COOKIELIST)
	fmt.Printf("elist: %v, err: %v\n", elist, err)
}

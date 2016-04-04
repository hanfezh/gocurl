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
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, url)
	// easy.Setopt(curl.OPT_VERBOSE, 1)
	easy.Setopt(curl.OPT_HEADERFUNCTION, myHeaderFunc)
	easy.Setopt(curl.OPT_WRITEFUNCTION, myWriteFunc)

	err := easy.Perform()
	fmt.Printf("Perform return %v\n", err)

	eurl, err := easy.Getinfo(curl.INFO_EFFECTIVE_URL)
	fmt.Printf("eurl: %v, err: %v\n", eurl, err)

	elist, err := easy.Getinfo(curl.INFO_COOKIELIST)
	fmt.Printf("elist: %v, err: %v\n", elist, err)
}

package main

import (
	"fmt"
	"github.com/ufengzhu/gocurl/curl"
	"os"
	"time"
)

func main() {
	multi := curl.NewMulti()
	easy := curl.NewEasy()

	defer multi.Cleanup()
	defer easy.EasyCleanup()

	// easy.EasySetopt(curl.OPT_URL, "http://www.example.com/")
	easy.EasySetopt(curl.OPT_URL, "http://www.google.com/")
	easy.EasySetopt(curl.OPT_VERBOSE, 1)

	/* add the individual transfers */
	multi.AddHandle(easy)
	defer multi.RemoveHandle(easy)

	/* we start some action by calling perform right away */
	running, err := multi.Perform()
	fmt.Printf("running: %d, err: %v\n", running, err)

	for err == nil && running > 0 {
		/* wait for activity, timeout or "nothing" */
		num, err := multi.Wait(1000)
		// fmt.Printf("num: %d, err: %v\n", num, err)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Wait failed: %v\n", err)
			break
		}

		// timeout
		if num == 0 {
			time.Sleep(100 * time.Millisecond)
		}

		running, err = multi.Perform()
		// fmt.Printf("running: %d, err: %v\n", running, err)
	}
}

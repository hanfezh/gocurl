package main

import (
	"fmt"
	"github.com/ufengzh/gocurl"
	"os"
	"time"
)

func main() {
	multi := gocurl.NewMulti()
	easy := gocurl.NewEasy()

	defer multi.Cleanup()
	defer easy.Cleanup()

	// easy.Setopt(gocurl.OPT_URL, "http://www.example.com/")
	easy.Setopt(gocurl.OPT_URL, "http://www.google.com/")
	easy.Setopt(gocurl.OPT_VERBOSE, 1)

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

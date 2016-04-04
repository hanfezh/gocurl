package main

import (
	"fmt"
	"github.com/ufengzhu/gocurl"
	"time"
)

func main() {
	multi := gocurl.NewMulti()
	defer multi.Cleanup()

	easy1 := gocurl.NewEasy()
	easy2 := gocurl.NewEasy()
	defer easy1.Cleanup()
	defer easy2.Cleanup()

	easy1.Setopt(gocurl.OPT_URL, "http://www.example.com/")
	// easy1.Setopt(gocurl.OPT_URL, "http://www.bing.com/")
	easy1.Setopt(gocurl.OPT_VERBOSE, 1)

	easy2.Setopt(gocurl.OPT_URL, "http://www.google.com/")
	easy2.Setopt(gocurl.OPT_VERBOSE, 1)

	multi.AddHandle(easy1)
	multi.AddHandle(easy2)
	defer multi.RemoveHandle(easy1)
	defer multi.RemoveHandle(easy2)

	num, err := multi.Perform()
	// fmt.Printf("num: %d, err: %v\n", num, err)
	for err == nil && num > 0 {
		ms, err := multi.Timeout()
		// fmt.Printf("ms: %d, err: %v\n", ms, err)
		if err != nil {
			fmt.Printf("Timeout failed: %v\n", err)
			break
		}
		if ms <= 0 {
			// one second
			ms = 1000
		}

		fds, err := multi.Select(ms)
		if err != nil {
			fmt.Printf("Select failed: %v\n", err)
			break
		}
		if fds < 1 {
			fmt.Printf("No fd was ready to read or write, fds = %d\n", fds)
			time.Sleep(1000 * time.Millisecond)
		}

		num, err = multi.Perform()
		// fmt.Printf("num: %d, err: %v\n", num, err)
	}

	fmt.Printf("End Perform\n")
}

package main

import "fmt"
import "os"
import "github.com/ufengzhu/gocurl/curl"

type config struct {
	traceAscii bool
}

func dump(text string, stream *os.File, data []byte, nohex bool) {
	var width int = 0x10

	size := len(data)

	if nohex {
		/* without the hex output, we can fit more on screen */
		width = 0x40
	}

	fmt.Fprintf(stream, "%s, %10.10d bytes (0x%8.8x)\n", text, size, size)

	for i := 0; i < size; i += width {
		fmt.Fprintf(stream, "%4.4x: ", i)

		if !nohex {
			/* hex not disabled, show it */
			for c := 0; c < width; c++ {
				if i+c < size {
					fmt.Fprintf(stream, "%02x ", data[i+c])
				} else {
					fmt.Fprintf(stream, "   ")
				}
			}
		}

		for c := 0; (c < width) && (i+c < size); c++ {
			/* check for 0D0A; if found, skip past and start a new line of output */
			if nohex && (i+c+1 < size) && data[i+c] == 0x0D && data[i+c+1] == 0x0A {
				i += (c + 2 - width)
				break
			}

			if (data[i+c] >= 0x20) && (data[i+c] < 0x80) {
				fmt.Fprintf(stream, "%c", data[i+c])
			} else {
				fmt.Fprintf(stream, "%c", '.')
			}

			/* check again for 0D0A, to avoid an extra \n if it's at width */
			if nohex && (i+c+2 < size) && data[i+c+1] == 0x0D && data[i+c+2] == 0x0A {
				i += (c + 3 - width)
				break
			}
		}

		/* newline */
		fmt.Fprintf(stream, "\n")
	}

	stream.Sync()
}

func myTrace(info int, data []byte, userdata interface{}) int {
	var text string

	switch info {
	case curl.INFO_TEXT:
		fmt.Fprintf(os.Stderr, "== Info: %s", string(data))
		return 0

	default: /* in case a new one is introduced to shock us */
		return 0

	case curl.INFO_HEADER_OUT:
		text = "=> Send header"

	case curl.INFO_DATA_OUT:
		text = "=> Send data"

	case curl.INFO_SSL_DATA_OUT:
		text = "=> Send SSL data"

	case curl.INFO_HEADER_IN:
		text = "<= Recv header"

	case curl.INFO_DATA_IN:
		text = "<= Recv data"

	case curl.INFO_SSL_DATA_IN:
		text = "<= Recv SSL data"
	}

	conf := userdata.(*config)
	dump(text, os.Stderr, data, conf.traceAscii)
	return 0
}

func main() {
	/* enable ascii tracing */
	conf := config{traceAscii: true}

	easy := curl.NewEasy()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_DEBUGFUNCTION, myTrace)
	easy.Setopt(curl.OPT_DEBUGDATA, &conf)

	/* the DEBUGFUNCTION has no effect until we enable VERBOSE */
	easy.Setopt(curl.OPT_VERBOSE, 1)

	/* example.com is redirected, so we tell libcurl to follow redirection */
	easy.Setopt(curl.OPT_FOLLOWLOCATION, 1)

	easy.Setopt(curl.OPT_URL, "http://example.com/")
	// easy.Setopt(curl.OPT_URL, "http://www.google.com")

	err := easy.Perform()
	if err != nil {

		/* Check for errors */
		fmt.Printf("Perform failed: %v\n", err)
	}
}

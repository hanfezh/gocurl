package main

import "fmt"
import "os"
import "bufio"
import "regexp"

func handleOpt(re *regexp.Regexp, line string, output *os.File) bool {
	substr := re.FindStringSubmatch(line)
	if len(substr) != 4 {
		// fmt.Printf("FindStringSubmatch failed: %v\n", substr)
		return false
	}

	// fmt.Printf("%v, %v, %v\n", substr[1], substr[2], substr[3])
	fmt.Fprintf(output, "    CURLOPT_%-20v = C.CURLOPT_%v\n", substr[1], substr[1])
	return true
}

func handleFile(inPath string, outPath string) int {
	// CINIT(WRITEDATA, OBJECTPOINT, 1),
	optStr := `\s*CINIT\(\s*([_\dA-Z]+)\s*,\s*(LONG|OBJECTPOINT|STRINGPOINT|FUNCTIONPOINT|OFF_T)\s*,\s*([\d]+)\s*\)\s*,\s*`
	optRe, err := regexp.Compile(optStr)
	if err != nil {
		fmt.Printf("Compile failed, %v\n", err)
		return -1
	}

	// CURLE_OK = 0,
	// CURLE_UNSUPPORTED_PROTOCOL,    /* 1 */
	// protoRe, err := regexp.Compile(`^\s*CURLE_([A-Z]+)\s*,`)
	codeRe, err := regexp.Compile(`^\s*(CURLE_[_A-Z]+)\s*(=\s*[\d]+)?,`)
	if err != nil {
		fmt.Printf("Compile failed, %v\n", err)
		return -1
	}

	input, err := os.Open(inPath)
	if err != nil {
		fmt.Printf("Open failed, %v\n", err)
		return -1
	}
	reader := bufio.NewReader(input)

	var opts []string
	var codes []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if optRe.Match(line) {
			substr := optRe.FindStringSubmatch(string(line))
			if len(substr) == 4 {
				opts = append(opts, substr[1])
			}
			// fmt.Printf("%v, %v, %v\n", substr[1], substr[2], substr[3])
		} else if codeRe.Match(line) {
			substr := codeRe.FindStringSubmatch(string(line))
			codes = append(codes, substr[1])
		}
	}

	writer, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("Create failed, %v\n", err)
		return -1
	}

	// Init output
	writer.WriteString("package curl\n")
	writer.WriteString("// #include <curl/curl.h>\n")
	writer.WriteString("import \"C\"\n\n")

	// Handle CURLOPT_XXX
	writer.WriteString("// CURLOPT_XXX\n")
	writer.WriteString("const (\n")
	for _, opt := range opts {
		fmt.Fprintf(writer, "    CURLOPT_%-24v = C.CURLOPT_%v\n", opt, opt)
	}
	writer.WriteString(")\n\n")

	// Handle CURLE_XXX
	writer.WriteString("// CURLE_XXX\n")
	writer.WriteString("const (\n")
	for _, code := range codes {
		fmt.Fprintf(writer, "    %-30v = C.%v\n", code, code)
	}
	writer.WriteString(")\n")

	return 0
}

func main() {
	var inPath string = "/usr/include/curl/curl.h"
	var outPath string = "curl/curl.go"

	fmt.Printf("Args: type = %T, len = %d, value = %v\n", os.Args, len(os.Args), os.Args)
	if len(os.Args) > 1 {
		inPath = os.Args[1]
	}

	handleFile(inPath, outPath)
}

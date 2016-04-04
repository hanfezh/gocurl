package main

import "fmt"
import "os"
import "bufio"
import "regexp"

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
	// codeRe, err := regexp.Compile(`^\s*CURLE_([A-Z]+)\s*,`)
	codeRe, err := regexp.Compile(`^\s*CURL(E_[_A-Z]+)\s*(=\s*[\d]+)?,`)
	if err != nil {
		fmt.Printf("Compile failed, %v\n", err)
		return -1
	}

	// #define CURLPROTO_HTTP   (1<<0)
	protRe, err := regexp.Compile(`^\s*#define\s+CURL(PROTO_[_\dA-Z]+)\s+\(`)
	if err != nil {
		fmt.Printf("Compile failed, %v\n", err)
		return -1
	}

	// CURLINFO_EFFECTIVE_URL    = CURLINFO_STRING + 1,
	// infoRe, err := regexp.Compile(`^\s*(CURLINFO_[_\dA-Z]+)\s*=\s*CURLINFO_(STRING|LONG|DOUBLE|SLIST|SOCKET)\s*`)
	infoRe, err := regexp.Compile(`^\s*CURL(INFO_[_\dA-Z]+)\s*=\s*CURLINFO_(STRING|LONG|DOUBLE|SLIST|SOCKET)\s*\+\s*\d+,`)
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
	var protos []string
	var infos []string

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
		} else if protRe.Match(line) {
			substr := protRe.FindStringSubmatch(string(line))
			protos = append(protos, substr[1])
		} else if infoRe.Match(line) {
			substr := infoRe.FindStringSubmatch(string(line))
			infos = append(infos, substr[1])
		}
	}

	writer, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("Create failed, %v\n", err)
		return -1
	}

	// Init output
	writer.WriteString("package gocurl\n")
	writer.WriteString("// #include <curl/curl.h>\n")
	writer.WriteString("import \"C\"\n\n")

	// Handle CURLOPT_XXX
	writer.WriteString("// OPT_XXX\n")
	writer.WriteString("const (\n")
	for _, opt := range opts {
		fmt.Fprintf(writer, "    OPT_%-24v = C.CURLOPT_%v\n", opt, opt)
	}
	writer.WriteString(")\n\n")

	// Handle CURLE_XXX
	writer.WriteString("// E_XXX\n")
	writer.WriteString("const (\n")
	for _, code := range codes {
		fmt.Fprintf(writer, "    %-30v = C.CURL%v\n", code, code)
	}
	writer.WriteString(")\n\n")

	// Handle CURLPROTO_XXX
	writer.WriteString("// PROTO_XXX\n")
	writer.WriteString("const (\n")
	for _, proto := range protos {
		fmt.Fprintf(writer, "    %-30v = C.CURL%v\n", proto, proto)
	}
	writer.WriteString(")\n\n")

	// Handle CURLINFO_XXX
	writer.WriteString("// INFO_XXX\n")
	writer.WriteString("const (\n")
	for _, info := range infos {
		fmt.Fprintf(writer, "    %-30v = C.CURL%v\n", info, info)
	}
	writer.WriteString(")\n\n")

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

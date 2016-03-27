package curl

// #cgo LDFLAGS: -lcurl
/*
#include <stdio.h>
#include <stdlib.h>
#include <curl/curl.h>

static CURLcode curl_easy_setopt_long(CURL *handle, CURLoption option, long param) {
	return curl_easy_setopt(handle, option, param);
}

static CURLcode curl_easy_setopt_off_t(CURL *handle, CURLoption option, off_t param) {
	return curl_easy_setopt(handle, option, param);
}

static CURLcode curl_easy_setopt_str(CURL *handle, CURLoption option, const char *param) {
	return curl_easy_setopt(handle, option, param);
}

static CURLcode curl_easy_setopt_file(CURL *handle, CURLoption option, FILE *param) {
	return curl_easy_setopt(handle, option, param);
}

static CURLcode curl_easy_setopt_ptr(CURL *handle, CURLoption option, void *param) {
	return curl_easy_setopt(handle, option, param);
}

extern size_t goWriteCallback(char *buffer, size_t size, size_t nmemb, void *userdata);

static size_t curl_write_func_wrap(char *buffer, size_t size, size_t nmemb, void *userdata)
{
	// printf("buffer = %p, size = %lu, nmemb = %lu\n", buffer, size, nmemb);
	return goWriteCallback(buffer, size, nmemb, userdata);
}

static void *curl_write_func()
{
	return (void *)&curl_write_func_wrap;
}
*/
import "C"
import "fmt"
import "unsafe"

const (
	// CURLPROTO_ defines are for the CURLOPT_*PROTOCOLS options
	CURLPROTO_HTTP   = C.CURLPROTO_HTTP
	CURLPROTO_HTTPS  = C.CURLPROTO_HTTPS
	CURLPROTO_FTP    = C.CURLPROTO_FTP
	CURLPROTO_FTPS   = C.CURLPROTO_FTPS
	CURLPROTO_SCP    = C.CURLPROTO_SCP
	CURLPROTO_SFTP   = C.CURLPROTO_SFTP
	CURLPROTO_TELNET = C.CURLPROTO_TELNET
	CURLPROTO_LDAP   = C.CURLPROTO_LDAP
	CURLPROTO_LDAPS  = C.CURLPROTO_LDAPS
	CURLPROTO_DICT   = C.CURLPROTO_DICT
	CURLPROTO_FILE   = C.CURLPROTO_FILE
	CURLPROTO_TFTP   = C.CURLPROTO_TFTP
	CURLPROTO_IMAP   = C.CURLPROTO_IMAP
	CURLPROTO_IMAPS  = C.CURLPROTO_IMAPS
	CURLPROTO_POP3   = C.CURLPROTO_POP3
	CURLPROTO_POP3S  = C.CURLPROTO_POP3S
	CURLPROTO_SMTP   = C.CURLPROTO_SMTP
	CURLPROTO_SMTPS  = C.CURLPROTO_SMTPS
	CURLPROTO_RTSP   = C.CURLPROTO_RTSP
	CURLPROTO_RTMP   = C.CURLPROTO_RTMP
	CURLPROTO_RTMPT  = C.CURLPROTO_RTMPT
	CURLPROTO_RTMPE  = C.CURLPROTO_RTMPE
	CURLPROTO_RTMPTE = C.CURLPROTO_RTMPTE
	CURLPROTO_RTMPS  = C.CURLPROTO_RTMPS
	CURLPROTO_RTMPTS = C.CURLPROTO_RTMPTS
	CURLPROTO_GOPHER = C.CURLPROTO_GOPHER
	CURLPROTO_SMB    = C.CURLPROTO_SMB
	CURLPROTO_SMBS   = C.CURLPROTO_SMBS
	CURLPROTO_ALL    = C.CURLPROTO_ALL
)

type CURL struct {
	ptr unsafe.Pointer
	// curl_slist
	headers       []unsafe.Pointer
	writeData     interface{}
	writeFunction *func([]byte, interface{}) int
}

var curlMap = make(map[unsafe.Pointer]*CURL)

// curlMap := make(map[uintptr]*CURL)

func (curl *CURL) EasyInit() int {
	curl.ptr = C.curl_easy_init()
	fmt.Printf("curl.ptr: %T %v\n", curl.ptr, curl.ptr)
	if curl.ptr == nil {
		return -1
	}

	// curl.headers = make([]unsafe.Pointer)
	curlMap[curl.ptr] = curl
	return 0
}

func (curl *CURL) EasySetopt(opt int, arg interface{}) int {
	// if arg == nil {
	// 	return -1
	// }

	switch {
	// case opt == CURLOPT_VERBOSE:
	// 	onoff := arg.(int)
	// 	C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_VERBOSE, C.long(onoff))

	// case opt == CURLOPT_HEADER:
	// 	onoff := arg.(int)
	// 	C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_HEADER, C.long(onoff))

	// case opt == CURLOPT_NOPROGRESS:
	// 	onoff := arg.(int)
	// 	C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_NOPROGRESS, C.long(onoff))

	// case opt == CURLOPT_NOSIGNAL:
	// 	onoff := arg.(int)
	// 	C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_NOSIGNAL, C.long(onoff))

	// case opt == CURLOPT_WILDCARDMATCH:
	// 	onoff := arg.(int)
	// 	C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_WILDCARDMATCH, C.long(onoff))

	case opt == CURLOPT_WRITEDATA:
		curl.writeData = arg

	case opt == CURLOPT_WRITEFUNCTION:
		fun := arg.(func([]byte, interface{}) int)
		curl.writeFunction = &fun
		ptr := C.curl_write_func()
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_WRITEFUNCTION, ptr)
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_WRITEDATA, curl.ptr)

	case opt == CURLOPT_URL:
		url := C.CString(arg.(string))
		defer C.free(unsafe.Pointer(url))
		C.curl_easy_setopt_str(curl.ptr, C.CURLOPT_URL, url)
		// C.curl_easy_setopt_file(curl.ptr, C.CURLOPT_WRITEDATA, C.stdout)

	// case CURLOPT_PATH_AS_IS:
	// 	leaveit := arg.(int)
	// 	C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_PATH_AS_IS, C.long(leaveit))

	case opt == CURLOPT_PROTOCOLS:
		bitmask := arg.(int)
		C.curl_easy_setopt_long(curl.ptr, C.CURLOPT_PROTOCOLS, C.long(bitmask))

	case opt == CURLOPT_HTTPHEADER:
		var list *C.struct_curl_slist = nil

		headers := arg.([]string)
		if len(headers) < 1 {
			break
		}
		for _, header := range headers {
			hdr := C.CString(header)
			defer C.free(unsafe.Pointer(hdr))
			fmt.Printf("header: %T, %v\n", hdr, hdr)
			list = C.curl_slist_append(list, hdr)
		}
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_HTTPHEADER, unsafe.Pointer(list))
		curl.headers = append(curl.headers, unsafe.Pointer(list))

	case opt >= C.CURLOPTTYPE_OFF_T:
		val := C.off_t(0)
		switch arg.(type) {
		case int:
			val = C.off_t(arg.(int))
		case int64:
			val = C.off_t(arg.(int64))
		case uint64:
			val = C.off_t(arg.(uint64))
		default:
			fmt.Printf("Not implemented, %T, %v\n", arg, arg)
		}
		C.curl_easy_setopt_off_t(curl.ptr, C.CURLoption(opt), val)

	case opt >= C.CURLOPTTYPE_FUNCTIONPOINT:

	// case opt >= CURLOPTTYPE_STRINGPOINT:
	case opt >= C.CURLOPTTYPE_OBJECTPOINT:

	case opt >= C.CURLOPTTYPE_LONG:
		val := C.long(0)
		switch arg.(type) {
		case int:
			val = C.long(arg.(int))
		case bool:
			if arg.(bool) {
				val = 1
			}
		default:
			fmt.Printf("Not implemented, %T, %v\n", arg, arg)
		}
		C.curl_easy_setopt_long(curl.ptr, C.CURLoption(opt), val)

	default:
		fmt.Printf("Invalid option: %d\n", opt)
		return -1
	}

	return 0
}

func (curl *CURL) EasyPerform() int {
	fmt.Printf("%T %v\n", curl.ptr, curl.ptr)
	return int(C.curl_easy_perform(curl.ptr))
}

func (curl *CURL) EasyCleanup() {
	fmt.Printf("EasyCleanup headers: len = %d\n", len(curl.headers))
	for _, header := range curl.headers {
		fmt.Printf("EasyCleanup header: %T, %v\n", header, header)
		// C.curl_slist_free_all((*C.struct_curl_slist)header)
		C.curl_slist_free_all((*C.struct_curl_slist)(header))
	}
	C.curl_easy_cleanup(curl.ptr)
	curl.ptr = nil
}

//export goWriteCallback
func goWriteCallback(buffer *C.char, size C.size_t, nmemb C.size_t, userdata unsafe.Pointer) C.size_t {
	// fmt.Printf("userdata: %T, %v\n", userdata, userdata)
	curl := curlMap[userdata]
	// fmt.Printf("curl: %T, %v\n", curl, curl)
	buf := C.GoBytes(unsafe.Pointer(buffer), C.int(size*nmemb))
	return C.size_t((*curl.writeFunction)(buf, curl.writeData))
}

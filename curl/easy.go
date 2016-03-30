package curl

// #cgo LDFLAGS: -lcurl
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
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

static CURLcode curl_easy_setopt_ptr(CURL *handle, CURLoption option, void *param) {
	return curl_easy_setopt(handle, option, param);
}

static CURLcode curl_easy_getinfo_str(CURL *handle, CURLINFO info, char **str) {
	return curl_easy_getinfo(handle, info, str);
}

static CURLcode curl_easy_getinfo_long(CURL *handle, CURLINFO info, long *val) {
	return curl_easy_getinfo(handle, info, val);
}

static CURLcode curl_easy_getinfo_double(CURL *handle, CURLINFO info, double *val) {
	return curl_easy_getinfo(handle, info, val);
}

static CURLcode curl_easy_getinfo_slist(CURL *handle, CURLINFO info, struct curl_slist **val) {
	return curl_easy_getinfo(handle, info, val);
}

extern size_t goWriteCallback(char *buffer, size_t size, size_t nmemb, void *userdata);
extern size_t goReadCallback(char *buffer, size_t size, size_t nmemb, void *instream);
extern size_t goHeaderCallback(char *buffer, size_t size, size_t nmemb, void *userdata);

static size_t curl_write_func_wrap(char *buffer, size_t size, size_t nmemb, void *userdata)
{
	// printf("buffer = %p, size = %lu, nmemb = %lu\n", buffer, size, nmemb);
	return goWriteCallback(buffer, size, nmemb, userdata);
}

static size_t curl_read_func_wrap(char *buffer, size_t size, size_t nmemb, void *instream)
{
	// printf("buffer = %p, size = %lu, nmemb = %lu\n", buffer, size, nmemb);
	return goReadCallback(buffer, size, nmemb, instream);
}

static size_t curl_header_func_wrap(char *buffer, size_t size, size_t nmemb, void *userdata)
{
	return goHeaderCallback(buffer, size, nmemb, userdata);
}

static void *curl_write_func()
{
	return (void *)&curl_write_func_wrap;
}

static void *curl_read_func()
{
	return (void *)&curl_read_func_wrap;
}

static void *curl_header_func()
{
	return (void *)&curl_header_func_wrap;
}
*/
import "C"
import "fmt"
import "unsafe"

const (
	OPTTYPE_LONG        = C.CURLOPTTYPE_LONG
	OPTTYPE_OBJECTPOINT = C.CURLOPTTYPE_OBJECTPOINT
	// OPTTYPE_STRINGPOINT   = C.CURLOPTTYPE_STRINGPOINT
	OPTTYPE_FUNCTIONPOINT = C.CURLOPTTYPE_FUNCTIONPOINT
	OPTTYPE_OFF_T         = C.CURLOPTTYPE_OFF_T
)

const (
	INFO_STRING = C.CURLINFO_STRING
	INFO_LONG   = C.CURLINFO_LONG
	INFO_DOUBLE = C.CURLINFO_DOUBLE
	INFO_SLIST  = C.CURLINFO_SLIST
	// INFO_SOCKET  = C.CURLINFO_SOCKET
	INFO_MASK    = C.CURLINFO_MASK
	INFO_TYPEMAK = C.CURLINFO_TYPEMASK
)

type CURL struct {
	ptr unsafe.Pointer
	// curl_slist
	headers    []unsafe.Pointer
	writeData  interface{}
	readData   interface{}
	headerData interface{}
	writeFunc  *func([]byte, interface{}) int
	readFunc   *func([]byte, interface{}) int
	headerFunc *func([]byte, interface{}) int
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
	if arg == nil {
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLoption(opt), unsafe.Pointer(nil))
		return 0
	}

	switch {
	case opt == OPT_WRITEDATA:
		curl.writeData = arg

	case opt == OPT_WRITEFUNCTION:
		fun := arg.(func([]byte, interface{}) int)
		curl.writeFunc = &fun
		ptr := C.curl_write_func()
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_WRITEFUNCTION, ptr)
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_WRITEDATA, curl.ptr)

	case opt == OPT_READDATA:
		curl.readData = arg

	case opt == OPT_READFUNCTION:
		fun := arg.(func([]byte, interface{}) int)
		curl.readFunc = &fun
		ptr := C.curl_read_func()
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_READFUNCTION, ptr)
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_READDATA, curl.ptr)

	case opt == OPT_HEADERDATA:
		curl.headerData = arg

	case opt == OPT_HEADERFUNCTION:
		fun := arg.(func([]byte, interface{}) int)
		curl.headerFunc = &fun
		ptr := C.curl_header_func()
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_HEADERFUNCTION, ptr)
		C.curl_easy_setopt_ptr(curl.ptr, C.CURLOPT_HEADERDATA, curl.ptr)

	case opt >= OPTTYPE_OFF_T:
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

	case opt >= OPTTYPE_FUNCTIONPOINT:
		fmt.Printf("Not implemented, %T, %v\n", arg, arg)

	// case opt >= OPTTYPE_STRINGPOINT:
	case opt >= OPTTYPE_OBJECTPOINT:
		// OPT_URL
		switch arg.(type) {
		case string:
			cstr := C.CString(arg.(string))
			defer C.free(unsafe.Pointer(cstr))
			C.curl_easy_setopt_str(curl.ptr, C.CURLoption(opt), cstr)

		case []string:
			// e.g. OPT_HTTPHEADER
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

		default:
			fmt.Printf("Not implemented, %T, %v\n", arg, arg)
		}

	case opt >= OPTTYPE_LONG:
		// OPT_VERBOSE
		// OPT_HEADER
		// OPT_NOPROGRESS
		// OPT_NOSIGNAL
		// OPT_WILDCARDMATCH
		// OPT_PROTOCOLS
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

func (curl *CURL) EasyGetinfo(info int) (ret interface{}, err error) {
	switch info & INFO_TYPEMAK {
	case INFO_STRING:
		var str *C.char
		code := C.curl_easy_getinfo_str(curl.ptr, C.CURLINFO(info), &str)
		if code == C.CURLE_OK {
			return C.GoString(str), nil
		}

	case INFO_LONG:
		var val C.long
		code := C.curl_easy_getinfo_long(curl.ptr, C.CURLINFO(info), &val)
		if code == C.CURLE_OK {
			return int(val), nil
		}

	case INFO_DOUBLE:
		var val C.double
		code := C.curl_easy_getinfo_double(curl.ptr, C.CURLINFO(info), &val)
		if code == C.CURLE_OK {
			return float64(val), nil
		}

	case INFO_SLIST:
		var list *C.struct_curl_slist = nil
		code := C.curl_easy_getinfo_slist(curl.ptr, C.CURLINFO(info), &list)
		if code == C.CURLE_OK {
			var tmp *C.struct_curl_slist = list
			var ret []string

			for tmp != nil {
				ret = append(ret, C.GoString(tmp.data))
				tmp = tmp.next
			}
			C.curl_slist_free_all(list)

			return ret, nil
		}

	// case INFO_SOCKET:

	default:
		err = fmt.Errorf("Invalid info: %d", info)
		return nil, err
	}

	return nil, fmt.Errorf("Failed to get info: %d", info)
}

//export goWriteCallback
func goWriteCallback(buffer *C.char, size C.size_t, nmemb C.size_t, userdata unsafe.Pointer) C.size_t {
	// fmt.Printf("userdata: %T, %v\n", userdata, userdata)
	curl := curlMap[userdata]
	// fmt.Printf("curl: %T, %v\n", curl, curl)
	buf := C.GoBytes(unsafe.Pointer(buffer), C.int(size*nmemb))
	return C.size_t((*curl.writeFunc)(buf, curl.writeData))
}

//export goReadCallback
func goReadCallback(buffer *C.char, size C.size_t, nmemb C.size_t, instream unsafe.Pointer) C.size_t {
	curl := curlMap[instream]
	// fmt.Printf("curl: %T, %v\n", curl, curl)
	var buf []byte
	len := (*curl.readFunc)(buf, curl.readData)
	str := C.CString(string(buf))
	defer C.free(unsafe.Pointer(str))
	C.memcpy(unsafe.Pointer(buffer), unsafe.Pointer(str), C.size_t(len))
	return C.size_t(len)
}

//export goHeaderCallback
func goHeaderCallback(buffer *C.char, size C.size_t, nmemb C.size_t, userdata unsafe.Pointer) C.size_t {
	// fmt.Printf("userdata: %T, %v\n", userdata, userdata)
	curl := curlMap[userdata]
	// fmt.Printf("curl: %T, %v\n", curl, curl)
	buf := C.GoBytes(unsafe.Pointer(buffer), C.int(size*nmemb))
	return C.size_t((*curl.headerFunc)(buf, curl.headerData))
}

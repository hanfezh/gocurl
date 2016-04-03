package curl

// #cgo CFLAGS: -I/usr/include
// #cgo LDFLAGS: -lcurl
// #include <curl/curl.h>
import "C"

// import "fmt"
import "unsafe"

type VersionInfoData struct {
	Age           int    /* age of the returned struct */
	Version       string /* LIBCURL_VERSION */
	VersionNum    uint   /* LIBCURL_VERSION_NUM */
	Host          string /* OS/host/cpu/machine when configured */
	Features      int    /* bitmask, see defines below */
	SslVersion    string /* human readable string */
	SslVersionNum int    /* not used anymore, always 0 */
	LibzVersion   string /* human readable string */
	/* protocols is terminated by an entry with a NULL protoname */
	Protocols []string

	/* The fields below this were added in CURLVERSION_SECOND */
	Ares    string
	AresNum int

	/* This field was added in CURLVERSION_THIRD */
	Libidn string

	/* These field were added in CURLVERSION_FOURTH */

	/* Same as '_libiconv_version' if built with HAVE_ICONV */
	IconvVerNum int

	LibsshVersion string /* human readable string */
}

func Version() string {
	cstr := C.curl_version()
	return C.GoString(cstr)
}

func VersionInfo() VersionInfoData {
	info := C.curl_version_info(C.CURLVERSION_NOW)

	data := VersionInfoData{}
	data.Age = int(info.age)
	data.Version = C.GoString(info.version)
	data.VersionNum = uint(info.version_num)
	data.Host = C.GoString(info.host)
	data.Features = int(info.features)
	data.SslVersion = C.GoString(info.ssl_version)
	data.SslVersionNum = int(info.ssl_version_num)
	data.LibzVersion = C.GoString(info.libz_version)

	ptr := info.protocols
	for *ptr != nil {
		data.Protocols = append(data.Protocols, C.GoString(*ptr))
		// fmt.Printf("ptr: %p %p %d %d\n", ptr, *ptr, unsafe.Sizeof(ptr), unsafe.Sizeof(*ptr))
		ptr = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(ptr)))
	}

	data.Ares = C.GoString(info.ares)
	data.AresNum = int(info.ares_num)
	data.Libidn = C.GoString(info.libidn)
	data.IconvVerNum = int(info.iconv_ver_num)
	data.LibsshVersion = C.GoString(info.libssh_version)

	return data
}

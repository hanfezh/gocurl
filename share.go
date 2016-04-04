package gocurl

// #cgo CFLAGS: -I/usr/include
// #cgo LDFLAGS: -lcurl
/*
#include <curl/curl.h>

static CURLSHcode curl_share_setopt_long(CURL *handle, CURLSHoption option, long param) {
	return curl_share_setopt(handle, option, param);
}

static CURLSHcode curl_share_setopt_ptr(CURL *handle, CURLSHoption option, void *param) {
	return curl_share_setopt(handle, option, param);
}
*/
import "C"
import "fmt"
import "unsafe"

const (
	SHOPT_NONE       = C.CURLSHOPT_NONE       /* don't use */
	SHOPT_SHARE      = C.CURLSHOPT_SHARE      /* specify a data type to share */
	SHOPT_UNSHARE    = C.CURLSHOPT_UNSHARE    /* specify which data type to stop sharing */
	SHOPT_LOCKFUNC   = C.CURLSHOPT_LOCKFUNC   /* pass in a 'curl_lock_function' pointer */
	SHOPT_UNLOCKFUNC = C.CURLSHOPT_UNLOCKFUNC /* pass in a 'curl_unlock_function' pointer */
	SHOPT_USERDATA   = C.CURLSHOPT_USERDATA   /* pass in a user data pointer used in the lock/unlock callback functions */
	SHOPT_LAST       = C.CURLSHOPT_LAST       /* never use */
)

// Different data locks for a single share.
const (
	LOCK_DATA_NONE = C.CURL_LOCK_DATA_NONE
	/**
	 * CURL_LOCK_DATA_SHARE is used internally to say that
	 * the locking is just made to change the internal state of the share
	 * itself.
	 */
	LOCK_DATA_SHARE       = C.CURL_LOCK_DATA_SHARE
	LOCK_DATA_COKKIE      = C.CURL_LOCK_DATA_COOKIE
	LOCK_DATA_DNS         = C.CURL_LOCK_DATA_DNS
	LOCK_DATA_SSL_SESSION = C.CURL_LOCK_DATA_SSL_SESSION
	LOCK_DATA_CONNECT     = C.CURL_LOCK_DATA_CONNECT
	LOCK_DATA_LAST        = C.CURL_LOCK_DATA_LAST
)

type Share struct {
	handle unsafe.Pointer
}

type ShareError C.CURLSHcode

func (code ShareError) Error() string {
	str := C.GoString(C.curl_share_strerror(C.CURLSHcode(code)))
	fmt.Printf("Share error[%d]: %s\n", code, str)
	return fmt.Sprintf("Share error[%d]: %s", code, str)
}

func codeToShError(code C.CURLSHcode) error {
	if code != C.CURLSHE_OK {
		return ShareError(code)
	}

	return nil
}

func NewShare() *Share {
	sh := C.curl_share_init()
	if sh == nil {
		return nil
	}

	return &Share{handle: sh}
}

func (share *Share) Setopt(opt int, arg interface{}) error {
	if arg == nil {
		ret := C.curl_share_setopt_ptr(share.handle, C.CURLSHoption(opt), unsafe.Pointer(nil))
		return codeToShError(ret)
	}

	// TODO
	// case SHOPT_LOCKFUNC:
	// case SHOPT_UNLOCKFUNC:
	// case SHOPT_USERDATA:

	switch opt {
	case SHOPT_SHARE, SHOPT_UNSHARE:
		switch arg.(type) {
		case int:
			val := C.long(arg.(int))
			ret := C.curl_share_setopt_long(share.handle, C.CURLSHoption(opt), val)
			return codeToShError(ret)

		default:
			return fmt.Errorf("Not support arg: %d, %T, %v", opt, arg, arg)
		}

	default:
		return fmt.Errorf("Not supported opt: %d", opt)
	}

	return nil
}

func (share *Share) Cleanup() error {
	ret := C.curl_share_cleanup(share.handle)
	return codeToShError(ret)
}

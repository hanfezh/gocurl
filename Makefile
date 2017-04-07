# Default install directory of libcurl
CFLAGS   := -I/usr/include
LDFLAGS  := -lcurl
# ufeng's install directory of libcurl
# CFLAGS   := -I/home/ufeng/local/include
# LDFLAGS  := -L/home/ufeng/local/lib -lcurl
PACKAGE  := github.com/ufengzh/gocurl

all: gocurl

gocurl:
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go build $(PACKAGE)

test:
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go test $(PACKAGE) -o gocurl_test

install:
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go install $(PACKAGE) 

demo:
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go run examples/simple.go
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go run examples/debug.go

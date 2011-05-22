include $(GOROOT)/src/Make.inc

TARG=github.com/davecheney/xattr
GOFILES=\
	xattr.go\
	xattr_$(GOOS).go\
	syscall_$(GOOS).go\

include $(GOROOT)/src/Make.pkg

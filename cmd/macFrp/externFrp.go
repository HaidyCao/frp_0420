package main

/*
 */
import "C"
import (
	"github.com/HaidyCao/frp_0420/cmd/frpc/sub"
	"github.com/HaidyCao/frp_0420/cmd/frps/frps"
	"github.com/HaidyCao/golib/crypto"
)

//export RunFrps
func RunFrps(cfgFilePath *C.char) C.int {
	path := C.GoString(cfgFilePath)
	crypto.DefaultSalt = "frp"

	if err := frps.RunFrps(path); err != nil {
		println(err.Error())
		return C.int(0)
	}
	return C.int(1)
}

//export RunFrpc
func RunFrpc(cfgFilePath *C.char) C.int {
	path := C.GoString(cfgFilePath)
	crypto.DefaultSalt = "frp"

	if err := sub.RunFrpc(path); err != nil {
		println(err.Error())
		return C.int(0)
	}
	return C.int(1)
}

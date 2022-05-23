package pack1

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include "number.h"
import "C"

//export add
func add(a, b C.int) C.int {
	return a + b
}

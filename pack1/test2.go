package pack1

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include "number.h"
import "C"
import "fmt"

func Test2() {
	v0, err0 := C.number_div(10, 2)
	fmt.Println(v0, err0)
}

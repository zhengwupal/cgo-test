package pack1

// #include <stdio.h>
// #include <stdlib.h>
// extern int goAdd(int, int);
// extern char* returnString(void);
//
// static char* cString(void) {
//     char *ci = returnString();
//     printf("ci: %s\n", ci); // 字符串用%s和字符串头指针
//     return ci;
// }
//
// static int* cAdd(int a, int b) {
//     int i = goAdd(a, b);
//     printf("i: %d\n", i);
//     int *p = &i;
//     return p;
// }
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/zhengwupal/cgo-test/logs"
)

//export goAdd
func goAdd(a1, b1 C.int) C.int {
	c := a1 + b1
	fmt.Println("2222", reflect.TypeOf(c))
	// fmt.Println("3", reflect.TypeOf(&c))
	// fmt.Println("3.1", reflect.TypeOf((*C.int)(unsafe.Pointer(&c))))
	return c
}

//export returnString
func returnString() *C.char {
	gostring := "hello world11111111111111111111"
	fmt.Println("s", reflect.TypeOf(C.CString(gostring)))
	logs.SugarLogger.Infof("test7 returnString gostring: %s", gostring)
	return C.CString(gostring)
}

func Test7() {
	co := C.cString()
	cos := C.GoString(co)
	C.free(unsafe.Pointer(co))
	fmt.Println(cos)

	logs.SugarLogger.Infof("test7 cos: %s", cos)

	var a, b int = 5, 6
	fmt.Println("1", reflect.TypeOf(a))
	fmt.Println("1.1", reflect.TypeOf(C.int(a)))
	// var i *C.int = C.cAdd(C.int(a), C.int(b))
	i := C.cAdd(C.int(a), C.int(b))
	// fmt.Println("4", reflect.TypeOf(int(i))) // c的int类型强转go的int类型
	fmt.Println("4", reflect.TypeOf(i))
	ii := (*int32)(unsafe.Pointer(i))
	fmt.Println("5", reflect.TypeOf(ii))
	fmt.Println(*ii)

	// fmt.Println("4", reflect.TypeOf(i))
	// fmt.Println("5", reflect.TypeOf((*int)(unsafe.Pointer(i))))
	// fmt.Println("6", reflect.TypeOf(*(*int)(unsafe.Pointer(i))))
	// ab := *(*int)(unsafe.Pointer(i))
	// fmt.Println(ab)
}

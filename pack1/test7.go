package pack1

// #include <stdio.h>
// extern int* goAdd(int, int);
// extern char* returnString(void);
//
// static char* cString(void) {
//     char *ci = returnString();
//     printf("ci: %s\n", ci); // 字符串用%s和字符串头指针
//     return ci;
// }
//
// static int* cAdd(int a, int b) {
//     int *i = goAdd(a, b);
//     printf("i: %d\n", *i);
//     return i;
// }
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

//export goAdd
func goAdd(a, b C.int) *C.int {
	c := a + b
	fmt.Println("2", reflect.TypeOf(c))
	fmt.Println("3", reflect.TypeOf(&c))
	fmt.Println("3.1", reflect.TypeOf((*C.int)(unsafe.Pointer(&c))))
	return &c
}

//export returnString
func returnString() *C.char {
	gostring := "hello world11111111111111111111"
	fmt.Println("s", reflect.TypeOf(C.CString(gostring)))
	return C.CString(gostring)
}

func Test7() {
	co := C.cString()
	fmt.Println(C.GoString(co))

	var a, b int = 5, 6
	fmt.Println("1", reflect.TypeOf(a))
	fmt.Println("1.1", reflect.TypeOf(C.int(a)))
	var i *C.int = C.cAdd(C.int(a), C.int(b))
	// i := C.cAdd(C.int(a), C.int(b))
	// fmt.Println("4", reflect.TypeOf(int(i))) // c的int类型强转go的int类型
	// fmt.Println(i)
	fmt.Println("4", reflect.TypeOf(i))
	fmt.Println("5", reflect.TypeOf((*int)(unsafe.Pointer(i))))
	fmt.Println("6", reflect.TypeOf(*(*int)(unsafe.Pointer(i))))
	fmt.Println(*(*int)(unsafe.Pointer(i)))
}

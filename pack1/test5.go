package pack1

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include <stdlib.h>
//#include "number.h"
import "C"
import (
	"fmt"
	"sync"
	"syscall"
	"unsafe"

	"github.com/zhengwupal/cgo-test/logs"
)

//export go_debug_log
func go_debug_log(msg *C.char) {
	LOG_STD("C function, ", C.GoString(msg))
	logs.SugarLogger.Infof("Test5 go_debug_log msg: %s", C.GoString(msg))
}

//export go_debug_log_char
func go_debug_log_char(msg *C.char, arg *C.char) {
	logs.SugarLogger.Infof("Test5 go_debug_log_char msg: %s --- %s", C.GoString(msg), C.GoString(arg))
}

func produce(jobs chan<- int, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := C.int(idx)
	v0, err0 := C.number_div(10, i)
	// v1, err1 := C.number_div(10, 2)
	cs := C.CString("hello")
	result := C.foo(cs)
	str := C.GoString(result)
	C.free(unsafe.Pointer(cs))
	C.free(unsafe.Pointer(result))

	if err0 != nil {
		fmt.Printf("Producer %v sending message: %d -- %s %s\n", idx, v0, err0, str)
	}
	fmt.Printf("Producer %v sending message: %d %s\n", idx, v0, str)
	jobs <- idx
}

func consume(jobs <-chan int, done chan<- bool) {
	for msg := range jobs {
		fmt.Printf("Consumed message \"%v\"\n", msg)
		logs.SugarLogger.Infof("Test5 produce %d End", msg)
	}
	done <- true
}

type GlobalConfig struct {
	PageSize     int
	MaxOpenFiles int
}

const (
	d_DEFAULT_PAGE_SIZE      = 4096
	d_DEFAULT_MAX_OPEN_FILES = 1024
)

func getDefaultGlobalConfig() GlobalConfig {
	return GlobalConfig{
		PageSize:     d_DEFAULT_PAGE_SIZE,
		MaxOpenFiles: d_DEFAULT_MAX_OPEN_FILES,
	}
}

func InitMaxOpenFiles() bool {
	return SetMaxOpenFiles(uint64(getDefaultGlobalConfig().MaxOpenFiles), uint64(getDefaultGlobalConfig().MaxOpenFiles))
}

func SetMaxOpenFiles(max, cur uint64) bool {
	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		LOG_STD("Error Getting Rlimit ", err)
		return false
	}
	LOG_STD(rLimit)
	DBG("Rlimit Start", rLimit)

	rLimit.Max = max
	rLimit.Cur = cur
	LOG_STD(rLimit.Max, rLimit.Cur)

	DBG("Set success, Rlimit Final", rLimit)
	return true
}

func Init() bool {
	SetLogName("log.server")
	err := InitLog(0)
	if err != nil {
		LOG_STD("Init log failed, error: ", err)
		return false
	}

	if !InitMaxOpenFiles() {
		LOG_STD("Set max open files failed.")
		return false
	}

	return true
}

func Test5() {
	if !Init() {
		ERR("Init failed.")
		return
	}

	logs.SugarLogger.Infof("Test5 Start")

	jobs := make(chan int, 100)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		logs.SugarLogger.Infof("Test5 produce %d Start", i)
		go produce(jobs, i, &wg)
	}

	go consume(jobs, done)

	wg.Wait()
	close(jobs)
	<-done
}

package oanic

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
)

// 产生panic时的调用栈打印
func PrintPanicStack() {
	if x := recover(); x != nil {
		log.Error(x)
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			log.Errorf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}

	}
}

func pp() (err error) {
	fmt.Println("pp")
	err = fmt.Errorf("%s", "我死了")
	panic(err)
	//return
}

func TestPanic(t *testing.T) {
	go func() {
		defer PrintPanicStack()
		i := 10
		for i != 0 {
			fmt.Println(i)
			fmt.Println(pp())
			<-time.After(1 * time.Second)
			i--
		}
	}()
	for {

	}
}

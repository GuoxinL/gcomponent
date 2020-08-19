/*
   Created by guoxin in 2020/1/10 2:15 下午
*/
package tools

import (
	"errors"
	"fmt"
	"github.com/GuoxinL/gcomponent/components/logging"
	"reflect"
	"runtime/debug"
	"time"
)

func Catch() {
	if err := recover(); err != nil {
		// 获得协程Id
		//buf := make([]byte, 1<<16)
		//runtime.Stack(buf, true)
		//获取上层调用者信息
		//str := fmt.Sprintf("异常堆栈信息：\n%v\n",string(debug.Stack()))
		//str += fmt.Sprintf("异常堆栈信息：%v\n", err)
		//for skip := 0; ; skip++ {
		//	if pc, file, line, ok := runtime.Caller(skip); !ok {
		//		break
		//	} else {
		//		funcInfo := runtime.FuncForPC(pc)
		//		funcName := "unknown"
		//		if funcInfo != nil {
		//			funcName = funcInfo.Name()
		//		}
		//		str += fmt.Sprintf("%v()\n\t%v:%v\n", funcName, file, line)
		//	}
		//}
		logging.Error0("异常堆栈信息：\n%v\n", string(debug.Stack()))
	}
	time.Sleep(1 * time.Second)
}

type TryCatch struct {
	errChan chan interface {
	}
	catches      map[reflect.Type]func(err error)
	defaultCatch func(err error)
}

func (t TryCatch) Try(block func()) TryCatch {
	t.errChan = make(chan interface{})
	t.catches = map[reflect.Type]func(err error){}
	t.defaultCatch = func(err error) {}
	go func() {
		defer func() {
			t.errChan <- recover()
		}()
		block()
	}()
	return t
}

func (t TryCatch) CatchAll(block func(err error)) TryCatch {
	t.defaultCatch = block
	return t
}

func (t TryCatch) Catch(e error, block func(err error)) TryCatch {
	errorType := reflect.TypeOf(e)
	t.catches[errorType] = block
	return t
}

func (t TryCatch) Run() {
	t.Finally(func() {})
}

/**
如果不需要Finally()请使用Run()
*/
func (t TryCatch) Finally(block func()) TryCatch {
	err := <-t.errChan
	if err != nil {
		of := reflect.TypeOf(err)
		catch := t.catches[of]
		kind := of.Kind()
		fmt.Println(kind)
		switch kind {
		case reflect.String:
			err2 := errors.New(err.(string))
			if catch != nil {
				catch(err2)
			} else {
				t.defaultCatch(err2)
			}
		default:
			if catch != nil {
				catch(err.(error))
			} else {
				t.defaultCatch(err.(error))
			}
		}
		logging.Error0("Panic Message: %v, Stack: %v\n", err, string(debug.Stack()))
	}
	block()
	return t
}

type ExampleError struct {
	error
}

package exReflect

import (
	"fmt"
	"reflect"
	"runtime"
	cstring "superplace/extend/string"
	cerr "superplace/logger/error"
)

var (
	nilFuncInfo = FuncInfo{}
)

type FuncInfo struct {
	Name       string
	Type       reflect.Type
	Value      reflect.Value
	InArgs     []reflect.Type
	InArgsLen  int
	OutArgs    []reflect.Type
	OutArgsLen int
}

func GetFuncInfo(fn interface{}) (FuncInfo, error) {
	if fn == nil {
		return nilFuncInfo, cerr.FuncIsNil
	}

	typ := reflect.TypeOf(fn)

	if typ.Kind() != reflect.Func {
		return nilFuncInfo, cerr.FuncTypeError
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	funcName := cstring.CutLastString(fullName, ".", "-")
	var inArgs []reflect.Type
	for i := 0; i < typ.NumIn(); i++ {
		t := typ.In(i)
		inArgs = append(inArgs, t)
	}

	var outArgs []reflect.Type
	for i := 0; i < typ.NumOut(); i++ {
		t := typ.Out(i)
		outArgs = append(outArgs, t)
	}

	funcInfo := FuncInfo{
		Name:       funcName,
		Type:       typ,
		Value:      reflect.ValueOf(fn),
		InArgs:     inArgs,
		InArgsLen:  typ.NumIn(),
		OutArgs:    outArgs,
		OutArgsLen: typ.NumOut(),
	}

	return funcInfo, nil
}
func GetFuncName(fn interface{}) string {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic(fmt.Sprintf("[fn = %v] is not func type.", fn))
	}

	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return cstring.CutLastString(fullName, ".", "-")
}

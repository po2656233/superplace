package exReflect

import (
	"fmt"
	cstring "github.com/po2656233/superplace/extend/string"
	cerr "github.com/po2656233/superplace/logger/error"
	"path"
	"reflect"
	"runtime"
	"strings"
)

var (
	nilFuncInfo = FuncInfo{}
)

type FuncInfo struct {
	Name       string
	FileName   string
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
	fileName := cstring.CutLastString(fullName, ".", ".")
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
		FileName:   fileName,
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

func GetPackName(face interface{}) string {
	pkName := reflect.TypeOf(face).PkgPath()
	list := strings.Split(pkName, "/")
	return list[len(list)-1]
}
func GetFuncPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}
	return absPath
}

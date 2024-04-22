package exReflect

import (
	"fmt"
	"reflect"
)

func ReflectTry(f reflect.Value, args []reflect.Value, handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("-------------panic recover---------------")
			if handler != nil {
				handler(err)
			}
		}
	}()
	f.Call(args)
}

func GetStructName(v interface{}) string {
	return reflect.Indirect(reflect.ValueOf(v)).Type().Name()
}

////GetInvokeFunc reflect function convert to MethodInfo
//func GetInvokeFunc(name string, fn facade{}) (*face.MethodInfo, error) {
//	if name == "" {
//		return nil, cerr.Error("func name is nil")
//	}
//
//	if fn == nil {
//		return nil, cerr.Errorf("func is nil. name = %s", name)
//	}
//
//	typ := reflect.TypeOf(fn)
//	val := reflect.ValueOf(fn)
//
//	if typ.Kind() != reflect.Func {
//		return nil, cerr.Errorf("name = %s is not func type.", name)
//	}
//
//	var inArgs []reflect.Type
//	for i := 0; i < typ.NumIn(); i++ {
//		t := typ.In(i)
//		inArgs = append(inArgs, t)
//	}
//
//	var outArgs []reflect.Type
//	for i := 0; i < typ.NumOut(); i++ {
//		t := typ.Out(i)
//		outArgs = append(outArgs, t)
//	}
//
//	invoke := &face.MethodInfo{
//		Type:    typ,
//		Value:   val,
//		InArgs:  inArgs,
//		OutArgs: outArgs,
//	}
//
//	return invoke, nil
//}

func IsPtr(val interface{}) bool {
	if val == nil {
		return false
	}

	return reflect.TypeOf(val).Kind() == reflect.Ptr
}

func IsNotPtr(val interface{}) bool {
	if val == nil {
		return false
	}

	return reflect.TypeOf(val).Kind() != reflect.Ptr
}

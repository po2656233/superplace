// Package exSlice code from: https://github.com/beego/beego/blob/develop/core/utils/slice.go
package exSlice

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	cstring "github/po2656233/superplace/extend/string"
	cutils "github/po2656233/superplace/extend/utils"
)

func Int32In(v int32, sl []int32) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

func Int64In(v int64, sl []int64) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

// StringIn checks given string in string slice or not.
func StringIn(v string, sl []string) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return 0, false
}

func StringInSlice(v string, sl []string) bool {
	_, ok := StringIn(v, sl)
	return ok
}

// InInterface checks given facade in facade slice.
func InInterface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// RandList generate an int slice from min to max.
func RandList(minValue, maxValue int) []int {
	if maxValue < minValue {
		minValue, maxValue = maxValue, minValue
	}

	length := maxValue - minValue + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += minValue
	}
	return list
}

// Merge merges facade slices to one slice.
func Merge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

// Reduce generates a new slice after parsing every value by reduce function
func Reduce(slice []interface{}, a func(interface{}) interface{}) (destSlice []interface{}) {
	for _, v := range slice {
		destSlice = append(destSlice, a(v))
	}
	return
}

// Rand returns random one from slice.
func Rand(a []interface{}) (b interface{}) {
	randNum := rand.Intn(len(a))
	b = a[randNum]
	return
}

// Sum sums all values in int64 slice.
func Sum(intslice []int64) (sum int64) {
	for _, v := range intslice {
		sum += v
	}
	return
}

// Filter generates a new slice after filter function.
func Filter(slice []interface{}, a func(interface{}) bool) (filterSlice []interface{}) {
	for _, v := range slice {
		if a(v) {
			filterSlice = append(filterSlice, v)
		}
	}
	return
}

// Diff returns diff slice of slice1 - slice2.
func Diff(slice1, slice2 []interface{}) (diffSlice []interface{}) {
	for _, v := range slice1 {
		if !InInterface(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// Intersect returns slice that are present in all the slice1 and slice2.
func Intersect(slice1, slice2 []interface{}) (diffSlice []interface{}) {
	for _, v := range slice1 {
		if InInterface(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

// Chunk separates one slice to some sized slice.
func Chunk(slice []interface{}, size int) (chunkSlice [][]interface{}) {
	if size >= len(slice) {
		chunkSlice = append(chunkSlice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkSlice = append(chunkSlice, slice[i:end])
		end += size
	}
	return
}

// Range generates a new slice from begin to end with step duration of int64 number.
func Range(start, end, step int64) (intSlice []int64) {
	for i := start; i <= end; i += step {
		intSlice = append(intSlice, i)
	}
	return
}

// Pad prepends size number of val into slice.
func Pad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

func Uniques[T comparable](slices ...[]T) []T {
	keys := map[T]struct{}{}

	for _, slice := range slices {
		for _, s := range slice {
			keys[s] = struct{}{}
		}
	}

	var uniqueSlice []T

	for t := range keys {
		uniqueSlice = append(uniqueSlice, t)
	}

	return uniqueSlice
}

// Unique cleans repeated values in slice.
func Unique[T comparable](slice ...T) []T {
	return Uniques[T](slice)
}

// Shuffle shuffles a slice.
func Shuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

func StringToInt(strSlice []string) []int {
	var intSlice []int

	for _, s := range strSlice {
		if cutils.IsNumeric(s) {
			val, ok := cstring.ToInt(s)
			if ok {
				intSlice = append(intSlice, val)
			}
		}
	}

	return intSlice
}

func StringToInt32(strSlice []string) []int32 {
	var intSlice []int32

	for _, s := range strSlice {
		if cutils.IsNumeric(s) {
			val, ok := cstring.ToInt32(s)
			if ok {
				intSlice = append(intSlice, val)
			}
		}
	}

	return intSlice
}

func StringToInt64(strSlice []string) []int64 {
	var intSlice []int64

	for _, s := range strSlice {
		if cutils.IsNumeric(s) {
			val, ok := cstring.ToInt64(s)
			if ok {
				intSlice = append(intSlice, val)
			}
		}
	}

	return intSlice
}

// IsSlice checks whether given value is array/slice.
// Note that it uses reflect internally implementing this feature.
func IsSlice(value interface{}) bool {
	rv := reflect.ValueOf(value)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

func IsEmptyWithString(p []string) bool {
	for _, s := range p {
		if strings.TrimSpace(s) == "" {
			return true
		}
	}
	return false
}

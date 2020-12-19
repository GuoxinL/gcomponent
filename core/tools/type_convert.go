/*
   Created by guoxin in 2020/11/1 12:08 上午
*/
package tools

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

type StrConvert func(in string) (out interface{}, err error)

// TypeConvert container
var TypeConvert map[reflect.Kind]StrConvert

func init() {
    TypeConvert = make(map[reflect.Kind]StrConvert)
    TypeConvert[reflect.Bool] = BoolConvert
    TypeConvert[reflect.Int] = IntConvert
    TypeConvert[reflect.Int8] = Int8Convert
    TypeConvert[reflect.Int16] = Int16Convert
    TypeConvert[reflect.Int32] = Int32Convert
    TypeConvert[reflect.Int64] = Int64Convert
    TypeConvert[reflect.Uint] = UintConvert
    TypeConvert[reflect.Uint8] = Uint8Convert
    TypeConvert[reflect.Uint16] = Uint16Convert
    TypeConvert[reflect.Uint32] = Uint32Convert
    TypeConvert[reflect.Uint64] = Uint64Convert
    TypeConvert[reflect.Float32] = Float32Convert
    TypeConvert[reflect.Float64] = Float64Convert

    TypeConvert[reflect.String] = StringConvert

    TypeConvert[reflect.Slice] = SliceConvert
    TypeConvert[reflect.Array] = ArrayConvert
    TypeConvert[reflect.Map] = MapConvert
    TypeConvert[reflect.Struct] = StructConvert

    // unimplemented
    TypeConvert[reflect.Uintptr] = UintptrConvert
    TypeConvert[reflect.Complex64] = Complex64Convert
    TypeConvert[reflect.Complex128] = Complex128Convert
    TypeConvert[reflect.Chan] = ChanConvert
    TypeConvert[reflect.Func] = FuncConvert
    TypeConvert[reflect.Interface] = InterfaceConvert
    TypeConvert[reflect.Ptr] = PtrConvert

}

func BoolConvert(in string) (out interface{}, err error) {
    return parseBool(in)
}

func IntConvert(in string) (out interface{}, err error) {
    return strconv.Atoi(in)
}

func Int8Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseInt(in, 0, 8)
    out = int8(r)
    return
}

func Int16Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseInt(in, 0, 16)
    out = int16(r)
    return
}

func Int32Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseInt(in, 0, 32)
    out = int32(r)
    return
}

func Int64Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseInt(in, 0, 64)
    out = int64(r)
    return
}

func UintConvert(in string) (out interface{}, err error) {
    r, err := strconv.ParseUint(in, 0, 32)
    out = uint(r)
    return
}

func Uint8Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseUint(in, 0, 8)
    out = uint8(r)
    return
}

func Uint16Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseUint(in, 0, 16)
    out = uint16(r)
    return
}

func Uint32Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseUint(in, 0, 32)
    out = uint32(r)
    return
}

func Uint64Convert(in string) (out interface{}, err error) {
    r, err := strconv.ParseUint(in, 0, 64)
    out = uint64(r)
    return
}

func UintptrConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func Float32Convert(in string) (out interface{}, err error) {
    float, err := strconv.ParseFloat(in, 32)
    out = float32(float)
    return
}

func Float64Convert(in string) (out interface{}, err error) {
    float, err := strconv.ParseFloat(in, 64)
    out = float64(float)
    return
}

func Complex64Convert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func Complex128Convert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func ArrayConvert(in string) (out interface{}, err error) {
    out = strings.Split(in, `,`)
    return
}

func ChanConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func FuncConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func InterfaceConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func MapConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func PtrConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func SliceConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func StringConvert(in string) (out interface{}, err error) {
    out = in
    return
}

func StructConvert(in string) (out interface{}, err error) {
    panic("unimplemented")
}

func parseBool(str string) (value bool, err error) {
    switch str {
    case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "y", "ON", "on", "On":
        return true, nil
    case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "n", "OFF", "off", "Off":
        return false, nil
    }
    return false, fmt.Errorf("parsing \"%s\": invalid syntax", str)
}

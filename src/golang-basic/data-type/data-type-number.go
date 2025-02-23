package datatype

import (
	"fmt"
	"math"
)

/*
Alias
byte -> uint8
rune -> int32
int -> Minimal int32
uint -> Minimal uint32
*/

func GetTypeNumber() string {
	return fmt.Sprint(
		"Int8 : Min ", math.MinInt8, ", Max ", math.MaxInt8, "\n",
		"Int16: Min ", math.MinInt16, ", Max ", math.MaxInt16, "\n",
		"Int32: Min ", math.MinInt32, ", Max ", math.MaxInt32, "\n",
		"Int64: Min ", math.MinInt64, ", Max ", math.MaxInt64, "\n",
		"UInt8 : Min ", 0, ", Max ", math.MaxUint8, "\n",
		"UInt16: Min ", 0, ", Max ", math.MaxUint16, "\n",
		"UInt32: Min ", 0, ", Max ", math.MaxUint32, "\n",
		"UInt64: Min ", 0, ", Max ", GetMaxUnassignedInt64(), "\n",
		"Float32: ", math.MaxFloat32, "\n",
		"Float64: ", math.MaxFloat64, "\n",
		"Float32 Non Zero: ", math.SmallestNonzeroFloat32, "\n",
		"Float64 Non Zero: ", math.SmallestNonzeroFloat64)
}

func GetMaxUnassignedInt64() uint64 {
	return math.MaxUint64
}

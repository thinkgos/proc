package cache

import "testing"

func Test_Incr_Int(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint", 1, DefaultExpiration)
	err := tc.Incr("tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}
}

func Test_Incr_Int8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint8", int8(1), DefaultExpiration)
	err := tc.Incr("tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("tint8 is not 3:", x)
	}
}

func Test_Incr_Int16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint16", int16(1), DefaultExpiration)
	err := tc.Incr("tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("tint16 is not 3:", x)
	}
}

func Test_Incr_Int32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint32", int32(1), DefaultExpiration)
	err := tc.Incr("tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("tint32 is not 3:", x)
	}
}

func Test_Incr_Int64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint64", int64(1), DefaultExpiration)
	err := tc.Incr("tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("tint64 is not 3:", x)
	}
}

func Test_Incr_Uint(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint", uint(1), DefaultExpiration)
	err := tc.Incr("tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x.(uint) != 3 {
		t.Error("tuint is not 3:", x)
	}
}

func Test_Incr_Uintptr(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuintptr", uintptr(1), DefaultExpiration)
	err := tc.Incr("tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("tuintptr is not 3:", x)
	}
}

func Test_Incr_Uint8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint8", uint8(1), DefaultExpiration)
	err := tc.Incr("tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("tuint8 is not 3:", x)
	}
}

func Test_Incr_Uint16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint16", uint16(1), DefaultExpiration)
	err := tc.Incr("tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("tuint16 is not 3:", x)
	}
}

func Test_Incr_Uint32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint32", uint32(1), DefaultExpiration)
	err := tc.Incr("tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("tuint32 is not 3:", x)
	}
}

func Test_Incr_Uint64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint64", uint64(1), DefaultExpiration)
	err := tc.Incr("tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("tuint64 is not 3:", x)
	}
}

func Test_Incr_Float32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float32", float32(1.5), DefaultExpiration)
	err := tc.Incr("float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func Test_Incr_Float64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float64", float64(1.5), DefaultExpiration)
	err := tc.Incr("float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func Test_IncrFloat_Float32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float32", float32(1.5), DefaultExpiration)
	err := tc.IncrFloat("float32", 2)
	if err != nil {
		t.Error("Error incrementfloating:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func Test_IncrFloat_Float64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float64", float64(1.5), DefaultExpiration)
	err := tc.IncrFloat("float64", 2)
	if err != nil {
		t.Error("Error incrementfloating:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func Test_Decr_Int(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int", int(5), DefaultExpiration)
	err := tc.Decr("int", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int")
	if !found {
		t.Error("int was not found")
	}
	if x.(int) != 3 {
		t.Error("int is not 3:", x)
	}
}

func Test_Decr_Int8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int8", int8(5), DefaultExpiration)
	err := tc.Decr("int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("int8 is not 3:", x)
	}
}

func Test_Decr_Int16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int16", int16(5), DefaultExpiration)
	err := tc.Decr("int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("int16 is not 3:", x)
	}
}

func Test_Decr_Int32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int32", int32(5), DefaultExpiration)
	err := tc.Decr("int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("int32 is not 3:", x)
	}
}

func Test_Decr_Int64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int64", int64(5), DefaultExpiration)
	err := tc.Decr("int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("int64 is not 3:", x)
	}
}

func Test_Decr_Uint(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint", uint(5), DefaultExpiration)
	err := tc.Decr("uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint")
	if !found {
		t.Error("uint was not found")
	}
	if x.(uint) != 3 {
		t.Error("uint is not 3:", x)
	}
}

func Test_Decr_Uintptr(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uintptr", uintptr(5), DefaultExpiration)
	err := tc.Decr("uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("uintptr is not 3:", x)
	}
}

func Test_Decr_Uint8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint8", uint8(5), DefaultExpiration)
	err := tc.Decr("uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("uint8 is not 3:", x)
	}
}

func Test_Decr_Uint16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint16", uint16(5), DefaultExpiration)
	err := tc.Decr("uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("uint16 is not 3:", x)
	}
}

func Test_Decr_Uint32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint32", uint32(5), DefaultExpiration)
	err := tc.Decr("uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("uint32 is not 3:", x)
	}
}

func Test_Decr_Uint64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint64", uint64(5), DefaultExpiration)
	err := tc.Decr("uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("uint64 is not 3:", x)
	}
}

func Test_Decr_Float32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float32", float32(5.5), DefaultExpiration)
	err := tc.Decr("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3:", x)
	}
}

func Test_Decr_Float64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float64", float64(5.5), DefaultExpiration)
	err := tc.Decr("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func Test_DecrFloat_Float32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float32", float32(5.5), DefaultExpiration)
	err := tc.DecrFloat("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3:", x)
	}
}

func Test_DecrFloat_Float64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float64", float64(5.5), DefaultExpiration)
	err := tc.DecrFloat("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func Test_IncrInt(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint", 1, DefaultExpiration)
	n, err := tc.IncrInt("tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}
}

func Test_IncrInt8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint8", int8(1), DefaultExpiration)
	n, err := tc.IncrInt8("tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("tint8 is not 3:", x)
	}
}

func Test_IncrInt16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint16", int16(1), DefaultExpiration)
	n, err := tc.IncrInt16("tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("tint16 is not 3:", x)
	}
}

func Test_IncrInt32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint32", int32(1), DefaultExpiration)
	n, err := tc.IncrInt32("tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("tint32 is not 3:", x)
	}
}

func Test_IncrInt64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tint64", int64(1), DefaultExpiration)
	n, err := tc.IncrInt64("tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("tint64 is not 3:", x)
	}
}

func Test_IncrUint(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint", uint(1), DefaultExpiration)
	n, err := tc.IncrUint("tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x.(uint) != 3 {
		t.Error("tuint is not 3:", x)
	}
}

func Test_IncrUintptr(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuintptr", uintptr(1), DefaultExpiration)
	n, err := tc.IncrUintptr("tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("tuintptr is not 3:", x)
	}
}

func Test_IncrUint8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint8", uint8(1), DefaultExpiration)
	n, err := tc.IncrUint8("tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("tuint8 is not 3:", x)
	}
}

func Test_IncrUint16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint16", uint16(1), DefaultExpiration)
	n, err := tc.IncrUint16("tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("tuint16 is not 3:", x)
	}
}

func Test_IncrUint32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint32", uint32(1), DefaultExpiration)
	n, err := tc.IncrUint32("tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("tuint32 is not 3:", x)
	}
}

func Test_IncrUint64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("tuint64", uint64(1), DefaultExpiration)
	n, err := tc.IncrUint64("tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("tuint64 is not 3:", x)
	}
}

func Test_IncrFloat32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float32", float32(1.5), DefaultExpiration)
	n, err := tc.IncrFloat32("float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3.5 {
		t.Error("Returned number is not 3.5:", n)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func Test_IncrFloat64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float64", float64(1.5), DefaultExpiration)
	n, err := tc.IncrFloat64("float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3.5 {
		t.Error("Returned number is not 3.5:", n)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func Test_DecrInt8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int8", int8(5), DefaultExpiration)
	n, err := tc.DecrInt8("int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("int8 is not 3:", x)
	}
}

func Test_DecrInt16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int16", int16(5), DefaultExpiration)
	n, err := tc.DecrInt16("int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("int16 is not 3:", x)
	}
}

func Test_DecrInt32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int32", int32(5), DefaultExpiration)
	n, err := tc.DecrInt32("int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("int32 is not 3:", x)
	}
}

func Test_DecrInt64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int64", int64(5), DefaultExpiration)
	n, err := tc.DecrInt64("int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("int64 is not 3:", x)
	}
}

func Test_DecrUint(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint", uint(5), DefaultExpiration)
	n, err := tc.DecrUint("uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint")
	if !found {
		t.Error("uint was not found")
	}
	if x.(uint) != 3 {
		t.Error("uint is not 3:", x)
	}
}

func Test_DecrUintptr(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uintptr", uintptr(5), DefaultExpiration)
	n, err := tc.DecrUintptr("uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("uintptr is not 3:", x)
	}
}

func Test_DecrUint8(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint8", uint8(5), DefaultExpiration)
	n, err := tc.DecrUint8("uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("uint8 is not 3:", x)
	}
}

func Test_DecrUint16(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint16", uint16(5), DefaultExpiration)
	n, err := tc.DecrUint16("uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("uint16 is not 3:", x)
	}
}

func Test_DecrUint32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint32", uint32(5), DefaultExpiration)
	n, err := tc.DecrUint32("uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("uint32 is not 3:", x)
	}
}

func Test_DecrUint64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint64", uint64(5), DefaultExpiration)
	n, err := tc.DecrUint64("uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("uint64 is not 3:", x)
	}
}

func Test_DecrFloat32(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float32", float32(5), DefaultExpiration)
	n, err := tc.DecrFloat32("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3 {
		t.Error("float32 is not 3:", x)
	}
}

func Test_DecrFloat64(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("float64", float64(5), DefaultExpiration)
	n, err := tc.DecrFloat64("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3 {
		t.Error("float64 is not 3:", x)
	}
}

func Test_IncrOverflowInt(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("int8", int8(127), DefaultExpiration)
	err := tc.Incr("int8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get("int8")
	int8 := x.(int8)
	if int8 != -128 {
		t.Error("int8 did not overflow as expected; value:", int8)
	}
}

func Test_IncrOverflowUint(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint8", uint8(255), DefaultExpiration)
	err := tc.Incr("uint8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get("uint8")
	uint8 := x.(uint8)
	if uint8 != 0 {
		t.Error("uint8 did not overflow as expected; value:", uint8)
	}
}

func Test_DecrUnderflowUint(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("uint8", uint8(0), DefaultExpiration)
	err := tc.Decr("uint8", 1)
	if err != nil {
		t.Error("Error decrementing int8:", err)
	}
	x, _ := tc.Get("uint8")
	uint8 := x.(uint8)
	if uint8 != 255 {
		t.Error("uint8 did not underflow as expected; value:", uint8)
	}
}

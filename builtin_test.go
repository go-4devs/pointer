package pointer

import "testing"

func TestBool(t *testing.T) {
	var x bool
	if *ToBool(x) != x {
		t.Fail()
	}
}

func TestByte(t *testing.T) {
	var x byte
	if *ToByte(x) != x {
		t.Fail()
	}
}

func TestComplex128(t *testing.T) {
	var x complex128
	if *ToComplex128(x) != x {
		t.Fail()
	}
}

func TestComplex64(t *testing.T) {
	var x complex64
	if *ToComplex64(x) != x {
		t.Fail()
	}
}

func TestFloat32(t *testing.T) {
	var x float32
	if *ToFloat32(x) != x {
		t.Fail()
	}
}

func TestFloat64(t *testing.T) {
	var x float64
	if *ToFloat64(x) != x {
		t.Fail()
	}
}

func TestInt(t *testing.T) {
	var x int
	if *ToInt(x) != x {
		t.Fail()
	}
}

func TestInt16(t *testing.T) {
	var x int16
	if *ToInt16(x) != x {
		t.Fail()
	}
}

func TestInt32(t *testing.T) {
	var x int32
	if *ToInt32(x) != x {
		t.Fail()
	}
}

func TestInt64(t *testing.T) {
	var x int64
	if *ToInt64(x) != x {
		t.Fail()
	}
}

func TestInt8(t *testing.T) {
	var x int8
	if *ToInt8(x) != x {
		t.Fail()
	}
}

func TestRune(t *testing.T) {
	var x rune
	if *ToRune(x) != x {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var x string
	if *ToString(x) != x {
		t.Fail()
	}
}

func TestUint(t *testing.T) {
	var x uint
	if *ToUint(x) != x {
		t.Fail()
	}
}

func TestUint16(t *testing.T) {
	var x uint16
	if *ToUint16(x) != x {
		t.Fail()
	}
}

func TestUint32(t *testing.T) {
	var x uint32
	if *ToUint32(x) != x {
		t.Fail()
	}
}

func TestUint64(t *testing.T) {
	var x uint64
	if *ToUint64(x) != x {
		t.Fail()
	}
}

func TestUint8(t *testing.T) {
	var x uint8
	if *ToUint8(x) != x {
		t.Fail()
	}
}

func TestUintptr(t *testing.T) {
	var x uintptr
	if *ToUintptr(x) != x {
		t.Fail()
	}
}

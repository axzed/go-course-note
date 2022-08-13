package day1

import (
	"testing"
)

const (
	UNKNOWN = iota
	FEMALE
	MALE
)

func TestValueRef(t *testing.T) {
	i := "this is a"
	j := &i
	*j = "this is j"
	t.Log(i, *j)
}

func TestValueCopy(t *testing.T) {
	i := "this is a"
	j := i
	t.Log(i, j)
	t.Log(&i, &j)
}

func TestConstOneLine(t *testing.T) {
	const c1, c2 string = "c1", "c2"
	t.Log(c1, c2)
}

func TestConsttMulti2(t *testing.T) {
	const (
		PI     float32 = 3.14
		PILong float32 = 3.14159
	)
	t.Log(PI, PILong)
}

func TestConsttMulti(t *testing.T) {
	const (
		PI     float32 = 3.14
		PILong float32 = 3.14159
	)
	t.Log(PI, PILong)
}

// 变量作用域
func TestScope(t *testing.T) {
	{
		var a string
		a = "hello"
		t.Log(a)
	}
	// 作用域之外无法使用
	// t.Log(a)
}

func TestConst(t *testing.T) {
	const (
		a = "sdf"
	)
	// a = "bbb"
	t.Log(a)
}

func TestShort(t *testing.T) {
	a, b, c := "value1", 10, 0.01
	t.Log(a, b, c)
}

func TestMergeMulti(t *testing.T) {
	var (
		a string  = "value1"
		b int     = 10
		c float32 = 0.01
	)
	t.Log(a, b, c)
}

func TestMerge(t *testing.T) {
	var a string = "one line"
	t.Log(a)
}

func TestMulti(t *testing.T) {
	var (
		a string
		b int
	)
	a, b = "string", 10
	t.Log(a, b)
}

func TestVariable(t *testing.T) {
	var a string
	a = "test"
	t.Log(a)
}

func TestFloat64(t *testing.T) {
	var a float64
	t.Log(a)
}

func TestStringNull(t *testing.T) {
	var a string
	t.Log(a)
}

func TestCharNull(t *testing.T) {
	var a rune
	t.Log(a)
}

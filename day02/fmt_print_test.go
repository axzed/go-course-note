package day2

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBasic(t *testing.T) {
	username := "boy"
	fmt.Printf("welcome, %s", username)
}

func TestNumber(t *testing.T) {
	a := 255
	fmt.Printf("二进制: %b\n", a)
	fmt.Printf("八进制: %o\n", a)
	fmt.Printf("十进制: %d\n", a)
	fmt.Printf("十六进制: %x\n", a)
	fmt.Printf("大写十六进制: %X\n", a)

	fmt.Printf("十六进制: %d\n", Hex2Dec("4E2D"))
	fmt.Printf("字符: %c\n", 20013)
	fmt.Printf("Unicode格式: %U\n", '衣') // U+4E2D
	fmt.Printf("%x", 0b1000100001100011)
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

func Hex2Bin(val string) string {
	return fmt.Sprintf("%b", Hex2Dec(val))
}

func Bin2Hex(val string) string {
	ui, err := strconv.ParseUint(val, 2, 64)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%x", ui)
}

func TestFloat(t *testing.T) {
	fmt.Printf("%e", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%E", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%f", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%g", 12675757563.5345432567)
	fmt.Println()
	fmt.Printf("%G", 12675757563.5345432567)
	fmt.Println()
}

func TestBool(t *testing.T) {
	fmt.Printf("%t", true)
	fmt.Println()
}

func TestString(t *testing.T) {
	str := "I'm a boy"
	fmt.Printf("%s", str)
	fmt.Println()
	fmt.Printf("%q", str)
	fmt.Println()
	fmt.Printf("%x", str)
	fmt.Println()
	fmt.Printf("%X", str)
	fmt.Println()
}

func TestPointer(t *testing.T) {
	a := "I'm a boy"
	b := &a
	fmt.Printf("%p", b)
	fmt.Println()
}

func TestErr(t *testing.T) {
	// fmt.Printf("this is  %s", 10)
}

func TestPos(t *testing.T) {
	f := 1010.0101
	s := "hey boy!"
	fmt.Printf("%v", f)
	fmt.Println()

	fmt.Printf("%v", s)
	fmt.Println()
}

type User struct {
	Name string
	Age  int
}

func TestV(t *testing.T) {
	user := User{"laowang", 33}
	fmt.Printf("%v", user) // Go默认形式 {laowang 33}
	fmt.Println()
	fmt.Printf("%#v", user) //类型+值对象 day2.User{Name:"laowang", Age:33}
	fmt.Println()
	fmt.Printf("%+v", user) //输出字段名和字段值形式 {Name:laowang Age:33}
	fmt.Println()
	fmt.Printf("%T", user) //值类型的Go语法表示形式, day2.User
	fmt.Println()
	fmt.Printf("%%")
}

func TestWidth(t *testing.T) {
	fmt.Printf("|%s|", "aa") // 不设置宽度
	fmt.Println()
	fmt.Printf("|%5s|", "aa") // 5个宽度,  默认+， 右对齐
	fmt.Println()
	fmt.Printf("|%-5s|", "aa") // 5个宽度, 左对齐
	fmt.Println()

	fmt.Printf("|%.2s|", "xxxx") // 最大宽度为5，超出的部分会被截断
}

func TestWithdFill(t *testing.T) {
	fmt.Printf("|%05s|", "aa")
	fmt.Println()
}

func TestFloatWidth(t *testing.T) {
	a := 54.123456
	fmt.Printf("|%f|", a)
	fmt.Println()
	fmt.Printf("|%5.1f|", a)
	fmt.Println()
	fmt.Printf("|%-5.1f|", a)
	fmt.Println()
	fmt.Printf("|%05.1f|", a)
	fmt.Println()
}

func TestWidthMore(t *testing.T) {
	fmt.Printf("|%2s|", "中国")
	fmt.Println()
	fmt.Printf("|%2s|", "ab")
	fmt.Println()
}

// 1000 1000 0110 0011
// 0110 0111 0000 1101  // 4949027062337
// 0101 0101 1001 1100
// 0110 1011 0010 0010
// 0111 1010 0111 1111
// 0100 1110 0010 1101
// 0101 0110 1111 1101
// 0111 1110 1010 0010

func TestInt(t *testing.T) {
	intA := []int{011, 0o110011100001101}
	for _, v := range intA {
		fmt.Println(v)
	}

}

func TestPrintln(t *testing.T) {
	i, err := fmt.Println(1)
	fmt.Println(i, err)
}

package day3

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings1(t *testing.T) {
	a := "hello"
	fmt.Println([]byte(a)) // [104 101 108 108 111]
}

func TestStrings2(t *testing.T) {
	a := "hello"
	b := []byte(a)
	b[0] = 'x'
	fmt.Println(string(b))
}

func TestStrings3(t *testing.T) {
	a := "hello"
	fmt.Println(len(a), a[0], a[1:3])
}

func TestString4(t *testing.T) {
	fmt.Println(strings.Compare("ab", "cd"))
	fmt.Println(strings.EqualFold("ab", "AB"))
}

func TestStrings6(t *testing.T) {
	a := "hello"
	b := []byte(a)
	fmt.Println(b, len(a), len(b))
}

func TestStrings7(t *testing.T) {
	fmt.Println(len("谷歌中国"), []byte("谷歌中国"))
	for _, v := range []byte("谷歌中国") {
		fmt.Printf("%b\n", v)
	}

	fmt.Printf("%c\n", 0b1000110000110111)
}

func TestString8(t *testing.T) {
	data := []byte{232, 176, 183}
	for i, v := range data {
		if i == 0 {
			fmt.Printf("%b\n", v&0b1111)
			continue
		}

		fmt.Printf("%b\n", v&0b111111)
	}

	code := 0b1000<<12 + 0b110000<<6 + 0b110111
	fmt.Printf("%b -> %c\n", code, code)
}

func TestString5(t *testing.T) {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(len("谷歌中国"))
	fmt.Println(strings.Count("谷歌中国", ""))
}

func TestStringSplite(t *testing.T) {
	str := "abc,DEF,MQP"
	fmt.Println(strings.Split(str, ","))
	fmt.Println(strings.SplitN(str, ",", 2))

	fmt.Println(strings.Split(str, ""))
}

func TestStringIndex(t *testing.T) {
	str := "abc,DEF,MQP,abc,DEF,MQP"
	fmt.Println(strings.Index(str, "DEF"))
	fmt.Println(strings.LastIndex(str, "DEF"))
}

func TestStringMap(t *testing.T) {
	str := "hello"
	new := strings.Map(func(c rune) rune {
		if c == 'h' {
			return 'm'
		}
		return c
	}, str)

	fmt.Println(new)
}

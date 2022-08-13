package homework

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Book struct {
	Title  string
	Author string
	Page   uint
	Tag    []string
}

func TestHomeWork1(t *testing.T) {
	// 空结构体 不占用内存空间, 不申请内存
	fmt.Println(unsafe.Sizeof(struct{}{}))

	// 取得结构体实例的 内存地址
	b := Book{Tag: []string{"abc", "def", "hjk"}}
	fmt.Println(&b)
	p1 := &b

	// offset 怎么计算
	fmt.Println(unsafe.Offsetof(p1.Tag)) // 40 <成员变量的指针> ----> 访问成员变量的值

	// 结构体实例地址, +  成员变量的偏移量 =  成员变量的内存地址
	// uintptr --->  Pointer ----> (*[]string) ---> *()

	p := (*[]string)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Offsetof(b.Tag)))
	fmt.Println(*p)

	fmt.Printf("%p\n", &b.Tag[0])
	h := (*reflect.SliceHeader)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Offsetof(b.Tag)))
	fmt.Printf("%p\n", h)
	ptr := (*string)(unsafe.Pointer(h.Data))
	fmt.Printf("%p -> %v\n", ptr, *ptr)

	fmt.Printf("%p\n", &b.Tag[1])
	ptr1 := (*string)(unsafe.Pointer(h.Data + 16))
	fmt.Printf("%p -> %v\n", ptr1, *ptr1)

	fmt.Printf("%p\n", &b.Tag[2])
	ptr2 := (*[]byte)(unsafe.Pointer(h.Data + 16))
	fmt.Printf("%p -> %v\n", ptr2, *ptr2)
}

// Student 保存的学员信息
// Student 名称(Name) 学号(Number) 科目(Subjects) 成绩(Scores)
type Student struct {
	Name     string   // 名称
	Number   uint16   // 学号  2 ^ 16
	Subjects []string // 数学  语文  英语
	Score    []int    //  88   99   77
}

// Class 保存的是班级的信息
type Class struct {
	Name     string     // 班级名称
	Number   uint8      // 班级编号
	Subjects []string   // 数学  语文  英语
	Students []*Student // 班级学员, []int --> [10, 20, 30]  []*int ---> [0xaabb, 0xccc, oxddd]
}

//
func (c *Class) AvgScore() []int {
	// slice   for range
	// 获取每一个学员的第一个成绩 就是 数据成绩
	return nil
}

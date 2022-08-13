# 作业

1. 通过内存地址访问Tag的值
```go
type Book struct {
	Title  string
	Author string
	Page   uint
	Tag    []string
}

// b := &Book{Tag: []string{"abc", "def", "hjk"}}
// 根据结构体的内存地址, 计算出Tag的内存地址, 并访问
```

思路: 指针 与 offset

2. 使用二维切片表示一组学生的各科成绩，计算这组学生的学科平均分

```go
//  数学   语文   英语   
//   88     88    90
//   66
//   ...   
//   avg     
scores = [][]int{
	{88, 88, 90},
	{66, 99, 94},
	{75, 84, 98},
	{93, 77, 66},
}
```

3. 使用结构体表示班级和学生，请计算每个班级学科平均分

```go
// Student 名称(Name) 学号(Number) 科目(Subjects) 成绩(Scores)
// Class   名称(Name) 编号(Number) 学员(Students)
// Class   实现一个平均值的方法
```
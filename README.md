# myGoMemos

## Array Slice

```Go
var thisIsArray int = [3]int{1,2,3} // Array必须定义好长度
```
```go
var thisIsSlice int = []int{} // Slice不必定义长度和内容
```

Two way to create a slice:

```Go
package main

import "fmt"

func main() {
    original := []int{10, 20, 30, 40, 50}
    
    // Method 1: Direct slicing
    slice1 := original[0:2] // 用一个Slice的一部分定义一个Slice
    
    // Method 2: Using the append function and a loop
    slice2 := make([]int, 0) // 用make定义一个Slice
    for i := 0; i < 2; i++ {
        slice2 = append(slice2, original[i])
    }
    
    // Print both slices
    fmt.Println("Slice 1:", slice1) // Output will be [10 20]
    fmt.Println("Slice 2:", slice2) // Output will be [10 20] as well
}
```

## Map
struct的数据结构是预先定义好的,无法动态地增加或删除键值对。而map则非常灵活，可以随时增加缩减

struct擅长描述数据实体,而map则更适合管理一个值的集合。


## make function

如果事先知道要定义的slice或者map有多少项目，就可以在定义时使用make函数，对容量作出定义，可以让Go预先分配对应的内存。

```go
var SliceByMake = make([]string, 2, 5) // make(type, len, cap) 类型，长度，容量
```

## For

在Go里，只有for循环

语法为

```go

for index, value := range userNames {

}
```
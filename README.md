# myGoMemos

## Pointer

```go
var a int //变量
var b_pt *int = &a // 指针

fmt.Print(&a) // a的地址 等于b_pt
fmt.Print(*b_pt) // 等于a的值 

func ScanPara(para *int){} //参数为一个指针类型
```

在Go中，需要修改原数据的操作，用Pointer比较好
类似Scan函数 Scan( &para )

在操作大型Struct或Array时，用数组可以提升效率

## IO

在Go语言中，进行输入输出（I/O）操作涉及到多个包，主要是`io`和`os`包。这些操作包括读取输入、写入输出到文件或标准输出（如控制台），以及处理网络I/O等。这里将介绍一些基础的I/O操作，包括文件的读取和写入、以及标准输入输出的处理。

### 1. 文件读取和写入
对文件进行读写操作通常需要使用`os`包来打开或创建文件，并使用`io`或`bufio`包来读取或写入文件。

#### 写入文件
以下是一个简单的例子，演示如何在Go中创建并写入文件：

```go
package main

import (
    "os"
    "log"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    _, err = file.WriteString("Hello, world!\n")
    if err != nil {
        log.Fatal(err)
    }
}
```
这段代码使用`os.Create`函数创建一个名为`example.txt`的文件，如果文件创建成功，则使用`WriteString`方法写入一些文本。`defer file.Close()`确保在函数退出前关闭文件。

#### 读取文件
读取文件的一个简单示例：

```go
package main

import (
    "bufio"
    "os"
    "log"
    "fmt"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```
这段代码使用`os.Open`打开文件，然后使用`bufio.Scanner`逐行读取文件内容。

### 2. 网络I/O
进行网络I/O通常涉及到`net`包。例如，创建一个简单的HTTP请求：

```go
package main

import (
    "io/ioutil"
    "net/http"
    "log"
    "fmt"
)

func main() {
    resp, err := http.Get("http://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}
```
这段代码发送一个HTTP GET请求到`example.com`，并打印返回的HTML内容。

这些是Go语言中最常见的几种I/O操作方式。每种方式都有其适用场景，例如文件I/O对于处理本地文件很有用，网络I/O对于网络应用程序开发是必需的，标准输入输出通常用于简单的交互式程序。


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

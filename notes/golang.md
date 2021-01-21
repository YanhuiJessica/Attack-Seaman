# Notes of Go

## Packages

- 所有 Go 程序由 packages 组成，类比 Python 程序的文件名
- 需要在程序头定义
    ```go
    package main
    ```

## Imports

```go
// Method 1
import "fmt"
import "math"

// Method 2, Recommended
import (
    "fmt"
    "math"
)
```

## Exported names

- 导出名必须首字母大写
- 导入一个 package 后，只能引用其导出名，如：`main.Pi`

## Functions

- 先参数名，后参数类型
    ```go
    func exp(x int, y int) int {
        return x + y
    }
    ```
- 当连续多个参数类型相同时，除最后一个，其他类型声明都可省略
    ```go
    func exp(x, y int) int {
        return x + y
    }
    ```
- 一个函数可以有任意数量的返回值
    ```go
    func swap(x, y string) (string, string) {
        return y, x
    }
    ```
- 函数的返回值可以命名，并被视作定义在函数顶部的变量
    ```go
    func split(sum int) (x, y int) {
        x = sum * 4 / 9
        y = sum - x
        return  // 无传参的返回语句返回已命名的返回值，不推荐在长函数中使用
    }
    ```

## Variables

```go
package main

import "fmt"

// 和 import 一样，也可以分开写
var (
	c int
    python, java bool
    s = "Type omitted!"
)

// 不能在函数外使用短变量声明 :=
// 函数外，所有语句应以关键字开头（var，func 等）

func main() {
    var i int
    j := 6.66
	fmt.Println(i, j, c, s, python, java)
}
```
- 未赋初值的变量将赋零值

    Type | Zero Value
    -|-
    Numeric|`0`
    Boolean|`false`
    String|`""`

### Basic types

```go
bool

string

// 除非特殊情况，直接使用 int
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)    // 复数
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

### Type conversions

- 类型转换只能是**显式**地
- `T(v)`将`v`的类型转换为`T`

```go
i := 42
f := float64(i)
u := uint(f)
```
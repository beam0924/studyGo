## 引言

### GO 语言介绍

* Go语言设计动机

C++太难学而且编译太慢

GO是21世纪的C语言

特点：语法简单；开发效率高；执行性能好

* 解释型语言和编译型语言的区别

![image-20211203221544642](https://static01.imgkr.com/temp/3a5fa0c1de774d988619e3d067a1b766.png)

### Go语言开发环境搭建

在进行go开发的时候，我们的代码保存在`$GOPATH/src`目录下面，在工程经过`go build`、`go install`、`go get`等指令之后，会将下载的第三方包源代码文件放在`$GOPATH/src`目录下，产生的二进制可执行文件放在`$GOPATH/bin`目录下，生成的中间缓存文件会被保存在`$GOPATH/pkg`目录下。

用Git这类版本控制工具时，只需要添加`$GOPATH/src`目录下源代码即可。



推荐的个人组织结构如下：

![IMG_837A1F3AFFF0-1](![image-20211203221544642](https://static01.imgkr.com/temp/7493d414987a4419aca97960bd36833f.jpeg)

更高级的方法推荐如下：

![IMG_1496](https://static01.imgkr.com/temp/a8d66adeca6c491aa189758fd4fecef5.jpg)

企业如下：

![IMG_5A45EF96AD86-1](https://static01.imgkr.com/temp/478c0b9a09604f62ad93cda8d9fbc0d7.jpeg)





### 编译

使用`go build`

1.在项目目录下执行`go build`

2.在任意目录下`go build`，需要在后面加上待编译项目的路径（项目路径从`GOPATH/src`后面写起,编译之后的可执行文件就保存在当前目录下）

3.`go build -o hello`自定义可执行文件的名字

4.`go run xxx.go` 像执行脚本文件一样执行Go代码

5.`go install xxx.go` 分两步：先编译得到一个可执行文件，再把可执行文件拷贝到`GOPATH/bin`下

6.可以通过设置`CGO_ENABLED | GOOS | GOARCH`来实现跨平台编译

7.`go install`：主要用来生成库和工具。一是编译包文件（无main包），将编译后的包文件放到 pkg 目录下（`$GOPATH/pkg`）。二是编译生成可执行文件（有main包），将可执行文件放到 bin 目录（`$GOPATH/bin`）。

8.要编译可执行文件，必须要有main包和main函数也就是入口函数。



### 注释

* 单行注释

  ```go
  //我是注释
  ```

* 多行注释

  ```go
  /*
  多行注释
  */
  ```

  



### 变量申明

* 函数外只能进行标识符（变量、常量、函数、类型）的声明，不能有其他操作语句

* 关键字：编程语言中预先定义好的具有特殊含义的标识符。关键字和保留字都不建议作为变量名

* Go语言一共有25个关键字，37个保留字

* Go语言中的变量必须先申明后使用

* 申明变量：`var 变量名 变量类型`

* Go语言中非全局变量申明了必须使用，不然编译不过去

* 没赋值的初始化变量有个初始值string类型是“”，int类型是0，bool类型是false

* Go语言推荐小驼峰式命名

* Go语言支持类型推导，编译器可根据等号后边的值来推导变量的类型完成初始化。

  ```go
  var name = "Qimi"
  var age = 18
  ```

* Go语言还支持简短变量申明，**只能在函数中使用**

  ```go
  s3 := "heiheihei"
  ```

* 匿名变量：在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量，用一个下划线`_`表示，匿名函数不占用命名空间，不会分配内存

* 函数外的每个语句都必须以关键字开始（var、const、func等）

* 同一个作用域中不能重复声明同一个变量



### 常量

* 常量定义和变量非常类似，只是把var换成const

* 常量定义之后必须赋值

* 常量在程序运行期间不能改变值

* 批量声明常量时，如果某一行声明后没有赋值，默认和上一行一致(下面代码输出为200)

  ```go
  const (
  	statusOK = 200
  	notFound
  )
  func main() {
  	fmt.Println(notFound)
  }
  
  //output:
  //200
  ```

* `iota`是go语言中的常量计数器，只能在常量的表达式中使用。`iota`在const关键字出现时被重置为0，const中每新增一行常量声明将使`iota`计数一次（`iota`可以理解为const语句块中的行索引）。使用`iota`可以简化定义，在定义枚举时非常有用。

  ```go
  const (
    			n1 = iota //0
    			n2        //1
    			n3				//2
    			n4				//3
  )
  ```

  ```go
  const (
  	a1, a2 = iota + 1, iota + 2
  	a3, a4 = iota + 1, iota + 2
  )
  
  func main() {
  	fmt.Println(a1, a2, a3, a4)
  }
  ```

  输出为`1 2 2 3`

```
//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
```



## GO语言基本数据类型

### 整型

* 整型按长度可分为`int8 int16 int32 int64 `对应的无符号整型`uint8 uint16 uint32 uint64`；其中`int32`对应C语言中的`short`型，`int64`对应C语言中的`long`型

* `int`在32位操作系统上就是`int32`，64位操作系统上就是`int64`

* `uintptr`无符号整型，用于存放一个指针

* 八进制&十六进制示例：

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	//十进制
  	var a int = 10
  	fmt.Printf("%d \n", a) //10
  	fmt.Printf("%b \n", a)  //1010  无法直接定义二进制数，占位符%b表示二进制
  
  	//八进制 常用于Linux文件权限
  	var b int = 077
  	fmt.Printf("%o \n", b) //77
  
  	//十六进制 常用于表示地址
  	var c int = 0xff
  	fmt.Printf("%x \n", c) //ff
  	fmt.Printf("%X \n", c) //FF
  }
  
  // output:
  // (base) BEAMWANG-MB0:02var beam$ go run test2.go
  // 10 
  // 1010 
  // 77 
  // ff 
  // FF 
  ```

* 用`fmt.Printf("%T \n", i)`查看变量`i`的类型



### 浮点型

* Go语言支持两种浮点型数：`float32`和`float64`,`float32`的浮点数的最大范围 约为`3.4e38`，可以使用常量定义：`math.MaxFloat32`。`float64`的浮点数的最大范围为`1.9e308`，可以用另外一个常量定义：`math.MaxFloat64`。
* 默认Go语言中的小数都是`float64`，`float32`类型只能显示声明，例如`f2 := float32(1.23456)`。
* Go是静态类型的语言，所以不同类型的变量之间不能直接互相赋值。



### 复数

有`complex64`和`complex128`两种

```go
package main

import "fmt"

func main() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}


// output:
// (base) BEAMWANG-MB0:02var beam$ go run test2.go
// (1+2i)
// (2+3i)
```



### 布尔值

Go语言中以`bool`类型进行布尔型数据声明，布尔型数据只有`true(真)`和`false(假)`两个值

要注意的是：

1. 布尔类型变量的默认值是`false`；
2. Go语言中不允许将整数强制转换成布尔型；
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

```go
package main

import "fmt"

func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T \n", b1)
	fmt.Printf("%T \n", b2)
	fmt.Printf("%v \n", b2)
}

// output:
// (base) BEAMWANG-MB0:02var beam$ go run test2.go
// bool 
// bool 
// false 
```



### 字符串

Go语言中的字符串是以原生数据类型出现的，内部实现使用`UTF-8`编码，字符串的值为双引号`双引号("")`中的内容（不可以是单引号，单引号包裹的是字符），可以在Go语言的源码中直接添加非ASCII码字符。

* 多行字符串定义，使用反引号原样输出，比如说路径可以不用加转译字符

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	s1 := `
  	世情薄
  	人情恶
  	雨送黄昏花易落
  	`
  	fmt.Println(s1)
  }
  
  //output:
  //(base) BEAMWANG-MB0:string beam$ go run main.go
  //
  //        世情薄
  //        人情恶
  //        雨送黄昏花易落
  //
  ```

* 字符串的常用操作：

  ```go
  len(str)  求长度
  + 或 fmt.Sprintf		拼接字符串
  strings.Split		分割
  strings.contains		判断是否包含
  strings.HasPrefix,strings.HasSuffix		前缀/后缀判断
  strings.Index(),strings.LastIndex()		子串出现的位置
  strings.Join(a[]string, sep string)		join操作
  ```

  ```go
  package main
  
  import (
  	"fmt"
  	"strings"
  )
  
  func main() {
  	s1 := "PATH:/Users/beam/Desktop/秋招"
  	//字符串相关的操作
  	//字符串长度打印
  	fmt.Println(len(s1))
  
  	//字符串拼接
  	name := "理想"
  	world := "大帅比"
  	ss := name + world
  	fmt.Println(ss)
  	ss1 := fmt.Sprintf("%s%s", name, world)
  	fmt.Println(ss1)
  
  	//字符串分割
  	ret := strings.Split(s1, "/")
  	fmt.Println(ret)
  
  	//判断包含
  	fmt.Println(strings.Contains(ss, "理想"))
  	fmt.Println(strings.Contains(ss, "理性"))
  
  	//前缀和后缀判断
  	fmt.Println(strings.HasPrefix(ss, "理想"))
  	fmt.Println(strings.HasSuffix(ss, "理想"))
  
  	//判断子串出现的位置
  	s2 := "abcdeb"
  	fmt.Println(strings.Index(s2, "c"))
  	fmt.Println(strings.Index(s2, "b"))
  	fmt.Println(strings.LastIndex(s2, "b"))
  
  	//拼接操作
  	fmt.Println(strings.Join(ret, "+"))
  }
  
  // output:
  // 31
  // 理想大帅比
  // 理想大帅比
  // [PATH: Users beam Desktop 秋招]
  // true
  // false
  // true
  // false
  // 2
  // 1
  // 5
  // PATH:+Users+beam+Desktop+秋招
  ```



### 字符

* Go语言的字符有两种：
  * `uint8`类型，或者叫`byte`型，代表了`ASCII`码的一个字符。
  * `rune `类型，代表一个`UTF-8`字符。
* 当需要处理中文、日文或者其他符合字符时，则需使用`rune `类型，`rune `的本质是一个`int32`
* 字符默认是`rune `类型，也就是`int32`。如果非要用`byte`型，则需要强制类型转换：`c := byte('H')`，这样`c`的类型就变成了`uint8`
* `byte`和`rune `其实就是类型别名，`byte`=`uint8`类型，`rune`=`int32`类型。



### 修改字符串

要修改字符串，需要先将其转换成`[]ruen`或`[]byte`，完成之后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。**字符串类型不能直接修改**

```go
package main

import "fmt"

func main() {
	s1 := "pig"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'b'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// output:
// big
// 红萝卜
```

### 类型转换

基本语法为`T(表达式)`，其中`T`表示要转换的类型，表达式包括变量和函数返回值等。

```go
n := 10
var f float64
f = float64(n)
```

### fmt总结

```go
package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n)
	//查看变量的值
	fmt.Printf("%v\n", n)
	//查看二进制
	fmt.Printf("%b\n", n)
	//查看十进制
	fmt.Printf("%d\n", n)
	//查看八进制
	fmt.Printf("%o\n", n)
	//查看十六进制
	fmt.Printf("%x\n", n)
	var s = "Hello, Tencent!"
	//查看字符串的值
	fmt.Printf("%s\n", s)
	//查看字符串的值（通用版）
	fmt.Printf("%v\n", s)
	//查看字符串的值（添加标识符）
	fmt.Printf("%#v\n", s)
}

// output:
// (base) BEAMWANG-MB0:fmt beam$ go run main.go
// int
// 100
// 1100100
// 100
// 144
// 64
// Hello, Tencent!
// Hello, Tencent!
// "Hello, Tencent!"
```



## GO语言流程控制



### if...else...

```go
package main

import "fmt"

//if条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	} else {
		fmt.Println("该写暑假作业了")
	}

	//多个条件判断
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习！")
	}
}

// output:
// 澳门首家线上赌场开业了
// 青年
```

```go
if age :=19; age > 18 { //age变量此时只在if条件判断语句中生效，减少了内存占用
		fmt.Println("澳门首家线上赌场开业了")
	} else {
		fmt.Println("该写暑假作业了")
	}
```



### for循环

Go语言中的所有循环类型均可以用`for`关键字来完成

for循环的基本格式如下：

```go
for 初始语句; 条件判断语句; 结束语句{
		循环体语句
}
```

```go
package main

import "fmt"

//for循环
func main() {
	//基本格式
	for i := 10; i < 15; i++ {
		fmt.Println(i)
	}

	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	//变种2
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//无限循环
	//for {
	//fmt.Println("123")
	//}

	//for range循环
	s := "Hello腾讯"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}

// output:
// 10
// 11
// 12
// 13
// 14
// 5
// 6
// 7
// 8
// 9
// 10
// 0 H
// 1 e
// 2 l
// 3 l
// 4 o
// 5 腾
// 8 讯
```



* 用`break`和`continue`控制循环

```go
package main

import "fmt"

//流程控制之跳出for循环
func main() {
	//当i=5时，跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("over")

	//当i=5时，跳过本次for循环（不执行for循环内部的打印语句，继续下一次循环）
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("over")
}
```



### 作业：打印九九乘法表

```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}


// output:
// 1 * 1 = 1
// 2 * 1 = 2       2 * 2 = 4
// 3 * 1 = 3       3 * 2 = 6       3 * 3 = 9
// 4 * 1 = 4       4 * 2 = 8       4 * 3 = 12      4 * 4 = 16
// 5 * 1 = 5       5 * 2 = 10      5 * 3 = 15      5 * 4 = 20      5 * 5 = 25
// 6 * 1 = 6       6 * 2 = 12      6 * 3 = 18      6 * 4 = 24      6 * 5 = 30      6 * 6 = 36
// 7 * 1 = 7       7 * 2 = 14      7 * 3 = 21      7 * 4 = 28      7 * 5 = 35      7 * 6 = 42      7 * 7 = 49
// 8 * 1 = 8       8 * 2 = 16      8 * 3 = 24      8 * 4 = 32      8 * 5 = 40      8 * 6 = 48      8 * 7 = 56      8 * 8 = 64
// 9 * 1 = 9       9 * 2 = 18      9 * 3 = 27      9 * 4 = 36      9 * 5 = 45      9 * 6 = 54      9 * 7 = 63      9 * 8 = 72      9 * 9 = 81
```



### switch语句

使用`switch`可以方便地对大量的值进行条件判断

```go
package main

import "fmt"

//switch 简化大量的判断(一个变量和具体的值作比较)
func main() {
	var n = 5
	if n == 1 {
		fmt.Println("拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小指")
	} else {
		fmt.Println("无效数字")
	}

	//switch简化上面代码
	switch n {
	case 1:
		fmt.Println("拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效数字")
	}
}
```

* `switch`语句中可以定义局部变量，并且`case`后面可以跟多个值，用逗号隔开。

  ```go
  func testSwitch() {
  	switch n := 7; n {
  	case 1, 3, 5, 7, 9:
  		fmt.Println("奇数")
  	case 2, 4, 6, 8:
  		fmt.Println("偶数")
  	default:
  		fmt.Println(n)
  	}
  }
  ```

* `case`分支语句之后还可以跟表达式，这时`switch`语句后面无需再跟判断变量。

  ```go
  func testSwitch2() {
  	age := 30
  	switch {
  	case age < 25:
  		fmt.Println("好好学习吧")
  	case age >= 25 && age < 35:
  		fmt.Println("好好工作吧")
  	case age > 60:
  		fmt.Println("好好享受吧")
  	default:
  		fmt.Println("活着真好")
  	}
  }
  ```

* `fallthrough`语法可以执行满足条件`case`的下一个`case`，这个是为兼容C语言中的`case`设计的。

  ```go
  func testSwitch3() {
  	s := 'a'
  	switch {
  	case s == 'a':
  		fmt.Println("a")
  		fallthrough
  	case s == 'b':
  		fmt.Println("b")
  	case s == 'c':
  		fmt.Println("c")
  	default:
  		fmt.Println("...")
  	}
  }
  
  // output:
  // a
  // b
  ```



### goto 跳转到指定标签

跳出多层for循环：

```go
package main

import "fmt"

func main() {
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break //只跳出内层for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出外层for循环
		}
	}

	//goto + label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto LABEL //调到我们指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
LABEL: //label标签
	fmt.Println("over")
}

// output:
// 0-A
// 0-B
// 0-A
// 0-B
// over
```

* `break`和`continue`也可以配合标签进行使用。
* 标签语法影响代码可读性，尽量少用。



## Go语言运算符



### 算数运算符

```go
+ 相加
- 相减
* 相乘
/ 相除
% 求余
```

`变量名++`和`变量名--`是单独的语句，不能放在等号的后边赋值；`++`和`--`也不是运算符。

### 关系运算符

```go
==
!=
>
>=
<
<=

//返回bool值
```

* Go语言是强类型的，相同类型的变量才能比较

### 逻辑运算符

```go
&&  逻辑与
||  逻辑或
!   逻辑非
```

### 位运算符

```go
& 参与运算的两数各对应的二进位相与。（两位均为1才为1）
| 参与运算的两数各对应的二进位相或。（两位有一位为1就为1）
^ 参与运算的两数各对应的二进位相异或。（两位不一样时则为1）
<< 左移n位就是乘以2的n次方。"a<<b"是把a的各二进位全部左移b位，高位丢弃，低位补零
>> 右移n位就是除以2的n次方。"a>>b"是把a的各二进位全部右移b位。
```

常用于IP地址和权限位的计算

### 赋值运算符

```go
= 
+=
-=
*=
/=
%=
<<=
>>=
&=
|=
^=
```



## Go语言复合数据类型

### 数组（array）

数组是同一种数据类型的集合，在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。数组定义的基本语法如下：

```go
var 数组变量名 [元素数量]T
//eg:
var a [3]int
```

* 数组的长度必须是常量，并且长度是数组类型的一部分，一旦定义，长度不能发生改变。`[5]int`和`[10]int`是不同的类型。
* 数组用下标进行访问，下标从`0`开始，最后一个下标是`len-1`，访问越界会panic。
* 数组支持`==`和`!=`操作符，因为内存总是被初始化过的。
* `[n]*T`表示指针数组，`*[n]T`表示数组指针。
* 多维数组只有最外层可以使用`[...]`

```go
package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [4]bool
	//数组的长度是数组类型的一部分，所以这里的a1和a2既不可以比较也不可以赋值
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是零值(bool值就是false，整形和浮点型都是0，字符串就是“”)
	fmt.Println(a1, a2)
	//初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//初始化方式2
	//根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a10)
	//初始化方式3
	//根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//遍历方法1：根据索引来遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//遍历方法2：for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多维数组
	//[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1 //深拷贝
	b1[0] = 100
	fmt.Println(b1, b2)
}
```

Output :

```go
a1:[3]bool a2:[4]bool
[false false false] [false false false false]
[true true true]
[0 1 2 3 4 5 6 7 8]
[1 0 0 0 2]
北京
上海
深圳
0 北京
1 上海
2 深圳
[[1 2] [3 4] [5 6]]
[1 2]
1
2
[3 4]
3
4
[5 6]
5
6
[100 2 3] [1 2 3]
```



### 切片（slice）

* 切片的引入是为了打破数组长度固定这一局限性；
* 切片是一个拥有相同类型元素的可变长度的序列，它是基于数组类型做的一层封装，支持自动扩容；
* 切片是一个引用类型，内部结构包括**地址(指针)**、**长度**和**容量**，一般用于快速地操作一块数据集合；

* 切片申明的基本语法如下：

  ```go
  var name[]T
  ```

  `name`表示变量名；`T`表示切片中的数据类型

* 切片拥有自己的长度和容量，我们可以用内置的`len()`函数求长度，`cap()`函数求切片的容量；

* 切片指向了一个底层数组，切片的长度就是它元素的个数，切片的容量是底层数组从切片的第一个元素到最后一个元素的数量；

* 切片是一个引用类型，他们都指向了底层的一个数组；

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	var s1 []int    // 定义一个存放int类型元素的切片
  	var s2 []string // 定义一个存放string类型元素的切片
  	fmt.Println(s1, s2)
  	fmt.Println(s1 == nil)
  	fmt.Println(s2 == nil)
  
  	//初始化
  	s1 = []int{1, 2, 3}
  	s2 = []string{"北京", "上海", "深圳"}
  	fmt.Println(s1, s2)
  
  	//长度和容量
  	fmt.Println(len(s1), len(s2))
  	fmt.Println(cap(s1), cap(s2))
  
  	//2，由数组得到切片
  	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
  	s3 := a1[0:4] //基于一个数组切割，左闭右开
  	fmt.Println(s3)
  	s4 := a1[:4]
  	s5 := a1[3:]
  	s6 := a1[:]
  	fmt.Println(s4, s5, s6)
  
  	//切片的容量是指底层数组的容量
  	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
  	//底层数组从切片的第一个元素到最后的元素数量
  	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
  	//切片再切片
  	s8 := s5[3:]
  	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
  
  	//切片是一个引用类型，他们都指向了底层的一个数组
  	a1[6] = 1200 //修改了底层数组的值
  	fmt.Println("s5", s5)
  	fmt.Println("s8", s8)
  }
  ```

  Output :

  ```go
  [] []
  true
  true
  [1 2 3] [北京 上海 深圳]
  3 3
  3 3
  [1 3 5 7]
  [1 3 5 7] [7 9 11 13] [1 3 5 7 9 11 13]
  len(s3):4 cap(s3):7
  len(s5):4 cap(s5):4
  len(s8):1 cap(s8):1
  s5 [7 9 11 1200]
  s8 [1200]
  ```

* 除了基于数组来创建切片，我们还可以使用内置的`make()`函数来创建切片，格式如下：

  ```go
  make([]T, size, cap)
  ```

  其中`T`是切片的数据类型，`size`是切片中元素的数量，`cap`是切片的容量。

  **切片的就是一个框，框住了一块连续的内存。**属于引用类型，真正的数据都是保存在底层的数组里面的。

* 切片之间是不能比较的，我们不能用`==`操作符来判断两个切片是否含有全部相等元素。切片唯一合法的比较操作是和`nil`比较，一个`nil`值的切片并没有底层数组，并且长度和容量都是0。但是，一个长度和容量都是0的切片不一定是`nil`，具体示例如下：

  ```go
  var s1 []int               //len(s1)=0;cap(s1)=0;s1==nil
  s2 := []int{}              //len(s2)=0;cap(s2)=0;s2!=nil
  s3 := make([]int, 0)       //len(s3)=0;cap(s3)=0;s3!=nil
  ```

* 要判断一个切片是否为空，要用`len(s) == 0`来判断，不应该使用`s == nil`来判断。

* 拷贝前后的两个切片共享底层数组，对一个切片的修改会影响另一个切片的内容。**浅拷贝**
* 切片的遍历方式和和数组是一致的，支持索引遍历和`for range`遍历。

```go
package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	//切片的复制
	s3 := []int{1, 3, 5}
	s4 := s3 //s3和s4都指向了同一个底层数组
	fmt.Println(s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	//切片的遍历
	//索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	//for range循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}
}
```

Output :

```go
s1=[0 0 0 0 0] len(s1)=5 cap(s1)=10
s2=[] len(s2)=0 cap(s2)=10
[1 3 5]
[1000 3 5] [1000 3 5]
1000
3
5
0 1000
1 3
2 5
```

### 切片的进阶操作-append() 

* Go语言的内建函数`append()`可以为切片动态地添加元素。当切片的底层数组不能容纳新增的元素时，切片就会自动按照一定的规则进行扩容，此时该切片指向的底层数组就会更换。扩容操作往往发生在`append()`函数调用时。

* `append()`函数可以自动初始化切片：

  ```go
  var s []int
  s = append(s, 1) //自动初始化切片
  ```

```go
package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州"
	//错误的写法，会导致编译错误，索引越界

	// 调用append函数，必须用原来的切片变量接收返回值
	s1 = append(s1, "广州") //append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) //...表示把前面的变量拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
```

Output :

```go
s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
s1=[北京 上海 深圳 广州] len(s1)=4 cap(s1)=6
s1=[北京 上海 深圳 广州 杭州 成都] len(s1)=6 cap(s1)=6
s1=[北京 上海 深圳 广州 杭州 成都 武汉 西安 苏州] len(s1)=9 cap(s1)=12
```



### 切片的进阶操作-copy() 

用`copy()`函数来对切片深拷贝

```go
package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1 //赋值  浅拷贝
	var a3 = make([]int, 3, 5)
	copy(a3, a1) //copy 深拷贝
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)
}
```

Output :

```go
[1 2 3] [1 2 3] [1 2 3]
[100 2 3] [100 2 3] [1 2 3]
```



### 从切片中删除元素

Go语言中没有删除切片元素的专用方法，我们可以使用切片自身的特性来删除元素。

```go
	//从切片中删除元素
	a := []int{30, 31, 33, 34, 35, 36, 37}
	//要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 34 35 36 37]
```

总结一下就是：要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:,index], a[index+1]...)`

此时切片的底层数组不变，删除元素的内存发生了覆盖。

* 练习题

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	var b = make([]int, 5, 10)
  	fmt.Println(b)
  	for i := 0; i < 10; i++ {
  		b = append(b, i)
  	}
  	fmt.Println(b)
  }
  
  ```

  ```go
  // output：
  [0 0 0 0 0]
  [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
  ```

* 内置函数对切片排序

  ```go
  var c = []int{3, 1, 5, 8, 4, 2}
  sort.Ints(c)
  fmt.Println(c)
  //output: [1 2 3 4 5 8]
  ```



### 指针

* Go语言中不存在指针操作，只需要记住两个符号：

  1. `&`取地址

  2. `*`根据地址取值

* Go语言中的指针只能读不能写

```go
	var a *int //nil point
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)
	
// output:
// <nil>
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10a4edc]

// goroutine 1 [running]:
// main.main()
//         /Users/beam/LearnGO/src/beamwang/day02/pointer/main.go:20 +0x2dc
// exit status 2
```

* 用`new`申请内存地址，得到指定类型的指针，并且指针对应的值为该类型的零值。

```go
//new函数申请一个内存地址
	var a = new(int)
  fmt.Println(a)
	fmt.Println(*a) //初始化零值
	*a = 100
	fmt.Println(*a)

// output:
// 0xc0000ae020
// 0
// 100
```

* `make`也是用于内存分配的，区别于`new`，它只用于`slice`、`map`以及`chan`的内存创建，而且返回的类型就是这三个类型本身，而不是它们的指针类型。

* `make`和`new`的区别：

  1. `make`和`new`都是用来申请内存的；

  2. `new`很少用，一般给基本的数据类型申请内存，比如int、string...，返回的对应类型的指针(`* int、*string`)；

  3. `make`是用来给`slice`、`map`以及`chan`申请内存的，返回的是对应的这三个类型本身；

### map

Go语言中提供的映射关系的容器叫`map`，其内部用`散列表(hash)`实现。

* `map`是一种无序的基于`key-value`的数据结构，Go语言中的`map`是引用类型，必须初始化才能使用；

* `map`的定义语法为：`map[KeyType]ValueType`。其中`KeyType`表示键的类型，`ValueType`表示键对应的值的类型；

* 判断某个`key`是否存在，方法如下：

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	var m1 map[string]int
  	fmt.Println(m1 == nil)        //还没有初始化，没有在内存中开辟空间
  	m1 = make(map[string]int, 10) //要估算该map的容量，避免程序运行期间动态扩容，不然影响执行效率
  	m1["理想"] = 18
  	m1["姬无命"] = 35
  	fmt.Println(m1)
  	fmt.Println(m1["理想"])
  	fmt.Println(m1["娜扎"]) //如果不存在这个key，拿到对应类型的零值
  	v, ok := m1["娜扎"]     //约定成俗，用ok接收返回bool值
  	if !ok {
  		fmt.Println("查无此key")
  	} else {
  		fmt.Println(v)
  	}
  }
  
  // output:
  // true
  // map[姬无命:35 理想:18]
  // 18
  // 0
  // 查无此key
  ```

* 遍历map，可用`for range`语句，方法如下：

  ```go
  	//map的遍历
  	for k, v := range m1 {
  		fmt.Println(k, v)
  	}
  	//只遍历key
  	for k := range m1 {
  		fmt.Println(k)
  	}
  	//只遍历value
  	for _, v := range m1 {
  		fmt.Println(v)
  	}
  ```

* 删除某一个键值对，用`delete()`函数，格式为`delete(map, key)`，删除不存在的key，就不做操作;

* 按照key值顺序打印字典内容:

  ```go
  package main
  
  import (
  	"fmt"
  	"math/rand"
  	"sort"
  	"time"
  )
  
  func main() {
  	rand.Seed(time.Now().UnixNano()) //随机初始化种子
  
  	var scoreMap = make(map[string]int, 200)
  
  	for i := 0; i < 20; i++ {
  		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
  		value := rand.Intn(100)          //生成0~99的随机整数
  		scoreMap[key] = value
  	}
  	//取出map中的所有key存入切片keys
  	var keys = make([]string, 0, 200)
  	for key := range scoreMap {
  		keys = append(keys, key)
  	}
  	//对切片进行排序
  	sort.Strings(keys)
  	//安装排序后的key遍历map
  	for _, key := range keys {
  		fmt.Println(key, scoreMap[key])
  	}
  }
  ```

  Output :

  ```
  stu00 23
  stu01 37
  stu02 44
  stu03 97
  stu04 65
  stu05 84
  stu06 48
  stu07 8
  stu08 12
  stu09 69
  stu10 51
  stu11 48
  stu12 23
  stu13 88
  stu14 28
  stu15 9
  stu16 13
  stu17 70
  stu18 15
  stu19 18
  ```

* 元素为map类型的切片：

  ```go
  package main
  
  import "fmt"
  
  //map和slice组合
  func main() {
  	//元素类型为map的切片
  	var s1 = make([]map[int]string, 10, 10)
  	s1[0] = make(map[int]string, 1) //对内部的map初始化
  	s1[0][10] = "沙河"
  	fmt.Println(s1)
  }
  
  // output:
  // [map[10:沙河] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
  ```

* 值为切片类型的map：

  ```go
  package main
  
  import "fmt"
  
  //map和slice组合
  func main() {
  	//值为切片类型的map
  	var m1 = make(map[string][]int, 10)
  	m1["北京"] = []int{10, 20, 30}
  	fmt.Println(m1)
  }
  
  // output:
  // map[北京:[10 20 30]]
  ```



* 作业题

  ```go
  package main
  
  import (
  	"fmt"
  	"strings"
  	"unicode"
  )
  
  func main() {
  	//1.判断字符串中汉字的数量
  	//难点是判断一个字符时汉字
  	s1 := "Hello沙河有沙안녕하세요"
  	count := 0
  	for _, c := range s1 {
  		if unicode.Is(unicode.Han, c) {
  			count++
  		}
  	}
  	fmt.Println(count)
  
  	//单词出现的次数统计
  	s2 := "how do you do"
  	lst := strings.Split(s2, " ")
  	m1 := make(map[string]int, 10)
  	for _, w := range lst {
  		if _, ok := m1[w]; !ok { //这种写法ok变量作用域只在if语句中，节省了内存
  			m1[w] = 1
  		} else {
  			m1[w] += 1
  		}
  	}
  	fmt.Println(m1)
  }
  
  // output:
  // 4
  // map[do:2 how:1 you:1]
  ```

  



## 函数

### 函数基本形式

Go语言中支持函数、匿名函数和闭包，用`func`关键字定义，具体格式如下：

```go
func 函数名(参数)(返回值){
	函数体
}
```

* 返回值可以命名也可以不命名；

* 常见函数种类举例：

  ```go
  package main
  
  import "fmt"
  
  //函数的定义
  func sum(i int, j int) int {
  	return i + j
  }
  
  // 没有返回值的函数
  func f1(x int,y int){
  	fmt.Println(x+y)
  }
  
  //没有参数也没有返回值的函数
  func f2(){
  	fmt.Println("f2")
  }
  //没有参数但是有返回值的函数
  func f3() int{
  	return 3
  }
  //命名返回值就当于提前声明一个变量，并且不需要显式返回（return后面的变量名可以省略）
  func f4(x int, y int) (res int) {
  	res = x + y
  	return 
  }
  //多个返回值的函数
  func f5() (int, int) {
  	return 1, 2
  }
  //参数的类型简写(当参数中连续多个参数的类型一致时，我们可以将非最后一个的参数类型省略)
  func f6(x, y int, m, n string, i, j bool) int {
  	return x + y
  }
  //可变长参数的函数，可变长的参数必须放在函数参数的最后
  func f7(x string,y ...int){
  	fmt.Println(x)
  	fmt.Println(y) //y是int类型的切片（...后面指定的）
  }
  
  func main() {
  	fmt.Println(sum(3, 4))
  }
  ```

* Go语言中没有默认参数的设计；

* 在一个命名的函数中不能够再申明命名函数，比如在`main()`函数中就不能再定义其他命名函数；

  

### defer语句

* Go语言的`defer`语句会将其后跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

* `defer`语句多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接等）；

* 在Go语言的函数中，`return`语句在底层不是**原子操作**，它分为给返回值赋值和RET指令两步。而`defer`语句的执行时机在返回值赋值操作后，RET指令执行前。

  ```go
  import "fmt"
  
  func deferdome() {
  	fmt.Println("start")
  	defer fmt.Println("嘿嘿嘿") //defer将该语句延迟到函数即将返回的时候再执行
  	defer fmt.Println("哈哈哈") //一个函数中可以有多个defer语句
  	defer fmt.Println("嘻嘻嘻") //多个defer语句按照先进先出顺序延时执行
  	fmt.Println("end")
  }
  
  func main() {
  	deferdome()
  }
  
  // output:
  // start
  // end
  // 嘻嘻嘻
  // 哈哈哈
  // 嘿嘿嘿
  ```

  ```go
  package main
  
  import "fmt"
  
  func f1() int {
  	x := 5
  	defer func() {
  		x++ //修改的是x不是返回值
  	}()
  	return x
  }
  
  func f2() (x int) {
  	defer func() {
  		x++
  	}()
  	return 5 //返回值=x
  }
  
  func f3() (y int) {
  	x := 5
  	defer func() {
  		x++
  	}()
  	return x //返回值 = y = x = 5
  }
  
  func f4() (x int) {
  	defer func(x int) {
  		x++ //改变的是函数的副本变量x
  	}(x)
  	return 5 //返回值 = x = 5
  }
  
  func f5() (x int) {
  	defer func(x *int) {
  		(*x)++
  	}(&x)
  	return 5
  }
  
  func main() {
  	fmt.Println(f1())
  	fmt.Println(f2())
  	fmt.Println(f3())
  	fmt.Println(f4())
  	fmt.Println(f5())
  }
  
  // output:
  // 5
  // 6
  // 5
  // 5
  ```

  ```go
  package main
  
  import "fmt"
  
  func calc(index string, a, b int) int {
  	ret := a + b
  	fmt.Println(index, a, b, ret)
  	return ret
  }
  
  func main() {
  	a := 1
  	b := 2
  	defer calc("1", a, calc("10", a, b)) // 参数的calc()会先计算出返回值，参数a的值也已经传入
  	a = 0
  	defer calc("2", a, calc("20", a, b))
  	b = 1
  }
  
  // output:
  // 10 1 2 3
  // 20 0 2 2
  // 2 0 2 2
  // 1 1 3 4
  ```

  

### 函数作用域

* 全局变量是定义在函数外部的变量，它在程序的整个运行周期内都有效。在函数中可以访问到全局变量。

  ```go
  package main
  
  import "fmt"
  
  var x = 100 //定义一个全局变量
  
  //定义一个函数
  func f1() {
  	// 函数中查找变量的顺序：
  	// 1.先在函数内部查找
  	// 2.找不到就往函数的外面查找，一直找到全局
  	fmt.Println(x)
  }
  func main() {
  	f1()
  }
  
  // output:
  // 100
  ```

* 局部变量又分为两种：函数内定义的和语句块定义的。

* 函数内定义的变量无法在该函数外使用，如果局部变量和全局变量重名，优先访问局部变量。

* 语句块定义的局部变量，我们通常在`if`条件判断、`for`循环，`switch`语句上使用这种定义变量的方式。



### 函数类型和变量

* 函数可以作为一种自定义变量类型，这个类型包括它的参数和返回值的数量和类型；

* Go语言中，函数可以作为变量和返回值；

  ```go
  package main
  
  import "fmt"
  
  func f1() {
  	fmt.Println("Hello, 沙河")
  }
  func f2() int {
  	fmt.Println("Hello, 沙河")
  	return 10
  }
  
  //函数也可以作为参数的类型
  func f3(x func() int) {
  	ret := x()
  	fmt.Println(ret)
  }
  
  func f4(x, y int) int {
  	return x + y
  }
  
  //函数还可以作为返回值
  func ff(a, b int) int {
  	return a + b
  }
  
  func f5(x func() int) func(int, int) int {
  	return ff
  }
  
  func main() {
  	a := f1
  	fmt.Printf("%T\n", a)
  	b := f2
  	fmt.Printf("%T\n", b)
  	f3(f2)
  	fmt.Printf("%T\n", f4)
  	f7 := f5(b)
  	fmt.Printf("%T\n", f7)
  }
  ```

  

### 匿名函数

函数作为返回值的时候，在Go语言中函数内部不能再定义命名函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```
func(参数)(返回值){
	函数体
}
```

```go
package main

import "fmt"

func main() {
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 29)

	//如果只是调用一次的函数，我们还可以简写成立即执行函数
	func() {
		fmt.Println("Hello,沙河")
	}()
	func(a, b int) {
		fmt.Println(a + b)
		fmt.Println("Hello,沙河")
	}(100, 200)
}

// output:
// 39
// Hello,沙河
// 300
// Hello,沙河
```



### 闭包

* 闭包指的是一个函数和与其相关的引用环境组合而成的实体。

* 闭包的原理基于两点：1.函数可以作为返回值 2.函数内部查找变量的顺序，现在自己内部找，找不到忘外层找；

  ```go
  //应用场景1：包装作为参数传递的函数，使其可以绕过参数类型检测
  package main
  
  import "fmt"
  
  func f1(f func()) {
  	fmt.Println("This is f1")
  	f()
  }
  
  func f2(x, y int) {
  	fmt.Println("This is f2")
  	fmt.Println(x + y)
  }
  
  //定义一个函数f3对f2进行包装，是f2可以传入f1
  func f3(f func(int, int), m, n int) func() {
  	tmp := func() {
  		f(m, n)
  	}
  	return tmp
  }
  
  func main() {
  	f1(f3(f2, 100, 200))
  }
  
  // output:
  // This is f1
  // This is f2
  // 300
  ```

  ```go
  //应用场景2：通过函数传参生成自己想要的函数
  package main
  
  import (
  	"fmt"
  	"strings"
  )
  
  func makeSuffixFunc(suffix string) func(string) string {
  	return func(name string) string {
  		if !strings.HasSuffix(name, suffix) {
  			return name + suffix
  		}
  		return name
  	}
  }
  
  func main() {
  	jpgFunc := makeSuffixFunc(".jpg")
  	txtFunc := makeSuffixFunc(".txt")
  	fmt.Println(jpgFunc("test"))
  	fmt.Println(txtFunc("test"))
  }
  
  // output:
  // test.jpg
  // test.txt
  ```

  ```go
  //应用场景3：多个函数共享同一个局部变量，比如例中的base = 10
  package main
  
  import "fmt"
  
  func calc(base int) (func(int) int, func(int) int) {
  	add := func(i int) int {
  		base += i
  		return base
  	}
  	sub := func(i int) int {
  		base -= i
  		return base
  	}
  	return add, sub
  }
  
  func main() {
  	f1, f2 := calc(10)
  	fmt.Println(f1(1), f2(2))
  	fmt.Println(f1(3), f2(4))
  	fmt.Println(f1(5), f2(6))
  
  }
  
  // output:
  // 11 9
  // 12 8
  // 13 7
  ```

  

### 内置函数

* `close`：主要用来关闭`channel`；

* `len`：用来求长度，比如`string`、`array`、`slice`、`map`、`channel`；

* `new`：用来分配内存，主要用来分配值类型，比如`int`、`struct`。返回的是指针；

* `make`：用来分配内存，主要用来分配引用类型，比如`chan`、`map`、`slice`；

* `append`：用来追加元素到数组、`slice`中；

* `panic`和`recover`：用来做错误处理；

* Go语言中目前是没有异常机制的，但是使用`panic/recover`模式来处理错误。`panic`可以在如何地方引发，但`recover`只有在`defer`调用的函数中有效；

* `recover()`一定要搭配`defer`使用，`defer`一定要在可能引发`panic`的语句之前定义；

  ```go
  package main
  
  import "fmt"
  
  func funcA() {
  	fmt.Println("a")
  }
  func funcB() {
  	//刚刚打开数据库连接
  	defer func() {
  		err := recover()
  		fmt.Println(err)
  		fmt.Println("释放数据库连接...")
  	}()
  	panic("出现严重的错误！！！") //程序崩溃退出
  	fmt.Println("b")
  }
  func funcC() {
  	fmt.Println("c")
  }
  func main() {
  	funcA()
  	funcB()
  	funcC()
  }
  
  // output:
  // a
  // 出现严重的错误！！！
  // 释放数据库连接...
  // c
  ```



## 其他知识点

### fmt标准库

* `fmt`包实现了类似C语言`printf`和`scanf`的格式化`I/O`。主要分为向外输出内容和获取输入内容两大部分；

* `fmt.Scan`从标准输入扫描文本，读取由空白符分割的值保存到传递给本函数的参数中，换行符视为空白符；

* `fmt.Scanln`自动换行符；

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	fmt.Print("长沙")
  	fmt.Print("腾讯")
  	fmt.Println()
  	fmt.Println("----------")
  	fmt.Println("长沙")
  	fmt.Println("腾讯")
  	fmt.Println("----------")
  	//fmt.printf 格式化打印
  	var m1 = make(map[string]int, 10)
  	m1["长沙"] = 1
  	fmt.Printf("%v\n", m1)
  	fmt.Printf("%#v\n", m1)
  	baifenbi := 90
  	fmt.Printf("%d%%\n", baifenbi)
  	fmt.Printf("%t\n", true)
  	//整数->字符
  	fmt.Printf("%q\n", 65)
  	//浮点数和复数
  	fmt.Printf("%b\n", 3.1415926535)
  	//字符串
  	fmt.Printf("%q\n", "海能卑下众水归") //字符串双引号表示
  	fmt.Printf("%5.2s\n", "海能卑下众水归")
  	// var s string
  	// fmt.Scan(&s)
  	// fmt.Println("用户输入的内容是：", s)
  	var (
  		name  string
  		age   int
  		class string
  	)
  	fmt.Scanf("%s %d %s\n", &name, &age, &class)
  	fmt.Println(name, age, class)
  }
  
  // ouput:
  // 长沙腾讯
  // ----------
  // 长沙
  // 腾讯
  // ----------
  // map[长沙:1]
  // map[string]int{"长沙":1}
  // 90%
  // true
  // 'A'
  // 7074237751826244p-51
  // "海能卑下众水归"
  //    海能
  // 汪国庆 20 5班
  // 汪国庆 20 5班
  ```

  


# GoLearning

## 1. 导言

- 没有对象 没有继承多态 没有泛型 没有try/catch
- 有接口 函数式编程 CSP并发模型(gorountine + channel) 

### 1-1 概览

- 基本语法
- 面向接口
- 函数式编程
- 工程化
- 并发编程
- 实战项目

## 2. 基础语法

### 2-1 变量定义

#### 使用var关键字

- ```go
  var a,b,c bool
  var s1,s2 string = "hello", "world"
  ```

- 放在函数内，或直接放在包内

- var()集中定义变量

#### 让编译器自动决定类型

- ```go
  var a,b,c = 12,true,"hello"
  ```

#### :=定义变量

- ```go
  a,b,i := true, 2, "he"
  ```

- 只能在函数内使用

### 2-2 内建变量类型

- bool string int byte(8位) rune(字符32位) float32 float64 complex64 complex128

- 强制类型转换（必须做）

### 2-3 常量与枚举

#### 常量

- const数值可作为各种类型使用

#### 枚举类型

- const() 块定义

### 2-4 条件语句

#### if

- if条件不需要括号

- if的条件可以赋值
- if条件赋值的变量作用域在这个if语句里

#### switch

- switch 自动break，使用fallthrough阻止break
- switch后可以没有表达式

### 2-5 循环

- for 条件不需要括号
- for的初始条件，结束条件， 递增表达式都可以省略

- for 省略初始条件，相当于while

```go
func convertToBin(n int) string {
	result := ""
	for ;n>0 ;n/=2  {
		lsb := n%2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
```

- for 可以省略初始条件和递增条件，相当于while

```go
for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
```

- 死循环：都省略

### 2-6 函数

- 函数可返回多个值
- 返回多个值可以取名字
- 函数式编程  -- 函数的参数也可以是函数

```go
func apply(op func(int,int) int, a,b int) int {
	p:=reflect.ValueOf(op).Pointer()
	opname := runtime.FuncForPC(p).Name()
	fmt.Printf("%s \n", opname)
	return op(a,b)
}
```

- 可变参数列表

```go
func sum(numbers ...int) int{
	s := 0
	for i:= range numbers{
		s += numbers[i]
	}
	return s
}
```

### 2-7 指针

- 指针不能运算

#### 参数传递

- go只有值传递

## 3. 内建容器

### 3-1 数组

```go
var arr1 [5]int
arr2 := [3]int{1,3,5}
arr3 := [...]int{2,3,4,5}
var grid [4][5]int
```

#### 遍历

```go
for i,v := range arr3{
		fmt.Println(i,v)
	}

```

#### 数组是值类型

- 调用函数 数组作为参数 会**拷贝**数组
- golang一般不直接用数组，而是**切片**

### 3-2 slice 切片   

- slice本身没有数据，是对底层的array的一个view

#### reslice

#### slice的扩展

- ptr 、 len 、 cap
- slice可以向后扩展，不可以向前扩展

### 3-3 slice操作

#### 1. 添加元素

- 添加元素时如果超越cap，系统会重新分配更大的底层数组
- 垃圾回收 √
- 由于值传递的关系，必须接收append的返回值

### 3-4 map

- 无序 hashmap

#### 1. key

- map使用哈希表，必须可以比较相等

- 除了slice map function 的内建类型都可以作为key

#### 2. map操作

- 创建：make(map[string]int)
- 获取元素：m[key]
- key不存在时，不会报错，获得value类型的初始值
- 用value, ok := m[key]来判断是否存在key
- 用delete删除一个key

#### 3. map遍历

- 使用range遍历key，或者遍历key，value对
- 不保证遍历顺序，如需顺序，需要手动对key进行排序
- 使用len获得元素

### 3-5 寻找最长不含有重复字符的子串

- lastOccured[x] 不存在，或者<start -> 无需操作
- lastOccured[x] >= start -> 更新start
- 更新lastOccured[x]，更新maxLength

### 3-6 字符和字符串的处理

#### rune 相当于go的char

- 使用range 遍历 pos，rune对
- 使用utf8.RuneCountInString
- 使用len获得字节长度
- 使用[]byte获得字节

#### 字符串其他操作

- strings包里面

## 4. 面向对象

- go只支持封装，不支持继承多态（面向接口编程）
- 没有class 只有**struct**

### 4-1 结构体和方法

- 不论是地址还是结构本身，一律使用 . 来访问成员

- 使用自定义工厂函数，返回了局部变量地址

```go
func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

```

- 定义方法:显示定义和命名方法接收者

```go
func (node *treeNode) setValue(value int) { //传引用
	node.value = value
}

```

- 使用指针作为方法接收者
  - 只有使用指针作为方法接收者才可以改变结构内容
  - nil指针也可以调用方法
- 值接收者 vs 指针接收者

  - 要改变内容必须使用指针接收者
  - 结构过大也要使用指针接收者 (拷贝)
  - **值接收者**go语言特有

### 4-2 封装

- 名字使用CamelCase
- 首字母大写：public
- 首字母小写：private 
- 针对包来说 main是入口
- 每个目录一个包
- main包包含可执行入口
- 为结构定义的方法必须放在同一个包内
- 可以是不同文件

### 4-3 拓展已有类型

- 如何扩充系统类型或者别人的类型
  - 定义别名
  - 使用组合

### 4-4 gopath 以及 目录结构

#### 1. go get

- gopm

## 5. 接口

### 5-1 duck typing

- duck typing
  - 描述事物外部行为而非内部结构

- go属于结构化类型系统

### 5-2 接口定义和实现

- 接口由**使用者**定义
- 接口的实现是隐式的
- 只实现接口里的方法

### 5-3 接口的值类型

#### 1. 接口变量里有什么

- 实现者的类型
- 实现者的值/指针

#### 2. 查看接口变量

- interface{} 表示任何类型

### 5-4 接口的组合



### 5-5 常用系统接口

- Stringer
- Reader
- Writer



## 6. 函数式编程

### 6-1 函数式编程

#### 1. 函数式编程 vs 函数指针

- 函数是一等公民：参数、变量、返回值都可以是函数
- 高阶函数
- 函数 -> 闭包

#### 2. 正统函数式编程

- 不可变性：不能有状态，只有常量和函数
- 函数只能有一个参数

#### 3. 闭包

- 函数体 
- 局部变量
- 自由变量

### 6-2 例题一



## 7.资源管理与出错处理

### 7-1 defer调用

- 确保调用在函数结束时发生 
- defer里面相当于有一个栈，先进后出
- 不怕return 和 panic

- 参数在defer语句时计算

#### 1. 何时使用defer调用

- Open/Close
- Lock/Unlcok
- PrintHeader/PrintFooter

### 7-2 错误处理概念

```go
file, err := os.Open("abc.txt")
if err != nil{
    if pathError, ok := err.(*os.PathError); ok {
        fmt.Println(pathError.Err)
    } else {
        fmt.Println)("unknown error", err)
    }
}

```



### 7-3 统一错误处理

### 7-4 panic 和 recover

#### panic

- 停止当前函数执行
- 一直向上返回，执行每一层defer
- 如果没遇到recover，程序退出

#### recover

- 仅在defer调用中使用
- 获取panic的值
- r := recover()

### 7-5 统一出错处理2



## 8 测试

### 8-1 测试

#### 传统测试 vs 表格驱动测试

##### 表格优点

- 分离的测试数据与测试逻辑
- 明确的出错信息
- 可以部分失败

### 8-2 代码覆盖率 和 性能测试

- benchmark *Testing.B

 



## 9 并发编程

### 9-1 goroutine

#### 协程 Coroutine

- 轻量级线程

- **非抢占式**多任务处理，由协程主动交出控制权
- 编译器层面的多任务，go有调度器
- 多个协程可能在一个或多个线程上运行:happy:

### 9-2 go的调度器

#### 普通函数 vs 协程

- 函数调用是协程的特例

#### goroutine 定义

- 在任何函数前加上 **go**关键字
- 不需要在定义时区分函数是否异步
- 调度器在合适的点进行切换 

#### goroutine可能切换的点

- I/O, select
- channel
- 等待锁 
- 函数调用
- Runtime.Gosched()

## 10 channel

### 10-1 channel

- channel 是 goroutine之间的交互
- buffered channel
- range
- 理论基础：CSP Communication Sequential Process

### 10-2 使用channel等待任务结束



## 11 其他标准库

### 11-1 http




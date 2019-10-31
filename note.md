# GoLearning
## 1. 基础语法

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

- 强制类型转换（必须做）

### 2-3 常量与枚举

#### 常量

- const数值可作为各种类型使用

#### 枚举类型

- const() 块定义

### 2-4 条件语句

#### if

- if的条件可以赋值
- if条件赋值的变量作用域在这个if语句里

#### switch

- switch 自动break，使用fallthrough阻止break
- switch后可以没有表达式

### 2-5 循环

- for 可以省略初始条件，相当于while

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

- 调用函数会**拷贝**数组

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

- 除了slice map function 的内建类型都可以作为key

### 3-5 寻找最长不含有重复字符的子串

- lastOccured[x] 不存在，或者<start -> 无需操作
- lastOccured[x] >= start -> 更新start
- 更新lastOccured[x]，更新maxLength

### 3-6 字符和字符串的处理

- rune 相当于go的char

## 4. 面向对象

- go只支持封装，不支持继承多态（面向接口编程）
- 没有class 只有struct

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


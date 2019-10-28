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


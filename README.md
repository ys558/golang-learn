# go 语言学习笔记
### 1. 代码风格
---
1.1 不强制规则：

1.1 运算符左右需要空格：
```go
var a string = "a变量"
```
1.2 小驼峰式命名:
```go
var nameVariable
```

1.3 强制规则：左括号必须紧接着语句不换行：

```go
import "fmt"

func main() {
    fmt.Println("hello world!")
}
```

1.4 在命令行格式化代码：`go fmt <文件名>` 如：

```shell
go fmt main.go
```

### [2. 变量 常量](https://github.com/ys558/golang-learn/blob/master/02var%26const/Variable%26Contant.go)

### [3. 基本数据类型](https://github.com/ys558/golang-learn/blob/master/03-baseDataType/main.go)

### [4. 基本数据类型的转换](https://github.com/ys558/golang-learn/blob/master/04-baseDatasConvert/main.go)
### [5. 运算操作符]()
---
import时包名可以写为`import (. "fmt")`, 用.代替 包名fmt

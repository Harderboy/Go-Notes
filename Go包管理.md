# Go 包管理

包（package）是多个 Go 源码的集合，是一种高级的代码复用方案，像 fmt、os、io 等这样具有常用功能的内置包在 Go语言中有 150 个以上，它们被称为标准库，大部分（一些底层的除外）内置于 Go 本身。

知识点：

- 包的作用
- 如何导入包
- 导入包变量：导入同一个包的变量、导入不同包的变量
- init()函数的作用
- 程序初始化的过程：调用包的顺序、包中变量初始化顺序

TODO：

- [ ] 总结使用过程遇到的坑

## 知识点总结

**main包**

Go 语言的入口 main() 函数所在的包（package）叫 main，main 包想要引用别的代码，需要import导入！

**包的特性如下：**

- **一个目录下的同级文件归属一个包**。
- 包名可以与其目录不同名。
- 包名为 main 的包为应用程序的入口包，编译源码没有 main 包时，将无法编译输出可执行的文件。

**对引用自定义包需要注意以下几点：**

- **使用 import 语句导入包时，使用的是包所属文件夹的名称，使用绝对路径（起始于工程根目录）**；
- 包中的函数名第一个字母要大写，否则无法在外部调用；
- **自定义包的包名不必与其所在文件夹的名称保持一致，但为了便于维护，建议保持一致**；
- 调用自定义包时使用 `包名.函数名` 的方式。

**import 使用**

要引用其他包，可以使用 import 关键字，可以单个导入或者批量导入，常用的导入方式有以下几种：

1. 通常导入

    ```Go
    // 单个导入
    import "package"
    // 批量导入
    import (
    "package1"
    "package2"
    )
    ```

2. 点操作我们有时候会看到如下的方式导入包

    ```go
    import(
        . "fmt"
    ) 
    ```

    这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调

    用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`

3. 起别名

    别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字。导入时，可以为包定义别名，语法演示：

    ```go
    import (
    p1 "package1"
    p2 "package2"
    )
    // 使用时：别名操作，调用包函数时前缀从原来的包名变成了我们起的别名
    // 可以用于包名冲突的情况，比如https://www.jianshu.com/p/07ffc5827b26中的例子
    p1.Method()
    ```

4. `_`操作如果仅仅需要导入包时执行初始化操作，并不需要使用包内的其他函数，常量等资源。则可以在导入包时，匿名导入。

    这个操作经常是让很多人费解的一个操作符，请看下面这个import：

    ```go
    import (
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
    )
    ```

    `_`操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的`init函数`。也就是说，使用下划线作为包的别名，会仅仅执行`init()`。

**注：导入的包的路径名，使用绝对路径（起始于工程根目录）。**

导入的包分为本地依赖包、网络依赖包，

- 如果需要导入的本地包是其他模块，需要在go.mod中使用replace将依赖包替换成绝对路径，参考 **[如何使用go module导入本地包](https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/)**、**[golang包管理](http://masaka.tech/golang%E5%8C%85%E7%AE%A1%E7%90%86/)**
- 如果 `GO111MODULE=on`即存在`go.mod`时，导入同一模块下的普通包，导入路径为：`模块名称/包所在目录`（引入相对该module的路径），比如`import "go-demo/numberdemo"`，其中`go-demo` 为module名称，`numberdemo` 为要引入的包所在的目录，如果`GO111MODULE=off`或者`GO111MODULE=auto`时，且项目代码不在`$GOPATH/src`下时，支持通过相对路径来导入包，如导入当前目录下的utils包，只需`import "./utils"`即可，注意`"./utils"`指的是要导入的包所在路径（文件夹/目录），包名可能不是utils，代码中使用对应包名即可。（文件夹名和包名最好一致是最佳实践（好习惯））
- 其他常见错误，参考 [Golang Package 与 Module 简介](https://www.jianshu.com/p/07ffc5827b26)

参考：

- [Go语言包package管理](https://www.jianshu.com/p/1a2d471bd71c)
- [五分钟理解golang的init函数](https://zhuanlan.zhihu.com/p/34211611)
- **[如何使用go module导入本地包](https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/)**
- **[golang包管理](http://masaka.tech/golang%E5%8C%85%E7%AE%A1%E7%90%86/)**
- [Go语言自定义包](http://c.biancheng.net/view/5123.html)
- [GO语言基础进阶教程：包的使用](https://zhuanlan.zhihu.com/p/71822746)
- ["package XXX is not in GOROOT" when building a Go project](https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project)
- [Golang Package 与 Module 简介](https://www.jianshu.com/p/07ffc5827b26)

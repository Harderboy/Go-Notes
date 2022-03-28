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

- **使用 import 语句导入包时，使用的是包所属文件夹的名称**；
- 包中的函数名第一个字母要大写，否则无法在外部调用；
- **自定义包的包名不必与其所在文件夹的名称保持一致，但为了便于维护，建议保持一致**；
- 调用自定义包时使用 `包名.函数名` 的方式。

参考：

- [Go语言包package管理](https://www.jianshu.com/p/1a2d471bd71c)
- [五分钟理解golang的init函数](https://zhuanlan.zhihu.com/p/34211611)
- [如何使用go module导入本地包](https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/)
- [Go语言自定义包](http://c.biancheng.net/view/5123.html)
- [GO语言基础进阶教程：包的使用](https://zhuanlan.zhihu.com/p/71822746)
- ["package XXX is not in GOROOT" when building a Go project](https://stackoverflow.com/questions/61845013/package-xxx-is-not-in-goroot-when-building-a-go-project)

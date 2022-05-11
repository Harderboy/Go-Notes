//go 语言中的一些内置的集合类型就在暗中使用指针

/*
map在被赋值或者作为参数传递的时候就不会被复制
map是一种隐式指针
这种写法就多此一举 func demolish(planets *map[string][string])
map的键和值都可以是指针类型
需要将指针指向map的情况并不多见
*/
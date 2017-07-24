package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}

/*
    命名返回值

    Go 的返回值可以被命名, 并且就像在函数体开头声明的变量那样使用

    返回值的名称应对具有一定意义, 可以作为文档使用。

    没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作裸返回

    直接返回语句仅应当用在像例子中的短函数中。在长函数中绘影响代码可读性
*/

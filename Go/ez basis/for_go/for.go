package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 45

}

/*
   for

   Go 只有一种循环结构, for 循环

   基本的 for 循环包含三个由分号分开的组成部分：
        初始化语句： 在第一个循环执行前被执行
        循环条件表达式： 每轮迭代开始前被求值
        后置语句： 每轮迭代后被执行

    初始化语句一般是一个短变量声明, 这里声明的变量仅在整个 for 循环里语句可见

    如果条件表达式的值为 false , 迭代将禁止

    注意： 不像C, Java, 或者 JS 等语言, for语句的三个组成部分并不需要用括号括起来,
         但循环体必须用 { }括起来
*/
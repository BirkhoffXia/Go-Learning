【8-1-作用域】
=====================================================================
[作用域]
函数会开辟一个局部作用域，其中定义的标识符仅能在函数之中使用，也称为标识符在函数中的可见范围
这种对标识符约束的可见范围称为作用域。
1.语句块作用域
2.显式的块作用域：在任何一个大括号中定义的标识符，其作用域只能在这对大括号中
3.universe块：宇宙块，意思就是全局块，不过是语言内建的。预定义的标识符就在这个全局环境中
							因此bool、int、nil、true、false、iota、append等标识符全局可见，随处可用。
4.包块：每一个package包含该包所有源文件，形成的作用域。有时在包中顶层代码定义标识符，也称为全局标识符。
				所有包内定义全局标识符，包内可见。包的顶层代码中标识符首字母大写则导出，从而包外可见，使用时也要加上包名。例如 fmt.Prinf()。
5.函数块：函数声明的时候使用了花括号，所以整个函数体就是一个显式代码块。这个函数就是一个块作用域.

package main

import "fmt"

//包级常量，变量定义，只能使用const，var定义
const a = 100

var b = 200

// c:= 300 //不可以使用短格式
var d = 400

func showB() int {
	return b
}

func main() {
	fmt.Println(1, a) //1 100
	// fmt.Println(1.1, &a) //注意常量不可寻址，这是对常量的保护 invalid operation: cannot take address of a (untyped int constant 100)

	var a = 500
	fmt.Println(2, a, &a) //2 500 0xc0000180b0 重新定义a为变量可以访问

	//对b的操作，思考以下 是否b变了？
	fmt.Println(3, b, &b) //3 200 0xc87340
	b = 600
	fmt.Println(3.1, b, &b) //3.1 600 0x9f8340 地址没有发生变化 值给替换为600
	b := 700
	fmt.Println(3.2, b, &b)   //3.2 700 0xc000122098	因为重新定义了 地址发生变化值发生变化
	fmt.Println(3.3, showB()) // 3.3 600 创建栈帧时已经把全局的200地址记录了 之后200改为600了 就是600 而不是main中的b=700

	{
		const j = 'A'
		var k = "magedu"
		t := true
		a = 700
		b := 800
		fmt.Println(4, a, b, d, j, k, t) //4 700 800 400 65 magedu true
		{
			x := 900
			fmt.Println(4.1, a, b, d, j, k, t, x) //4.1 700 800 400 65 magedu true 900
		}
		// fmt.Println(4.2, x) //无作用域 块外不可见
	}
	// fmt.Println(4.3, j, k, t) //无作用域 错误
	fmt.Println(4.4, a, b) //700 700

	for i, v := range []int{1, 3, 5} {
		fmt.Printf("Index = %d,Type = %T,value = %[2]d\n", i, v)
	}
	// Index = 0,Type = int,value = 1
	// Index = 1,Type = int,value = 3
	// Index = 2,Type = int,value = 5
	// fmt.Println(i,v) //i,v不可见，他们zaifor的作用域中
}

[总结]
标识符作用域
	标识符对外不可见，在标识符定义所在作用域外是看不到标识符的
	使用标识符，自己这一层定义的标识符优先，如果没有，就向外层找同名标识符--自己优先，由近及远
	标识符对内可见，在内部的局部作用域中，可以使用外部定义的标识符--向内穿透
	包级标识符
		在所在包内，都可见
		跨包访问，包级标识符必须大写开头，才能导出到包外，可以在包外使用xx包名.varName 方。式访问。例如 fmt.print()

【8-2-匿名函数】
=====================================================================
匿名函数
	没有名字的函数
	用途:主要用在高阶函数中，函数化编程strings.Map，非常高级，逻辑的高度抽象
高阶函数
	某一个形参类型是函数
		处理逻辑外移，留给使用者
	某一个返回值类型是函数
	上面2个条件满足其一

package main

import (
	"fmt"
	"strings"
)

func add(x, y int) int {
	return x + y
}

// 可以定义类型
type MyFunc func(a, b int) int

func calc1(x, y int, fn MyFunc) int {
	return fn(x, y)
}

// %T func(a,b int) int 函数的类型，签名
func calc(x, y int, fn func(a, b int) int) int {
	return fn(x, y)
}

func main() {
	//需要对2个int整数做某种计算
	fmt.Println(calc(4, 5, add)) //9
	//通过匿名函数进行调用
	fmt.Println(calc(4, 5, func(a, b int) int { return a - b })) //-1

	//strings.Map
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "abc")) //bcd
}

【8-3-递归函数】
=====================================================================
递归要求
	递归一定要有退出条件，递归调用一定要执行到这个退出条件。没有退出条件的递归调用，就是无限调用
	递归调用的深度不宜过深
	Go语言不可能让函数无限调用，栈空间终会耗尽
		goroutine stack exceeds 1000000000-byte limit

package main

import (
	"fmt"
)

func fibLoop(n int) int { //循环
	switch {
	case n < 0:
		panic("Panic : N is Negetive")
	case n == 0:
		return 0
	case n < 3:
		return 1
	}
	//n >=3
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b

}

// 递归版本 1.采用递推公式
func fib1(n int) int {
	if n < 3 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 2.循环层次变成递归函数层次
func fib2(n, a, b int) int {
	if n < 3 {
		return b
	}
	return fib2(n-1, b, a+b)
}

func main() {
	for i := 0; i < 11; i++ {
		fmt.Println(i, fibLoop(i))
	}
	// 	0 0
	// 1 1
	// 2 1
	// 3 2
	// 4 3
	// 5 5
	// 6 8
	// 7 13
	// 8 21
	// 9 34

	//1.
	value := fib1(10)
	fmt.Println(value)

	//2.
	value = fib2(10, 1, 1)
	fmt.Println(value) //55
}

[递归效率]
以上3个斐波那契数列实现，请问那个效率高?递归效率一定低吗?哪个版本好?
	递归版本1效率极低，是因为有大量重复计算。
	递归版本2采用了递归函数调用层次代替循环层次，效率还不错，和循环版效率差不多
递归版2和循环版谁好?
	循环版好些，因为递归有深度限制，再一个函数调用开销较大。
[总结]
	递归是一种很自然的表达，符合逻辑思维
	递归相对运行效率低，每一次调用函数都要开辟栈帧
	递归有深度限制，如果递归层次太深，函数连续压栈，栈内存就可能溢出了如果是有限次数的递归，可以使用递归调用
		或者使用循环代替，循环代码稍微复杂一些，但是只要不是死循环，可以多次迭代直至算出结果
	绝大多数递归，都可以使用循环实现
	即使递归代码很简洁，但是能不用则不用递归

【8-4-嵌套函数和闭包】
=====================================================================
package main

import "fmt"

func outer() {
	c := 65
	inner := func() {
		fmt.Println("1 inner c=", c)
	}
	// fmt.Println(inner)     //0x709a20
	inner()                      //	1 inner c= 65
	fmt.Println("2 outer c=", c) //	2 outer c= 65
}
func main() {
	outer()
}


package main

import "fmt"

func outer() {
	c := 65
	inner := func() {
		c = 97 //此时97是上面c的重新赋值
		fmt.Println("1 c=", c, &c)
		c := 0x31
		fmt.Println("3 c=", c, &c) //此时的c是局部变量 后面销毁了

	}
	// fmt.Println(inner)     //0x709a20
	inner()                    //1 c= 65
	fmt.Println("2 c=", c, &c) //2 c= 65

}
func main() {
	outer()
}

// 1 c= 97 0xc000122058
// 3 c= 49 0xc000122078
// 2 c= 97 0xc000122058


[闭包] - GO了解以下即可 Python和JS需要装饰器需要
	嵌套函数
	自由变量：不在本地作用域创建的局部变量，在外层的函数作用域中创建的局部变量
	闭包：内存函数用到了外层函数中定义的自由变量，条件成了，形成了闭包
	**闭包形成以后有变量逃逸情况，如下。
	栈
		栈回收了一部分内存栈帧消亡
	堆(垃圾检测)
		只要有人记得你，你 引用计数不为0，不是垃圾
		gc清理
package main

import "fmt"

func outer() func() { //高阶函数
	c := 65
	fmt.Println("1 c=", c, &c)
	inner := func() {
		fmt.Println("2 c=", c, &c)
	}
	fmt.Println("inner address :", inner)
	return inner
}
func main() {
	var fn = outer()
	fmt.Println("fn address :", fn) //fn 记录了outer本地调用临时创建了inner函数本体，被谁fn利用，引用计数
	fn()
}

	第15行调用outer函数并返回inner函数对象，并使用标识符fn记住了它。outer函数执行完了，其栈帧上的局部变量应该释放，包括inner函数，因为它也是局部的。
		但是，c、inner对应的值都不能释放，因为fn要用。所以这些值不能放在栈上，要放到堆上。在Go语言中，这称为变量逃逸，逃逸到堆上
	在某个时刻，fn函数调用时，需要用到c，但是其内部没有定义c，它是outer的局部变量，如果这个c早已随着outer的调用而释放，
		那么fn函数调用一定出现错误，所以，这个outer的c不能释放但是outer已经调用完成了，怎么办?闭包，让inner函数记住自由变量c(逃逸到堆上的内存地址)
		
【8-5-defer】
=====================================================================
defer应用场景:资源释放。例如文件打开后要关闭、网络连接后要断开、获得锁用完后的释放等
              以上场景中，获得资源后紧跟着写defer语句，以确保函数正常退出或panic时，能够释放资源。

defer意思是推迟、延迟。语法很简单，就在正常的语句前加上defer就可以了。
defer执行时机:在某函数中使用defer语句，会使得defer后跟的语句进行延迟处理
							当该函数即将返回时，或发生panic时，defer后语句开始执行。
	            注意os.Exit不是这两种情况，不会执行defer
defer执行顺序:同一个函数可以有多个defer语句，依次加入调用栈中(LIFO)，函数返回或panic时，从栈顶依次执行defer后语句。
	执行的先后顺序和注册的顺序正好相反，也就是后注册的先执行defer后的语句必须是一个函数或方法的调用。

defer fn() 以后去执行fn ，fn调用会延后
函数中任意位置
执行顺序
	按照注册的顺序，注册到执行栈，到这执行 LIFO 后进先出

package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
	// 	start
	// end
	// 3
	// 2
	// 1
}

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("我从2之后崩溃了")      //代码从此不再向后执行
	log.Fatal("从2之后fatal") //os.Exst(1)
	defer fmt.Println(2)
	fmt.Println("end")
	// start
	// 2
	// 1
	// panic: 我从2之后崩溃了
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	count := 1
	defer fmt.Println(count) //1 注册时注入，计算 计算好马上记录下来
	count++
	defer fmt.Println(count) //2
	count++
	defer fmt.Println(count) //3
	fmt.Println("end")
	// start
	// end
	// 3
	// 2
	// 1
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	count := 1
	defer func() { 
		fmt.Println(count) 
	}() //无参函数，这个函数不需要注册，最后返回count = 3
	count++
	defer fmt.Println(count)
	count++
	defer fmt.Println(count)
	fmt.Println("end")
	// start
	// end
	// 3
	// 2
	// 3
}

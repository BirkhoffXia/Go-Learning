【1-Go开发环境】
##Install ： Go:1.20.7 （https://go.dev/dl/）

##配置环境变量
	GOPATH 
	GOROOT
	GOPROXY
		https://mirrors.aliyun.com/goproxy/,direct
		https://proxy.golang.com.cn,direct(如果要装插件 改为这个环境变量即可)
		go env
	GOMODCACHE=C:\Users\BIRKHOFF ALW\go\pkg\mod

##本地开发环境
set GOMODCACHE=C:\Users\BIRKHOFF\go\pkg\mod
set GOPATH=C:\Users\BIRKHOFF\go
set GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
set GOROOT=D:\Software Install\Go

##Git 
	
##VScode
        Go 插件
	GOPROXY: https://proxy.golang.com.cn,direct(如果要装插件 改为这个环境变量即可)
	cmd:安装可能需要梯子
		go install -v golang.org/x/tools/gopls@latest 
		go install -v github.com/cweill/gotests/gotests@v1.6.0
		go install -v github.com/fatih/gomodifytags@v1.16.0 
		go install -v github.com/go-delve/delve/cmd/dlv@latest

##写第一个GO程序
package main

import "fmt"

func main() {
	fmt.Printf("BIRKHOFF FIRST GO PROCESS")
}

CTRL + J : 启动命令行

##启动编译可执行文件
PS E:\goprojects> go build main.go
PS E:\goprojects> .\main.exe
BIRKHOFF FIRST GO PROCESS
PS E:\goprojects> 

##表示编译并执行代码
go run main.go

##F5执行 run-debug模式 会报错
#表示没有找到文件
Build Error: go build -o e:\goprojects\__debug_bin1722392392.exe -gcflags all=-N -l .
go: go.mod file not found in current directory or any parent directory; see 'go help modules' (exit status 1)
#需要初始化、管理第三方包依赖的
E:\goprojects>go mod init test 产生go.mod
go: creating new go.mod: module test
go: to add module requirements and sums:
        go mod tidy
#再看F5 

##outline map插件-方便为了查看源码函数

##Golang postfix插件-方便后期开发方便 

=========================================================================================================================
【2-标识符和常量】
# TODO TREE

#var是变量 const是常量无法改变会报错
	var a = 2000
	a = 4000

	const b = 319
	b = 1220

#单引号只能写一个字符

package main

import (
	"fmt"
)

func main() {
	fmt.Println("BIRKHOFF FIRST GO PROCESS")
	fmt.Printf("\"abc\": %v\n", "abc") //postfix

	// TODO 此模块下次再确认
	// NOTE 提醒
	// DEPECATED 未来此版本函数废弃

	var height = 175
	var weight = 83
	//var bmi = 0
	if height >= 175 {
		var bmi = height / weight
		fmt.Println(weight)
		fmt.Println(height)
		fmt.Println(bmi)
	}

	//字面常量：为了方便变成建立的表示方式
	//常量，值，恒定不变的量
	//无类型常量untyped constant 、缺省类型为bool、rune、int、float64、complex128或者字符串
	// 字符类型 rune ''单引号只能一个字符
	// rune 字符类型本质是证书 就是int32 4个字节的整数
	fmt.Println('\u6d4b')
	fmt.Println('\n')
	fmt.Println('\x31')
	fmt.Println('1')
	//字符串可以里有多个字符
	fmt.Println("")
	fmt.Println("BIRHOFF")
	fmt.Println("\x61d\x63") //abc
	fmt.Println("\u6d4b试")   //测试
	fmt.Println("\n")

	//批量赋值 、类型推导 字面常量缺省值赋值给左边
	// const (
	// 	a int8 = 100
	// 	b      = 1220
	// 	c      = "abc"
	// )
	// fmt.Println(a, b, c)

	//itoa：用在成批定义的时候，一旦在成批定义时用到了iota，它就相当于行索引
	// const a = iota
	// const b = iota
	// fmt.Println(a) //0
	// fmt.Printf("%T", b) // int
	// const (
	// 	a = iota
	// 	b //b=iota 如果和上一行公式一样，可以省略
	// 	c //c=iota
	// )
	// fmt.Println(a) //0
	// fmt.Println(b) //1
	// fmt.Println(c) //2
	// const (
	// 	SUN = iota
	// 	MON
	// 	TUE
	// 	WES
	// )
	// fmt.Println(SUN, MON, TUE, WES) // 0 1 2 3
	//iota看作是行索引
	//_：是空表标识符 为匿名变量赋值 其值会被抛弃 因后续代码不能使用匿名变量的值，也不能使用匿名变量为其它变量赋值 是一个黑洞 合法标识符
	// const (
	// 	x = 100
	// 	y
	// 	a = iota
	// 	b
	// 	c
	// 	_
	// 	d = 1220
	// 	e
	// 	f
	// 	g = iota
	// )
	// fmt.Println(a, b, c, d, e, f, x, y, g) //2 3 4 1220 1220 1220 100 100 9
}
=========================================================================================================================
【2.1-标识符】
[本质]
	指代，指向任何合法值
	程序员编程时用的，写代码的时候
	编译后，标识符是什么?所有的值都要放在内存中，标识符指向这个值，标识符不见了，因为都换成了内存地址
	但是 标识符是指定，是名称，不是指针类型，标识符类型就是指代的值的类型
	指针类型是Go中的类型
	标识符最终是为了找到那个值
	
var p *int //p是标识符，p的类型是值的类型，*int nil，未来要和*int类型的数据建立关联关系，指向这个值
var p a_type //p是标识符，指向某个类型的值，p编译后编程内存地址

package main

import "fmt"

/*
//首字母、x首字母小写，整个main包里可见，包内全局变量x
//Y首字母大写，不但在main包内可见，还可以在包外可见
//Go中，我们定义的标识符，最大的全局是哪里?包
//x,Y都是main包的全局变量，只不过包外只能使用Y，使用方法 main.Y
var x = 100   //变量定义，包内顶层代码，全局变量，最好初始化
const Y = 200 //全局常量
*/

//z:=300 //短格式变量定义，不能全局使用，只能用在函数里面 快捷写法的

func main() {
	/*
		//打印全局变量
		fmt.Println(x, Y) //100 200
		//Println 首字母大写，表示标识符是包外可见的fmt这个包的全局标识符
	*/

	/*如果一个变量标识符不初始化，go会替我们初始化为该类型的零值、称为 零值可用*/
	//var a 错 不知道类型、不知道值、因为标识符需要有个指代
	// var a int      //可以 但是没有使用、c c++等、制作了标识符定义但没有赋初始值称为声明
	// fmt.Println(a) //0

	// var b string   //初始化 声明
	// fmt.Println(b) //空

	// var c bool     //一旦确定了声明的标识符类型，go为了不出错 提供了一个 该类型的 零值
	// fmt.Println(c) //false

	//const d int //常量标识符定义必须初始化

	/*
		//Go 不支持再次赋值不一样的类型
		var test = 100
		fmt.Println(test)
		fmt.Printf("test = %T\n", test)
		// test = "HEYE" //不可以类型变了 编译会报错
		// fmt.Printf("test = %T", test)
	*/

	/*
		//var a int,b int //这样写法不行
		//var a,b int ,c string //不行 只能同一类型
		//var a,b int // a,b都是int类型
	*/

	/*
		//零值可以用 : 只声明了变量的类型 没有赋初值(初始化) Go提供该类型
		//bool false int 0 float64 0 string ""
		//var a int
		//var b string
	*/

	/*
		//带初始化
		// var a = 100
		// var b int = 200         // var b int = int(200)
		// var c, d int = 300, 400 //没有类型推导 var c, d = 300, 400 int型
		// var j, k int8 = 30, 40  // var j,k int8 = int8(30),int8(40)
		// var e, f = 500, "abc"   //类型推导了
		// var g int,f string = 600, "abcdefg" //错误
		//fmt.Println(a, b, c, d, e, f, j, k)
	*/

	/*
		var t = nil //不可以这样写
		fmt.Println(t)
	*/

	//批量写
	/*
		var ( //变量有业务相关性
			a           = 100
			b    string = "abc"
			c, d        = 200, "abc"
			g, h int    = 400, 500
		)
		fmt.Println(a, b, c, d, g, h)
	*/

	/*
		//下划线 可以接受数值为黑洞 但不能打印
		var a, _ = 200, 400
		fmt.Println(a)    //200
		fmt.Println(a, _) //cannot use _ as value or type
	*/

	/*
		//短格式，海象符
		a := 100
		b := "abc"
		fmt.Println(a, b) //100 abc
	*/

	/*
		//短格式 这是声明，初始化   a:=100 等于 var a = 100
		a, b := 100, "abc"
		fmt.Println(a, b) //100 abc
		// b := "xyz"        //error 重复定义
		//var b = "xyz"     //重复
	*/

	/*
		//已经赋值不可以在短格式赋值
		var a int
		a := 200
	*/

	/*
		//此赋值不对 少:  :=
		a, b = 100, "xyz"
		fmt.Println(a, b)
	*/

	/*
		//这样可以 因为已经声明勒 再次赋值
		a, b := 100, "xyz"
		fmt.Println(a, b)//100 xyz
		a, b = 200, "bva"
		fmt.Println(a, b) //200 bva
	*/

	/*
		//同类型可以交换
		a, b := 100, 300
		a, b = b, a
		fmt.Println(a, b) //300 100
	*/

	//使用中间变量
	a, b := 100, 300
	c := a
	a = b
	b = c
	fmt.Println(a, b) //300 100

}
=========================================================================================================================
【2.2 - 整数和进制】
0x32 转为十进制 3*16 + 2 = 50
50的2进制
128 64 32 16 8 4 2 1   
00110010
十六进制 4个一组
0011 0010

8421 8421
0011 0010
-----------
    3   2
0x32


package main

import "fmt"

func main() {
	/*
		//bool 型就是 true 和 false
		a := false
		fmt.Println(a + 100) //一个a 是 bool不能够进行加数字
	*/

	/*
		a := 100          //int 64bits 8 Bytes有符号
		var b int64 = 200 //int64 64bits 8 Bytes有符号 var b int64 = int64(200)
		fmt.Println(a, b) // 100 200
		//fmt.Println(a + b) 不允许因为a int 、b int64 类型不一致不能相加
		fmt.Println(b + int64(a)) //int64(a)
		fmt.Println(a + 200) //300
		fmt.Println(b + 300) //也可以 因为300 无类型字面常量 可以进行隐式转换
	*/

	/*
		var a int8 = 127 // 可以
		fmt.Println(a)
		//var b int8 = 128 //不可以 超了范围勒 因为是有符号 有符号最大127 0111 1111 前面的0为正数
		//fmt.Println(b)
		var c uint8 = 128 // 可以因为是无符号 ：1111 1111 前面1为占位符 可以最大为255
		fmt.Println(c)
		//var d uint8 = -1 // uint8 无符号不能是负数
		//fmt.Println(d)
	*/

	/*
		var a = 1 * 2.3 //为什么可以？ 因为右边使用的都是无类型常量untyped constant，它会在上下文中隐式转换。Go为了方便，不能过于死板，要减少程序员转换类型的负担，在无类型常量上做了一些贴心操作。
		fmt.Println(a)
		//占位符 %T type 、%f float 、%v value 调用该类型默认打印格式
		fmt.Printf("Type = %T , Equals = %v\n", a, a) //Format Type = float64 , Equals = 2.300000

		var b = 1
		//fmt.Println(b * 2.4) //不允许 b是int型 2.4是float32不行
		fmt.Println(float32(b) * 2.4)             //可以 进行转换=> 2.4
		fmt.Printf("Type = %T\n", float32(b)*2.4) //Type = float32
	*/

	var a = 1
	fmt.Printf("a's type = %T \n", a) //int

	//强制类型转换:把一个值从一个类型强制显式转换到另一种类型，有可能转换失败。
}

=========================================================================================================================
【2.3 - 字符串及格式化】
package main

func main() {
	/*
		var a = 3.1415926
		fmt.Printf("%T %v\n", a, a)         //float64 3.1415926
		fmt.Printf("%T %f\n", a, a)         //float64 3.141593
		fmt.Printf("%T %.3f\n", a, a)       //float64 3.142
		fmt.Printf("%T ,,%10.3f,,\n", a, a) //float64 ,,     3.142,,  10显示宽带 小数点留3位 小数点占1位
	*/

	/*
		//转义字符:每一个都是字符 rune类型-int32 可以作为单独字符使用，也可以作为字符串中的一个字符
		//rune:符文是int32的别名，在所有方面都与int32等效。按照惯例，它用于区分字符值和整数值
		var a rune = '\'' //一个字符，rune int32 4bytes 整数
		//参数索引
		fmt.Printf("Type = %T %[1]v\n", a)    //Type = int32 39 参数索引
		fmt.Printf("Type = %[1]T %[1]v\n", a) //使用同一个参数a 否则%[2]v 就超出范围勒 所以用同一个参数
		//fmt.Printf("Type = ..... %[n]T %v\n", a,......n ) 默认%v = %[n+1]v

		var b = 200
		fmt.Printf("%[2]d %[2]d\n", a, b)          //200 200 表示使用b没有使用到a a没有用到
		fmt.Printf("%[2]d %[1]d %v\n", a, b, 1000) //200 39 200 第二参数b、第一个参数a、%v表示1+1 就是第二个参数b的值

		var x rune = 97
		fmt.Printf("%T,%[1]d,%[1]c\n", x) //int32,97,a %c char字符方式打印表示通过将97转换为ASCII码表进行输出
	*/

	/*
		'\n' int32 4bytes整数 10 0xa
		'\r\n' 错误 因为单引号只能一个字符
		"\n" 字符组成的序列,string
		"\r\n" 2个字符串组成
		字符串
		字符串组成的序列
		"","\n" "abcd\r\nabc\txyz123"
		123 int型字面常量
		"123" 字符串，字符1、2、3 3个字符构成序列
	*/

	/*
		var x = "abc\nxyz" //标准输出、标准错误输出
		fmt.Print(x)
		var y = `abc	123` //7个字符 abc\t123
		fmt.Println(y)
		var z = `abc\t123` //8个 反引号不支持转义符号
		fmt.Println(z)     //abc\t123
	*/

	/*
		var x = "abc\nxyz\n123"
		fmt.Printf("%v\n", x)
		fmt.Printf("%s\n", x)
		fmt.Printf("%q\n", x) //%q 等于 "%s" "abc\nxyz\n123"

		y := 97
		fmt.Printf("%d %[1]x %#[1]x %[1]b %#[1]b %[1]c %[1]q\n", y) //97 61 0x61 1100001 0b1100001 a 'a'

		//%U 把一个整数用Unicode格式打印。
		fmt.Printf("%U,%x,%c\n", 27979, 27979, 27979) //U+6D4B,6d4b,测
	*/

	/*
		//Sprint:相当于Print，不过输出为string。
		//Sprintln:相当于Println，不过输出为string
		//Sprintf:相当于Printf，不过输出为string
		var x = "abc\nxyz\n123"
		var s = fmt.Sprintf("{%s,%c}", x, 100) //String
		fmt.Println(s + "\n456")
		{abc
		xyz
		123,d}
		456
	*/
}
=========================================================================================================================
【2.4 - 操作符】
	/*
			fmt.Println(15&5, 15&^5) //5 10
			//& 做乘法*
			1111
			0101
		    ------
			0101

			//&^ 上面为1下面为0-漏1 相同1为0 按照y有1的位清空x对应位
			1111
			0101
			------
		    1010 =10
	*/

	//比较运算符构成的比较表达式 返回值bool类型 同类型比较
	fmt.Println(1 == '1') //false 是否 返回bool 1=="abc" 不同类型不匹配不能操作

	//逻辑运算符 &&、||、！
	//fmt.Println(1 && 0, "abc" && true) //操作数Go中必须使用bool
	fmt.Println(true && true, false && false, 5 > 3 && 1 < 2) //true false true
	//&& || 有短路 第一个数false 就后面不计算了

	//赋值运算符 = += -= *= /= %= 》= 《= &= &^= ^= |=

	// 无三目运算符

	//指针 ： 类型、类型保存什么数据？大整数，内存地址
	a := 1000
	b := &a
	fmt.Printf("%T a address = %[1]v\nb address = %p\n", &a, b) //%p 指针本质是大整数，门牌号码  2个地址一样
	//*int a address = 0xc0000a6090,b address = 0xc0000a6090
	c := *b
	fmt.Println(c)              //1000
	fmt.Printf("%T %[1]v\n", c) //int 1000
	fmt.Println(a == c)         //true
	fmt.Printf("a address = %p\nb address = %p\nc address = %p\n", &a, b, &c)
	//a address = 0xc0000a6090
	//b address = 0xc0000a6090
	//c address = 0xc0000a6098 a c2个地址不一样

	var p1 *int                        //存int类型的值的地址
	fmt.Printf("%T %[1]p %[1]v\n", p1) //*int 0x0 <nil>(空指针很危险)
	p1 = &a
	fmt.Printf("%T %[1]p %[1]v\n", p1) //*int 0xc0000a6090 0xc0000a6090
	fmt.Println(*p1)                   //1000
	*p1 = 2000
	fmt.Println(*p1)  //2000
	fmt.Println(a, c) //2000 1000 a也变成了2000 c没有变

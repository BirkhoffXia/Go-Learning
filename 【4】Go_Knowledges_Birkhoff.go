【4-1-顺序表操作原理】
=====================================================================


【4-2-链接表操作原理】
=====================================================================


【4-3-数值处理和Scan】
=====================================================================
package main

import (
	"fmt"
	"math"
)

func main() {
	// / 截取整数部分
	fmt.Println(1/2, 3/2, 5/2)    //0 1 2
	fmt.Println(-1/2, -3/2, -5/2) //0 -1 -2
	fmt.Println("=================================")
	//math.Ceil 向上取整走 像大的数走
	fmt.Println(math.Ceil(2.01), math.Ceil(2.5), math.Ceil(2.8))    //3 3 3
	fmt.Println(math.Ceil(-2.01), math.Ceil(-2.5), math.Ceil(-2.8)) //-2 -2 -2
	fmt.Println("=================================")
	//math.Floor 向下取整
	fmt.Println(math.Floor(2.01), math.Floor(2.5), math.Floor(2.8))    //2 2 2
	fmt.Println(math.Floor(-2.01), math.Floor(-2.5), math.Floor(-2.8)) //-3 -3 -3
	fmt.Println("=================================")
	//math.Round 四舍五入 超过0.5 进一位 否则去掉一位
	fmt.Println(math.Round(2.01), math.Round(2.5), math.Round(2.8))                  //2 3 3
	fmt.Println(math.Round(-2.01), math.Round(-2.5), math.Round(-2.8))               //-2 -3 -3
	fmt.Println(math.Round(0.5), math.Round(-1.5), math.Round(2.8), math.Round(3.5)) //1 -2 3 4
	fmt.Println("=================================")

	//Other math.*
	fmt.Println(math.Abs(-2.7))                               //绝对值2.7
	fmt.Println(math.E, math.Pi)                              //常数2.718281828459045 3.141592653589793
	fmt.Println(math.MaxInt16, math.MinInt16)                 //常数，极值32767 -32768
	fmt.Println(math.Log10(100), math.Log2(8))                //对数2 3
	fmt.Println(math.Max(1, 2), math.Min(-2, 3))              //最大值，最小值2 -2
	fmt.Println(math.Pow(2, 3), math.Pow10(3))                //幂8 1000
	fmt.Println(math.Mod(5, 2), 5%2)                          //取模1 1
	fmt.Println(math.Sqrt(2), math.Sqrt(3), math.Pow(2, 0.5)) //开方1.4142135623730951 1.7320508075688772 1.4142135623730951
	fmt.Println("=================================")

	// Scan:空白字符分割，回车提交。换行符当做空白字符
	var s1, s2 string
	n, err := fmt.Scan(&s1, &s2) //以后Scan要加 & 指针
	fmt.Println(err)             //err 如果是nil说明没有错误
	if err != nil {
		panic(err) //如果走到这 说明有错误 使用panic函数 报错原因
	}
	fmt.Println(n) //能够走到这里 err是nil,没有错误
	fmt.Printf("%T %[1]v ,%T %[2]v", s1, s2)
	fmt.Println("=================================")

	//Scan 想获得整数
	var i1, i2 int
	n, err = fmt.Scan(&i1, &i2) //这里直接用= 因为上面已经赋值过了
	if err != nil {
		panic(err) //panic: expected integer  panic: type not a pointer: int
	}
	fmt.Println(n) //能够走到这里 err是nil,没有错误
	fmt.Printf("%T %[1]v ,%T %[2]v", i1, i2)

}

	//Scanf 最好使用空格分隔符  , ; 分割不行 可能会出错 %s,%d ,会给%s匹配 可以写成%d,%s
	//[情况：1]推荐空格分割符 
	var s1, s2 string
	n, err := fmt.Scanf("%s %s", &s1, &s2) //
	if err != nil {
		panic(err) //panic: input does not match format
	}
	fmt.Println(n)
	fmt.Printf("%T %[1]v,%T %[2]v\n", s1, s2)
	fmt.Println("=======================================")
	
	//[情况：2]
	var s1 int
	var s2 string
	n, err := fmt.Scanf("%d,%s", &s1, &s2) //%d,%s可以
	if err != nil {
		panic(err) //panic: input does not match format
	}
	fmt.Println(n)
	fmt.Printf("%T %[1]v,%T %[2]v\n", s1, s2) //输入123,xks => int 123,string xks
//[情况：3]
	var s1 string
	var s2 int
	n, err := fmt.Scanf("%s,%d", &s1, &s2) //%s,%d不可以
	if err != nil {
		panic(err) //panic: input does not match format
	}
	fmt.Println(n)
	fmt.Printf("%T %[1]v,%T %[2]v\n", s1, s2)
	
【4-4-数组】
=====================================================================
Go中是顺序表
	可以索引，从0开始，不支持负索引
	长度不可变，定义时必须指定长度，不可增删元素
	内容可变，元素可以被更新
	值类型：值就是数组，对数组来说，很多操作如果基于数组会进行数据整个复制copy，产生数组的副本

package main

import "fmt"

//定义一个函数
func showAddr(arr [3]int) [3]int { //(arr [3]int) 入参  [3]int：出参
	fmt.Printf("arr %p %p,%v\n", &arr, &arr[0], arr) //arr 0xc000012258 0xc000012258,[100 200 300]
	return arr
}

func showAddrPointer(arr *[3]int) *[3]int {
	fmt.Printf("arr pointer %p %p,%v\n", arr, &arr[0], arr) //arr pointer 0xc0000122b8 0xc0000122b8,&[100 200 300]
	//Go语言允许 使用指针加索引
	arr[0] = 200                                            //这里改了等于a16[0]改为200
	fmt.Printf("arr pointer %p %p,%v\n", arr, &arr[0], arr) //arr pointer 0xc0000122b8 0xc0000122b8,&[200 200 300]
	return arr
}

func main() {
	//定义
	//[]int:切片类型 ，[3]int是数组类型，int表示元素类型
	//[]表示切片类型， [常量整数值]数组，3表示长度为3
	var a1 [3]int //数组定义，由于没有初始化，所以是零值
	//var a0 [] 错误
	//var a0 []int //可以这个是切片类型
	fmt.Printf("%T,%[1]v\n", a1) //[3]int,[0 0 0]
	a1[0] = 100
	a1[1] = 200
	a1[2] = 300
	fmt.Printf("%T,%[1]v\n", a1) //[3]int,[100 200 300]
	fmt.Println("============================================")

	a1 = [3]int{300, 400, 500}   // {}int类型，{}表示字面量定义
	fmt.Printf("%T,%[1]v\n", a1) //[3]int,[300 400 500]
	a1 = [3]int{1}               //如果后面没写默认都是0
	fmt.Printf("%T,%[1]v\n", a1) //[3]int,[1 0 0]
	fmt.Println("============================================")

	const length = 5 // 这里使用var变量不允许
	a3 := [length]int{}
	fmt.Printf("%T %[1]v\n", a3) //[5]int [0 0 0 0 0]
	fmt.Println("============================================")

	var a4 = [3]int{}            // 省略类型 类型推导
	fmt.Printf("%T %[1]v\n", a4) //[3]int [0 0 0]
	fmt.Println("============================================")

	var a5 = [...]int{1}         //让机器算有几个元素，不能扩展了 一旦定义了长度不能变 内容可以变
	fmt.Printf("%T %[1]v\n", a5) //[1]int [1]
	fmt.Println("============================================")

	//想定义数组最后一位 指定索引
	var a6 = [5]int{1, 4: 7}
	fmt.Printf("%T %[1]v\n", a6) //[5]int [1 0 0 0 7]
	fmt.Println("============================================")

	//多维数组
	var a7 = [2][3]int{{}, {2}}  //[2]第一维度 行，行中每一个元素[3]int,int元素类型
	fmt.Printf("%T %[1]v\n", a7) //[2][3]int [[0 0 0] [2 0 0]]
	fmt.Println("============================================")

	//...只能写在第一维度
	var a8 = [...][3]int{{2: -1}, {2}, {3}, {4}} //使用... 推算 是四维数组
	fmt.Printf("%T %[1]v\n", a8)                 //[4][3]int [[0 0 -1] [2 0 0] [3 0 0] [4 0 0]]
	fmt.Println("============================================")

	var a9 [3]int
	fmt.Println(a9, len(a9)) //[0 0 0] 3

	//遍历数组
	var a10 = [3]int{100, 200, 300}
	for i := 0; i < len(a10); i++ {
		fmt.Println(i, a10[i])
	}
	fmt.Println("============================================")

	for i, v := range a10 {
		fmt.Println(i, v, a10[i])
	}
	// 0 100 100
	// 1 200 200
	// 2 300 300
	fmt.Println("============================================")

	// 数组的地址 就是第一个元素的地址 地址是连续开辟的 16进制 每个占8个字节
	// 数组必须在编译时就确定长度，之后不能改变长度
	// 数组首地址就是数组地址
	// 所有元素一个接一个顺序存储在内存中
	// 元素的值可以改变，但是元素地址不变
	var a11 [3]int // 内存开辟空间存放长度为3的数组，零值填充
	for i := 0; i < len(a11); i++ {
		fmt.Println(i, a11[i], &a11[i])
	}
	// 0 0 0xc0000121b0
	// 1 0 0xc0000121b8
	// 2 0 0xc0000121c0
	fmt.Printf("%p %p,%v\n", &a11, &a11[0], a11) //0xc0000121b0 0xc0000121b0,[0 0 0]
	a11[0] = 1000
	fmt.Printf("%p %p,%v\n", &a11, &a11[0], a11) //0xc0000121b0 0xc0000121b0,[1000 0 0]
	fmt.Println("============================================")

	// 每个元素间隔16个字节，为什么?"abc"是16个字节?这说明什么? 答：实际 通过索引指向指针位置找到数据
	// 存1个字符串用了16个字节，2^128种状态，字符串中一个字符至少占用1个字节 16个字节=128位
	// 数组首地址就是数组地址
	// 所有元素顺序存储在内存中
	// 元素的值可以改变，但是元素地址不变
	var a12 = [3]string{"abc", "def", "xyz"} // 内存开辟空间存放长度为3的数组
	for i := 0; i < len(a12); i++ {
		fmt.Println(i, a12[i], &a12[i])
	}
	fmt.Printf("%p %p,%v\n", &a12, &a12[0], a12)
	a12[0] = "oooooo"
	fmt.Printf("%p %p，%v\n", &a12, &a12[0], a12)
	// 0 abc 0xc000114510
	// 1 def 0xc000114520
	// 2 xyz 0xc000114530
	// 0xc000114510 0xc000114510,[abc def xyz]
	// 0xc000114510 0xc000114510，[oooooo def xyz]
	fmt.Println("============================================")

	//Go语言复制了一份 多了一个副本copy 但是地址不一样 内容保持一致
	var a13 = [3]int{100, 200, 300}
	fmt.Printf("%p %p,%v\n", &a13, &a13[0], a13) //0xc0000aa138 0xc0000aa138,[100 200 300]

	a14 := a13                                   //这说明，Go语言在这些地方对数组进行了值拷贝，都生成了一份副本。
	fmt.Printf("%p %p,%v\n", &a14, &a14[0], a14) //0xc000126150 0xc000126150,[100 200 300]

	a15 := showAddr(a13)                         //函数传过去的 又是一个副本copy
	fmt.Printf("%p %p,%v\n", &a15, &a15[0], a15) //0xc000012258 0xc000012258,[100 200 300]
	fmt.Println("============================================")

	//减少副本copy-使用指针
	var a16 = [3]int{100, 200, 300}
	fmt.Printf("a16 %p %p,%v\n", &a16, &a16[0], a16)       //a16 0xc0000122b8 0xc0000122b8,[100 200 300]
	a17 := showAddrPointer(&a16)                           //将a16的地址传到函数中 return也是个指针 这个指针指向就是a16地址 = 操作a16
	fmt.Printf("return a17 %p %p,%v\n", a17, &a17[0], a17) //return a17 0xc0000122b8 0xc0000122b8,&[200 200 300]
	fmt.Println("---------------------------------------------")
	fmt.Printf("a16 %p %p,%v\n", &a16, &a16[0], a16) //a16 0xc0000122b8 0xc0000122b8,[200 200 300]
	a17[2] = 999
	fmt.Printf("a16 %p %p,%v\n", &a16, &a16[0], a16) //a16 0xc0000122b8 0xc0000122b8,[200 200 999]

}

【4-5-切片内存模型】
=====================================================================
基于顺序表实现，基于GO的数组实现
有索引，习惯从0开始，可索引
长度可变
内容可变，元素可以被更新
引用类型
底层基于数组
值类型？不是，引用类型，地址，其他语言中，如果遇到了引用类型，就可以认为大家操作的是同一个内存的数据
	 传址
	所谓的传址(址copy)，切片的结构体(Header，标头值)的完全的复制，底层数组没有copy

物理实现
	底层数据存储使用Go的数组
	定义一个结构体
		指针 Go的底层数组
			len len(切片)
			cap cap(切片)

	不同的切片，可以共用同一个底层数组，也可以不共用
	go没有提供顺序表实现切片除尾部append之外的方法
		网传，开辟一个内存空间，插入点之前复制过来，补现有插入元素，再把之后copy过来
		为什么这么干?

值copy
	值类型，数组是自身完全copy，自身的副本引用类型，切片是header的copy
	字符串
		有header，16字节
		表现却像值类型，兼容其他语言的习惯，就像int一样				


[定义]
package main

import "fmt"

func main() {
	//定义
	var s0 []int                 // 零值nil ,可以直接使用
	fmt.Println(s0)              //[]
	fmt.Printf("%T %[1]v\n", s0) //[]int []
	print(s0)                    //[0/0]0x0
	fmt.Println(s0 == nil)       //	true

	s0 = append(s0, 1)
	fmt.Println(s0 == nil, s0) //	false [1]

	var s1 = []int{}       //字面量定义cap = len = 0 赋值初值了 不等于nil
	fmt.Println(s1 == nil) //false

	s2 := []int{1, 3, 5} //cap = len =3
	fmt.Println(s2)      //[1 3 5]

	//make
	var s3 = make([]int, 0)    //len = cap = 0 []int
	fmt.Println(s3 == nil, s3) //false []

	var s4 = make([]int, 3, 5) //len=3 cap=5 ,[]int
	fmt.Println(s4 == nil, s4) //false [0 0 0]

	//容器尾部添加 append
	fmt.Println(append(s0, 1, 2, 3, 4, 5)) //[1 1 2 3 4 5]
 
	//切片本质是对底层数组一个连续片段的引用。此片段可以是整个底层数组，也可以是由起始和终止索引标识的一些数组项的子集。
}

[切片-原理]-结构体
package main

import "fmt"

func main() {
	var s0 = []int{1, 11, 111, 1111} //cap len 4
	fmt.Println(len(s0), cap(s0))    //len 、cap 从结构体中取值
	fmt.Printf("&s0 address = %p,&s0[0] address = %p,&s0[1] address = %p\n", &s0, &s0[0], &s0[1])
	//&s0 address = 0xc000008078,&s0[0] address = 0xc0000161a0,&s0[1] address = 0xc0000161a8
	//&s0    address = 0xc000008078 切片结构体的地址 s0指代0xc000008078这个结构体(array->0xc0000161a0，len，cap)
	//&s0[0] address = 0xc0000161a0 底层数组的首元素的地址
	//&s0[1] address = 0xc0000161a8 连续内存占8个字节
}

[切片-长度&容量-扩容]

	/*
		append一定返回一个新的切片，但本质上来说返回的是新的Header
		append可以增加若干元素
			如果增加元素时，当前长度+新增个数<=cap则不扩容
				原切片使用原来的底层数组，返回的新切片也使用这个底层数组
				返回的新切片有新的长度
				原切片长度不变
			如果增加元素时，当前长度 + 新增个数 > cap则需要扩容
				生成新的底层数组，新生成的切片使用该新数组，将就元素复制到新的数组，其后追加新元素
				原切片底层数组、长度、容量不变
	*/
	
package main

import "fmt"

func main() {
	var s1 = []int{1, 3, 5, 7} //cap len 4
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)

	s1 = make([]int, 2, 5) //重新构造一个切片 结构体的地址不变，首元素的地址发生变化
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	// s1被重新赋值，但是s1地址没有变化
	// s1 0xc00009a060,0xc0000d0020, l=4 ,c=4 ,[1 3 5 7]
	// s1 0xc00009a060,0xc0000be0c0, l=2 ,c=5 ,[0 0]
	fmt.Println("---------------------------------------------------------------")

	s1 = append(s1, 200) //append 返回新的header信息，覆盖 append就是给slice追加使用
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	s2 := append(s1, 1, 2) //append 返回新的header信息，使用新的变量存储
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("s2 %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	// 目前没有超过容量，底层共用同一个数组，但是，对底层数组使用的片段不一样
	// s1 0xc00009a060,0xc0000be0c0, l=3 ,c=5 ,[0 0 200]
	// s1 0xc00009a060,0xc0000be0c0, l=3 ,c=5 ,[0 0 200]
	// s2 0xc00009a0c0,0xc0000be0c0, l=5 ,c=5 ,[0 0 200 1 2]
	fmt.Println("---------------------------------------------------------------")

	s3 := append(s1, -1)
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("s2 %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Printf("s3 %p,%p, l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	// 目前三个切片底层同一个数组，只不过长度不一样
	// s1 0xc00009a060,0xc0000be0c0, l=3 ,c=5 ,[0 0 200]
	// s2 0xc00009a0c0,0xc0000be0c0, l=5 ,c=5 ,[0 0 200 -1 2]
	// s3 0xc00009a108,0xc0000be0c0, l=4 ,c=5 ,[0 0 200 -1]
	fmt.Println("---------------------------------------------------------------")

	s4 := append(s3, 3, 4, 5)
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("s2 %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Printf("s3 %p,%p, l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	fmt.Printf("s4 %p,%p, l=%-2d,c=%-2d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	//底层数组变了，容量也增加了 容量翻倍5*2=10 
	// s1 0xc00009a060,0xc0000be0c0, l=3 ,c=5 ,[0 0 200]
	// s2 0xc00009a0c0,0xc0000be0c0, l=5 ,c=5 ,[0 0 200 -1 2]
	// s3 0xc00009a108,0xc0000be0c0, l=4 ,c=5 ,[0 0 200 -1]
	// s4 0xc00009a168,0xc0000ac050, l=7 ,c=10,[0 0 200 -1 3 4 5]
	fmt.Println("---------------------------------------------------------------")

	s5 := append(s4, 6, 7, 8, 9)
	fmt.Printf("s1 %p,%p, l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("s2 %p,%p, l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Printf("s3 %p,%p, l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	fmt.Printf("s4 %p,%p, l=%-2d,c=%-2d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	fmt.Printf("s5 %p,%p, l=%-2d,c=%-2d,%v\n", &s5, &s5[0], len(s5), cap(s5), s5)
	// s1 0xc00009a060,0xc0000be0c0, l=3 ,c=5 ,[0 0 200]
	// s2 0xc00009a0c0,0xc0000be0c0, l=5 ,c=5 ,[0 0 200 -1 2]
	// s3 0xc00009a108,0xc0000be0c0, l=4 ,c=5 ,[0 0 200 -1]
	// s4 0xc00009a168,0xc0000ac050, l=7 ,c=10,[0 0 200 -1 3 4 5]
	// s5 0xc00009a1e0,0xc0000ce0a0, l=11,c=20,[0 0 200 -1 3 4 5 6 7 8 9]
	fmt.Println("---------------------------------------------------------------")

}

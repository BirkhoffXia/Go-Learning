【7-1-哈希表原理】
=====================================================================
Map
	*Go采用冲突解决方案是拉链法
	
	数据结构、字典、映射(key-value)，hash 表 
		"name" => "Tom"
		"age" => 20
		"height"=> 175
	映射
		x--function-->y  y=f(x)，一对一，x->x'->x''变形
				js 箭头函数，x=>x*x
哈希表
	hash算法
		y = hash(key)
			y称为hash值，散列值，y值是个大整数，习惯输出为8x十六进制的字符串
			y可以当做门牌号码
			key "abc" "\x61\x62\x63"  0x616263  => 1234
			"abb" => "\x61\x62\x62"   0x616262  => 513
		冲突
			y的冲突:可以找到不同的k1、k2，计算出的y值一样
			y -> 桶index 余数也一样
				y一样，余数一样，冲突
				y不一样，余数一样，冲突
			冲突解决方案2种
				开地址法:当前桶index有人了，不能挤一起，按照某种规则选择一个新的无人占用的桶
					向后找空闲桶
					按平方值找空闲桶
				拉链法
					Go语言多数库采用的方案，不同的实现略有不同
					桶是一个数据结构，存储一个指针，指针指向链接表的Head
	map
		底层是hash表
		长度可变，value可变，key不可变(不同key对应不同的hash值)
		key是无序的
			key微小的变化，导致y或桶index有巨大的变化
		key去重
			同一个key只能对应一个value
		可索引？ 不是线性表，不可索引
			m1["abc"]
			
			
		Header
			B 2^B是桶总数
				2^(B+1) 扩桶容，增量，翻倍
			len(count) ，kv对的个数  O(1)
			buckets 指针 指向底层hash表
			oldbuckets 指针 指向老的
	增
		f(k1)-> index1，把kv对存入桶中
		f(k2)-> indexl，冲突
			挤一挤，挤在一个桶里面的链表中，为了匹配，每一个节点存储k和v
			链接表，尾部加一个新节点kv对
		新增时，就可以知道Load factor负载因子，len(map)/桶总数(2^B) > 6.5，扩容了
		len + 1
	查
		f(k1)-> indexl
			固定步骤，key-》bucket index o(l)
				如果桶的链表很小(还可以做一些优化)，我们就可以认为定位一个key对应的value就是0(1)
				如果通过key找到value对map来说是最好的
			该桶的链表遍历 O(n)
				链表n越大，表示什么?该桶的冲突太多了
					数据倾斜了，很多很多key计算都得到这个桶index，概率极低
					更多是下面这种情况，所有桶 链表都长了，这说明 整个hash的kv的总数太多了，显得桶少了
						要扩容了，扩桶
				找到k1的key，取出kv对
				没有找到
			定位
				key O(1)
				value定位，用不上hash了，遍历所有kv对，去找value，不能这么用，禁止，及其没有效率
	改
		定位后，发现key已存在，修改链表上的k对应的value
			m1["abc"] = 1000
	
	删除
		定位后
			kv在，删除，从链表中删除
				每8个kv对放在一个节点中(顺序表)，溢出才会next指向下一个节点
				删除kv对应位置情况，导致顺序表空洞
			不在，不删
			
	集合set
		Go没有提供
		只有key
			map的value	特殊处理
		去重，key唯一
			
	扩容策略
		翻倍扩容：桶不够，负载因子高
			如果不扩容，桶的链表就太长了，对hash表的搜索变成了对线性表的搜索，降级位O(n)
			Go扩容的问题
				翻倍扩容，一定是新值开辟新的桶数组空间
				数据迁移
					新老桶数组共存一段时间
					渐进式 迁移
						当写入时，老桶数组还在，就去上面找至多2条桶链表迁移以下
						减少一次性迁移所有数据带来的性能问题 
				index问题
					会混乱马？不会，为什么？2^B有关
		等量扩容
			删除导致链表的空洞，规整，链表重排
			
		
什么是相同的key？
hash值相同则一定相同么？
冲突的key有什么异同？
	相同的key一定计算得到相同hash值，桶index一样，表中存储k和v，k是否相等，同一条kv记录
	不同的key
		相同的hash值，hash冲突、index冲突，同一个桶的链表中不同kv对
		不同hash值
			index冲突，同一个桶的链表中不同kv对
			index不冲突，不同的桶的链表中分别存储
		
【7-2-map及操作】
=====================================================================
[映射]
映射Map，也有语言称为字典dict。
	长度可变
	存储的元素是key-value对(键值对)，value可变
	key无序不重复
	不可索引，需要通过key来访问
	不支持零值可用，也就是说，必须要用make或字面常量构造
	引用类型
	哈希表

请问以下谁的遍历效率最高，可以多选？ 都不是 都不高只要遍历，都要所有元素不重复拿一遍，n越大越耗时
	A：数组 B：切片 C:链表 D：Map E：A&B&C 
树、图、线性表、哈希表、只要将元素按照某种顺序不重复拿一编

package main

import "fmt"

func main() {
	/*
		var m1 map[string]int        //nil 很危险，map不是零值可用 不支持零值可用
		fmt.Printf("%T %[1]v\n", m1) //map[string]int map[]
		fmt.Println(m1, m1 == nil)   //true
		m1["t"] = 200 //panic 不可以  不可以直接在空值中赋值
	*/
	//1 字面量
	var m0 = map[string]int{}           //安全，没有一个键值对而已 该类型的map实例对象 有底层的哈希桶数组B
	fmt.Printf("%T %[1]v\n", m0)        //map[string]int map[]
	fmt.Println(m0, m0 == nil, len(m0)) //map[] false 0
	m0["a"] = 1234
	fmt.Println(m0, len(m0)) //map[a:1234] 1

	var m1 = map[string]int{
		"a": 11,
		"b": 22,
		"c": 33, //Go要求这里以逗号结尾
	}
	fmt.Printf("%T %[1]v\n", m1) //map[string]int map[a:11 b:22 c:33]
	fmt.Println(m1, len(m1))     //map[a:11 b:22 c:33] 3

	//2 make
	m2 := make(map[int]string)
	fmt.Println(m2, len(m2), m2 == nil) //map[] 0 false

	//分配足够容量来容纳100个元素，长度为0。为了减少扩容，可以提前给出元素隔宿
	m3 := make(map[int]string, 100)
	fmt.Println(m3, len(m3), m3 == nil) //map[] 0 false

	//新增或修改
	var m4 = make(map[string]int)
	m4["a"] = 11                        //key不存在，则创建新的key和value对
	m4["a"] = 22                        //key已经存在，则覆盖value
	fmt.Println(m4, len(m4), m4 == nil) //map[a:22] 1 false

	//查找 使用map一般需要使用key来查找，时间复杂度为O(1) 用key来访问map最高效的方式
	fmt.Println(m4["a"]) //22
	fmt.Println(m4["b"]) //0 找不到，返回value值的类型的零值
	if v, ok := m4["b"]; ok {
		fmt.Println(v, m4["b"])
	} else {
		fmt.Println("key不存在") //key不存在
	}
	if v, ok := m4["a"]; ok {
		fmt.Println(v, m4["a"]) //22 22
	} else {
		fmt.Println("key不存在")
	}

	//移除
	delete(m4, "a")                     //存在，从桶的链表中删除kv对
	fmt.Println(m4, len(m4), m4 == nil) //map[] 0 false
	delete(m4, "b")                     //不存在，删除操作，也不会panic

	//遍历
	var m5 = map[string]int{
		"a": 111,
		"b": 222,
		"c": 333,
	}
	for k, v := range m5 {
		fmt.Println(k, v)
	}
	// a 111
	// b 222
	// c 333
	/*map 的key 是无序的 千万不要从遍历结果来推测其内部顺序*/
	//Go 中的map并不是并发安全的数据结构，如果出现并发读同时有并发写，或并发写同时有并发读，会抛出异常
}
【7-3-排序】
=====================================================================
序列order、sequence
	元素，明确前后关系
	排成一队，线性的

排序sort
	按照某个指标进行大小排序，升序和降序。排队而且按照大小个排好了
	有线性数据结构，可以sort了
	性能问题
		非常耗时的操作
				差一点的0(n^2) 2层for
				好一点 0(nlogn)
				最好的情况o(n)
				
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 切片排序
	// 针对int、string有快捷方法Ints、strings
	a := []int{-1, 23, 5, 9, 7}
	//sort.Sort(sort.IntSlice(a))
	//sort.IntSlice(a)强制类型转换以施加接口方法
	//sort.Intslice(a)强制类型转换，得到增强的切片类型 type Intslice[]int
	// 为IntSlice扩增方法，因为[]int不能扩增方法，例如增加元素比较大小的方法
	sort.Ints(a)   //就地修改原切片的底层数组
	fmt.Println(a) //默认升序 	[-1 5 7 9 23]

	b := []string{"xyz", "a", "abc", "Ab", "x"}
	sort.Strings(b)
	fmt.Println(b) //	[Ab a abc x xyz]

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a) //	[23 9 7 5 -1]

	//二分查找
	//二分查找，必须是升序
	//二分查找的 前提是 排过序
	c := []int{-1, 23, 5, 9, 7}
	sort.Ints(c)
	i := sort.SearchInts(c, 6)
	fmt.Println(i) //2   插入第二个索引位置
}

什么是相同的key?hash值相同则key一定相同吗?冲突的key有什么异同?
	有冲突的key就是相同的key吗?也就是说，如果2个key计算的hash值相同就是同一个key吗?
		key计算的hash值相同只能说明hash冲突，如果key也相等，才能说明是用一个key。同一个key计算的hash值一定一样，但是hash冲突不一定是同一个key。
		
		
【7-4-函数基本概念】
=====================================================================
package main

import "fmt"

//func 函数的定义 定义时
//function 功能，封装成一个函数，因为它能完成一项或多项功能
//Add是函数名，标识符 ，首字母大写对包外可见 替换成地址,地址执行内存中一个位置，编译好的函数指令
//标识符后面第一个()定义，形式参数，简称形参，占位符，未来(调用时)要用会给具体值
//函数可以没有返回值，也可以有n个返回值
func add(x int, y int) int { //函数体，若干语句构成，最终编译成指令
	fmt.Println("add called", x, y, "\"x + y\" = ", x+y)
	return x + y
}

func main() { //入口函数，决定了一个可执行文件，标注一个入口地址
	fmt.Printf("%T\n", add)      //func(int, int)
	fmt.Printf("%T\n", 100)      //int
	fmt.Printf("%T\n", "string") //string
	add(1220, 319)               //add called 函数的调用(执行)，函数标识符后面跟上一个小括号，调用时
	//调用时，通过add指向的地址找到函数的指令，执行这些指令
	//一个函数定义好之后，就可以调用，什么时候调用?根据需要，调用次，也可以n次
	add(1, 2)
	add(2, 3)
	add(3, 4)

	// 调用时送入和形参对应的参数，他不是占位的，实实在在的参数，实际参数，简称实参
	money := add(4, 5)
	fmt.Printf("money = %d\n", money) //money = 9
}

【7-5-函数调用原理】
=====================================================================
特别注意，函数定义只是告诉你有一个函数可以用，但这不是函数调用执行其代码。至于函数什么时候被调用，不知道。一定要分清楚定义和调用的区别。
函数调用相当于运行一次函数定义好的代码，函数本来就是为了复用，试想你可以用加法函数，我也可以用加法函数，你加你的，我加我的，应该互不干扰的使用函数。
为了实现这个目标，函数调用的一般实现，都是把函数压栈(LIFO)，每一个函数调用都会在栈中分配专用的栈帧，局部变量、实参、返回值等数据都保存在这里。

上面的代码，首先调用main函数，main压栈，接着调用add(4,5)时，add函数压栈，压在main的栈帧之上
	add调用return，将add返回值保存在main栈帧的本地变量out上，add栈帧消亡，回到main栈帧

【7-6-函数返回值】
=====================================================================
[函数类型]
package main

import "fmt"

func fn1()              {}
func fn2(i int) int     { return 100 }
func fn3(j int) (r int) { return 200 }

func main() {
	fmt.Printf("%T\n", fn1)
	fmt.Printf("%T\n", fn2)
	fmt.Printf("%T\n", fn3)
}

// func()
// func(int) int
// func(int) int

返回值
	可以返回0个或多个值
	可以在函数定义中写好返回值参数列表
		可以没有标识符，只写类型。但是有时候不便于代码阅读，不知道返回参数的含义。
		可以和形参一样，写标识符和类型来命名返回值变量，相邻类型相同可以合并写
		如果返回值参数列表中只有一个返回参数值类型，小括号可以省略。
		以上2种方式不能混用，也就是返回值参数要么都命名，要么都不要命名
	return
		return之后的语句不会执行，函数将结束执行
		如果函数无返回值，函数体内根据实际情况使用return。
		return后如果写值，必须写和返回值参数类型和个数一致的数据。
		return后什么都不写那么就使用返回值参数列表中的返回参数的值

【7-7-函数形参】
=====================================================================
[形参]
可以无形参，也可以多个形参
不支持形式参数的默认值
形参时局部变量

func fn1()                {} //无形参
func fn2(int)             {} //有一个int形参，但是没法用它，不推荐
func fn3(x int)           {} //单参函数
func fn4(x int,y int)     {} //多参函数
func fn5(x,y int,z string){} //相邻形参类型相同，可以写到一起

fn1()
fn2(5)
fn3(10)
fn4(4),5
fn(7,8,"OK")

[可变参数]
可变参数 variadic 其他语言也有类似的被称为剩余参数，但Go语言有所不同

package main

import "fmt"

func fn6(nums ...int) { //可变形参 0~n
	fmt.Printf("%T %[1]v,%d,%d\n", nums, len(nums), cap(nums))
}

func fn7(x int, y int, others ...int) { //可变形参 0~n
	fmt.Println(x, y, "|||", others, len(others), cap(others))
}

func main() {
	//*可变参数收集实参到一个切片中
	fn6(1)    //	[]int [1],1,1
	fn6(3, 5) //[]int [3 5],2,2
	//构造一个新的切片和其底层数组，把7、8、9放进底层数组中
	fn6(7, 8, 9) //[]int [7 8 9],3,3

	//[]int{11,22,33}的标头值复制进去覆盖nums的header，共用底层数组
	fn6([]int{11, 22, 33}...) //[]int [11 22 33],3,3
	// fn6([]int{11, 22, 33} //错误

	//*如果有可变参数，那它必须位于参数列表中最后
	fn7(22, 33)         //22 33 ||| [] 0 0
	fn7(22, 33, 44)     //22 33 ||| [44] 1 1
	fn7(22, 33, 44, 55) //22 33 ||| [44 55] 2 2

	//1 fn7()     							     NO
	//2 fn7(1)   							     NO
	//3 fn7(1,2) 							     OK
	//4 fn7(1,2,3,4) 						     OK
	//5 fn7(1,[]int{2,3,4}...) 				     NO
	//6 fn7(1,2,[]int{3,4}...) 				     OK
	//7 fn7(1,2,[]int{3,4,5}...)                 OK
	//8 fn7(1,2,[]int{3,4,5}...,6,7,8)           NO
	//9 fn7(1,2,[]int{3,4,5}...,[]int{6,7,8}...) NO
	//10 fn7(1,2,[]int{3,4,5})                   NO 要传可变参数要传 正确类型参数或者加...  不能传切片	
}


*可变参数收集实参到一个切片中
*如果有可变参数，那它必须位于参数列表中最后
	func fn7(x,y int,nums ...int,z strings){} 这是错误的


[总结]
可以看出，可变参数限制较多
	直接提供对应实参，会被封装成一个新的切片
	可以使用使用切片传递的方式 切片..，但是这种方式只能单独为可变形参提供实参，因为这是实参切片的header的复制
	这和Python、JavaScript中的参数解构不一样，也确实没有它们灵活方便。
	
Q:
func fn6(nums []int){}和func fn6(nums ...int){} 调用时，都可以使用切片，那有什么区别呢?形参使用切片类型还是可变参数呢?

func fn6(nums []int){} ：只能传切片  复制了Header 共用底层数组
	fn6([]int{})
	fn6([]int{1})
	fn6([]int{1,11})

func fn6(nums ...int){} ：
	fn6() :空切片
	fn6(1,2,3,4):4个切片
	fn6([]int{}...):共用底层数组 Header复制一份
	fn6([]int{2}...)
	fn6([]int{2,22}...)
	fn6([]int{})//错误
	fn6(1,[]int{}...)//错误 要同类型
【3-1-分支】
=====================================================================
package main

import "fmt"

func main() {

	if true { //必须是bool表达式
		fmt.Println("BIRKHOFF")
	} else {

	}

	/*
		//switch case fallthrough
		//case写什么，取决于 switch后面写的类型 == 多分支 一个分支成立 其他分支不执行
		//多分支不支持穿透，fallthrough只能保证穿到下一层
		var a int = 200
		switch a {
		case 1:
			fmt.Println("=1")
		case 200: // a>200 错误的语法
			fmt.Println("200")
			fallthrough
		default:
			fmt.Println("不等于1，不等于200")

		}
	*/

	//switch后面没写 默认是bool型
	//switch后面的值就是要判断的，case后的类型要和这个值一致
	//var a = 100
	//if、switch等语句中临时用短格式定义的局部变量，他们的可见范围(作用域)仅只能用在当前语句中，如果嵌套，会向内穿透，在内部可以看到可以用

	switch /*bool型，true*/ a := 1; true { //此时a:=1 只能在switch中使用
	case a > 1:
		fmt.Println(">1", a)
	case a < 1:
		fmt.Println("<1", a)
	default:
		fmt.Println("=1", a)
	}

}

【3-2-循环for】
=====================================================================
package main

import "fmt"

func main() {
	//for、while、until、do，Go中只有一个for
	//for {} = for ; ; {} = for ; true ; {} = for true;{} 死循环
	//for i := 0; i < 5; i += 3 {
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	//break : 打破、破坏性的 for本趟不能正常执行完，

	/*
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ {
				fmt.Println(i, j)
			}
			fmt.Println("----------------------------")
		}
	*/
	/*
		for i := 1; i <= 9; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d * %d = %d  ", i, j, i*j)
			}
			fmt.Printf("\n")
		}
	*/

	//break END
	//goto END
	//END
	//fmt.Println("abc")

	/*
		for range 高级for
			主要是遍历容器内的元素的
			array、slice、map、channel、string
			for k,v := range map
			for v := range channel
	*/

	//for range
	//遍历字符串，将按照字符类型，即使是中文也是 按照rune一个一个取字符，rune就是int32.所以里面就是整数，整数是Unicode的码点
	//码点查表找到一个汉字字符
	for i, v := range "abc测试" {
		fmt.Printf("%T %[1]v ,%T %[2]v %[2]x %[2]q\n", i, v) //rune int32 ascii unicode
	}
	//int 0 ,int32 97 61 'a'
	//int 1 ,int32 98 62 'b'
	//int 2 ,int32 99 63 'c'
	//int 3 ,int32 27979 6d4b '测'
	//int 6 ,int32 35797 8bd5 '试'

	/*
		//int32 是4个字节 但是'测'占用3个字节
		[线性数据结构]
		string：汉字在utf-8中需要3个字节组成的数字表达
		for i,v:= range 字符串
		按照字符遍历，返回一个个rune的字符，int32 4Bytes整数，Unicode码点
		i索引，按照字节的索引，汉字占3个字节，底层字符串存储时对于非ascii，采用utf-8，汉字占3个字节
		for range 把utf8汉字编码替我们转成了unicode码
	*/

	fmt.Println("测试")           //内部存储utf-8
	fmt.Println("\xe6\xb5\x8b") //对应测字
	fmt.Println("\xe8\xaf\x95") //对应试字

	//for i,_ := range arr //如果第二个值不用可以简写	for i:= range arr
	//for _,v :=range arr 只取第二个值
	arr := [3]int{11, 33, 55}
	for i, v := range arr {
		fmt.Printf("%d:%d = %d\n", i, v, arr[i])
	}
	//0:11 = 11
	//1:33 = 33
	//2:55 = 55
}

【3-3-随机数】
=====================================================================
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	/*
		标准库"math/rand'
		我们使用的是伪随机数，是内部写好的公式计算出来的。这个公式运行提供一个种子，有这个种子作为起始值开始计算。
		src:= rand.NewSource(100)，使用种子100创建一个随机数源
		rand.New(rand.NewSource(time.Now().UnixNano())，利用当前时间的纳秒值做种子
		r10:= rand.New(src)，使用源创建随机数生成器
		r10.Intn(5)，返回[0,5)的随机整数
	*/

	/*Go 1.20+ 以下写法随机数 自动随机种子 直接拿来用*/
	//rand.Intn(n int)
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(rand.Intn(5) - 5) //返回[-5,5)的随机整数 算法推算的
		}
	*/

	/*1.20.0之前需要使用随机种子来随机数*/
	//OS有个参数 randautoseed随机种子
	os.Setenv("GODEBUG", "randautoseed=0")
	for i := 0; i < 5; i++ {
		fmt.Println(i, rand.Intn(8)) //rand 内部就是默认用1作为seed
	}
	// 0 1
	// 1 7
	// 2 7
	// 3 3
	// 4 1
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	src := rand.NewSource(1) //使用种子1创建一个随机数源
	r1 := rand.New(src)      //随机数产生器，驱动src源
	r5 := rand.New(rand.NewSource(5))
	for i := 0; i < 5; i++ {
		fmt.Println(i, r1.Intn(8), r5.Intn(8))
	}
	// 0 1 2
	// 1 7 4
	// 2 7 1
	// 3 3 0
	// 4 1 3
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	// 只要每一回把种子变了就行
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //利用当前时间的纳秒值做种子
	for i := 0; i < 5; i++ {
		fmt.Println(i, r.Intn(8))
	}
	// 0 2
	// 1 7
	// 2 2
	// 3 5
	// 4 7
}

【3-4-线性表存储原理】
=====================================================================
线性数据结构
	数学上逻辑概念：一组数据 有序 order 每个数据称为元素 元素有前后关系
	线性表
		有索引 习惯从0开始 可索引
		物理实现：顺序表、链接表
	内存编址线性的
	顺序表:字符串string、数组array、切片slice
		开辟一段连续的内存空间、元素一个挨着一个保存在这个内存空间中
			如果连续空间没有了，怎么办？ 需要GC垃圾回收，清理出来、如果完全没有 崩溃需要找原因为什么耗掉了这么多内存
				GC很耗时，Stop theworld STW，不能干活
				减少垃圾产生量，而不是不用内存  
					内存中所有你使用的空间，请你把它多用一会
		序列的元素的先后顺序，使用内存地址的顺序表达
		Cap容量：能够容纳几个元素
		len长度：现有元素个数 len<=cap
		
		很多语言中，数组的首地址就是第一个元素的地址(Go语言也是)
		
		C：Create 新增 
			 尾部追加append 代价非常小 推荐的方式
			 插入insert：
			 		 中间插入，占据插入点，插入点及之后的元素全部向后挪动
			 		 首部插入，变成队首，所有元素全部后移
			 如果数据少，数据规模小，随便你，随便完 但是如果规模大 考虑效率了
			 len + 1 
			 			可能出发扩容 len > cap
			 				不搬家，向后扩容
			 				要搬家，新址直接开辟足够的空间，复制所有元素过来
			 					旧址标记为拆(垃圾)，适当的时候出发GC
			 					计算机不是真的拆，标记为垃圾，就可以覆盖
			 			尽量一次分配好分配足够  
		R：Read   读取
				定位
					可以使用索引定位
						首地址+索引*元素类型占的字节数
							O(2) 2常量：代表做这件事总是只需要2步做完，不管元素有多少个,和元素规模无关
							O(1)
						时间复杂度O(?) 做事情需要几步，步数越多，一般就耗时		
				使用内容定位，该做还是做，但是能少做不多做，频繁做不合适
					需要遍历
						最优第一下就找到了，最差遍历完了还没找到
							O(1)              O(n)
				遍历
					容器内的元素不重复的拿一遍，O(n)					 
		U：Update 更新
				定位问题
				一旦定位，覆盖值
		D：Delete 删除
		   	len - 1
		   	pop 删除且返回给你用
			 	尾部删除pop，代价最小
			 	Remove
			 		中间删除，删除点其后元素向前挪动
			 		首部删除，所有元素都要向前挪动
			 	要注意规模大时，效率问题
		
		*新增、删除尾部操作较为合适、使用索引访问是最高效的
	
	链接表
		不需要连续的内存空间
		为每个元素单独开辟空间，用额外的内存保存下一个或上一个元素地址
		(value,next *(下一个元素的地址指针*)) 单向链表
		(value,next*,pre*) 双向链表
		手拉手模型、串珠模型
		len长度、元素实际个数
		cap容量 和 len一般相等

有序和排序
	有序order 你前我后 所有人排成一列
	排序sort 按照某个指标进行大小排序 升序和降序 排队而且按照大小个排
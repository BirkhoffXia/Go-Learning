【6-1-字符和整数】
=====================================================================
type rune = int32 //rune是int32的别名，4个字节，可以是Unicode字节
type byte = uint8 //byte是uint8的别名，1个字节
 
byte:兼容ASCII码字符，用byte类型即uint8别名，占用1个字节
rune:汉字等字符，unicode编码，用rune类型，即int32别名，占用4个字节
一个字符字面量使用单引号引起来

*字符串是若干个字符组成，在内存中使用utf-8编码。而rune只能保存一个字符，输出是unicode，不是UTF-8
'?' 1个字符，定义一个字符的
一个字符对应几个整数？1个
一个字符对应的整数占用几个字节？rune 4个字节 但是整数可能使用全部字节
'ab'错误： '\x61b'错误
'\x6d\x4b' 错误
var d int32 = 0x6d4b // 0x6d4b int类型

package main

import "fmt"

func main() {
	var c = '测' //字面量表达法，计算机内存存储时二进制编码，rune int32 4bytes Unicode码点
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", c)
	//int32,27979,测,'测',6d4b

	c = 'a' //无类型字面常量 类型没变 可以
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", c)
	//int32,97,a,'a',61

	// var d byte = '测' //一个字节放不下 rune类型
	var d byte = 0x61
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", d)
	//uint8,97,a,'a',61

	var e = 0x61
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", e)
	//int,97,a,'a',61 类型字面常量是int 4个字节

	var f byte = '\x61' //转移字符 16进制
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", f)
	//uint8,97,a,'a',61

	var g rune = 27979
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", g)
	//int32,27979,测,'测',6d4b

	var h rune = 0x6d4b
	fmt.Printf("%T,%[1]d,%[1]c,%[1]q,%[1]x\n", h)
	//int32,27979,测,'测',6d4b

	// var f rune = '\x6d\x4b'//单引号2个字节 不行
}


【6-2-字符串】
=====================================================================
字符串
	由多个字符组成的，用顺序表
			每个字符要编码，UTF-8编码得到的所有字符的字节有序序列
			字节序列
			UTF-8编码
					英文：1个字节
					中文：3个字节
	字符串表现不像切片，更像 int之类值类型
	Header 16个字节
			指针8个字节 指向底层顺序表 固定大小 数组[5]byte
			长度8个字节 字符串的字符个数？？？？？
					len(string)  时间复杂度:O(1) ，总字节数
	字面量表达法
			"" 空串，类型是string
 			"a" string    UTF-8，1个字节
			"测" string   UTF-8，3个字节
			"测试" string UTF-8，6个字节
	字面常量，只读，不可变，不可以修改元素，不可以更改长度
	可以索引
		"abc"[0] 时间复杂度O(1)
		用索引取，索引取值范围[0,len(string)-1]
		按照字节取
				中文有可能取一部分
		可以切么？对字符串切片
		遍历
				C风格for按照字节遍历   [按照字节遍历]
				for range 按照字符遍历 [按照字符遍历]
						byte -> int32 rune
						utf-8 中文编码 3bytes -> unicode码点 -> int32 rune 
						
package main

import "fmt"

func main() {
	s := "magedu"                   //UTF-8
	fmt.Printf("%T\n", s[len(s)-1]) //Byte uint8
	s = "马哥"                        //UTF-8 6Bytes
	fmt.Printf("%T,%[1]v\n", s[0])  //Byte uint8,233

	/*子串*/
	s = "magedu马哥"
	fmt.Printf("%T,%[1]v\n", s[1:])
	fmt.Printf("%T,%[1]s\n", s[1:])
	//string,agedu马哥
	//string,agedu马哥

	/*遍历*/
	s = "magedu马哥"
	//C 风格
	for i := 0; i < len(s); i++ { //英文:ASCII 中文：3个字节
		fmt.Printf("i = %d,%T,s[i] = %[2]d %[2]c\n", i, s[i])
	}
	// i = 1,uint8,s[i] = 97 a
	// i = 2,uint8,s[i] = 103 g
	// i = 3,uint8,s[i] = 101 e
	// i = 4,uint8,s[i] = 100 d
	// i = 5,uint8,s[i] = 117 u
	// i = 6,uint8,s[i] = 233 é
	// i = 7,uint8,s[i] = 169 ©
	// i = 8,uint8,s[i] = 172 ¬
	// i = 9,uint8,s[i] = 229 å
	// i = 10,uint8,s[i] = 147 
	// i = 11,uint8,s[i] = 165 ¥
	
	//按字符来，中文，byte -> int32，utf8 -> unicode -> rune(int32)
	for i, v := range s { 
		fmt.Printf("i = %d,%T;s[i] = %[2]d %[2]c;v = %T %[3]d %[3]c\n", i, s[i], v)
	}
	// i = 0,uint8;s[i] = 109 m;v = int32 109 m
	// i = 1,uint8;s[i] = 97 a;v = int32 97 a
	// i = 2,uint8;s[i] = 103 g;v = int32 103 g
	// i = 3,uint8;s[i] = 101 e;v = int32 101 e
	// i = 4,uint8;s[i] = 100 d;v = int32 100 d
	// i = 5,uint8;s[i] = 117 u;v = int32 117 u
	// i = 6,uint8;s[i] = 233 é;v = int32 39532 马
	// i = 9,uint8;s[i] = 229 å;v = int32 21733 哥
}
	
【6-3-字符串操作】
=====================================================================
[字符串拼接]
	/*
		简单拼接字符串常用+、fmt.Sprintf。如果手里正好有字符串的序列，
		可以考虑Join。如果反复多次拼接，strings.Builder是推荐的方式。
		bytes.Buffer用法同strings.Builder。
	*/
	
package main

import (
	"fmt"
	"strings"
)

func main() {
	s0 := "magedu.com"
	s1 := "马哥教育"
	fmt.Printf("%s %p;%s %p\n", s0, &s0, s1, &s1)
	//magedu.com 0xc00005e270;马哥教育 0xc00005e280

	/* + 拼接 正好有字符串 */
	s2 := s0 + s1
	fmt.Printf("%s %p\n", s2, &s2) //magedu.com马哥教育 0xc00005e2b0

	/*strings.Join 手里正好由字符串切片*/
	s3 := strings.Join([]string{s0, s1}, "------")
	fmt.Printf("%s %p\n", s3, &s3) // magedu.com------马哥教育 0xc0000882a0

	/*Sprintf*/
	s4 := fmt.Sprintf("%s======%s", s0, s1)
	fmt.Printf("%s %p\n", s4, &s4) //magedu.com======马哥教育 0xc0001022c0

	////需要多次才能拼成一个字符串，builder更合适
	var builder strings.Builder
	builder.Write([]byte(s0))
	builder.WriteByte('-')
	builder.WriteString("-")
	builder.WriteRune('-')
	builder.WriteString(s1)
	fmt.Println(builder.String())
	//magedu.com---马哥教育
}

[查询] :  时间复杂度O(n) 少用 能不用则不用，非必要不要用
[大小写]
ToLower:转换为小写
ToUpper:转换为大小
[前后缀]
HasPrefix:是否以子串开头
HasSuffix:是否以子串结尾
[移除]
[分割]
[替换]

package main

import (
	"fmt"
	"strings"
)

func main() {
	/*查询*/
	s := "www.magedu.com马哥教育"
	fmt.Println('马', "马"[0]) //39532 233 马占三个字节取第一个字节的整数位233

	fmt.Println(
		strings.Index(s, "马"),     //返回索引值 14
		strings.LastIndex(s, "马"), //从右向左找 找到还是正方向 14
		strings.IndexAny(s, "马m"), //4 返回找到的位置索引，chars多个字符组成的字符串，任意一个字符匹配立即返回
		//IndexAny返回s中字符中任何Unicode代码点的第一个实例的索引，如果s中不存在字符中的Unicode代码点，则返回-1。
		strings.IndexByte(s, 233), //14 一个字节一个字节找
		//IndexByte返回s中c的第一个实例的索引，如果s中不存在c，则返回-1。
		strings.IndexRune(s, 39532), //14
		//IndexRune返回Unicode代码点r的第一个实例的索引，如果s中不存在符文，则返回-1。如果r是utf8.RuneError，则返回任何无效UTF-8字节序列的第一个例子。
		strings.Contains(s, "马"), //true
		strings.Count(s, "m"),    //2
		strings.Count(s, "w"),    //3
		strings.Count(s, "ww"),   //1 一旦用到过了 就不再计数 指针指向第三个w了
		strings.Count(s, "www"),  //1
	)

	fmt.Println(
		strings.ToUpper(s),                //WWW.MAGEDU.COM马哥教育
		strings.HasPrefix(s, "edu"),       //false
		strings.HasSuffix(s[0:14], "com"), //true
	)

	fmt.Println(
		strings.TrimSpace("abcd \n\t"),                    //abcd 移除空白字符
		strings.TrimSpace("		\t\r\nab    cd   \n\t"),      //ab   cd
		strings.TrimPrefix("www.magedu.edu-马哥教育", "www."), //如果开头或结尾匹配，则去除。否则，返回原字符串副本
		//magedu.edu-马哥教育
		strings.TrimSuffix("www.magedu.edu-马哥教育", ".edu"), //如果开头或结尾匹配，则去除。否则，返回原字符串副本
		//www.magedu.edu-马哥教育
		strings.TrimLeft("abcdddeabeccc", "dbac"), //字符串开头的字符如果在字符集中，则全部移除，直到碰到第一个不在字符集中的字符为止
		//eabeccc
		strings.TrimRight("abcdddeabeccc", "acdb"), //字符串结尾的字符如果在字符集中，则全部移除，直到碰到第一个不在字符集中的字符为止
		//abcdddeabe
		strings.Trim("abcdddeabeccc", "bacd"), //字符串两头的字符如果在字符集中，则全部移除，直到左或右都碰到第一个不在的字符集中的字符为止
		//eabe
	)

	/*分割*/
	s1 := strings.Split(s, ".") //
	fmt.Println(s1, len(s1))    //[www magedu com马哥教育] 3
	fmt.Println("-----------------------------------")
	fmt.Println(strings.Split(s, "=")) //[www.magedu.com马哥教育]
	fmt.Println(strings.Split(s, ""))  //返回的都是字符串类型 返回的都是字符串切片
	//[w w w . m a g e d u . c o m 马 哥 教 育]
	s2 := strings.Split(s, "")
	for i, v := range s2 {
		fmt.Printf("%d %s %[2]T\n", i, v) //返回的都是字符串类型 返回的都是字符串切片
	}
	// 0 w string
	// 1 w string
	// 2 w string
	// 3 . string
	// 4 m string
	// 5 a string
	// 6 g string
	// 7 e string
	// 8 d string
	// 9 u string
	// 10 . string
	// 11 c string
	// 12 o string
	// 13 m string
	// 14 马 string
	// 15 哥 string
	// 16 教 string
	// 17 育 string
	
	fmt.Println(strings.SplitAfter(s, ".")) //切的元素跟在后面
	//[www. magedu. com马哥教育]
	fmt.Println(strings.SplitAfter(s, "="))
	//[www.magedu.com马哥教育]
	fmt.Println(strings.SplitAfter(s, ""))
	//[w w w . m a g e d u . c o m 马 哥 教 育]
	fmt.Println("-----------------------------------")
	
	fmt.Println(strings.SplitN(s, ".", 1)) //N 返回切片的元素的个数
	//[www.magedu.com马哥教育]
	fmt.Println(strings.SplitN(s, ".", 2)) //[www magedu.com马哥教育]
	fmt.Println("-----------------------------------")
	
	fmt.Println(strings.SplitAfterN(s, ".", 2)) //[www. magedu.com马哥教育]
	fmt.Println("-----------------------------------")
	
	fmt.Println(strings.Cut(s, ".")) //Cut 返回三个元素 最后一个返回bool表示是否成功切成功
	//www magedu.com马哥教育 true
	
}


[Map]
	数据结构，字典、映射(key-value)、hash表
	映射
		x--f-->y,y=f(x) 一对一
		
package main

import (
	"fmt"
	"strings"
)

func main() {
	/*替换*/
	/*
		Replace(s,old,new string,n int)string
			n<0 等价于ReplaceAll,全部替换
			n==0 或者 old==new 返回s
			n>0 最多替换n次、如果n超过找到old子串的次数x，也就只能替换x次了
			未到到替换初，就返回s
	*/
	s0 := "www..magedu.com马哥教育."
	fmt.Println(strings.ReplaceAll(s0, "..", ".")) //返回一个新字符串
	//www.magedu.com马哥教育.

	fmt.Println(strings.Replace(s0, "..", ".", 1)) //返回一个新字符串
	//www.magedu.com马哥教育.

	/*其他 Repeat Map*/
	fmt.Println(strings.Repeat("~", 30))
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	for i, v := range s0 {
		fmt.Printf("%d %d %[2]c\n", i, v)
	}
	// 0 119 w
	// 1 119 w
	// 2 119 w
	// 3 46 .
	// 4 46 .
	// 5 109 m
	// 6 97 a
	// 7 103 g
	// 8 101 e
	// 9 100 d
	// 10 117 u
	// 11 46 .
	// 12 99 c
	// 13 111 o
	// 14 109 m
	// 15 39532 马
	// 18 21733 哥
	// 21 25945 教
	// 24 32946 育
	// 27 46 .

	/*Mapping*/
	//按照给定处理每个rune字符的函数依次处理每个字符后，拼接成字符串返回，注意Map是一对一的映射，不能减少元素个数
	t := strings.Map(
		func(r rune /*入参*/) rune /*出参*/ {
			fmt.Printf("%T %[1]d\n", r)
			return r
		}, //匿名函数、回调函数
		s0, //被遍历的字符串 for-range 每次遍历一个字符rune
	)
	fmt.Printf("%T:%[1]s\n", t)
	// 0 119 w
	// 1 119 w
	// 2 119 w
	// 3 46 .
	// 4 46 .
	// 5 109 m
	// 6 97 a
	// 7 103 g
	// 8 101 e
	// 9 100 d
	// 10 117 u
	// 11 46 .
	// 12 99 c
	// 13 111 o
	// 14 109 m
	// 15 39532 马
	// 18 21733 哥
	// 21 25945 教
	// 24 32946 育
	// 27 46 .
	// string:www..magedu.com马哥教育.
}

【6-4-类型转换】
=====================================================================
[数值类型转换]
	低精度向高精度转换可以，高精度向低精度转换会损失精度
	无符号向有有符号转换，最高位是符号位
	byte和int可以互相转换
	float和int可以相互转换，float到int会丢失精度
	bool和init不能互相转换
	不同长度的int和float之间可以互相转换

package main

import "fmt"

func main() {
	var i int8 = -1
	var j uint8 = uint8(i)
	fmt.Println(i, j) //-1 255

	/*
		-1 byte : 计算机内部负数采用补码
			有符号
				源码：1000 0001
				反码：符号位不变，其余按位取反
				反码：1111 1110
				补码：反码+1
				补码：1111 1111

				源码转换：
					符号位不变，按位取反
					1000 0000 + 1
					1000 0001
					---------------
				    -        1
	*/

	//不可以直接对字面常量进行 修改 使用变量进行强制转换
	//cannot convert 3.14 (untyped float constant) to type int
	// fmt.Println(int(3.14)) //错误 不允许无类型float常量转到int
	var a = 3.14 //定义有类型变量转换就没有问题
	fmt.Printf("%T:%[1]v => %T %[2]d\n", a, int(a))
	//float64:3.14 => int 3

	//byte rune本质上就是整数和无类型常量可以直接计算，自动转换
	b := 'a'
	c := b + 1
	fmt.Printf("%T %[1]v\n", b)       //int32 97
	fmt.Printf("%T %[1]v %[1]c\n", c) //int32 98 b
	
	//但是，如果不使用无类型常量，有类型的计算如果类型不一致要报错，因为Go是对类型要求非常严苛的语言，要强制类型转换。
	/*
		b := 'a'
		e := 1
		c := b + e // rune和int类型不能加，必须转换。比如c:= int(b)+ e或c := b + rune(e)
		fmt.Println(c)
	*/
}

[类型别名和类型定义]
package main

import "fmt"

// type MyInt uint8 //不加等号 新类型定义 两个不同类型 我基于你 但是和你不一样
type MyInt = uint8 //别名

func main() {
	//type byte = uint8
	//byte只是存在于代码中，为了方便阅读或使用，编译完成后，不会有byte类型。
	var a byte = 'A'
	var b uint8 = 3
	fmt.Printf("%T %T\n", a, b)   //uint8 uint8
	fmt.Printf("%T %[1]c\n", a+b) //uint8 D

	var c MyInt = 1
	//a + c ?
	fmt.Printf("%T %[1]c\n", a+c) //uint8 B
	// fmt.Printf("%T %[1]c\n", a+byte(c)) //uint8 B
}

[字符串转换]
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转换 string(整数)：整数作为码点，ASCII Unicode找字符
	fmt.Println(string(126), string(27979)) //~ 测 utf-8 3个字节

	s1 := "126"
	fmt.Printf("%T %[1]q\n", strconv.Itoa(97)) //整数转换 string "97" 2个字节长度
	fmt.Println(strconv.Atoi(s1))              //126 <nil> 十进制理解126
	fmt.Println(strconv.ParseInt(s1, 10, 32))  //返回int64类型 十进制理解126       126 <nil>
	fmt.Println(strconv.ParseInt(s1, 16, 32))  //返回int64类型 十六进制理解0x126   294 <nil>

	s2 := "115.6"
	fmt.Println(strconv.ParseFloat(s2, 64)) //115.6 <nil>
	fmt.Println(strconv.ParseBool("true"))  //true <nil>
}

[字符串与字节序列转换]
string(一个整数)，强制类型转换一个整数，相当于把整数当unicode码，去查一个字符，最后返回字符串
string(整数序列)，强制类型转换一个整数序列，也是转成字符串

package main

import (
	"fmt"
)

func main() {
	//转换为序列
	//强制类型转 string => []byte；string => []rune
	//[]byte 字节序列  []rune:rune序列
	s1 := "abc"
	fmt.Println([]byte(s1)) //[]byte, [97 98 99] Header 24个字节 1个字符占1个字节共3个字节
	fmt.Println([]rune(s1)) //[]int32,[97 98 99] Header 24个字节 1个字符占4个字节共12个字节
	fmt.Println("--------------------------------------")

	s2 := "测试"              //6 bytes strings
	fmt.Println([]byte(s2)) //[]byte 6 bytes底层数组 [230 181 139 232 175 149]
	fmt.Println(s2[1])      //181

	//通过字节序列 按照 string进行转换 UTF-8 3个字节转换为中文
	fmt.Println(string([]byte{230, 181, 139}))                //string "测" 3bytes
	fmt.Println(string([]byte{230, 181, 139, 232, 175, 149})) //string "测试" 6bytes
	fmt.Println(string([]byte{230, 181}))                     //�

	fmt.Println([]rune(s2)) //[]rune ，2个元素，每个字符utf-8 => Unicode  //[27979 35797]
	fmt.Println(string([]rune{27979, 35797})) //string "测试" 6bytes utf-8 //测试
	fmt.Println(string([]byte{97, 49, 65})) //string 占3个字节 "a1A"
}

【6-5-哈希原理】
=====================================================================
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x := "abc"
	//16进制 每4位一段 256/4=64个字符
	h := sha256.New() //256/4=64个字符组成的字符串，字节序列 0~9a~f
	// h := md5.New()// 128/4=32个字符组成的字符串
	h.Write([]byte(x))
	y := h.Sum(nil)
	fmt.Printf("%T %[1]v\n", y) //[]byte
	for i, v := range y {
		fmt.Printf("%d,%x,%[2]d\n", i, v)
	}
	r := fmt.Sprintf("%x", y)
	fmt.Printf("%T %[1]s,len=%d\n", r, len(r))
	// []uint8 [186 120 22 191 143 1 207 234 65 65 64 222 93 174 34 35 176 3 97 163 150 23 122 156 180 16 255 97 242 0 21 173]
	// 0,ba,186
	// 1,78,120
	// 2,16,22
	// 3,bf,191
	// 4,8f,143
	// 5,1,1
	// 6,cf,207
	// 7,ea,234
	// 8,41,65
	// 9,41,65
	// 10,40,64
	// 11,de,222
	// 12,5d,93
	// 13,ae,174
	// 14,22,34
	// 15,23,35
	// 16,b0,176
	// 17,3,3
	// 	18,61,97
	// 	19,a3,163
	// 	20,96,150
	// 	21,17,23
	// 	22,7a,122
	// 	23,9c,156
	// 	24,b4,180
	// 	25,10,16
	// 	26,ff,255
	// 	27,61,97
	// 	28,f2,242
	// 	29,0,0
	// 	30,15,21
	// 	31,ad,173
	// string ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad,len=64
}
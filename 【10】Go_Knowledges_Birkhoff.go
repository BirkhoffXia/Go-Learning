Viedo Date: 2023/10/21
Made By:	BIRKHOFF
Date:	2024-06-25

【10-1-面向对象】
=====================================================================
[面向对象]
	是我们认识世界的一种方法论
	对象：类(类型 class) 一类事务的抽象

	张三、李四吃鱼
		抽象
			人类
				名字(字段、属性、数据成员)
				吃(动作、操作、方法)
			鱼类
				吃(动作、操作、方法)
		实例(创建实例的过程为实例化) instance | object
			具体的人，它是抽象概念人类的实实在在的一个实例，它的名字为张三
			具体的人，它是抽象概念人类的实实在在的一个实例，它的名字为张三
			具体的鱼，它是抽象概念鱼类的实实在在的一个实例
			实例张三、实例李四 正在吃具体的那个鱼
			每一次实例化得到的实例都不同的

			单例模式：单实例，不管你怎么实例化，只能得到同一个实例
			工厂模式：批量生成实例的函数、给一些构造用的实参，返回一个个不同的实例
		动作
			不同类型之间的相互作用
			吃，具体的动作，具体的人发出了具体的这个动作
			动作应该属于具体的人吗？动作的定义是定义在类型上的，因为所有人都可以发出这个动作

	三要素
		封装
			通过结构体，可以把数据字段封装在内，还可以为结构体提供方法
			访问控制:暴露某些属性和方法，隐藏一些属性和方法
				属性、方法标识符首字母大写，实现了对包外可见的访问控制
				属性、方法标识符首字母小写，仅包内可见
				通过首字母大小写，在一定程度上实现了public、private的访问控制
		继承：
			Go语言没有提供继承的语法，实际上需要通过匿名结构体嵌入(组合)来实现类似效果
			子类从父类继承，那么父类的属性和方法就不需要再子类种重新定义
			子类可以对父类中没有属性和方法进行定义，那么这些属性和方法和父类无关
			覆盖override
				也称重写。注意不是重载overload。
				子类可以对父类已有的属性和方法进行重新的定义
		多态
			Go语言不能像]ava语言一样使用多态，但可以通过引入interface接口来解决。
			前提：继承、覆盖
			不同子类对父类有覆盖，导致不同子类调用覆盖的方法表现不同，表现出不同的态


	Go中面向对象
		封装
			type 类型 struct{
				属性1 类型
				属性2类型
			}

			func(u User) getName() string {}
	访问控制
		实现了一个简单的访问控制
		类型名大写，包内外都可见;小写，包内可用
		属性名小写，包内可见;属性大写且类型名大写，包外可见
		方法名小写，包内可见;方法大写且类型名大写，包外可见
		成员名小写，包内可见;成员大写且类型名大写，包外可见
		其他语言，public、private、protected

	构造函数
		普通函数 NewXxxx
		EG：
		type Animal struct {
			name string
			age  int
		}

		func NewAnimal(name string, age int) *Animal {
			return &Animal{name, age}
		}

		func NewDefaultAnimal() *Animal {
			return &Animal{"nobody", 1}
		}

	继承
		父子关系
			匿名属性成员
		父类的属性和方法都可以继承下来，不需要重复写代码

	覆盖、重写
		继承自父类的方法或属性，重新定义
			1、完全覆盖:子类的实现和父类完全不一样，不需要父类的实现
			2、锦上添花:继承并发扬，子类的实现需要依赖父类，一般情况下，需要先调用父类方法
				 子类的自己实现调用父类的方法会对属性做一些操作
		覆盖后的访问原则：优先自己，就近原则
					EG：
					package main

					import (
						"fmt"
					)

					type Animal struct {
						name string
						age  int
					}
					type Cat struct {
						Animal
						color string
					}

					// 定义一个方法
					func (*Animal) run() {
						fmt.Println("Animal run -----------")
					}
					func NewAnimal(name string, age int) *Animal {
						return &Animal{name, age}
					}
					func NewDefaultAnimal() *Animal {
						return &Animal{"nobody", 1}
					}

					// 定义Cat独有swim
					func (*Cat) swim() {
						fmt.Println("Cat swim -----------")
					}

					// 重写Anim run()
					// func (*Cat) run() { //override 完全覆盖
					//
					//		fmt.Println("Cat run -----------")
					//	}

					// override 继承并发扬 实现自己的需求
					func (c *Cat) run() {
						c.Animal.run() // 继承并发扬
						fmt.Println("Cat run **************")
					}

					func main() {
						cat := new(Cat)
						fmt.Println(cat.name, cat.age, cat.color)
						// cat.run()        //继承 猫没定义run()之前 继承父类：Animal run -----------
						cat.run()        //猫重写run() Cat run -----------
						cat.Animal.run() // Animal run ----------- 不影响父类定义run()
						cat.swim()       //Cat swim -----------

						fmt.Println("--------------------------")
						a := NewAnimal("XKS", 30)
						a.run() //Animal run -----------
					}

	多态
		面向对象编程中非常好用，灵活
		前提:父子继承、覆盖
		不同子类对父类有覆盖，导致不同子类调用覆盖的方法表现不同，表现出不同的态
EG：
package main

import "fmt"

type Animal struct {
	name string
	age  int
}
type Cat struct {
	Animal
	color string
}
type Dog struct {
	Animal
	color string
}

func (*Animal) run() {
	fmt.Println("Animal run -----------")
}

func (c *Cat) run() {
	c.Animal.run()
	fmt.Println("Cat run **************")
}

func (d *Dog) run() {
	fmt.Println("Dog run ##############")
}

// 使用接口
// Animal、Cat、Dog结构体都实现了Runner接口
type Runner interface {
	run()
}

/*此函数 支持Cat 和 Dog run() 因为都实现了run()函数*/
/*
    Q: 为什么这里可以传入一个 接口为参数 test(a Runner)？
    A: 这Cat、Dog 2种类型都实现了该接口，实例既是结构体类型又是接口类型。
			 Go语言中，接口类型可以被任何实现接口定义的类型来赋值。
				这样就使得可以将不同的类型传入同的函数，并保证它们都能正确地执行相同的操作。
				在你所提供的代码中，test 函数期望一个 Runner 接口作为参数，
				而任何实现了 run() 方法的类型都可以被作 Runner 接口。
*/

func test(a Runner) {
	fmt.Printf("a: %T\n", a) //a: *main.Dog 、a: *main.Cat
	a.run()
}

func main() {
	//var a Animal = Cat{} Go不能这样赋值
	//a.run()
	c := new(Cat)
	d := new(Dog)
	test(c)
	test(d)

	fmt.Printf("%T,%p\n", c, &c) //*main.Cat,0xc00000a028
	fmt.Printf("%T,%p\n", d, &d) //*main.Dog,0xc00000a030
}
// Animal run -----------
// Cat run **************
// Dog run ##############

【10-2-排序精讲】
=====================================================================
排序
	前提:排成一队，构造线性数据结构，往往使用顺序表
	遍历:所有元素都要参与
	对比:需要维度，身高、年龄、体重
	互换:调整位置
	非线性数据结构能不能排序? 不能直接排序，线性化过程，元素不重复的拿出来排成一排

	边界：使用索引，索引不要超界
	对比: 两两比较，需要维度，身高、年龄、体重
	互换:两两调整位置
	排序算法
		交互排序:冒泡法
		选择排序、堆排序选择排序:
		插入排序:插入排序
		快速排序:递归排序
	时间复杂度
		无限期望o(n)
		大多数o(n^2)
		好一点o(nlogn)
		排序效率都不高，能少做最好
	应用
		TOPN 前10个

	非线性
		线性化(切片)

[结构体排序]
参照sort.Ints()的实现
func Ints(x []int) { Sort(IntSlice(x)) }

func Sort(data Interface) {...}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int //IntSlice实现了Interface接口 所以IntSlice既是[]int类型也是Interface接口类型

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x IntSlice) Sort() { Sort(x) }


EG:
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type StudentSlice []Student

func (a StudentSlice) Len() int           { return len(a) }
func (a StudentSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StudentSlice) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	tom := Student{"tom", 20}
	jerry := Student{"jerry", 10}

	students := []Student{tom, jerry} //线性化
	fmt.Println(students)

	//func Sort(data Interface) {  StudentSlice 实现了Interface中的三个函数Len\Swap\Less
	//升序
	// sort.Sort(StudentSlice(students)) //接口类型的 必须实现那三个方法
	//降序
	sort.Sort(sort.Reverse(StudentSlice(students))) //[{tom 20} {jerry 10}]
	fmt.Println(students)

	// [{tom 20} {jerry 10}]
	// [{jerry 10} {tom 20}]
}

[切片排序]
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	students := []Student{
		Student{"tom", 20},
		Student{"jerry", 10},
		Student{"ben", 8},
	}
	fmt.Println(students)

	//对Slice进行排序
	// sort.Slice()
	//Slice(x any, less func(i, j int) bool)
	sort.Slice()
	sort.Slice(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	fmt.Println(students)
}

[Map排序]-key排序
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"
	fmt.Println(m) //map[0:b 1:a 2:c]

	//To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

// map[0:b 1:a 2:c]
// Key: 0 Value: b
// Key: 1 Value: a
// Key: 2 Value: c

[Map排序]-value排序
package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Key   int
	Value string
}

func main() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"
	fmt.Println(m) //map[0:b 1:a 2:c]

	//通过把map 放到一个 结构体中
	entries := make([]Entry, len(m))
	i := 0
	for k, v := range m {
		entries[i] = Entry{k, v}
		i++
	}
	fmt.Println(entries) //[{1 a} {2 c} {0 b}]  此时还没排序

	//然后 自定函数进行比较
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Value < entries[j].Value
	}) //Value 升序
	fmt.Println(entries) //[{1 a} {0 b} {2 c}] 已排序

}

Q:排序的元素类型需要一样么？
A1:Len、Swap不是问题
	不同类型元素可以按照某种维度可以比较大小吗？
		看情况
		如果你能够对其中不同类型找到一个维度能比较大小就可以了

A2；元素类型不需要一样，只要能提供元素之间如何比较大小就行。但是多数情况下，都是相同类型元素比较大小。
		不同类型元素之间比较大小意义不大，极少使用。

【10-3-序列化概念】
=====================================================================
[序列化和反序列化]
Q:为什么要序列化?
A1:内存中的map、slice、array以及各种对象，如何保存到一个文件中?如果是自己定义的结构体的实例，如何保存到一个文件中?
如何从文件中读取数据，并让它们在内存中再次恢复成自己对应的类型的实例?
	要设计一套协议，按照某种规则，把内存中数据保存到文件中。文件是一个字节序列，所以必须把数据转换成字节序列，输出到文件。
	这就是序列化。 反之，从文件的字节序列恢复到内存并且还是原来的类型，就是反序列化。
A2:
		内存中数据结构非常复杂，掉电、进程崩溃导致数据丢失，会把数据存储或者网络传输，存储、传输都是按照字节来的
		所以要把内存中的 立体的数据结构转换成字节序列
		应用:网络传输、持久化(首先序列化然后存储到存储介质上)、进程间传递数据
	二进制序列化:更加高效
		序列化:内存中某种类型的数据结构(立体结构)=>二进制字节序列
		反序列化:二进制的字节序列=>内存中原来的类型的数据(立体结构)Protocol buffers、标准库、MessagePack等
	字符(文本)序列化-人类可读
		内存中某种类型的数据结构(立体结构)=>字符序列(字符串)
		字符序列(字符串)=>内存中原来的类型的数据(立体结构)
		JSON、XML等

【10-4-Json序列化】
=====================================================================
JSON
	浏览器是个软件
		HTML解析、渲染
		网景公司发明Javascript解释器，js引擎
		ECMA script，简称ES6

	WEB中，网络传输数据基于应用层协议HTTP协议，把一部分数据(某种类型)采用JsN来文本序列化，到了客户端浏览器中，js引擎度
	序列化json文本，还原到内存中，某种类型的数据
	浏览器广泛支持
	序列化内存中某类型数据=>字符串
	问题
		字符串表达字符串类型的数据、整型的数据、数组数据等
		Go、Java、c++、Python等语言中的不同的类型的数据 序列化为 Json格式的数据(大字符串)
		通过网络协议传输到客户端
		浏览器 Js引擎解析 大字符串解析为 Javascript的数据类型
		Python代码解析?  大字符串解析为Python的数据类型
		Java也是如此
		跨平台了
	标准库Json，反射效率比较低
		序列化方法Marshal ,Encode编码
		反序列化方法Unmarshal，Decode解码
	类型
		数值类型number -> float64
		对象类型定义 {"1":}

EG:
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//序列化
	var data = []any{
		100, 20.5, true, false, nil, "aabb", //基本类型
		[3]int{97, 98, 99},                  //Go array => js array
		[]int{65, 66, 67},                   //Go slice => js array
		map[string]int{"abc": 49, "aa": 50}, //Go map => js object
	}

	var target = make([][]byte, 0, len(data))
	for i, v := range data { //一个一个单独序列化，看变化
		b, err := json.Marshal(v)
		if err != nil {
			continue
		}
		fmt.Printf("%d %T: %[2]v => %T %[3]v %s\n", i, v, b, string(b))
		target = append(target, b)
	}
	// 0 int: 100 => []uint8 [49 48 48] 100
	// 1 float64: 20.5 => []uint8 [50 48 46 53] 20.5
	// 2 bool: true => []uint8 [116 114 117 101] true
	// 3 bool: false => []uint8 [102 97 108 115 101] false
	// 4 <nil>: <nil> => []uint8 [110 117 108 108] null
	// 5 string: aabb => []uint8 [34 97 97 98 98 34] "aabb"
	// 6 [3]int: [97 98 99] => []uint8 [91 57 55 44 57 56 44 57 57 93] [97,98,99]
	// 7 []int: [65 66 67] => []uint8 [91 54 53 44 54 54 44 54 55 93] [65,66,67]
	// 8 map[string]int: map[aa:50 abc:49] => []uint8 [123 34 97 97 34 58 53 48 44 34 97 98 99 34 58 52 57 125] {"aa":50,"abc":49}
	fmt.Println("------------------")
	// x, err := json.Marshal(data)
	// fmt.Printf("%s err = %v\n", x, err) //人类所能看到的
	//[100,20.5,true,false,null,"aabb",[97,98,99],[65,66,67],{"aa":50,"abc":49}] err = <nil>
	// fmt.Printf("%v err = %v\n", x, err) //计算机语言
	//[91 49 48 48 44 50 48 46 53 44 116 114 117 101 44 102 97 108 115 101 44 110 117 108 108 44 34 97 97 98 98 34 44 91 57 55 44 57 56 44 57 57 93 44 91 54 53 44 54 54 44 54 55 93 44 123 34 97 97 34 58 53 48 44 34 97 98 99 34 58 52 57 125 93] err = <nil

	/*反序列化*/
	for i, v := range target {
		var t any
		err := json.Unmarshal(v, &t) //这里注意是& 地址
		if err != nil {
			continue
		}
		fmt.Printf("%d: %T %[2]v %[2]s,%T %[3]v\n", i, v, t)
	}
	// 0: []uint8 [49 48 48] 100,float64 100
	// 1: []uint8 [50 48 46 53] 20.5,float64 20.5
	// 2: []uint8 [116 114 117 101] true,bool true
	// 3: []uint8 [102 97 108 115 101] false,bool false
	// 4: []uint8 [110 117 108 108] null,<nil> <nil>
	// 5: []uint8 [34 97 97 98 98 34] "aabb",string aabb
	// 6: []uint8 [91 57 55 44 57 56 44 57 57 93] [97,98,99],[]interface {} [97 98 99]
	// 7: []uint8 [91 54 53 44 54 54 44 54 55 93] [65,66,67],[]interface {} [65 66 67]
	// 8: []uint8 [123 34 97 97 34 58 53 48 44 34 97 98 99 34 58 52 57 125] {"aa":50,"abc":49},map[string]interface {} map[aa:50 abc:49]}
}

【10-5-结构图序列化和msgpack】
=====================================================================
[结构体序列化]

EG:
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//结构体序列化
	var data = Person{
		Name: "BIRKHOFF",
		Age:  20,
	}

	b, err := json.Marshal(data) //[]byte string
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)            //这是Person的实例  {Name:BIRKHOFF Age:20}
	fmt.Printf("%v, %s\n", b, string(b)) //这是字符串 [123 34 78 97 109 101 34 58 34 66 73 82 75 72 79 70 70 34 44 34 65 103 101 34 58 50 48 125], {"Name":"BIRKHOFF","Age":20}

	fmt.Println("=============================")

	//反序列化
	//[]byte(`{"Name":"BIRHOFF","Age":20}`) 这种写法时 字符标签 
	var b1 = []byte(`{"Name":"BIRHOFF","Age":20}`) //字符串，增加了些空格，js中对象也就是键值对
	//【1】知道目标类型时
	var p Person

	err = json.Unmarshal(b1, &p) //填充成功，通过指针填充结构体
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", p) //main.Person {Name:BIRHOFF Age:20}
	fmt.Println("=============================")

	//【2】不知道类型时
	var i interface{}
	err = json.Unmarshal(b1, &i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", i) //map[string]interface {} map[Age:20 Name:BIRHOFF]
}

[切片序列化]
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//结构体序列化
	var data = []Person{
		{Name: "AAA", Age: 20},
		{Name: "aaa", Age: 30},
	}

	b, err := json.Marshal(data) //[]byte string
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data) //这是Person的实例  [{Name:AAA Age:20} {Name:aaa Age:30}]
	fmt.Printf("%v, %s\n", b, string(b))
	//这是字符串
	//[91 123 34 78 97 109 101 34 58 34 65 65 65 34 44 34 65 103 101 34 58 50 48 125 44 123 34 78 97 109 101 34 58 34 97 97 97 34 44 34 65 103 101 34 58 51 48 125 93], [{"Name":"AAA","Age":20},{"Name":"aaa","Age":30}]
	fmt.Println("=============================")

	//反序列化

	//【1】不知道类型时
	var i interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", i) //[]interface {} [map[Age:20 Name:AAA] map[Age:30 Name:aaa]]

	fmt.Println("=============================")

	//【2】知道目标类型时
	var b1 = []byte(`[{"Name":"BIRHOFF","Age":20},{"Name":"HEYE","Age":30}]`)
	var j []Person
	err = json.Unmarshal(b1, &j)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", j) //[]main.Person [{Name:BIRHOFF Age:20} {Name:HEYE Age:30}]
}

[字符标签]
type Person struct {
	Name   string `json:"t"` //tag 属性 序列化，反序列化设置 输出时就是这个t名字
	Age    int    `json:"a,omitempty"`
	Weight int    `json:"-"` //不输出
}
结构体的字段可以增加标签tag，序列化、反序列化时使用
	在字段类型后，可以跟反引号引起来的一个标签，用json为key，value用双引号引起来写，key与value直接使用冒号，这个标签中不要加入多余空格，否则语法错误
		Name string`json:"name"，这个例子序列化后得到的属性名为name
			json表示json库使用
			双引号内第一个参数用来指定字段转换使用的名称，多个参数使用逗号隔开
		Name string json:"name,omitempty"，omitempty为序列化时忽略空值，也就是该字段不序列化
			空值为false、0、空数组、空切片、空map、空串、nil空指针、nil接口值
			空数组、空切片、空串、空map，长度len为0，也就是容器没有元素
		如果使用-，该字段将被忽略
			Name string `json:"_"、序列化后没有该字段，反序列化也不会转换该字段
			Name string `json:"_,",序列化后该字段显示但名为""，反序列化也会转换该字段
		多标签使用空格间隔
			Name string`json:"name,omitempty" msgpack:"myname"

[MessagePack]
MessagePack是一个基于二进制高效的对象序列化类库，可用于跨语言通信。 
它可以像JSON那样，在许多种语言之间交换结构对象。但是它比JSON更快速也更轻巧。
支持Python、Ruby、Java、C/C++、Go等众多语言。
宣称比Google Protocol Buffers还要快4倍。

go get github.com/vmihailenco/msgpack/v5
go: downloading github.com/vmihailenco/msgpack/v5 v5.3.5
go: downloading github.com/vmihailenco/tagparser/v2 v2.0.0
go: added github.com/vmihailenco/msgpack/v5 v5.3.5
go: added github.com/vmihailenco/tagparser/v2 v2.0.0

package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

type Person struct {
	Name string `json:"name" msgpack:"myname"` //tag 属性 序列化，反序列化设置 输出时就是这个t名字
	Age  int    `json:"age" msgpack:"myage"`
}

func main() {
	//结构体序列化
	var data = []Person{
		{Name: "Tom", Age: 16},
		{Name: "Jerry", Age: 21},
	}

	b, err := msgpack.Marshal(data) //[]byte string
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data) //这是Person的实例  [{Name:AAA Age:20} {Name:aaa Age:30}]
	fmt.Printf("%v, %s\n", b, string(b))
	fmt.Println("=============================")

	var j []Person
	err = msgpack.Unmarshal(b, &j)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T %+[1]v\n", j) //
}

// [{Name:Tom Age:16} {Name:Jerry Age:21}]
// [146 130 166 109 121 110 97 109 101 163 84 111 109 165 109 121 97 103 101 16 130 166 109 121 110 97 109 101 165 74 101 114 114 121 165 109 121 97 103 101 21],
// ���myname�Tom�myage��myname�Jerry�myage
// =============================
// []main.Person [{Name:Tom Age:16} {Name:Jerry Age:21}]



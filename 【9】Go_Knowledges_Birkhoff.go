Viedo Date: 2023/10/14
Made By:	BIRKHOFF 
Date:	2024-06-24   

【9-1-结构体及成员】
=====================================================================
Go语言的结构体有点像面向对象语言中的“类”，但不完全是，Go语言也没打算真正实现面向对象范式。

结构体可以扩展方法，因为结构体是类型
	方法method，本质上就是function
	User.getName //类型的调用 人类的名字 不合适
	u1.getName //User具体
	
	
		
Go中，可以为新的类型扩展方法
	新的，可以基于原来的类型生成新的类型，最好不要在原来类型上直接修改

[定义]
使用type定义结构体，可以把结构体看做类型使用。必须指定结构体的字段(属性)名称和类型
type User struct {
	id int
	name,addr string //多个字段类型相同可以合并写
	score float32
}

User不过是个标识符，是一个指代罢了
真正的类型定义是struct{}的部分

[初始化]
package main

import (
	"fmt"
)

type User struct {
	id         int
	name, addr string //多个字段类型相同可以合并写
	score      float32
}

func main() {
	// 初始化
	//1 零值 很常用，经常用
	var u1 User // u1是User的实例
	//%v打印结构体
	fmt.Printf("%T %[1]v", u1) //main.User {0   0}

	u1.id = 1
	fmt.Printf("%T %[1]v", u1) //main.User {1   0}
	fmt.Println("------------------------------------\n")
	fmt.Println(u1)         //{1   0}
	fmt.Printf("%v\n", u1)  //{1   0}                                    默认缺省格式打印
	fmt.Printf("%+v\n", u1) //{id:1 name: addr: score:0}                 增强打印
	fmt.Printf("%#v\n", u1) //main.User{id:1, name:"", addr:"", score:0} 详细打印

	//2 字面量初始化，推荐
	u2 := User{}            //字段为零值
	fmt.Printf("%#v\n", u2) //main.User{id:0, name:"", addr:"", score:0}

	//3 字面量初始化，field:value为字段赋值
	u3 := User{id: 100}
	fmt.Printf("%#v\n", u3) //main.User{id:100, name:"", addr:"", score:0}

	u4 := User{
		id:    319,
		score: 99,
		addr:  "Shanghai",
		name:  "BIRKHOFF",
	} //名称对应无所谓顺序
	fmt.Printf("%#v\n", u4) //main.User{id:319, name:"BIRKHOFF", addr:"Shanghai", score:99}

	u5 := User{1220, "HEYE", "SHANGHAI", 59} //无字段名称必须按照顺序给出全部字段值
	fmt.Printf("%#v\n", u5)                  //main.User{id:1220, name:"HEYE", addr:"SHANGHAI", score:59}
}

[可见性]
	Go包的顶层代码中，首字母大写的标识符，跨package包可见(导出)，否则只能本包内可见
	导出的结构体，package内外皆可见，同时，导出的结构体中的成员(属性、方法)要在包外也可见，则也需首字母大写

[访问]
	u5 := User{1220, "HEYE", "SHANGHAI", 59} //无字段名称必须按照顺序给出全部字段值
	fmt.Printf("%#v\n", u5)                  //main.User{id:1220, name:"HEYE", addr:"SHANGHAI", score:59}
	fmt.Println(u5.id, u5.name, u5.addr)     //1220 HEYE SHANGHAI

[修改]
	u5 := User{1220, "HEYE", "SHANGHAI", 59} //无字段名称必须按照顺序给出全部字段值
	u5.name = "HY"
	u5.score = 150
	fmt.Printf("%#v\n", u5) //main.User{id:1220, name:"HY", addr:"SHANGHAI", score:150}

[成员方法]
package main

import (
	"fmt"
)

type User struct {
	id         int
	name, addr string //多个字段类型相同可以合并写
	score      float32
}

// 定义一个普通的函数
func getName(u User) string {
	return fmt.Sprintf("%d : %s", u.id, u.name)
}

// 定义方法
// User 表示和User类型有关
// u称为receiver 表示和User的具体某个实例有关
func (u User) getName() string {
	return fmt.Sprintf("%d : %s", u.id, u.name)
}

func main() {
	u5 := User{1220, "HEYE", "SHANGHAI", 59}
	//调用普通函数方法
	fmt.Println(getName(u5)) //1220 : HEYE
	//调用方法
	fmt.Println(u5.getName()) //1220 : HEYE
}

-------写一个方法 自定义MyInt 自增1---------
package main

import "fmt"

type MyInt int

func (i MyInt) inc() int {
	t := int(i)
	t++
	return t
}

func main() {
	var j MyInt = 99
	fmt.Printf("%T %[1]v\n", j) //main.MyInt 99

	value := j.inc()                //
	fmt.Printf("%T %[1]v\n", value) //int 100
}


【9-2-结构体指针和值类型】
=====================================================================
指针
	p1 类型是Point，指代Point的某一个具体的实例
	p2:=&p1 类型是*Point，p2记录一个大整数是地址，通过地址找到p1这个标识符指向的实例
	p1.x 通过实例访问属性
	p2.x 通过指针找到实例访问属性x，，在其他语言中p2 -> x

package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p1 := Point{4, 5}
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:4 y:5},0xc0000180a0

	p2 := &p1                           //*Point 指向 p1这个实例的指针
	fmt.Printf("%T %+[1]v,%[1]p\n", p2) //*main.Point &{x:4 y:5},0xc0000180a0

	p3 := new(Point)                      //构建一个Point零值的实例，返回该实例的指针给p3
	fmt.Printf("%T %+[1]v,%p\n", p3, &p3) //*main.Point &{x:0 y:0},0xc0000ca020

	//指针用途
	p1.x = 100
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:100 y:5},0xc0000180a0 修改地址不变
	p2.x = 200
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:200 y:5},0xc0000a6070
	fmt.Printf("%T %+[1]v,%[1]p\n", p2)   //*main.Point &{x:200 y:5},0xc0000a6070
}

[多次值copy]
package main

import "fmt"

type Point struct {
	x, y int
}

func testCopy(t Point) Point {
	fmt.Printf("t %T %+[1]v,%p\n", t, &t) //*main.Point &{x:4 y:5},0xc0000180a0
	return t
}

func main() {
	p1 := Point{4, 5}
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:4 y:5},0xc0000180a0

	p2 := p1                              //完全副本
	fmt.Printf("%T %+[1]v,%p\n", p2, &p2) //main.Point {x:4 y:5},0xc0000180f0

	p3 := &p1
	fmt.Printf("%T %+[1]v,%p\n", p3, p3) //*main.Point &{x:4 y:5},0xc0000180a0
	fmt.Println("-----------------------------------------------")

	p4 := testCopy(p1)
	fmt.Printf("p1 %T %+[1]v,%p\n", p1, &p1) 
	fmt.Printf("p4 %T %+[1]v,%p\n", p4, &p4) 
	//会有多次copy
	// t main.Point {x:4 y:5},0xc000018150
	// p1 main.Point {x:4 y:5},0xc0000180a0
	// p4 main.Point {x:4 y:5},0xc000018140
}

[减少值copy 使用指针传参]
package main

import "fmt"

type Point struct {
	x, y int
}

func testCopy(t Point) *Point {
	fmt.Printf("t %T %+[1]v,%p\n", t, &t) //*main.Point &{x:4 y:5},0xc0000180a0
	return &t
}

func main() {
	p1 := Point{4, 5}
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:4 y:5},0xc0000180a0

	p4 := testCopy(p1)
	fmt.Printf("p1 %T %+[1]v,%p\n", p1, &p1)
	fmt.Printf("p4 %T %+[1]v,%p\n", p4, p4)
	//所以返回指针减少copy 返回的值和p4是一个地址
	// t main.Point {x:4 y:5},0xc0000180f0
	// p1 main.Point {x:4 y:5},0xc0000180a0
	// p4 *main.Point &{x:4 y:5},0xc0000180f0
}

【9-3-构造器】
=====================================================================
Go语言并没有从语言层面为结构体提供什么构造器，但是有时候可以通过一个函数为结构体初始化提供属性值，从而方便得到以恶结构体实例。习惯上，函数命名为NewXxx的形式

	严格来说，Go语言结构体并没有构造函数这样的语法不成文的约定
	可以没有，不限制
	因为某些结构体数据成员很多，而且初始化数据成员很麻烦，甚至初始化初使用者不知道z的方法给使用者提供方便
	Newxxx成为了结构体xxx的构造函数了
	函数，普通的函数，不是方法(因为没有receiver)
	Newxxx函数，没有实例而创造出新的实例

	重载overload
		Go没有重载函数同名
		形参个数、类型不一样支持重载
			func fl(x,y int) f1(4，5)
			func fl(x int)   f1(4)

package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func NewAnimal(name string, age int) Animal {
	a := Animal{name, age}
	fmt.Printf("NewAnimal : %+v,%p\n", a, &a) //NewAnimal : {name:BIRKHOFF age:30},0xc000008090
	return a
}

func main() {
	a := NewAnimal("BIRKHOFF", 30)
	//发现2个地址不一样 是副本 可以使用指针减少副本
	fmt.Printf("Main : %+v,%p\n", a, &a) //Main : {name:BIRKHOFF age:30},0xc000008078
}

#上例中，NewAnimal的返回值使用了值拷贝，增加了内存开销，习惯上返回值会采用指针类型，避免实例的拷贝。		
func NewAnimal(name string, age int) *Animal {
	a := Animal{name, age}
	fmt.Printf("NewAnimal : %+v,%p\n", a, &a) //NewAnimal : {name:BIRKHOFF age:30},0xc000118060
	return &a
}
	
【9-4-父子结构体】
=====================================================================
[匿名结构体]
type定义新类型 
	type 类型名 类型本体
	type 类型名	类型 基于类型定义新类型
	type 类型名=类型 类型的别名
	
var 定义新变量，变量有类型
	var 变量名 类型
	
	//得到是一个结构体类型的变量
	var pd struct {
		x, y int
	}
*匿名结构体，只是为了快速方便地得到一个结构体实例，而不是使用结构体创建N个实例。

package main

import "fmt"

//定义新类型
type Point struct {
	x, y int
}

func main() {
	//类型没有给出名称 这种定义叫做 匿名结构体
	//得到是一个结构体类型的变量
	var pd struct {
		x, y int
	}
	var i int
	var p1 Point
	// var p2 pd 错误 2个变量错误
	fmt.Printf("%T\n", p1) //main.Point
	fmt.Printf("%T\n", i)  //int
	fmt.Printf("%T\n", pd) //struct { x int; y int } 没有名字

	pd.x = 20
	pd.y = 30
	fmt.Printf("%#v\n", pd) //struct { x int; y int } 没有名字

	var pd2 = struct { //直接使用 要用 =
		x, y int
	}{8, 8}
	fmt.Printf("%#v\n", pd2) //struct { x int; y int }{x:8, y:8}
}


[匿名成员]
	成员属性，没有名字
	不推荐使用
	type Point struct {
		x int
		int //只有类型没有属性名的成员，匿名成员
		string
		bool
	}

package main

import "fmt"

//匿名成员
type Point struct {
	x   int
	int //只有类型没有属性名的成员，匿名成员
	string
	bool
}

func main() {
	p1 := new(Point)
	fmt.Printf("%T %+[1]v\n", p1)                 //*main.Point &{x:0 int:0 string: bool:false}
	fmt.Println(p1.x, p1.int, p1.string, p1.bool) //0 0  false
	p1.string = "xks"
	p1.bool = true
	fmt.Printf("%T %+[1]v\n", p1) //*main.Point &{x:0 int:0 string:xks bool:true}
}

[父子结构体]
	使用结构体嵌套实现类似面向对象父类子类继承(派生)的效果
	子结构体使用匿名成员能简化调用父结构体成员

EG:
动物类包括猫类，猫属于猫类，猫也属于动物类，某动物一定是动物类，但不能说某动物一定是猫类将上例中的Animal结构体
使用匿名成员的方式，嵌入到Cat结构体中，看看效果

package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type Cat struct { //子结构体 直接拥有了父结构体的属性
	//name string
	//age int
	Animal //匿名成员，父结构体，Go提供语法糖，结构体嵌套
	color  string
}

func main() {
	var c = new(Cat)             //c是Cat的实例
	fmt.Printf("%T %+[1]v\n", c) //*main.Cat &{Animal:{name: age:0} color:}
	c.name = "XKS"               //=c.Animal.name
	c.age = 30                   //=c.Animal.age
	c.color = "blue"             //=c.Animal.color
	fmt.Printf("%T %+[1]v\n", c) //*main.Cat &{Animal:{name:XKS age:30} color:blue}
}
*子结构体，为什么定义一个子类cat?父类并没有包含cat的所有特征，也没有包括Dog类所有特征



【9-5-receiver和深浅拷贝】
=====================================================================
[指针类型receiver]
Go语言中，可以为任意类型包括结构体增加方法，形式是 func Receiver 方法名 签名{函数体}
	这个receiver类似其他语言中的this或self.
	receiver必须是一个类型T实例或者类型T的指针，T不能是指针或接口。

[getter、setter]
package main

import "fmt"

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point { //构造函数不是方法 是一个普通函数
	return &Point{x, y}
}

/*getter*/
// 2种 receiver 实例或指针
//相当于实例.GetX() 指针.GetX() 调用该方法
func (p Point) GetX() int { //通过方法控制读或写 只是这里的receiver是实例
	//receiver是实例，通过实例调用有值拷贝，通过指针调用同样会值拷贝
	fmt.Printf("instance.GetX(), %p,%+v\n", &p, p)
	return p.x
}

func (p *Point) GetY() int { //通过方法控制读或写 只是这里的receiver是实例
	//receiver是指针，通过实例访问会取实例的指针，或通过指针访问
	fmt.Printf("pointer.GetY(), %p,%+v\n", p, p)
	return p.y
}

/*setter*/
// 2种 receiver 实例或指针
func (p Point) SetX(v int) {
	fmt.Printf("instance.SetX(), %p,%+v\n", &p, p)
	p.x = v
	fmt.Printf("instance.SetX(), %p,%+v\n", &p, p)

}

func (p *Point) SetY(v int) {
	fmt.Printf("pointer.SetY(), %p,%+v\n", p, p)
	p.y = v
	fmt.Printf("pointer.SetY(), %p,%+v\n", p, p)
}

func main() {
	// var p1 = NewPoint(4, 5) //p1 是指针
	var p1 = Point{4, 5}

	fmt.Printf("%T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y) //4
	fmt.Println("-------------------------")                  //4
	// fmt.Println(p1.GetX())                              //4
	fmt.Println("-------------------------") //4

	fmt.Println(p1.GetX(), (&p1).GetX())
	// instance.GetX(), 0xc0000180f0,{x:4 y:5} 2个值不同发生了值拷贝
	// instance.GetX(), 0xc000018120,{x:4 y:5}
	// 4 4
	fmt.Println("-------------------------") //4
	fmt.Println(p1.GetY(), (&p1).GetY())
	// pointer.GetY(), 0xc0000180a0,&{x:4 y:5} 2个地址相同，同一个地址
	// pointer.GetY(), 0xc0000180a0,&{x:4 y:5}
	// 5 5
	fmt.Println("-------------------------") //4

	//Setter
	p1.SetX(11)
	fmt.Printf("SetX() %T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y)
	(&p1).SetX(11)
	fmt.Printf("SetX() %T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y)
	//发现 通过receiver是实例 没有修改原来 因为修改了副本的数据 原本的数据没有修改
	// instance.SetX(), 0xc000018190,{x:4 y:5}
	// instance.SetX(), 0xc000018190,{x:11 y:5}
	// SetX() main.Point {x:4 y:5} 0xc0000180a0 , 4 5
	// instance.SetX(), 0xc000018200,{x:4 y:5}
	// instance.SetX(), 0xc000018200,{x:11 y:5}
	// SetX() main.Point {x:4 y:5} 0xc0000180a0 , 4 5
	fmt.Println("-------------------------") //4
	p1.SetY(22)
	fmt.Printf("SetY() %T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y)
	(&p1).SetY(222)
	fmt.Printf("SetY() %T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y)
	//receiver 作为指针返回 是同一个
	// pointer.SetY(), 0xc0000180a0,&{x:4 y:5}
	// pointer.SetY(), 0xc0000180a0,&{x:4 y:22}
	// SetY() main.Point {x:4 y:22} 0xc0000180a0 , 4 22

	// pointer.SetY(), 0xc0000180a0,&{x:4 y:22}
	// pointer.SetY(), 0xc0000180a0,&{x:4 y:222}
	// SetY() main.Point {x:4 y:222} 0xc0000180a0 , 4 222
}


[总结]
接收器receiver可以是类型T 也可以是 指针*T，定义的方法有什么区别？从上面可以看出 
	receiver是实例，方法内操作的是实例的副本，相当于实例的拷贝
		实例访问，方法内使用时该实例的副本
		指针访问，方法内使用的是指针指向的实例的副本
	reciver是指针，方法内操作的是实例的指针的副本，相当于实例的指针的拷贝
		实例访问，方法内使用的是该实例的指针
		指针访问，方法内使用的是该指针

**非常明显，如果是非指针接收器方法调用有值拷贝，操作的是副本，而指针接收器方法调用操作的是同一个内存的同一个实例。
	如果是操作大内存对象时，且操作同一个实例时，一定要采用指针接收器的方法。


[深浅拷贝]
shadow copy
	影子拷贝，也叫浅拷贝。遇到引用类型数据，仅仅复制一个引用而已
deep copy
	深拷贝，往往会递归复制一定深度

	注意，深浅拷贝说的是拷贝过程中是否发生递归拷贝，也就是说如果某个值是一个地址，是只复制这个地址，还是复制地址指向的内容。
	值拷贝是深拷贝，地址拷贝是浅拷贝，这种说法是错误的。因为地址拷贝只是拷贝了地址，因此本质上来讲也是值拷贝。

	Go语言中，引用类型实际上拷贝的是标头值，这也是值拷贝，并没有通过标头值中对底层数据结构的指针指向的内容进行复制，这就是浅拷贝。
	非引用类型的复制就是值拷贝，也就是再造一个副本，这也是浅拷贝。因为你不能说对一个整数值在内存中复制出一个副本，就是深的拷贝。
	像整数类型这样的基本类型就是一个单独的值，没法深入拷贝，根本没法去讲深入的事儿。
	简单讲，大家可以用拷贝文件是否对软链接跟进来理解。直接复制软链接就是浅拷贝，钻进软链接里面复制其内容就是深拷贝。
	*复杂数据结构，往往会有嵌套，有时嵌套很深，如果都采用深拷贝，那代价很高，
		所以，浅拷贝才是语言普遍采用的方案。


【9-6-接口概念】
=====================================================================
接口interface，和Java类似，是一组行为规范的集合，就是定义一组未实现的函数声明。谁使用接口就是参照接口的方法定义实现它们
type 接口名 interface{
		方法1 (参数列表1) 返回值列表1
		方法2 (参数列表2) 返回值列表2
		...
}
接口命名习惯在接口名后面加上er后缀
参数列表、返回值列表参数名可以不写
如果要在包外使用接口，接口名应该首字母大写，方法要在包外使用，方法名首字母也要大写
接口中的方法应该设计合理，不要太多
---
接口名，习惯以er结尾
一套方法的声明的组合，都不实现，谁用谁实现
	函数签名，形参，返回值都可以只有类型
	实现接口
		*** 对于某一个接口的方法全部实现，一个不落

可见性
Go语言种，建议，实现小接口
	接口的方法尽量的少
	用组合来实现大接口

约束方法
	方法是动作、操作、定义在结构体
	约束、结构体的方法成员
	结构体实现了某接口，遵守该接口的约定
		实现了该接口所有的方法
		
Go采用 非侵入设计 相对Java来说的
	Java
		class Xxx implements 接口1，接口2
	Go悄悄的实现

EG：
package main

import "fmt"

/*	如果一个结构体实现了一个接口声明的所有方法，就说结构体实现了该接口。
	一个结构体可以实现多个不同接口。
*/

//Person实现了Sporter接口,下面函数缺一不可
type Sporter interface {
	run()
	jump()
}

type Person struct {
	name string
	age  int
}

func (p *Person) run() {
	fmt.Println("run -----------")
}
func (p *Person) jump() {
	fmt.Println("jump -----------")
}
func (p *Person) swim() {
	fmt.Println("swimming ~~~~~~~~~")
}

/*Users 实现了Sporter接口*/
type User struct {
	id    int
	name  string
	score float32
}

func (p *User) run() {
	fmt.Println("User run -----------")
}
func (p *User) jump() {
	fmt.Println("User jump -----------")
}
func (p *User) walk() {
	fmt.Println("User swimming ~~~~~~~~~")
}

func main() {
	p1 := new(Person)
	p1.run()
	p1.jump()

	var s Sporter = p1 //var s Sporter = Sporter(p1)
	//p1 既是Person类型 也是 Sporter类型
	fmt.Printf("p1 = %T,s = %T\n", p1, s) //p1 = *main.Person,s = *main.Person
	s.jump()
	s.run()
	// s.swim() 不能调用这个方法 因为接口里没有这个方法
	// run -----------
	// jump -----------

	//s 既是User类型 也是 Sporter类型
	u := new(User)
	s = u                                //var s Sporter = Sporter(p1)
	fmt.Printf("p1 = %T,s = %T\n", s, u) //p1 = *main.User,s = *main.User
	s.jump()
	s.run()
	// User jump -----------
	// User run -----------
}
 
【9-7-接口嵌套】
=====================================================================
type Reader interface{
	Read(p []byte) (n int,err errot)
}

type Closer interface{
	Close() error
}

type ReadCloser interface{
	Reader
	Closer
}
ReadCloser接口是Reader、Closer接口组合而成，也就是说它拥有Read，Close方法声明

type runner interface{
	run()
}
type jumper interface{
	jump()
}
type Sporter interface{
	runner
	jumper
}

*  一般将接口分开定义 然后组合成一个大接口

【9-8-空接口和接口类型断言】
=====================================================================
[空接口]
	空接口，实际上是空接口类型，写作
		interface {}。
	为了方便使用，Go语言为它定义一个别名any类型
		type any =interface{}
	空接口，没有任何方法声明，因此，任何类型都无需显式实现空接口的方法，因为任何类型都满足这个空接口的要求。
	任何类型的值都可以看做是空接口类型。
、
EG：
package main

import "fmt"

func main() {
	var a = 500
	var b interface{}
	b = a
	fmt.Printf("%v,%[1]T ; %v,%[2]T\n", a, b) //500,int ; 500,int

	var c = "abcd"
	b = c
	fmt.Printf("%v,%[1]T ; %v,%[2]T\n", b, c) //abcd,string ; abcd,string

	//	b = []any{100, "xyz", [3]int{1, 2, 3}}
	//  b = []interface{}{100, "xyz", [3]int{1, 2, 3}}
	// type any = interface{}
	b = []any{100, "xyz", [3]int{1, 2, 3}}
	fmt.Printf("%v,%[1]T\n", b) //[100 xyz [1 2 3]],[]interface {}
}

[接口类型断言]
接口类型断言(Type Assertions)可以将接口转换成另外一种接口，也可以将接口转换成另外的类型。接口类型断言格式 
t := i.(T)
	i代表接口变量
	T表示转换目标类型
	t代表转换后的变量
		断言失败，也就是说i没有实现T接口的方法则panic
	t，ok:=i.(T)，则断言失败不panic，通过ok是true或false判断i是否是T类型接口

EG：
package main

import "fmt"

func main() {
	var b interface{} = 500
	// fmt.Println(b.(string)) //panic转换失败
	//panic: interface conversion: interface {} is int, not string
	if s, ok := b.(string); ok {
		fmt.Println("断言成功，值是", s)
	} else {
		fmt.Println("断言失败") //断言失败
	}
}

[type-switch]
package main

import "fmt"

func main() {
	var i interface{} = 500
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("整形", v) //
	default:
		fmt.Println("其他类型")
	}
}

[输出格式接口]
我们使用fmt.Print等函数时，对任意一个值都有一个缺省打印格式。本质上就是实现打印相关的接口。

package main

import "fmt"

type Person struct {
	name string
}

/*fmt包内部会对两个接口做处理
也就是你自己定义了String、GoString 就按照你自己定义方法来调用
否则就默认打印
*/

/*String Interface*/
//1.
// func (Person) String() string { //影响 缺省格式%v和详细格式%+v
// 	return "aaa"
// }
//2.指针作为receiver 只对指针有效 指针输出自定义格式
func (*Person) String() string { //影响 缺省格式%v和详细格式%+v
	return "aaa"
}

/*GoString Interface*/
//1.
// func (Person) GoString() string { //影响 %#v
// 	return "xks"
// }
//2.指针作为receiver 只对指针有效 指针输出自定义格式
func (*Person) GoString() string { //影响 %#v
	return "xks"
}
//普通函数
func (*Person) foo() string {
	return "foo"
}

func main() {
	p1 := Person{"Tom"}
	fmt.Println("1:", p1, &p1)          //{Tom} &{Tom} 缺省打印格式
	fmt.Printf("2: %+v,%+v\n", p1, &p1) //{name:Tom},&{name:Tom}
	fmt.Printf("3: %#v,%#v\n", p1, &p1) //main.Person{name:"Tom"},&main.Person{name:"Tom"}
	fmt.Println(p1.foo(), (&p1).foo())  //foo foo

	// 如果不定义 String GoString 使用系统内建方式
	// 1: {Tom} &{Tom}
	// 2: {name:Tom},&{name:Tom}
	// 3: main.Person{name:"Tom"},&main.Person{name:"Tom"}

	// 1.如果定义 String GoString 使用系统内建方式
	// 1: aaa aaa
	// 2: aaa,aaa
	// 3: xks,xks

	//2.如果指针作为receiver 只对指针有效 指针输出自定义格式
	// 1: {Tom} aaa
	// 2: {name:Tom},aaa
	// 3: main.Person{name:"Tom"},xks

	//可以使用类型断言进行判断
	var i interface{}
	i = p1
	//i.(fmt.Stringer)
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("整形", v)
	case Person:
		fmt.Println("Person type:", v)
	case fmt.Stringer: //接口类型
		fmt.Println("String Interface type:", v) //Person type: aaa
	default:
		fmt.Println("其他类型")
	}
}
[总结]
Stringer、GoStringer接口的方法，如果receiver如果是指针，只能对指针有作用;如果receiver如果是实例，实例、指针都有作用。 
普通方法receiver不管是实例还是指针，实例、指针都可以调用该方法时

【9-9-自定义错误和断电调试】
=====================================================================
异常处理
	type error {
		Error() string
	}

[实现源码 自定义error]-系统内置源码error.new(type)

package main

import "fmt"

//包内部使用
type errorString struct {
	s string
}

//首先是errorString的方法，同时也符合error接口的定义
func (e *errorString) Error() string {
	return e.s
}

//构造函数
// func New(text string) *errorString {
// 	return &errorString{s: text}
// }
func New(reason string) error { //error是内建函数
	return &errorString{s: reason} //既是errorString类型，也是error接口类型
}

func (e *errorString) foo() {
	fmt.Println("errorString结构体的foo方法")
}

func main() {
	e1 := New("Error:[1]")  //var e1 error = New("xxx")
	fmt.Printf("%T\n", e1)  //*main.errorString
	fmt.Println(e1.Error()) //Error:[1]
	if v, ok := e1.(*errorString); ok {
		v.foo() //errorString结构体的foo方法
	}

	// e1,e2 区别
	e2 := errorString{"Error:[2]"} // var e2 errorString
	fmt.Println(e2.Error())        //Error:[2]
}

[自定义Error]
package main

import (
	"errors"
	"fmt"
)

var ErrDisvisionByZero = errors.New("除零异常") //var ErrDisvisionByZero error类型

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDisvisionByZero
	}
	return a / b, nil
}

func main() {
	if v, err := div(5, 0); err == nil { //nil 无错误
		fmt.Println("v =", v)
	} else { //有错误
		// fmt.Println("err = ", err)
		// fmt.Println("有错误", err.Error(), v)
		fmt.Println("有错误", err) //fmt.Print* 内部对error特殊处理

	}
	//div(5, 2) => v = 2
	//div(5, 0) => 有错误 除零异常 0

	// fmt.Println(div(5, 2)) //2 <nil>

}

【9-10-panic和recover】
=====================================================================

panic是不好的，因为它发生时，往往会造成程序崩溃、服务终止等后果，所以没人希望它发生。但是如果在错误发生时，不及时panic而终止程序运行
	继续运行程序恐怕造成更大的损失，付出更加惨痛的代价。所以，有时候，panic导致的程序崩溃实际上可以及时止损，只能两害相权取其轻。
panic虽然不好，体验很差，但也是万不得已，可以马上暴露问题，及时发现和纠正问题。

panic产生
	runtime运行时错误导致抛出panic，比如数组越界、除零
	主动手动调用panic(reason)，这个reason可以是任意类型

panic执行
	逆序执行当前已经注册过的goroutine的defer链(recover从这里介入)
	打印错误信息和调用堆栈
	调用exit(2)结束整个进程

抛出异常的方法
	系统运行时抛出panic
	手动panic
如果panic没有处理，将向外排除，直到把程序


EG：
package main

import (
	"errors"
	"fmt"
)

func div(a, b int) int {
	defer fmt.Println("start")
	defer fmt.Println(a, b)
	r := a / b //这一行有可能panic
	fmt.Println("end") //panic了不会再执行了
	return r
}

func main() {
	fmt.Println(div(5, 0))
}

// 5 0
// start
// panic: runtime error: integer divide by zero
 
// goroutine 1 [running]: 下面是调用栈，div压着main
// main.div(0x5, 0x0)
// 	e:/goprojects/main.go:13 +0x2cc 出错的行号
// main.main()
// 	e:/goprojects/main.go:19 +0x25

[recover]
package main

import (
	"errors"
	"fmt"
	"runtime"
)

var ErrDisvisionByZero = errors.New("除零异常") //var ErrDisvisionByZero error类型

func div(a, b int) int {
	defer func() {
		err := recover()
		fmt.Printf("1 %+v, %[1]T\n", err)
	}()
	defer fmt.Println("start")
	defer fmt.Println(a, b)
	defer func() {
		err := recover() //一旦recover了，就相当于处理过错误
		fmt.Printf("2 %+v, %[1]T\n", err)
		switch v := err.(type) { //类型断言
		case runtime.Error:
			//在源码种 runtime/error.go No.75 行，还为errorString也实现了RuntimeError方法
			//也就是说，标准库runtime运行时错误，都是runtime.Error接口的
			fmt.Printf("原因：%T, %#[1]v\n", v)
		case []int:
			fmt.Println("原因：切片", v)
		}
		fmt.Println("离开recover处理")
	}()

	r := a / b //这一行有可能panic

	panic([]int{1, 3, 5})
	fmt.Println("end")
	return r
}

func main() {
	fmt.Println(div(5, 2), "!!!")
	fmt.Println("main exit")
}

/*1*/
//	fmt.Println(div(5, 0), "!!!")
// 2 runtime error: integer divide by zero, runtime.errorString
// 原因：runtime.errorString, "integer divide by zero"
// 离开recover处理
// 5 0
// start
// 1 <nil>, <nil>
// 0 !!!
// main exit

/*2*/
//	fmt.Println(div(5, 2), "!!!")
// 2 [1 3 5], []int
// 原因：切片 [1 3 5]
// 离开recover处理
// 5 2
// start
// 1 <nil>, <nil>
// 0 !!!
// main exit

上例中，一旦在某函数中panic，当前函数panic之后的语句将不再执行，开始执行defer。如果在defer中错误被recover后
	就相当于当前函数产生的错误得到了处理。当前函数执行完defer，当前函数退出执行，程序还可以从当前函数之后继续执行。

可以观察到panic和recover有如下
	有panic，一路向外抛出，但没有一处进行recover，也就是说没有地方处理错误，程序崩溃有painc，有recover来捕获，相当于错误被处理掉了
	当前函数defer执行完后，退出当前函数从当前函数之后继续执行
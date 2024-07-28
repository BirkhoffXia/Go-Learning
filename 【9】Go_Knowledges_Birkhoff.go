Viedo Date: 2023/10/14
Made By:	BIRKHOFF 
Date:	2024-06-24   

��9-1-�ṹ�弰��Ա��
=====================================================================
Go���ԵĽṹ���е���������������еġ��ࡱ��������ȫ�ǣ�Go����Ҳû��������ʵ���������ʽ��

�ṹ�������չ��������Ϊ�ṹ��������
	����method�������Ͼ���function
	User.getName //���͵ĵ��� ��������� ������
	u1.getName //User����
	
	
		
Go�У�����Ϊ�µ�������չ����
	�µģ����Ի���ԭ�������������µ����ͣ���ò�Ҫ��ԭ��������ֱ���޸�

[����]
ʹ��type����ṹ�壬���԰ѽṹ�忴������ʹ�á�����ָ���ṹ����ֶ�(����)���ƺ�����
type User struct {
	id int
	name,addr string //����ֶ�������ͬ���Ժϲ�д
	score float32
}

User�����Ǹ���ʶ������һ��ָ������
���������Ͷ�����struct{}�Ĳ���

[��ʼ��]
package main

import (
	"fmt"
)

type User struct {
	id         int
	name, addr string //����ֶ�������ͬ���Ժϲ�д
	score      float32
}

func main() {
	// ��ʼ��
	//1 ��ֵ �ܳ��ã�������
	var u1 User // u1��User��ʵ��
	//%v��ӡ�ṹ��
	fmt.Printf("%T %[1]v", u1) //main.User {0   0}

	u1.id = 1
	fmt.Printf("%T %[1]v", u1) //main.User {1   0}
	fmt.Println("------------------------------------\n")
	fmt.Println(u1)         //{1   0}
	fmt.Printf("%v\n", u1)  //{1   0}                                    Ĭ��ȱʡ��ʽ��ӡ
	fmt.Printf("%+v\n", u1) //{id:1 name: addr: score:0}                 ��ǿ��ӡ
	fmt.Printf("%#v\n", u1) //main.User{id:1, name:"", addr:"", score:0} ��ϸ��ӡ

	//2 ��������ʼ�����Ƽ�
	u2 := User{}            //�ֶ�Ϊ��ֵ
	fmt.Printf("%#v\n", u2) //main.User{id:0, name:"", addr:"", score:0}

	//3 ��������ʼ����field:valueΪ�ֶθ�ֵ
	u3 := User{id: 100}
	fmt.Printf("%#v\n", u3) //main.User{id:100, name:"", addr:"", score:0}

	u4 := User{
		id:    319,
		score: 99,
		addr:  "Shanghai",
		name:  "BIRKHOFF",
	} //���ƶ�Ӧ����ν˳��
	fmt.Printf("%#v\n", u4) //main.User{id:319, name:"BIRKHOFF", addr:"Shanghai", score:99}

	u5 := User{1220, "HEYE", "SHANGHAI", 59} //���ֶ����Ʊ��밴��˳�����ȫ���ֶ�ֵ
	fmt.Printf("%#v\n", u5)                  //main.User{id:1220, name:"HEYE", addr:"SHANGHAI", score:59}
}

[�ɼ���]
	Go���Ķ�������У�����ĸ��д�ı�ʶ������package���ɼ�(����)������ֻ�ܱ����ڿɼ�
	�����Ľṹ�壬package����Կɼ���ͬʱ�������Ľṹ���еĳ�Ա(���ԡ�����)Ҫ�ڰ���Ҳ�ɼ�����Ҳ������ĸ��д

[����]
	u5 := User{1220, "HEYE", "SHANGHAI", 59} //���ֶ����Ʊ��밴��˳�����ȫ���ֶ�ֵ
	fmt.Printf("%#v\n", u5)                  //main.User{id:1220, name:"HEYE", addr:"SHANGHAI", score:59}
	fmt.Println(u5.id, u5.name, u5.addr)     //1220 HEYE SHANGHAI

[�޸�]
	u5 := User{1220, "HEYE", "SHANGHAI", 59} //���ֶ����Ʊ��밴��˳�����ȫ���ֶ�ֵ
	u5.name = "HY"
	u5.score = 150
	fmt.Printf("%#v\n", u5) //main.User{id:1220, name:"HY", addr:"SHANGHAI", score:150}

[��Ա����]
package main

import (
	"fmt"
)

type User struct {
	id         int
	name, addr string //����ֶ�������ͬ���Ժϲ�д
	score      float32
}

// ����һ����ͨ�ĺ���
func getName(u User) string {
	return fmt.Sprintf("%d : %s", u.id, u.name)
}

// ���巽��
// User ��ʾ��User�����й�
// u��Ϊreceiver ��ʾ��User�ľ���ĳ��ʵ���й�
func (u User) getName() string {
	return fmt.Sprintf("%d : %s", u.id, u.name)
}

func main() {
	u5 := User{1220, "HEYE", "SHANGHAI", 59}
	//������ͨ��������
	fmt.Println(getName(u5)) //1220 : HEYE
	//���÷���
	fmt.Println(u5.getName()) //1220 : HEYE
}

-------дһ������ �Զ���MyInt ����1---------
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


��9-2-�ṹ��ָ���ֵ���͡�
=====================================================================
ָ��
	p1 ������Point��ָ��Point��ĳһ�������ʵ��
	p2:=&p1 ������*Point��p2��¼һ���������ǵ�ַ��ͨ����ַ�ҵ�p1�����ʶ��ָ���ʵ��
	p1.x ͨ��ʵ����������
	p2.x ͨ��ָ���ҵ�ʵ����������x����������������p2 -> x

package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p1 := Point{4, 5}
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:4 y:5},0xc0000180a0

	p2 := &p1                           //*Point ָ�� p1���ʵ����ָ��
	fmt.Printf("%T %+[1]v,%[1]p\n", p2) //*main.Point &{x:4 y:5},0xc0000180a0

	p3 := new(Point)                      //����һ��Point��ֵ��ʵ�������ظ�ʵ����ָ���p3
	fmt.Printf("%T %+[1]v,%p\n", p3, &p3) //*main.Point &{x:0 y:0},0xc0000ca020

	//ָ����;
	p1.x = 100
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:100 y:5},0xc0000180a0 �޸ĵ�ַ����
	p2.x = 200
	fmt.Printf("%T %+[1]v,%p\n", p1, &p1) //main.Point {x:200 y:5},0xc0000a6070
	fmt.Printf("%T %+[1]v,%[1]p\n", p2)   //*main.Point &{x:200 y:5},0xc0000a6070
}

[���ֵcopy]
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

	p2 := p1                              //��ȫ����
	fmt.Printf("%T %+[1]v,%p\n", p2, &p2) //main.Point {x:4 y:5},0xc0000180f0

	p3 := &p1
	fmt.Printf("%T %+[1]v,%p\n", p3, p3) //*main.Point &{x:4 y:5},0xc0000180a0
	fmt.Println("-----------------------------------------------")

	p4 := testCopy(p1)
	fmt.Printf("p1 %T %+[1]v,%p\n", p1, &p1) 
	fmt.Printf("p4 %T %+[1]v,%p\n", p4, &p4) 
	//���ж��copy
	// t main.Point {x:4 y:5},0xc000018150
	// p1 main.Point {x:4 y:5},0xc0000180a0
	// p4 main.Point {x:4 y:5},0xc000018140
}

[����ֵcopy ʹ��ָ�봫��]
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
	//���Է���ָ�����copy ���ص�ֵ��p4��һ����ַ
	// t main.Point {x:4 y:5},0xc0000180f0
	// p1 main.Point {x:4 y:5},0xc0000180a0
	// p4 *main.Point &{x:4 y:5},0xc0000180f0
}

��9-3-��������
=====================================================================
Go���Բ�û�д����Բ���Ϊ�ṹ���ṩʲô��������������ʱ�����ͨ��һ������Ϊ�ṹ���ʼ���ṩ����ֵ���Ӷ�����õ��Զ�ṹ��ʵ����ϰ���ϣ���������ΪNewXxx����ʽ

	�ϸ���˵��Go���Խṹ�岢û�й��캯���������﷨�����ĵ�Լ��
	����û�У�������
	��ΪĳЩ�ṹ�����ݳ�Ա�ܶ࣬���ҳ�ʼ�����ݳ�Ա���鷳��������ʼ����ʹ���߲�֪��z�ķ�����ʹ�����ṩ����
	Newxxx��Ϊ�˽ṹ��xxx�Ĺ��캯����
	��������ͨ�ĺ��������Ƿ���(��Ϊû��receiver)
	Newxxx������û��ʵ����������µ�ʵ��

	����overload
		Goû�����غ���ͬ��
		�βθ��������Ͳ�һ��֧������
			func fl(x,y int) f1(4��5)
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
	//����2����ַ��һ�� �Ǹ��� ����ʹ��ָ����ٸ���
	fmt.Printf("Main : %+v,%p\n", a, &a) //Main : {name:BIRKHOFF age:30},0xc000008078
}

#�����У�NewAnimal�ķ���ֵʹ����ֵ�������������ڴ濪����ϰ���Ϸ���ֵ�����ָ�����ͣ�����ʵ���Ŀ�����		
func NewAnimal(name string, age int) *Animal {
	a := Animal{name, age}
	fmt.Printf("NewAnimal : %+v,%p\n", a, &a) //NewAnimal : {name:BIRKHOFF age:30},0xc000118060
	return &a
}
	
��9-4-���ӽṹ�塿
=====================================================================
[�����ṹ��]
type���������� 
	type ������ ���ͱ���
	type ������	���� �������Ͷ���������
	type ������=���� ���͵ı���
	
var �����±���������������
	var ������ ����
	
	//�õ���һ���ṹ�����͵ı���
	var pd struct {
		x, y int
	}
*�����ṹ�壬ֻ��Ϊ�˿��ٷ���صõ�һ���ṹ��ʵ����������ʹ�ýṹ�崴��N��ʵ����

package main

import "fmt"

//����������
type Point struct {
	x, y int
}

func main() {
	//����û�и������� ���ֶ������ �����ṹ��
	//�õ���һ���ṹ�����͵ı���
	var pd struct {
		x, y int
	}
	var i int
	var p1 Point
	// var p2 pd ���� 2����������
	fmt.Printf("%T\n", p1) //main.Point
	fmt.Printf("%T\n", i)  //int
	fmt.Printf("%T\n", pd) //struct { x int; y int } û������

	pd.x = 20
	pd.y = 30
	fmt.Printf("%#v\n", pd) //struct { x int; y int } û������

	var pd2 = struct { //ֱ��ʹ�� Ҫ�� =
		x, y int
	}{8, 8}
	fmt.Printf("%#v\n", pd2) //struct { x int; y int }{x:8, y:8}
}


[������Ա]
	��Ա���ԣ�û������
	���Ƽ�ʹ��
	type Point struct {
		x int
		int //ֻ������û���������ĳ�Ա��������Ա
		string
		bool
	}

package main

import "fmt"

//������Ա
type Point struct {
	x   int
	int //ֻ������û���������ĳ�Ա��������Ա
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

[���ӽṹ��]
	ʹ�ýṹ��Ƕ��ʵ�������������������̳�(����)��Ч��
	�ӽṹ��ʹ��������Ա�ܼ򻯵��ø��ṹ���Ա

EG:
���������è�࣬è����è�࣬èҲ���ڶ����࣬ĳ����һ���Ƕ����࣬������˵ĳ����һ����è�ཫ�����е�Animal�ṹ��
ʹ��������Ա�ķ�ʽ��Ƕ�뵽Cat�ṹ���У�����Ч��

package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type Cat struct { //�ӽṹ�� ֱ��ӵ���˸��ṹ�������
	//name string
	//age int
	Animal //������Ա�����ṹ�壬Go�ṩ�﷨�ǣ��ṹ��Ƕ��
	color  string
}

func main() {
	var c = new(Cat)             //c��Cat��ʵ��
	fmt.Printf("%T %+[1]v\n", c) //*main.Cat &{Animal:{name: age:0} color:}
	c.name = "XKS"               //=c.Animal.name
	c.age = 30                   //=c.Animal.age
	c.color = "blue"             //=c.Animal.color
	fmt.Printf("%T %+[1]v\n", c) //*main.Cat &{Animal:{name:XKS age:30} color:blue}
}
*�ӽṹ�壬Ϊʲô����һ������cat?���ಢû�а���cat������������Ҳû�а���Dog����������



��9-5-receiver����ǳ������
=====================================================================
[ָ������receiver]
Go�����У�����Ϊ�������Ͱ����ṹ�����ӷ�������ʽ�� func Receiver ������ ǩ��{������}
	���receiver�������������е�this��self.
	receiver������һ������Tʵ����������T��ָ�룬T������ָ���ӿڡ�

[getter��setter]
package main

import "fmt"

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point { //���캯�����Ƿ��� ��һ����ͨ����
	return &Point{x, y}
}

/*getter*/
// 2�� receiver ʵ����ָ��
//�൱��ʵ��.GetX() ָ��.GetX() ���ø÷���
func (p Point) GetX() int { //ͨ���������ƶ���д ֻ�������receiver��ʵ��
	//receiver��ʵ����ͨ��ʵ��������ֵ������ͨ��ָ�����ͬ����ֵ����
	fmt.Printf("instance.GetX(), %p,%+v\n", &p, p)
	return p.x
}

func (p *Point) GetY() int { //ͨ���������ƶ���д ֻ�������receiver��ʵ��
	//receiver��ָ�룬ͨ��ʵ�����ʻ�ȡʵ����ָ�룬��ͨ��ָ�����
	fmt.Printf("pointer.GetY(), %p,%+v\n", p, p)
	return p.y
}

/*setter*/
// 2�� receiver ʵ����ָ��
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
	// var p1 = NewPoint(4, 5) //p1 ��ָ��
	var p1 = Point{4, 5}

	fmt.Printf("%T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y) //4
	fmt.Println("-------------------------")                  //4
	// fmt.Println(p1.GetX())                              //4
	fmt.Println("-------------------------") //4

	fmt.Println(p1.GetX(), (&p1).GetX())
	// instance.GetX(), 0xc0000180f0,{x:4 y:5} 2��ֵ��ͬ������ֵ����
	// instance.GetX(), 0xc000018120,{x:4 y:5}
	// 4 4
	fmt.Println("-------------------------") //4
	fmt.Println(p1.GetY(), (&p1).GetY())
	// pointer.GetY(), 0xc0000180a0,&{x:4 y:5} 2����ַ��ͬ��ͬһ����ַ
	// pointer.GetY(), 0xc0000180a0,&{x:4 y:5}
	// 5 5
	fmt.Println("-------------------------") //4

	//Setter
	p1.SetX(11)
	fmt.Printf("SetX() %T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y)
	(&p1).SetX(11)
	fmt.Printf("SetX() %T %+[1]v %p , %d %d\n", p1, &p1, p1.x, p1.y)
	//���� ͨ��receiver��ʵ�� û���޸�ԭ�� ��Ϊ�޸��˸��������� ԭ��������û���޸�
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
	//receiver ��Ϊָ�뷵�� ��ͬһ��
	// pointer.SetY(), 0xc0000180a0,&{x:4 y:5}
	// pointer.SetY(), 0xc0000180a0,&{x:4 y:22}
	// SetY() main.Point {x:4 y:22} 0xc0000180a0 , 4 22

	// pointer.SetY(), 0xc0000180a0,&{x:4 y:22}
	// pointer.SetY(), 0xc0000180a0,&{x:4 y:222}
	// SetY() main.Point {x:4 y:222} 0xc0000180a0 , 4 222
}


[�ܽ�]
������receiver����������T Ҳ������ ָ��*T������ķ�����ʲô���𣿴�������Կ��� 
	receiver��ʵ���������ڲ�������ʵ���ĸ������൱��ʵ���Ŀ���
		ʵ�����ʣ�������ʹ��ʱ��ʵ���ĸ���
		ָ����ʣ�������ʹ�õ���ָ��ָ���ʵ���ĸ���
	reciver��ָ�룬�����ڲ�������ʵ����ָ��ĸ������൱��ʵ����ָ��Ŀ���
		ʵ�����ʣ�������ʹ�õ��Ǹ�ʵ����ָ��
		ָ����ʣ�������ʹ�õ��Ǹ�ָ��

**�ǳ����ԣ�����Ƿ�ָ�����������������ֵ�������������Ǹ�������ָ��������������ò�������ͬһ���ڴ��ͬһ��ʵ����
	����ǲ������ڴ����ʱ���Ҳ���ͬһ��ʵ��ʱ��һ��Ҫ����ָ��������ķ�����


[��ǳ����]
shadow copy
	Ӱ�ӿ�����Ҳ��ǳ���������������������ݣ���������һ�����ö���
deep copy
	�����������ݹ鸴��һ�����

	ע�⣬��ǳ����˵���ǿ����������Ƿ����ݹ鿽����Ҳ����˵���ĳ��ֵ��һ����ַ����ֻ���������ַ�����Ǹ��Ƶ�ַָ������ݡ�
	ֵ�������������ַ������ǳ����������˵���Ǵ���ġ���Ϊ��ַ����ֻ�ǿ����˵�ַ����˱���������Ҳ��ֵ������

	Go�����У���������ʵ���Ͽ������Ǳ�ͷֵ����Ҳ��ֵ��������û��ͨ����ͷֵ�жԵײ����ݽṹ��ָ��ָ������ݽ��и��ƣ������ǳ������
	���������͵ĸ��ƾ���ֵ������Ҳ��������һ����������Ҳ��ǳ��������Ϊ�㲻��˵��һ������ֵ���ڴ��и��Ƴ�һ��������������Ŀ�����
	���������������Ļ������;���һ��������ֵ��û�����뿽��������û��ȥ��������¶���
	�򵥽�����ҿ����ÿ����ļ��Ƿ�������Ӹ�������⡣ֱ�Ӹ��������Ӿ���ǳ������������������渴�������ݾ��������
	*�������ݽṹ����������Ƕ�ף���ʱǶ�׺�����������������Ǵ��ۺܸߣ�
		���ԣ�ǳ�������������ձ���õķ�����


��9-6-�ӿڸ��
=====================================================================
�ӿ�interface����Java���ƣ���һ����Ϊ�淶�ļ��ϣ����Ƕ���һ��δʵ�ֵĺ���������˭ʹ�ýӿھ��ǲ��սӿڵķ�������ʵ������
type �ӿ��� interface{
		����1 (�����б�1) ����ֵ�б�1
		����2 (�����б�2) ����ֵ�б�2
		...
}
�ӿ�����ϰ���ڽӿ����������er��׺
�����б�����ֵ�б���������Բ�д
���Ҫ�ڰ���ʹ�ýӿڣ��ӿ���Ӧ������ĸ��д������Ҫ�ڰ���ʹ�ã�����������ĸҲҪ��д
�ӿ��еķ���Ӧ����ƺ�����Ҫ̫��
---
�ӿ�����ϰ����er��β
һ�׷�������������ϣ�����ʵ�֣�˭��˭ʵ��
	����ǩ�����βΣ�����ֵ������ֻ������
	ʵ�ֽӿ�
		*** ����ĳһ���ӿڵķ���ȫ��ʵ�֣�һ������

�ɼ���
Go�����֣����飬ʵ��С�ӿ�
	�ӿڵķ�����������
	�������ʵ�ִ�ӿ�

Լ������
	�����Ƕ����������������ڽṹ��
	Լ�����ṹ��ķ�����Ա
	�ṹ��ʵ����ĳ�ӿڣ����ظýӿڵ�Լ��
		ʵ���˸ýӿ����еķ���
		
Go���� ��������� ���Java��˵��
	Java
		class Xxx implements �ӿ�1���ӿ�2
	Go���ĵ�ʵ��

EG��
package main

import "fmt"

/*	���һ���ṹ��ʵ����һ���ӿ����������з�������˵�ṹ��ʵ���˸ýӿڡ�
	һ���ṹ�����ʵ�ֶ����ͬ�ӿڡ�
*/

//Personʵ����Sporter�ӿ�,���溯��ȱһ����
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

/*Users ʵ����Sporter�ӿ�*/
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
	//p1 ����Person���� Ҳ�� Sporter����
	fmt.Printf("p1 = %T,s = %T\n", p1, s) //p1 = *main.Person,s = *main.Person
	s.jump()
	s.run()
	// s.swim() ���ܵ���������� ��Ϊ�ӿ���û���������
	// run -----------
	// jump -----------

	//s ����User���� Ҳ�� Sporter����
	u := new(User)
	s = u                                //var s Sporter = Sporter(p1)
	fmt.Printf("p1 = %T,s = %T\n", s, u) //p1 = *main.User,s = *main.User
	s.jump()
	s.run()
	// User jump -----------
	// User run -----------
}
 
��9-7-�ӿ�Ƕ�ס�
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
ReadCloser�ӿ���Reader��Closer�ӿ���϶��ɣ�Ҳ����˵��ӵ��Read��Close��������

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

*  һ�㽫�ӿڷֿ����� Ȼ����ϳ�һ����ӿ�

��9-8-�սӿںͽӿ����Ͷ��ԡ�
=====================================================================
[�սӿ�]
	�սӿڣ�ʵ�����ǿսӿ����ͣ�д��
		interface {}��
	Ϊ�˷���ʹ�ã�Go����Ϊ������һ������any����
		type any =interface{}
	�սӿڣ�û���κη�����������ˣ��κ����Ͷ�������ʽʵ�ֿսӿڵķ�������Ϊ�κ����Ͷ���������սӿڵ�Ҫ��
	�κ����͵�ֵ�����Կ����ǿսӿ����͡�
��
EG��
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

[�ӿ����Ͷ���]
�ӿ����Ͷ���(Type Assertions)���Խ��ӿ�ת��������һ�ֽӿڣ�Ҳ���Խ��ӿ�ת������������͡��ӿ����Ͷ��Ը�ʽ 
t := i.(T)
	i����ӿڱ���
	T��ʾת��Ŀ������
	t����ת����ı���
		����ʧ�ܣ�Ҳ����˵iû��ʵ��T�ӿڵķ�����panic
	t��ok:=i.(T)�������ʧ�ܲ�panic��ͨ��ok��true��false�ж�i�Ƿ���T���ͽӿ�

EG��
package main

import "fmt"

func main() {
	var b interface{} = 500
	// fmt.Println(b.(string)) //panicת��ʧ��
	//panic: interface conversion: interface {} is int, not string
	if s, ok := b.(string); ok {
		fmt.Println("���Գɹ���ֵ��", s)
	} else {
		fmt.Println("����ʧ��") //����ʧ��
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
		fmt.Println("�ַ���")
	case int:
		fmt.Println("����", v) //
	default:
		fmt.Println("��������")
	}
}

[�����ʽ�ӿ�]
����ʹ��fmt.Print�Ⱥ���ʱ��������һ��ֵ����һ��ȱʡ��ӡ��ʽ�������Ͼ���ʵ�ִ�ӡ��صĽӿڡ�

package main

import "fmt"

type Person struct {
	name string
}

/*fmt���ڲ���������ӿ�������
Ҳ�������Լ�������String��GoString �Ͱ������Լ����巽��������
�����Ĭ�ϴ�ӡ
*/

/*String Interface*/
//1.
// func (Person) String() string { //Ӱ�� ȱʡ��ʽ%v����ϸ��ʽ%+v
// 	return "aaa"
// }
//2.ָ����Ϊreceiver ֻ��ָ����Ч ָ������Զ����ʽ
func (*Person) String() string { //Ӱ�� ȱʡ��ʽ%v����ϸ��ʽ%+v
	return "aaa"
}

/*GoString Interface*/
//1.
// func (Person) GoString() string { //Ӱ�� %#v
// 	return "xks"
// }
//2.ָ����Ϊreceiver ֻ��ָ����Ч ָ������Զ����ʽ
func (*Person) GoString() string { //Ӱ�� %#v
	return "xks"
}
//��ͨ����
func (*Person) foo() string {
	return "foo"
}

func main() {
	p1 := Person{"Tom"}
	fmt.Println("1:", p1, &p1)          //{Tom} &{Tom} ȱʡ��ӡ��ʽ
	fmt.Printf("2: %+v,%+v\n", p1, &p1) //{name:Tom},&{name:Tom}
	fmt.Printf("3: %#v,%#v\n", p1, &p1) //main.Person{name:"Tom"},&main.Person{name:"Tom"}
	fmt.Println(p1.foo(), (&p1).foo())  //foo foo

	// ��������� String GoString ʹ��ϵͳ�ڽ���ʽ
	// 1: {Tom} &{Tom}
	// 2: {name:Tom},&{name:Tom}
	// 3: main.Person{name:"Tom"},&main.Person{name:"Tom"}

	// 1.������� String GoString ʹ��ϵͳ�ڽ���ʽ
	// 1: aaa aaa
	// 2: aaa,aaa
	// 3: xks,xks

	//2.���ָ����Ϊreceiver ֻ��ָ����Ч ָ������Զ����ʽ
	// 1: {Tom} aaa
	// 2: {name:Tom},aaa
	// 3: main.Person{name:"Tom"},xks

	//����ʹ�����Ͷ��Խ����ж�
	var i interface{}
	i = p1
	//i.(fmt.Stringer)
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case string:
		fmt.Println("�ַ���")
	case int:
		fmt.Println("����", v)
	case Person:
		fmt.Println("Person type:", v)
	case fmt.Stringer: //�ӿ�����
		fmt.Println("String Interface type:", v) //Person type: aaa
	default:
		fmt.Println("��������")
	}
}
[�ܽ�]
Stringer��GoStringer�ӿڵķ��������receiver�����ָ�룬ֻ�ܶ�ָ��������;���receiver�����ʵ����ʵ����ָ�붼�����á� 
��ͨ����receiver������ʵ������ָ�룬ʵ����ָ�붼���Ե��ø÷���ʱ

��9-9-�Զ������Ͷϵ���ԡ�
=====================================================================
�쳣����
	type error {
		Error() string
	}

[ʵ��Դ�� �Զ���error]-ϵͳ����Դ��error.new(type)

package main

import "fmt"

//���ڲ�ʹ��
type errorString struct {
	s string
}

//������errorString�ķ�����ͬʱҲ����error�ӿڵĶ���
func (e *errorString) Error() string {
	return e.s
}

//���캯��
// func New(text string) *errorString {
// 	return &errorString{s: text}
// }
func New(reason string) error { //error���ڽ�����
	return &errorString{s: reason} //����errorString���ͣ�Ҳ��error�ӿ�����
}

func (e *errorString) foo() {
	fmt.Println("errorString�ṹ���foo����")
}

func main() {
	e1 := New("Error:[1]")  //var e1 error = New("xxx")
	fmt.Printf("%T\n", e1)  //*main.errorString
	fmt.Println(e1.Error()) //Error:[1]
	if v, ok := e1.(*errorString); ok {
		v.foo() //errorString�ṹ���foo����
	}

	// e1,e2 ����
	e2 := errorString{"Error:[2]"} // var e2 errorString
	fmt.Println(e2.Error())        //Error:[2]
}

[�Զ���Error]
package main

import (
	"errors"
	"fmt"
)

var ErrDisvisionByZero = errors.New("�����쳣") //var ErrDisvisionByZero error����

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDisvisionByZero
	}
	return a / b, nil
}

func main() {
	if v, err := div(5, 0); err == nil { //nil �޴���
		fmt.Println("v =", v)
	} else { //�д���
		// fmt.Println("err = ", err)
		// fmt.Println("�д���", err.Error(), v)
		fmt.Println("�д���", err) //fmt.Print* �ڲ���error���⴦��

	}
	//div(5, 2) => v = 2
	//div(5, 0) => �д��� �����쳣 0

	// fmt.Println(div(5, 2)) //2 <nil>

}

��9-10-panic��recover��
=====================================================================

panic�ǲ��õģ���Ϊ������ʱ����������ɳ��������������ֹ�Ⱥ��������û��ϣ������������������ڴ�����ʱ������ʱpanic����ֹ��������
	�������г��������ɸ������ʧ���������Ӳ�ʹ�Ĵ��ۡ����ԣ���ʱ��panic���µĳ������ʵ���Ͽ��Լ�ʱֹ��ֻ��������Ȩȡ���ᡣ
panic��Ȼ���ã�����ܲ��Ҳ���򲻵��ѣ��������ϱ�¶���⣬��ʱ���ֺ;������⡣

panic����
	runtime����ʱ�������׳�panic����������Խ�硢����
	�����ֶ�����panic(reason)�����reason��������������

panicִ��
	����ִ�е�ǰ�Ѿ�ע�����goroutine��defer��(recover���������)
	��ӡ������Ϣ�͵��ö�ջ
	����exit(2)������������

�׳��쳣�ķ���
	ϵͳ����ʱ�׳�panic
	�ֶ�panic
���panicû�д����������ų���ֱ��������


EG��
package main

import (
	"errors"
	"fmt"
)

func div(a, b int) int {
	defer fmt.Println("start")
	defer fmt.Println(a, b)
	r := a / b //��һ���п���panic
	fmt.Println("end") //panic�˲�����ִ����
	return r
}

func main() {
	fmt.Println(div(5, 0))
}

// 5 0
// start
// panic: runtime error: integer divide by zero
 
// goroutine 1 [running]: �����ǵ���ջ��divѹ��main
// main.div(0x5, 0x0)
// 	e:/goprojects/main.go:13 +0x2cc ������к�
// main.main()
// 	e:/goprojects/main.go:19 +0x25

[recover]
package main

import (
	"errors"
	"fmt"
	"runtime"
)

var ErrDisvisionByZero = errors.New("�����쳣") //var ErrDisvisionByZero error����

func div(a, b int) int {
	defer func() {
		err := recover()
		fmt.Printf("1 %+v, %[1]T\n", err)
	}()
	defer fmt.Println("start")
	defer fmt.Println(a, b)
	defer func() {
		err := recover() //һ��recover�ˣ����൱�ڴ��������
		fmt.Printf("2 %+v, %[1]T\n", err)
		switch v := err.(type) { //���Ͷ���
		case runtime.Error:
			//��Դ���� runtime/error.go No.75 �У���ΪerrorStringҲʵ����RuntimeError����
			//Ҳ����˵����׼��runtime����ʱ���󣬶���runtime.Error�ӿڵ�
			fmt.Printf("ԭ��%T, %#[1]v\n", v)
		case []int:
			fmt.Println("ԭ����Ƭ", v)
		}
		fmt.Println("�뿪recover����")
	}()

	r := a / b //��һ���п���panic

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
// ԭ��runtime.errorString, "integer divide by zero"
// �뿪recover����
// 5 0
// start
// 1 <nil>, <nil>
// 0 !!!
// main exit

/*2*/
//	fmt.Println(div(5, 2), "!!!")
// 2 [1 3 5], []int
// ԭ����Ƭ [1 3 5]
// �뿪recover����
// 5 2
// start
// 1 <nil>, <nil>
// 0 !!!
// main exit

�����У�һ����ĳ������panic����ǰ����panic֮�����佫����ִ�У���ʼִ��defer�������defer�д���recover��
	���൱�ڵ�ǰ���������Ĵ���õ��˴�����ǰ����ִ����defer����ǰ�����˳�ִ�У����򻹿��Դӵ�ǰ����֮�����ִ�С�

���Թ۲쵽panic��recover������
	��panic��һ·�����׳�����û��һ������recover��Ҳ����˵û�еط�������󣬳��������painc����recover�������൱�ڴ��󱻴������
	��ǰ����deferִ������˳���ǰ�����ӵ�ǰ����֮�����ִ��
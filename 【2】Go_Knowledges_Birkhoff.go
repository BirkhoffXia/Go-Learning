��1-Go����������
Go:1.20.7
	GOPATH 
	GOROOT
	GOPROXY
		https://mirrors.aliyun.com/goproxy/,direct
		https://proxy.golang.com.cn,direct(���Ҫװ��� ��Ϊ���������������)
		go env
	GOMODCACHE=C:\Users\BIRKHOFF ALW\go\pkg\mod

Git 
	
VScode
	GOPROXY: https://proxy.golang.com.cn,direct(���Ҫװ��� ��Ϊ���������������)
	cmd:
		go install -v golang.org/x/tools/gopls@latest 
		go install -v github.com/cweill/gotests/gotests@v1.6.0
		go install -v github.com/fatih/gomodifytags@v1.16.0 
		go install -v github.com/go-delve/delve/cmd/dlv@latest

package main

import "fmt"

func main() {
	fmt.Printf("BIRKHOFF FIRST GO PROCESS")
}

PS E:\goprojects> go build main.go
PS E:\goprojects> .\main.exe
BIRKHOFF FIRST GO PROCESS
PS E:\goprojects> 

#��ʾ���벢ִ�д���
go run main.go

##F5ִ�� run-debugģʽ �ᱨ��
#��ʾû���ҵ��ļ�
Build Error: go build -o e:\goprojects\__debug_bin1722392392.exe -gcflags all=-N -l .
go: go.mod file not found in current directory or any parent directory; see 'go help modules' (exit status 1)
#��Ҫ��ʼ���������������������
E:\goprojects>go mod init test ����go.mod
go: creating new go.mod: module test
go: to add module requirements and sums:
        go mod tidy
#�ٿ�F5 

##outline map���-����Ϊ�˲鿴Դ�뺯��
##Golang postfix���-������ڿ������� 

=========================================================================================================================
��2-��ʶ���ͳ�����
# TODO TREE

#var�Ǳ��� const�ǳ����޷��ı�ᱨ��
	var a = 2000
	a = 4000

	const b = 319
	b = 1220

#������ֻ��дһ���ַ�

package main

import (
	"fmt"
)

func main() {
	fmt.Println("BIRKHOFF FIRST GO PROCESS")
	fmt.Printf("\"abc\": %v\n", "abc") //postfix

	// TODO ��ģ���´���ȷ��
	// NOTE ����
	// DEPECATED δ���˰汾��������

	var height = 175
	var weight = 83
	//var bmi = 0
	if height >= 175 {
		var bmi = height / weight
		fmt.Println(weight)
		fmt.Println(height)
		fmt.Println(bmi)
	}

	//���泣����Ϊ�˷����ɽ����ı�ʾ��ʽ
	//������ֵ���㶨�������
	//�����ͳ���untyped constant ��ȱʡ����Ϊbool��rune��int��float64��complex128�����ַ���
	// �ַ����� rune ''������ֻ��һ���ַ�
	// rune �ַ����ͱ�����֤�� ����int32 4���ֽڵ�����
	fmt.Println('\u6d4b')
	fmt.Println('\n')
	fmt.Println('\x31')
	fmt.Println('1')
	//�ַ����������ж���ַ�
	fmt.Println("")
	fmt.Println("BIRHOFF")
	fmt.Println("\x61d\x63") //abc
	fmt.Println("\u6d4b��")   //����
	fmt.Println("\n")

	//������ֵ �������Ƶ� ���泣��ȱʡֵ��ֵ�����
	// const (
	// 	a int8 = 100
	// 	b      = 1220
	// 	c      = "abc"
	// )
	// fmt.Println(a, b, c)

	//itoa�����ڳ��������ʱ��һ���ڳ�������ʱ�õ���iota�������൱��������
	// const a = iota
	// const b = iota
	// fmt.Println(a) //0
	// fmt.Printf("%T", b) // int
	// const (
	// 	a = iota
	// 	b //b=iota �������һ�й�ʽһ��������ʡ��
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
	//iota������������
	//_���ǿձ��ʶ�� Ϊ����������ֵ ��ֵ�ᱻ���� ��������벻��ʹ������������ֵ��Ҳ����ʹ����������Ϊ����������ֵ ��һ���ڶ� �Ϸ���ʶ��
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
��2.1-��ʶ����
[����]
	ָ����ָ���κκϷ�ֵ
	����Ա���ʱ�õģ�д�����ʱ��
	����󣬱�ʶ����ʲô?���е�ֵ��Ҫ�����ڴ��У���ʶ��ָ�����ֵ����ʶ�������ˣ���Ϊ���������ڴ��ַ
	���� ��ʶ����ָ���������ƣ�����ָ�����ͣ���ʶ�����;���ָ����ֵ������
	ָ��������Go�е�����
	��ʶ��������Ϊ���ҵ��Ǹ�ֵ
	
var p *int //p�Ǳ�ʶ����p��������ֵ�����ͣ�*int nil��δ��Ҫ��*int���͵����ݽ���������ϵ��ָ�����ֵ
var p a_type //p�Ǳ�ʶ����ָ��ĳ�����͵�ֵ��p��������ڴ��ַ

package main

import "fmt"

/*
//����ĸ��x����ĸСд������main����ɼ�������ȫ�ֱ���x
//Y����ĸ��д��������main���ڿɼ����������ڰ���ɼ�
//Go�У����Ƕ���ı�ʶ��������ȫ��������?��
//x,Y����main����ȫ�ֱ�����ֻ��������ֻ��ʹ��Y��ʹ�÷��� main.Y
var x = 100   //�������壬���ڶ�����룬ȫ�ֱ�������ó�ʼ��
const Y = 200 //ȫ�ֳ���
*/

//z:=300 //�̸�ʽ�������壬����ȫ��ʹ�ã�ֻ�����ں������� ���д����

func main() {
	/*
		//��ӡȫ�ֱ���
		fmt.Println(x, Y) //100 200
		//Println ����ĸ��д����ʾ��ʶ���ǰ���ɼ���fmt�������ȫ�ֱ�ʶ��
	*/

	/*���һ��������ʶ������ʼ����go�������ǳ�ʼ��Ϊ�����͵���ֵ����Ϊ ��ֵ����*/
	//var a �� ��֪�����͡���֪��ֵ����Ϊ��ʶ����Ҫ�и�ָ��
	// var a int      //���� ����û��ʹ�á�c c++�ȡ������˱�ʶ�����嵫û�и���ʼֵ��Ϊ����
	// fmt.Println(a) //0

	// var b string   //��ʼ�� ����
	// fmt.Println(b) //��

	// var c bool     //һ��ȷ���������ı�ʶ�����ͣ�goΪ�˲����� �ṩ��һ�� �����͵� ��ֵ
	// fmt.Println(c) //false

	//const d int //������ʶ����������ʼ��

	/*
		//Go ��֧���ٴθ�ֵ��һ��������
		var test = 100
		fmt.Println(test)
		fmt.Printf("test = %T\n", test)
		// test = "HEYE" //���������ͱ��� ����ᱨ��
		// fmt.Printf("test = %T", test)
	*/

	/*
		//var a int,b int //����д������
		//var a,b int ,c string //���� ֻ��ͬһ����
		//var a,b int // a,b����int����
	*/

	/*
		//��ֵ������ : ֻ�����˱��������� û�и���ֵ(��ʼ��) Go�ṩ������
		//bool false int 0 float64 0 string ""
		//var a int
		//var b string
	*/

	/*
		//����ʼ��
		// var a = 100
		// var b int = 200         // var b int = int(200)
		// var c, d int = 300, 400 //û�������Ƶ� var c, d = 300, 400 int��
		// var j, k int8 = 30, 40  // var j,k int8 = int8(30),int8(40)
		// var e, f = 500, "abc"   //�����Ƶ���
		// var g int,f string = 600, "abcdefg" //����
		//fmt.Println(a, b, c, d, e, f, j, k)
	*/

	/*
		var t = nil //����������д
		fmt.Println(t)
	*/

	//����д
	/*
		var ( //������ҵ�������
			a           = 100
			b    string = "abc"
			c, d        = 200, "abc"
			g, h int    = 400, 500
		)
		fmt.Println(a, b, c, d, g, h)
	*/

	/*
		//�»��� ���Խ�����ֵΪ�ڶ� �����ܴ�ӡ
		var a, _ = 200, 400
		fmt.Println(a)    //200
		fmt.Println(a, _) //cannot use _ as value or type
	*/

	/*
		//�̸�ʽ�������
		a := 100
		b := "abc"
		fmt.Println(a, b) //100 abc
	*/

	/*
		//�̸�ʽ ������������ʼ��   a:=100 ���� var a = 100
		a, b := 100, "abc"
		fmt.Println(a, b) //100 abc
		// b := "xyz"        //error �ظ�����
		//var b = "xyz"     //�ظ�
	*/

	/*
		//�Ѿ���ֵ�������ڶ̸�ʽ��ֵ
		var a int
		a := 200
	*/

	/*
		//�˸�ֵ���� ��:  :=
		a, b = 100, "xyz"
		fmt.Println(a, b)
	*/

	/*
		//�������� ��Ϊ�Ѿ������� �ٴθ�ֵ
		a, b := 100, "xyz"
		fmt.Println(a, b)//100 xyz
		a, b = 200, "bva"
		fmt.Println(a, b) //200 bva
	*/

	/*
		//ͬ���Ϳ��Խ���
		a, b := 100, 300
		a, b = b, a
		fmt.Println(a, b) //300 100
	*/

	//ʹ���м����
	a, b := 100, 300
	c := a
	a = b
	b = c
	fmt.Println(a, b) //300 100

}
=========================================================================================================================
��2.2 - �����ͽ��ơ�
0x32 תΪʮ���� 3*16 + 2 = 50
50��2����
128 64 32 16 8 4 2 1   
00110010
ʮ������ 4��һ��
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
		//bool �;��� true �� false
		a := false
		fmt.Println(a + 100) //һ��a �� bool���ܹ����м�����
	*/

	/*
		a := 100          //int 64bits 8 Bytes�з���
		var b int64 = 200 //int64 64bits 8 Bytes�з��� var b int64 = int64(200)
		fmt.Println(a, b) // 100 200
		//fmt.Println(a + b) ��������Ϊa int ��b int64 ���Ͳ�һ�²������
		fmt.Println(b + int64(a)) //int64(a)
		fmt.Println(a + 200) //300
		fmt.Println(b + 300) //Ҳ���� ��Ϊ300 ���������泣�� ���Խ�����ʽת��
	*/

	/*
		var a int8 = 127 // ����
		fmt.Println(a)
		//var b int8 = 128 //������ ���˷�Χ�� ��Ϊ���з��� �з������127 0111 1111 ǰ���0Ϊ����
		//fmt.Println(b)
		var c uint8 = 128 // ������Ϊ���޷��� ��1111 1111 ǰ��1Ϊռλ�� �������Ϊ255
		fmt.Println(c)
		//var d uint8 = -1 // uint8 �޷��Ų����Ǹ���
		//fmt.Println(d)
	*/

	/*
		var a = 1 * 2.3 //Ϊʲô���ԣ� ��Ϊ�ұ�ʹ�õĶ��������ͳ���untyped constant������������������ʽת����GoΪ�˷��㣬���ܹ������壬Ҫ���ٳ���Աת�����͵ĸ������������ͳ���������һЩ���Ĳ�����
		fmt.Println(a)
		//ռλ�� %T type ��%f float ��%v value ���ø�����Ĭ�ϴ�ӡ��ʽ
		fmt.Printf("Type = %T , Equals = %v\n", a, a) //Format Type = float64 , Equals = 2.300000

		var b = 1
		//fmt.Println(b * 2.4) //������ b��int�� 2.4��float32����
		fmt.Println(float32(b) * 2.4)             //���� ����ת��=> 2.4
		fmt.Printf("Type = %T\n", float32(b)*2.4) //Type = float32
	*/

	var a = 1
	fmt.Printf("a's type = %T \n", a) //int

	//ǿ������ת��:��һ��ֵ��һ������ǿ����ʽת������һ�����ͣ��п���ת��ʧ�ܡ�
}

=========================================================================================================================
��2.3 - �ַ�������ʽ����
package main

func main() {
	/*
		var a = 3.1415926
		fmt.Printf("%T %v\n", a, a)         //float64 3.1415926
		fmt.Printf("%T %f\n", a, a)         //float64 3.141593
		fmt.Printf("%T %.3f\n", a, a)       //float64 3.142
		fmt.Printf("%T ,,%10.3f,,\n", a, a) //float64 ,,     3.142,,  10��ʾ��� С������3λ С����ռ1λ
	*/

	/*
		//ת���ַ�:ÿһ�������ַ� rune����-int32 ������Ϊ�����ַ�ʹ�ã�Ҳ������Ϊ�ַ����е�һ���ַ�
		//rune:������int32�ı����������з��涼��int32��Ч�����չ����������������ַ�ֵ������ֵ
		var a rune = '\'' //һ���ַ���rune int32 4bytes ����
		//��������
		fmt.Printf("Type = %T %[1]v\n", a)    //Type = int32 39 ��������
		fmt.Printf("Type = %[1]T %[1]v\n", a) //ʹ��ͬһ������a ����%[2]v �ͳ�����Χ�� ������ͬһ������
		//fmt.Printf("Type = ..... %[n]T %v\n", a,......n ) Ĭ��%v = %[n+1]v

		var b = 200
		fmt.Printf("%[2]d %[2]d\n", a, b)          //200 200 ��ʾʹ��bû��ʹ�õ�a aû���õ�
		fmt.Printf("%[2]d %[1]d %v\n", a, b, 1000) //200 39 200 �ڶ�����b����һ������a��%v��ʾ1+1 ���ǵڶ�������b��ֵ

		var x rune = 97
		fmt.Printf("%T,%[1]d,%[1]c\n", x) //int32,97,a %c char�ַ���ʽ��ӡ��ʾͨ����97ת��ΪASCII���������
	*/

	/*
		'\n' int32 4bytes���� 10 0xa
		'\r\n' ���� ��Ϊ������ֻ��һ���ַ�
		"\n" �ַ���ɵ�����,string
		"\r\n" 2���ַ������
		�ַ���
		�ַ�����ɵ�����
		"","\n" "abcd\r\nabc\txyz123"
		123 int�����泣��
		"123" �ַ������ַ�1��2��3 3���ַ���������
	*/

	/*
		var x = "abc\nxyz" //��׼�������׼�������
		fmt.Print(x)
		var y = `abc	123` //7���ַ� abc\t123
		fmt.Println(y)
		var z = `abc\t123` //8�� �����Ų�֧��ת�����
		fmt.Println(z)     //abc\t123
	*/

	/*
		var x = "abc\nxyz\n123"
		fmt.Printf("%v\n", x)
		fmt.Printf("%s\n", x)
		fmt.Printf("%q\n", x) //%q ���� "%s" "abc\nxyz\n123"

		y := 97
		fmt.Printf("%d %[1]x %#[1]x %[1]b %#[1]b %[1]c %[1]q\n", y) //97 61 0x61 1100001 0b1100001 a 'a'

		//%U ��һ��������Unicode��ʽ��ӡ��
		fmt.Printf("%U,%x,%c\n", 27979, 27979, 27979) //U+6D4B,6d4b,��
	*/

	/*
		//Sprint:�൱��Print���������Ϊstring��
		//Sprintln:�൱��Println���������Ϊstring
		//Sprintf:�൱��Printf���������Ϊstring
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
��2.4 - ��������
	/*
			fmt.Println(15&5, 15&^5) //5 10
			//& ���˷�*
			1111
			0101
		    ------
			0101

			//&^ ����Ϊ1����Ϊ0-©1 ��ͬ1Ϊ0 ����y��1��λ���x��Ӧλ
			1111
			0101
			------
		    1010 =10
	*/

	//�Ƚ���������ɵıȽϱ��ʽ ����ֵbool���� ͬ���ͱȽ�
	fmt.Println(1 == '1') //false �Ƿ� ����bool 1=="abc" ��ͬ���Ͳ�ƥ�䲻�ܲ���

	//�߼������ &&��||����
	//fmt.Println(1 && 0, "abc" && true) //������Go�б���ʹ��bool
	fmt.Println(true && true, false && false, 5 > 3 && 1 < 2) //true false true
	//&& || �ж�· ��һ����false �ͺ��治������

	//��ֵ����� = += -= *= /= %= ��= ��= &= &^= ^= |=

	// ����Ŀ�����

	//ָ�� �� ���͡����ͱ���ʲô���ݣ����������ڴ��ַ
	a := 1000
	b := &a
	fmt.Printf("%T a address = %[1]v\nb address = %p\n", &a, b) //%p ָ�뱾���Ǵ����������ƺ���  2����ַһ��
	//*int a address = 0xc0000a6090,b address = 0xc0000a6090
	c := *b
	fmt.Println(c)              //1000
	fmt.Printf("%T %[1]v\n", c) //int 1000
	fmt.Println(a == c)         //true
	fmt.Printf("a address = %p\nb address = %p\nc address = %p\n", &a, b, &c)
	//a address = 0xc0000a6090
	//b address = 0xc0000a6090
	//c address = 0xc0000a6098 a c2����ַ��һ��

	var p1 *int                        //��int���͵�ֵ�ĵ�ַ
	fmt.Printf("%T %[1]p %[1]v\n", p1) //*int 0x0 <nil>(��ָ���Σ��)
	p1 = &a
	fmt.Printf("%T %[1]p %[1]v\n", p1) //*int 0xc0000a6090 0xc0000a6090
	fmt.Println(*p1)                   //1000
	*p1 = 2000
	fmt.Println(*p1)  //2000
	fmt.Println(a, c) //2000 1000 aҲ�����2000 cû�б�
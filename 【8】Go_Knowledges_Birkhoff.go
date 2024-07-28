��8-1-������
=====================================================================
[������]
�����Ὺ��һ���ֲ����������ж���ı�ʶ�������ں���֮��ʹ�ã�Ҳ��Ϊ��ʶ���ں����еĿɼ���Χ
���ֶԱ�ʶ��Լ���Ŀɼ���Χ��Ϊ������
1.����������
2.��ʽ�Ŀ����������κ�һ���������ж���ı�ʶ������������ֻ������Դ�������
3.universe�飺����飬��˼����ȫ�ֿ飬�����������ڽ��ġ�Ԥ����ı�ʶ���������ȫ�ֻ�����
							���bool��int��nil��true��false��iota��append�ȱ�ʶ��ȫ�ֿɼ����洦���á�
4.���飺ÿһ��package�����ð�����Դ�ļ����γɵ���������ʱ�ڰ��ж�����붨���ʶ����Ҳ��Ϊȫ�ֱ�ʶ����
				���а��ڶ���ȫ�ֱ�ʶ�������ڿɼ������Ķ�������б�ʶ������ĸ��д�򵼳����Ӷ�����ɼ���ʹ��ʱҲҪ���ϰ��������� fmt.Prinf()��
5.�����飺����������ʱ��ʹ���˻����ţ������������������һ����ʽ����顣�����������һ����������.

package main

import "fmt"

//�����������������壬ֻ��ʹ��const��var����
const a = 100

var b = 200

// c:= 300 //������ʹ�ö̸�ʽ
var d = 400

func showB() int {
	return b
}

func main() {
	fmt.Println(1, a) //1 100
	// fmt.Println(1.1, &a) //ע�ⳣ������Ѱַ�����ǶԳ����ı��� invalid operation: cannot take address of a (untyped int constant 100)

	var a = 500
	fmt.Println(2, a, &a) //2 500 0xc0000180b0 ���¶���aΪ�������Է���

	//��b�Ĳ�����˼������ �Ƿ�b���ˣ�
	fmt.Println(3, b, &b) //3 200 0xc87340
	b = 600
	fmt.Println(3.1, b, &b) //3.1 600 0x9f8340 ��ַû�з����仯 ֵ���滻Ϊ600
	b := 700
	fmt.Println(3.2, b, &b)   //3.2 700 0xc000122098	��Ϊ���¶����� ��ַ�����仯ֵ�����仯
	fmt.Println(3.3, showB()) // 3.3 600 ����ջ֡ʱ�Ѿ���ȫ�ֵ�200��ַ��¼�� ֮��200��Ϊ600�� ����600 ������main�е�b=700

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
		// fmt.Println(4.2, x) //�������� ���ⲻ�ɼ�
	}
	// fmt.Println(4.3, j, k, t) //�������� ����
	fmt.Println(4.4, a, b) //700 700

	for i, v := range []int{1, 3, 5} {
		fmt.Printf("Index = %d,Type = %T,value = %[2]d\n", i, v)
	}
	// Index = 0,Type = int,value = 1
	// Index = 1,Type = int,value = 3
	// Index = 2,Type = int,value = 5
	// fmt.Println(i,v) //i,v���ɼ�������zaifor����������
}

[�ܽ�]
��ʶ��������
	��ʶ�����ⲻ�ɼ����ڱ�ʶ�������������������ǿ�������ʶ����
	ʹ�ñ�ʶ�����Լ���һ�㶨��ı�ʶ�����ȣ����û�У����������ͬ����ʶ��--�Լ����ȣ��ɽ���Զ
	��ʶ�����ڿɼ������ڲ��ľֲ��������У�����ʹ���ⲿ����ı�ʶ��--���ڴ�͸
	������ʶ��
		�����ڰ��ڣ����ɼ�
		������ʣ�������ʶ�������д��ͷ�����ܵ��������⣬�����ڰ���ʹ��xx����.varName ����ʽ���ʡ����� fmt.print()

��8-2-����������
=====================================================================
��������
	û�����ֵĺ���
	��;:��Ҫ���ڸ߽׺����У����������strings.Map���ǳ��߼����߼��ĸ߶ȳ���
�߽׺���
	ĳһ���β������Ǻ���
		�����߼����ƣ�����ʹ����
	ĳһ������ֵ�����Ǻ���
	����2������������һ

package main

import (
	"fmt"
	"strings"
)

func add(x, y int) int {
	return x + y
}

// ���Զ�������
type MyFunc func(a, b int) int

func calc1(x, y int, fn MyFunc) int {
	return fn(x, y)
}

// %T func(a,b int) int ���������ͣ�ǩ��
func calc(x, y int, fn func(a, b int) int) int {
	return fn(x, y)
}

func main() {
	//��Ҫ��2��int������ĳ�ּ���
	fmt.Println(calc(4, 5, add)) //9
	//ͨ�������������е���
	fmt.Println(calc(4, 5, func(a, b int) int { return a - b })) //-1

	//strings.Map
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "abc")) //bcd
}

��8-3-�ݹ麯����
=====================================================================
�ݹ�Ҫ��
	�ݹ�һ��Ҫ���˳��������ݹ����һ��Ҫִ�е�����˳�������û���˳������ĵݹ���ã��������޵���
	�ݹ���õ���Ȳ��˹���
	Go���Բ������ú������޵��ã�ջ�ռ��ջ�ľ�
		goroutine stack exceeds 1000000000-byte limit

package main

import (
	"fmt"
)

func fibLoop(n int) int { //ѭ��
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

// �ݹ�汾 1.���õ��ƹ�ʽ
func fib1(n int) int {
	if n < 3 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 2.ѭ����α�ɵݹ麯�����
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

[�ݹ�Ч��]
����3��쳲���������ʵ�֣������Ǹ�Ч�ʸ�?�ݹ�Ч��һ������?�ĸ��汾��?
	�ݹ�汾1Ч�ʼ��ͣ�����Ϊ�д����ظ����㡣
	�ݹ�汾2�����˵ݹ麯�����ò�δ���ѭ����Σ�Ч�ʻ�������ѭ����Ч�ʲ��
�ݹ��2��ѭ����˭��?
	ѭ�����Щ����Ϊ�ݹ���������ƣ���һ���������ÿ����ϴ�
[�ܽ�]
	�ݹ���һ�ֺ���Ȼ�ı������߼�˼ά
	�ݹ��������Ч�ʵͣ�ÿһ�ε��ú�����Ҫ����ջ֡
	�ݹ���������ƣ�����ݹ���̫���������ѹջ��ջ�ڴ�Ϳ����������������޴����ĵݹ飬����ʹ�õݹ����
		����ʹ��ѭ�����棬ѭ��������΢����һЩ������ֻҪ������ѭ�������Զ�ε���ֱ��������
	��������ݹ飬������ʹ��ѭ��ʵ��
	��ʹ�ݹ����ܼ�࣬�����ܲ������õݹ�

��8-4-Ƕ�׺����ͱհ���
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
		c = 97 //��ʱ97������c�����¸�ֵ
		fmt.Println("1 c=", c, &c)
		c := 0x31
		fmt.Println("3 c=", c, &c) //��ʱ��c�Ǿֲ����� ����������

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


[�հ�] - GO�˽����¼��� Python��JS��Ҫװ������Ҫ
	Ƕ�׺���
	���ɱ��������ڱ��������򴴽��ľֲ������������ĺ����������д����ľֲ�����
	�հ����ڴ溯���õ�����㺯���ж�������ɱ������������ˣ��γ��˱հ�
	**�հ��γ��Ժ��б���������������¡�
	ջ
		ջ������һ�����ڴ�ջ֡����
	��(�������)
		ֻҪ���˼ǵ��㣬�� ���ü�����Ϊ0����������
		gc����
package main

import "fmt"

func outer() func() { //�߽׺���
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
	fmt.Println("fn address :", fn) //fn ��¼��outer���ص�����ʱ������inner�������壬��˭fn���ã����ü���
	fn()
}

	��15�е���outer����������inner�������󣬲�ʹ�ñ�ʶ��fn��ס������outer����ִ�����ˣ���ջ֡�ϵľֲ�����Ӧ���ͷţ�����inner��������Ϊ��Ҳ�Ǿֲ��ġ�
		���ǣ�c��inner��Ӧ��ֵ�������ͷţ���ΪfnҪ�á�������Щֵ���ܷ���ջ�ϣ�Ҫ�ŵ����ϡ���Go�����У����Ϊ�������ݣ����ݵ�����
	��ĳ��ʱ�̣�fn��������ʱ����Ҫ�õ�c���������ڲ�û�ж���c������outer�ľֲ�������������c��������outer�ĵ��ö��ͷţ�
		��ôfn��������һ�����ִ������ԣ����outer��c�����ͷŵ���outer�Ѿ���������ˣ���ô��?�հ�����inner������ס���ɱ���c(���ݵ����ϵ��ڴ��ַ)
		
��8-5-defer��
=====================================================================
deferӦ�ó���:��Դ�ͷš������ļ��򿪺�Ҫ�رա��������Ӻ�Ҫ�Ͽ���������������ͷŵ�
              ���ϳ����У������Դ�������дdefer��䣬��ȷ�����������˳���panicʱ���ܹ��ͷ���Դ��

defer��˼���Ƴ١��ӳ١��﷨�ܼ򵥣��������������ǰ����defer�Ϳ����ˡ�
deferִ��ʱ��:��ĳ������ʹ��defer��䣬��ʹ��defer������������ӳٴ���
							���ú�����������ʱ������panicʱ��defer����俪ʼִ�С�
	            ע��os.Exit�������������������ִ��defer
deferִ��˳��:ͬһ�����������ж��defer��䣬���μ������ջ��(LIFO)���������ػ�panicʱ����ջ������ִ��defer����䡣
	ִ�е��Ⱥ�˳���ע���˳�������෴��Ҳ���Ǻ�ע�����ִ��defer�����������һ�������򷽷��ĵ��á�

defer fn() �Ժ�ȥִ��fn ��fn���û��Ӻ�
����������λ��
ִ��˳��
	����ע���˳��ע�ᵽִ��ջ������ִ�� LIFO ����ȳ�

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
	panic("�Ҵ�2֮�������")      //����Ӵ˲������ִ��
	log.Fatal("��2֮��fatal") //os.Exst(1)
	defer fmt.Println(2)
	fmt.Println("end")
	// start
	// 2
	// 1
	// panic: �Ҵ�2֮�������
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	count := 1
	defer fmt.Println(count) //1 ע��ʱע�룬���� ��������ϼ�¼����
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
	}() //�޲κ����������������Ҫע�ᣬ��󷵻�count = 3
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

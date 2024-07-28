Viedo Date: 2023/11/18
Made By:	BIRKHOFF
Date:	2024-07-01


��14-1-Э�̡�
=====================================================================
��������
	����Ҫ���ö��̣߳���һ���߳��У�ʵ��������Ϊ���Ƶ��л����������е�Ч��
	��һ���߳��ڣ�ÿһ�����������ڶ��߳��л��������е�Ч����������ֽ�Go-routine
		��Ϊ���ƣ�������Աд���룬�û�̬����
	
	Э�̿����ɶ��߳��𣿿����У����ų�
		
	����
	��python�У��ṩ��yield��䣬�����ú�����yield����ͣ��������yield�ǿ�����Աд�ģ������ɿ�����Ա�Լ����������ںδ� ��ͣ ִ��
		��ͣ�ˣ�����Ҫ�ָ���
	
	���̵߳�ÿһ���߳�Ҫִ��targetĿ��ִ�к������߳��Ǹɻ�ģ��߳�ȥִ��ָ��������ָ�Ҳ������ͣ�̣߳��������Ա���ͣ����������ͣ˭����? OS��������Ա���ܸ�Ԥ
		��ͣҲҪ�ָ�
		�ں�̬

[����]
1.����Э�̣����ʻ���Ҫ���߳�ô��
	A�������ж��߳�
2.����Э�̣��᲻�����̵߳��л���
	A:��Ȼ�ᣬ�̵߳��л���ܲ���
		OS����n���̣߳�������ĳ����ü����̣߳������������̵��߳� 
		����Э���ѵ�������Ͳ������ö��߳���ô��������
3.Э��Ҳ�Ǵ��룬����ϵͳ����С��λ��Э��ô��
	A:�̣߳�Э��Ҳ�����߳�������

�׶ˣ�
	һ�����Ƕ�����Э��ֻ��һ���߳������У�����һ��Э������ס�ˣ�û������yieldû���ó�����Ȩ�����治�ܽ�����

EG1��
def count():
    c=1
    for i in range(5):
        print(c)
        c+=1
x = count() #���û���쳣��û��return֮ǰ������������ô�� û��
#����û�з�����ִ����һ��printô�� ����
#��ͬһ���߳��У�˳��ִ�У�����
print(type(x),x)
print("######")

# 1
# 2
# 3
# 4
# 5
# <class 'NoneType'> None
# ######

EG2��
#����������;��ͣ
def count():
    c=1
    for i in range(5):
        print(c)
        yield c #�������yield���˺�������Ϊ���������������
        print('+' * 30)
        c+=1
x = count() #������ִ�У�������������һ�����������󷵻�
for t in x: t = next(x)
    print("for t == next(t) =>",t)
# 1
# for t == next(t) => 1
# ++++++++++++++++++++++++++++++
# 2
# for t == next(t) => 2
# ++++++++++++++++++++++++++++++
# 3
# for t == next(t) => 3
# ++++++++++++++++++++++++++++++
# 4
# for t == next(t) => 4
# ++++++++++++++++++++++++++++++
# 5
# for t == next(t) => 5
# ++++++++++++++++++++++++++++++
# ######

# print(type(x),x)
# next(x) #��һ�£�����yield��ͣ����������������return��Ч��
# # <class 'generator'> <generator object count at 0x000001B8731FDEB8>
# # 1
# # ######
# next(x) #��һ�¡����ϴ���ͣ��������ִ�У�ֱ�������������ٴ�yield
# print(next(x),"3 After")
# print(next(x),"4 After")
# print(next(x),"5 After")
# print(next(x),"6 After")

print("######")

# <class 'generator'> <generator object count at 0x000002B236B7DE08>
######

# <class 'generator'> <generator object count at 0x00000209BD8DDEB8>
# 1
# ++++++++++++++++++++++++++++++
# 2
# ++++++++++++++++++++++++++++++
# 3
# 3 3 After
# ++++++++++++++++++++++++++++++
# 4
# 4 4 After
# ++++++++++++++++++++++++++++++
# 5
# 5 5 After
# ######

#�������������ѡ��������
#����ʹ��next(g)��һ������һ��
#����ʹ��for��������������ߣ�ֱ����ͷ
#��������ֻ�����

EG3:
import string
import time


def count():
    c=1
    for i in range(5):
        print(c)
        yield c
        c+=1

def char():
    s = string.ascii_lowercase
    for c in s:
        print(c)
        yield c

#t1 t2����generator
t1 = count()
t2 = char()

tasks = [t1, t2]
while True: # 1.�����ѭ��
    pops = []
    for i,t in enumerate(tasks): # 2.�����б�
        r = next(t,None)
        if r is None:
            # task.pop(i) #�������Ա����ڵ��� ���ȱ仯 �ǳ�Σ��������
            pops.append(i)
        else:
            print("Main next() = ",next(t))
    for i in reversed(pops): #���Ƴ�tasks���� ������pop ������������
        tasks.pop(i)
    if len(tasks) == 0: #if not tasks:
        time.sleep(1)

# next(t1)
# next(t2)
# next(t1)
# next(t2)
# next(t2)
# next(t2)
print('+' * 30)

# ����ֻ��1���߳�
# ��ʵ�����ƶ��̵߳��л�Ч�������߳��л�Ч��������Ա���ɿ���
# ����������ʵ�ֵ����ɿ�����Ա���Ƶ��л�		

��14-2-GMPģ�͡�
=====================================================================
Go����Э���У��ǳ���Ҫ�ľ���Э�̵�����scheduler��������ѯ��netpoller��

GoЭ�̵����У���������Ҫ��ɫ:
	M:Machine Thread����ϵͳ�̳߳��󡢷�װ�����д������ն�Ҫ��ϵͳ�߳������У�Э������Ҳ�Ǵ��룬Ҳ������
	G:Goroutine��GoЭ�̡��洢��Э�̵�ִ��ջ��Ϣ��״̬���������ȡ���ʼջ��СԼΪ2~4k�����Ͽ��������Goroutine��������
	P:Go1.1�汾���룬Processor�����⴦����
		����ͨ����������GOMAXPROCS��runtime.GOMAXPROCS()���ã�Ĭ��ΪCPU������
		P���������������ɲ��е�G������
		P���Լ��Ķ���(����256)��������Ŵ�ִ�е�G
		M��P��Ҫ����һ������P�����е�G�����������߳���ִ��
		
		Goroutine ---> Processor(Queue) ---> Machine
		
		Goroutine
			��û��ʹ�ö��̣߳�						        һ���У�GOMAXPROCS����P�ĸ��������Ʋ���
			��û��ʹ�����ݽṹ��						      LRQ��GRQ
			��û��ʹ��Blocking systemcall��       ��,��M��Gһ���P���P��������M
			��û���õ�non-blocking IO systemcall���У�M��G���G�Ӹ�netpoller(Linux epoll) ����
			Э�̸��һ���߳��ڶ���������û��ռ��л�  
			
			Q�����GOMAXPROCSΪ1��˵��ʲô��
			A�����̣߳�Pֻ��1�� Mֻ��1�� Gֻ��һ�� �������е�ֻ��һ�� 

	������ѯ��Netpoller�ڲ�����ʹ����IO��·���úͷ�����IO���������ǿμ������е�select��ѭ����
	GO�Բ�ͬ����ϵͳMAC(kqueue)��Linux(epoll)��Windows(iocp)�ṩ��֧�֡�

��14-3-GoTcp��
=====================================================================
package main

/*net���ײ�ʹ�÷�����IO*/
import (
	"fmt"
	"log"
	"net"
)

func main() { //main��������Э���У�mainЭ�̡�Э����û���õ��̣߳��õ���

	laddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")
	if err != nil {
		log.Panic(err)
	}

	/*�󶨶˿ں͵�ַ*/
	server, err := net.ListenTCP("tcp4", laddr) //(*net.TCPListener, error)
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()

	/*�ײ������
	  ����ʹ���ߵ�ѧϰ��������װ����������һ��ʹ�þͿ���*/
	conn, err := server.Accept() //(net.Conn, error)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer) //���ԣ�������0��ʼ����д
	if err != nil {
		log.Panic(err)
	}
	data := buffer[:n]
	fmt.Printf("Debug: %v\n", n)
	//echo server
	conn.Write(data)
}


��14-4-GoЭ��ʹ�á�
=====================================================================
[Goroutine]
package main

/*net���ײ�ʹ�÷�����IO*/
import (
	"fmt"
	"runtime"
)

func add(x, y int) int {
	var c int
	defer fmt.Printf("1 Return %d\n", c)
	defer func() { fmt.Printf("2 Return %d\n", c) }()
	fmt.Printf("add Called: x=%d,y=%d\n", x, y)
	c = x + y
	return c
}

func main() { //main��������Э���У�mainЭ�̡�Э����û���õ��̣߳��õ���
	fmt.Println("Main Start", runtime.NumGoroutine())
	// go add(4, 5)
	add(4, 5)
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// add Called: x=4,y=5
// 2 Return 9
// 1 Return 0
// Main End 1

package main

/*net���ײ�ʹ�÷�����IO*/
import (
	"fmt"
	"runtime"
	"time"
)

func add(x, y int) int {
	var c int
	defer fmt.Printf("1 Return %d\n", c)
	defer func() { fmt.Printf("2 Return %d\n", c) }()
	fmt.Printf("add Called: x=%d,y=%d\n", x, y)
	c = x + y
	return c
}

func main() { //main��������Э���У�mainЭ�̡�Э����û���õ��̣߳��õ���
	fmt.Println("Main Start", runtime.NumGoroutine())
	go add(4, 5)
	time.Sleep(2 * time.Second) //�ſ���һ��
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// Main End 2

/*�ſ�time.Sleep*/
// Main Start 1
// add Called: x=4,y=5
// 2 Return 9
// 1 Return 0
// Main End 1
/*
Q:Ϊʲô?
A:��Ϊ������Э��������add����ôgo add(4,5)��һ��û�б�Ҫ�ȵ��������زŽ���
  ���Գ���ִ����һ�д�ӡMain Exit����ʱmain�������¿�����Go��������ʱҲ������һ��Э��
  main�����������У����Գ�Ϊmain goroutine(��Э��)��
  ������Э��һ��ִ�н���������̽�������������ȴ�δִ���������Э�̡�
  ��ô�������� time.s1eep(2)����һֱ�ȣ���β��������߳����ŵȴ�Э��ִ�н�����?�ȴ���
*/

[�ȴ���]
package main

/*net���ײ�ʹ�÷�����IO*/
import (
	"fmt"
	"runtime"
	"sync"
)

func add(x, y int, wg *sync.WaitGroup) int {
	defer wg.Done() //wg��������1
	var c int
	defer fmt.Printf("1 Return %d\n", c)
	defer func() { fmt.Printf("2 Return %d\n", c) }()
	fmt.Printf("add Called: x=%d,y=%d\n", x, y)
	c = x + y
	return c
}

func main() {
	var wg sync.WaitGroup //����ȴ���
	wg.Add(1)
	fmt.Println("Main Start", runtime.NumGoroutine())
	go add(4, 5, &wg)
	wg.Wait() //����������ֵΪ0 fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// add Called: x=4,y=5
// 2 Return 9
// 1 Return 0
// Main End 1


[����Э��]
һ��Э��A�д���������һ��Э��B��A������Э�̣�B��Ϊ��Э�̡�

��Э�̽���ִ�У���Э�̲������κ�Ӱ�졣��Ȼ��Э�̽���ִ�У�Ҳ����Ը�Э����ʲôӰ�졣
����Э��û��ʲô�ر��������ϵ�����Զ������С�
ֻ����Э�����⣬���������������

EG:
package main

/*net���ײ�ʹ�÷�����IO*/
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //����ȴ���
	fmt.Println("Main Start", runtime.NumGoroutine())
	count := 6
	wg.Add(count)

	go func() {
		fmt.Println("��Э�̿�ʼ,׼��������������Э��")
		defer func() {
			wg.Done() //ע��wg��������
			fmt.Println("��Э�̽�����~~~~")
		}()

		for i := 0; i < count-1; i++ {
			go func(id int) {
				defer wg.Done()
				fmt.Printf("��Э�� %d ������\n", id)
				time.Sleep((5 * time.Second))
				fmt.Printf("��Э�� %d ����\n", id)
			}(i)
		}
	}()
	fmt.Println("Now Goroutine: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// Now Goroutine:  2
// ��Э�̿�ʼ,׼��������������Э��
// ��Э�̽�����~~~~
// ��Э�� 4 ������
// ��Э�� 3 ������
// ��Э�� 1 ������
// ��Э�� 0 ������
// ��Э�� 2 ������
// ��Э�� 2 ����
// ��Э�� 0 ����
// ��Э�� 1 ����
// ��Э�� 3 ����
// ��Э�� 4 ����
// Main End 1
// ע:Э������Ƕ����ĺ�������������������������Ƕ����д������ֻ��Ϊ�˷�����ʾ��

��14-5-WebServerʵս֮Goroutine�桿
=====================================================================
package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

var html_response_body = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>BIRKHOFF</title>
</head>
<body>
    <h1>BIRKHOFF HTML TEST PAGE -- Web Server Goroutine</h1>
</body>
</html>`

//utf-8 string

var header = `HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: %d
X-Server: BIRKHOFF
content-type: text/html; charset=utf-8
Server: birkhoffxia.com

%s`

var response = strings.ReplaceAll(fmt.Sprintf(header, len(html_response_body), html_response_body), "\n", "\r\n")

func main() {
	fmt.Sprintf(header, len(html_response_body), html_response_body)

	laddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")
	if err != nil {
		log.Panic(err)
	}

	/*�󶨶˿ں͵�ַ*/
	server, err := net.ListenTCP("tcp4", laddr) //(*net.TCPListener, error)
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()

	go func() {
		for {
			conn, err := server.Accept() //(net.Conn, error)
			if err != nil {
				log.Panic(err)
			}
			go func(conn net.Conn) {
				defer conn.Close()
				buffer := make([]byte, 4096) //���û�����
				n, err := conn.Read(buffer)  //�ɹ����ؽ����˶����ֽ�
				if n == 0 {
					fmt.Printf("�ͻ���-[%s]-�����Ͽ���\n", conn.RemoteAddr().String())
					return
				}
				if err != nil {
					log.Panic(err)
					return
				}
				data := buffer[:n] //data http������ user-agent cookie �ύ������ ��ѯ�ַ��� method
				//URL method => handler function ���ɲ�ͬ��Ӧ���ķ���
				//handler�ڲ����Σ��̶�ǩ��
				//handler(*request, response_writer)
				// /=>h1(*request,writer) url=/ writer.Writer(data)
				// / POST => h2
				// /index.go => h3
				fmt.Println(len(data), string(data))
				fmt.Printf("Debug: %v\n", n)

				conn.Write([]byte(response))
			}(conn)
		}
	}()
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("Now Processes nums: %d\n", runtime.NumGoroutine())
	}
}


��14-6-ͨ����
=====================================================================
�ǻ���ͨ��:����Ϊ0��ͨ����Ҳ��ͬ��ͨ��������ͨ�����͵�һ��Ԫ��ʱ�����û�н��ղ���������������ֱ�������ա�ͬ������ʱ�����û�����ݱ����;�����������ֱ�������ݷ��͡�
����ͨ��:  ������Ϊ0��ͨ����ͨ��������������ͨ���������ݵĲ����ᱻ����;ͨ��Ϊ�գ��ٴӸ�ͨ���������ݵĲ����ᱻ������
package main

import (
	"fmt"
	"time"
)

func main() {
	/*nil ��д �������� �޷����*/
	// var c1 chan int //��ֵ���� ��nilͨ��������
	// fmt.Println(c1 == nil) //true

	/*�ǻ���ͨ��*/
	//���� ��������0 �ǻ���ͨ�����޻��塣û���˶�ȡ������Ϳ���ֱ�����ߣ������ȶ���û�����룬��ȡ�Ῠ��ֱ����������
	// var c1 = make(chan int, 0)
	//����ͨ�������������������Լ������벻���ģ������������������Ῠ:ֻҪ��Ԫ�أ��Ϳ��Զ���ֱ��ô��Ԫ�ض���ס�������Ԫ�ؽ��룬�Ͳ�����
	var c1 = make(chan int, 1)
	fmt.Println(len(c1), cap(c1), c1) //0 0 0xc000102060
	fmt.Println("111111111111111111")
	go func() {
		fmt.Println("333333333333333333")
		c1 <- 111
		fmt.Println("����Э�̸�����")
	}()
	fmt.Println("222222222222222222")
	time.Sleep(10 * time.Second)
	fmt.Printf("Channel Read : %d\n", <-c1)
 
	fmt.Println("444444444444444444")
}

[����ͨ��] ---> ����������ģ��
<-chan type ���ֶ����ʾֻ��һ��channel�����ã�˵������ֻ����
chan <- type ���ֶ����ʾֻ��һ��channel����д��˵������ֻд��

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(ch chan<- int) { //����ֻдchannel ���� ֻд��ֻҪ��ͨ������д��������
	for {
		ch <- rand.Intn(10)
		time.Sleep(time.Second * 1)
	}
}
func consume(ch <-chan int) { //����ֻ��channel ���ѣ�ֻ����ֻҪ��ͨ�����ж���������
	for {
		t := <-ch
		fmt.Println("�����ߣ���ͨ���õ���", t)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch = make(chan int) //�޻���ͨ��������ͬ����
	// for i := 0; i < 10; i++ {
	// 	go produce(ch)
	// }
	go produce(ch)
	go consume(ch)
	wg.Wait()
}

[ͨ���ر�]
ʹ��close(ch)�ر�һ��ͨ��
ֻ�з��ͷ����ܹر�ͨ����һ��ͨ���رգ������߲����������з������ݣ�����panic
ͨ���ر�����:���߽��������������ݿ��Ե�����
ͨ���ر�
	t��ok:= <-ch ��t:= <-ch ��ͨ���ж�ȡ����
	���������ȴ�ͨ���е����ݵĽ����ߣ�����ͨ�����رգ������߽�������������ȡ����ʧ��okΪfalse��������ֵ
	��������Ȼ���Է��ʹرյ�ͨ������������
		���ͨ���ڻ���ʣ�����ݣ�okΪtrue����������
		���ͨ����ʣ������ݱ������ˣ��������ղ�������okΪfalse��������ֵ
�Ѿ��رյ�ͨ�������ٴιر���panic����˲�Ҫ�ظ��ر�

EG��
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(ch chan<- int) { //����ֻдchannel ���� ֻд��ֻҪ��ͨ������д��������
	for i := 0; i < 3; i++ {
		ch <- rand.Intn(10)
		fmt.Println("---------------")
	}
}

func consume(ch <-chan int, wg *sync.WaitGroup) { //����ֻ��channel ���ѣ�ֻ����ֻҪ��ͨ�����ж���������
	defer wg.Done()
	time.Sleep(time.Second * 10)
	for {
		t, ok := <-ch
		if ok {
			fmt.Println("�����ߣ���ͨ���õ���", t)
		} else {
			fmt.Println("Channel�ر���", ok, t)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch = make(chan int, 8) //�޻���ͨ��������ͬ����

	go produce(ch)
	time.Sleep(time.Second)
	go consume(ch, &wg)
	time.Sleep(5 * time.Second)
	close(ch)
	wg.Wait()
}

// ---------------
// ---------------
// ---------------
// �����ߣ���ͨ���õ��� 2
// �����ߣ���ͨ���õ��� 5
// �����ߣ���ͨ���õ��� 7
// Channel�ر��� false 0

[ͨ������]
1��nilͨ��
	���͡����ա�����������
2������ġ�δ�رյ�ͨ��
	�൱��һ������Ԫ�ص�ͨ�����������꣬�����ڵ���һ��Ԫ�ص���

package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5) //���壬δ�ر�ͨ��
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)
	c1 <- 111
	c1 <- 222
	c1 <- 333
	fmt.Println(<-c1, "@@@")

	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //���������
}

3������ġ��رյ�ͨ��
	�رպ�ͨ�������ڽ����µ�Ԫ�أ���ô�൱�ڱ������޸�Ԫ��������������ͽ����ˡ�
EG��
package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5) //���壬δ�ر�ͨ��
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)
	c1 <- 111
	c1 <- 222
	c1 <- 333
	fmt.Println(<-c1, "@@@")

	close(c1) //�ر�ͨ���������ٽ�����
	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //���������
}

// cl:0,5,0xc000150000n111 @@@
// 222 ###
// 333 ###
// ========================

4���ǻ��塢δ�ر�ͨ��
	 �൱��һ������Ԫ�ص�ͨ�����������꣬�����ڵ���һ��Ԫ�ص��
EG:
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 5) //���壬δ�ر�ͨ��
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)

	go func() {
		count := 1
		for {
			time.Sleep(3 * time.Second)
			c1 <- count
			count++
		}
	}()

	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //���������
}
	package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 5) //���壬δ�ر�ͨ��
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)

	go func() {
		count := 1
		for {
			time.Sleep(3 * time.Second)
			c1 <- count
			count++
		}
	}()

	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //���������
}
// cl:0,5,0xc000120000n1 ###
// 2 ###
// 3 ###
// 4 ###
// 5 ###
// 6 ###
// 7 ###
// 8 ###
// 9 ###
// 10 ###
// 11 ###
// 12 ###

5.�ǻ��塢�ر�ͨ��
	�رպ�ͨ�������ڽ����µ�Ԫ�أ���ô�൱�ڱ������޸�Ԫ��������������ͽ����ˡ�
EG:
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) //�ǻ��壬δ�ر�ͨ��
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)

	go func() {
		defer close(c1)
		count := 1
		for {
			time.Sleep(3 * time.Second)
			c1 <- count
			count++
		}
	}()

	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //���������
}

��nilͨ����
	δ�ر�ͨ������ͬһ�����޵���������һֱ����ͨ����Ԫ�أ�û��Ԫ�ؾ�����
	�ѹر�ͨ���������ܼ����µ�Ԫ�أ������굱ǰͨ���ڵ�Ԫ�أ�������0��Ԫ�أ�Ȼ���������	

[��ʱ��]
EG��time.NewTicker(2 * time.Second)
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)
	for {
		fmt.Println(<-t.C) //ͨ��ÿ����2��ͽ���һ��
	}
}

// 2024-07-02 14:40:53.5397943 +0800 CST m=+2.015958201
// 2024-07-02 14:40:55.5419405 +0800 CST m=+4.018104401
// 2024-07-02 14:40:57.5273588 +0800 CST m=+6.003522701
// 2024-07-02 14:40:59.5287525 +0800 CST m=+8.004916401
// 2024-07-02 14:41:01.537747 +0800 CST m=+10.013910901
// 2024-07-02 14:41:03.5273134 +0800 CST m=+12.003477301
// 2024-07-02 14:41:05.5289956 +0800 CST m=+14.005159501
// 2024-07-02 14:41:07.5267152 +0800 CST m=+16.002879101

EG��time.NewTimer(2 * time.Second)

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(2 * time.Second)
	for {
		fmt.Println(<-t.C) //ͨ������2���ֻ�ܽ���һ��
	}
}

	
[ͨ������]
	channel���ˣ�������д:channel���ˣ���������������Ϊ0��ͨ���������Ϊ0��Ԫ�ؾ����������˵�ǰЭ��֮��ύ��CPU��ȥִ������Э�̣�
		ϣ������Э�̰����Լ����������main���������ˣ��������̽����ˡ�
	�����mainЭ���У�ִ���������ʱ���������ܣ�����Ѿ�û��������Э�̿���ִ�У�ֻʣ��Э���Լ������������ˣ����Լ����Լ�ɱ������һ��fatalerror deadlock
EG��
package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int) //�ǻ��壬δ�ر�ͨ��
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)
	c1 <- 111 //��ǰЭ�������������ܽ⣬����
}

**
	1.���ͨ����������mainЭ���з��������Ƿ�������Э���У���Э�̻���������ţ�Ҳ���ܷ���������������������mainЭ����һ��ֵ�õȴ���ϣ����
		���������ܰ���ʶ������������������κ�Э�̰�����Э�̽������״̬����ô��ʵ�ϸ���Э�̽����������Ѿ������ˡ�
	2.������Σ�����ܻᵼ�½��̻��ţ���ʵ����ĳЩЭ��δ����������������Ӧ�������õı���ϰ�ߣ������������ĳ��֡�
Viedo Date: 2023/11/18
Made By:	BIRKHOFF
Date:	2024-07-01


【14-1-协程】
=====================================================================
基本概念
	不需要利用多线程，在一个线程中，实现任务人为控制的切换，交替运行的效果
	在一个线程内，每一个任务类似于多线程切换交替运行的效果，起个名字叫Go-routine
		人为控制，开发人员写代码，用户态控制
	
	协程可以由多线程吗？可以有，不排斥
		
	核心
	在python中，提供了yield语句，可以让函数在yield处暂停下来。而yield是开发人员写的，可以由开发人员自己决定函数在何处 暂停 执行
		暂停了，还是要恢复的
	
	多线程的每一个线程要执行target目标执行函数，线程是干活的，线程去执行指定函数的指令，也可以暂停线程，函数可以被暂停，但是这暂停谁做的? OS，开发人员不能干预
		暂停也要恢复
		内核态

[问题]
1.有了协程，请问还需要多线程么？
	A：可以有多线程
2.有了协程，会不会有线程的切换？
	A:当然会，线程的切换你管不着
		OS管着n个线程，不管你的程序用几个线程，还是其他进程的线程 
		用了协程难道本程序就不可以用多线程了么？可以用
3.协程也是代码，操作系统的最小单位是协程么？
	A:线程，协程也得在线程中运行

弊端：
	一旦我们多任务协程只在一个线程中运行，其中一个协程阻塞住了，没有碰到yield没有让出控制权，交替不能进行了

EG1：
def count():
    c=1
    for i in range(5):
        print(c)
        c+=1
x = count() #如果没有异常，没有return之前，函数返回了么？ 没有
#函数没有返回能执行下一句print么？ 不能
#在同一个线程中，顺序执行，串行
print(type(x),x)
print("######")

# 1
# 2
# 3
# 4
# 5
# <class 'NoneType'> None
# ######

EG2：
#函数能在中途暂停
def count():
    c=1
    for i in range(5):
        print(c)
        yield c #如果有了yield，此函数将成为特殊的生成器函数
        print('+' * 30)
        c+=1
x = count() #不是真执行，而是立即构造一个生成器对象返回
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
# next(x) #拨一下，会在yield处停下来，出现了类似return的效果
# # <class 'generator'> <generator object count at 0x000001B8731FDEB8>
# # 1
# # ######
# next(x) #拨一下。从上次暂停处向后接替执行，直到函数结束或再次yield
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

#生成器对象就是选代器对象
#可以使用next(g)走一步，拨一下
#可以使用for连续驱动它向后走，直到尽头
#驱动方向只能向后

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

#t1 t2都是generator
t1 = count()
t2 = char()

tasks = [t1, t2]
while True: # 1.任务大循环
    pops = []
    for i,t in enumerate(tasks): # 2.任务列表
        r = next(t,None)
        if r is None:
            # task.pop(i) #对于线性表正在迭代 长度变化 非常危险有问题
            pops.append(i)
        else:
            print("Main next() = ",next(t))
    for i in reversed(pops): #待移除tasks索引 倒过来pop 否则索引超界
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

# 代码只有1个线程
# 有实现类似多线程的切换效果，多线程切换效果开发人员不可控制
# 但我们线性实现的是由开发人员控制的切换		

【14-2-GMP模型】
=====================================================================
Go语言协程中，非常重要的就是协程调度器scheduler和网络轮询器netpoller。

Go协程调度中，有三个重要角色:
	M:Machine Thread，对系统线程抽象、封装。所有代码最终都要在系统线程上运行，协程最终也是代码，也不例外
	G:Goroutine，Go协程。存储了协程的执行栈信息、状态和任务函数等。初始栈大小约为2~4k理论上开启百万个Goroutine不是问题
	P:Go1.1版本引入，Processor，虚拟处理器
		可以通过环境变量GOMAXPROCS或runtime.GOMAXPROCS()设置，默认为CPU核心数
		P的数量决定着最大可并行的G的数量
		P有自己的队列(长度256)，里面放着待执行的G
		M和P需要绑定在一起，这样P队列中的G才能真正在线程上执行
		
		Goroutine ---> Processor(Queue) ---> Machine
		
		Goroutine
			有没有使用多线程？						        一定有，GOMAXPROCS控制P的个数，控制并发
			有没有使用数据结构？						      LRQ、GRQ
			有没有使用Blocking systemcall？       有,连M和G一起和P解绑，P另找其他M
			有没有用到non-blocking IO systemcall？有，M和G解绑，G扔给netpoller(Linux epoll) 监听
			协程概念：一个线程内多个任务在用户空间切换  
			
			Q：如果GOMAXPROCS为1，说明什么？
			A：多线程，P只有1个 M只有1个 G只有一个 真正运行的只有一个 

	网络轮询器Netpoller内部就是使用了IO多路复用和非阻塞IO，类似我们课件代码中的select的循环。
	GO对不同操作系统MAC(kqueue)、Linux(epoll)、Windows(iocp)提供了支持。

【14-3-GoTcp】
=====================================================================
package main

/*net包底层使用非阻塞IO*/
import (
	"fmt"
	"log"
	"net"
)

func main() { //main函数跑在协程中，main协程。协程有没有用到线程？用到了

	laddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")
	if err != nil {
		log.Panic(err)
	}

	/*绑定端口和地址*/
	server, err := net.ListenTCP("tcp4", laddr) //(*net.TCPListener, error)
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()

	/*底层非阻塞
	  减少使用者的学习附带，封装的像阻塞的一样使用就可以*/
	conn, err := server.Accept() //(net.Conn, error)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer) //可以，从索引0开始覆盖写
	if err != nil {
		log.Panic(err)
	}
	data := buffer[:n]
	fmt.Printf("Debug: %v\n", n)
	//echo server
	conn.Write(data)
}


【14-4-Go协程使用】
=====================================================================
[Goroutine]
package main

/*net包底层使用非阻塞IO*/
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

func main() { //main函数跑在协程中，main协程。协程有没有用到线程？用到了
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

/*net包底层使用非阻塞IO*/
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

func main() { //main函数跑在协程中，main协程。协程有没有用到线程？用到了
	fmt.Println("Main Start", runtime.NumGoroutine())
	go add(4, 5)
	time.Sleep(2 * time.Second) //放开这一句
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// Main End 2

/*放开time.Sleep*/
// Main Start 1
// add Called: x=4,y=5
// 2 Return 9
// 1 Return 0
// Main End 1
/*
Q:为什么?
A:因为会启动协程来运行add，那么go add(4,5)这一句没有必要等到函数返回才结束
  所以程序执行下一行打印Main Exit。这时main函数无事可做，Go程序启动时也创建了一个协程
  main函数运行其中，可以称为main goroutine(主协程)。
  但是主协程一旦执行结束，则进程结束，根本不会等待未执行完的其它协程。
  那么，除了像 time.s1eep(2)这样一直等，如何才能让主线程优雅等待协程执行结束呢?等待组
*/

[等待组]
package main

/*net包底层使用非阻塞IO*/
import (
	"fmt"
	"runtime"
	"sync"
)

func add(x, y int, wg *sync.WaitGroup) int {
	defer wg.Done() //wg计数器减1
	var c int
	defer fmt.Printf("1 Return %d\n", c)
	defer func() { fmt.Printf("2 Return %d\n", c) }()
	fmt.Printf("add Called: x=%d,y=%d\n", x, y)
	c = x + y
	return c
}

func main() {
	var wg sync.WaitGroup //定义等待组
	wg.Add(1)
	fmt.Println("Main Start", runtime.NumGoroutine())
	go add(4, 5, &wg)
	wg.Wait() //阻塞计数器值为0 fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// add Called: x=4,y=5
// 2 Return 9
// 1 Return 0
// Main End 1


[父子协程]
一个协程A中创建了另外一个协程B，A称作父协程，B称为子协程。

父协程结束执行，子协程不会有任何影响。当然子协程结束执行，也不会对父协程有什么影响。
父子协程没有什么特别的依赖关系，各自独立运行。
只有主协程特殊，它结束程序结束。

EG:
package main

/*net包底层使用非阻塞IO*/
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //定义等待组
	fmt.Println("Main Start", runtime.NumGoroutine())
	count := 6
	wg.Add(count)

	go func() {
		fmt.Println("父协程开始,准备创建并启动子协程")
		defer func() {
			wg.Done() //注意wg的作用域
			fmt.Println("父协程结束了~~~~")
		}()

		for i := 0; i < count-1; i++ {
			go func(id int) {
				defer wg.Done()
				fmt.Printf("子协程 %d 运行中\n", id)
				time.Sleep((5 * time.Second))
				fmt.Printf("子协程 %d 结束\n", id)
			}(i)
		}
	}()
	fmt.Println("Now Goroutine: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Main End", runtime.NumGoroutine())
}

// Main Start 1
// Now Goroutine:  2
// 父协程开始,准备创建并启动子协程
// 父协程结束了~~~~
// 子协程 4 运行中
// 子协程 3 运行中
// 子协程 1 运行中
// 子协程 0 运行中
// 子协程 2 运行中
// 子协程 2 结束
// 子协程 0 结束
// 子协程 1 结束
// 子协程 3 结束
// 子协程 4 结束
// Main End 1
// 注:协程最好是独立的函数，而不是上例这样将函数嵌套着写。上例只是为了方便演示。

【14-5-WebServer实战之Goroutine版】
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

	/*绑定端口和地址*/
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
				buffer := make([]byte, 4096) //设置缓冲区
				n, err := conn.Read(buffer)  //成功返回接受了多少字节
				if n == 0 {
					fmt.Printf("客户端-[%s]-主动断开了\n", conn.RemoteAddr().String())
					return
				}
				if err != nil {
					log.Panic(err)
					return
				}
				data := buffer[:n] //data http请求报文 user-agent cookie 提交的数据 查询字符串 method
				//URL method => handler function 生成不同响应报文返回
				//handler内部传参，固定签名
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


【14-6-通道】
=====================================================================
非缓冲通道:容量为0的通道，也叫同步通道。这种通道发送第一个元素时，如果没有接收操作就立即阻塞，直到被接收。同样接收时，如果没有数据被发送就立即阻塞，直到有数据发送。
缓冲通道:  容量不为0的通道。通道已满，再往该通道发送数据的操作会被阻塞;通道为空，再从该通道接收数据的操作会被阻塞。
package main

import (
	"fmt"
	"time"
)

func main() {
	/*nil 读写 都会阻塞 无法解除*/
	// var c1 chan int //零值定义 ，nil通道不可用
	// fmt.Println(c1 == nil) //true

	/*非缓冲通道*/
	//长度 容量都是0 非缓冲通道，无缓冲。没有人读取，塞入就卡，直到读走；有人先读，没人塞入，读取会卡，直到有人塞入
	// var c1 = make(chan int, 0)
	//缓冲通道。容量不满，都可以继续塞入不卡的，如果已满，继续塞入会卡:只要有元素，就可以读，直到么有元素而卡住，如果有元素进入，就不卡了
	var c1 = make(chan int, 1)
	fmt.Println(len(c1), cap(c1), c1) //0 0 0xc000102060
	fmt.Println("111111111111111111")
	go func() {
		fmt.Println("333333333333333333")
		c1 <- 111
		fmt.Println("我子协程干完了")
	}()
	fmt.Println("222222222222222222")
	time.Sleep(10 * time.Second)
	fmt.Printf("Channel Read : %d\n", <-c1)
 
	fmt.Println("444444444444444444")
}

[单向通道] ---> 生产消费者模型
<-chan type 这种定义表示只从一个channel里面拿，说明这是只读的
chan <- type 这种定义表示只往一个channel里面写，说明这是只写的

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(ch chan<- int) { //定义只写channel 生产 只写，只要该通道具有写能力就行
	for {
		ch <- rand.Intn(10)
		time.Sleep(time.Second * 1)
	}
}
func consume(ch <-chan int) { //定义只读channel 消费，只读，只要该通道具有读能力就行
	for {
		t := <-ch
		fmt.Println("消费者，从通道拿到了", t)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch = make(chan int) //无缓冲通道，被迫同步了
	// for i := 0; i < 10; i++ {
	// 	go produce(ch)
	// }
	go produce(ch)
	go consume(ch)
	wg.Wait()
}

[通道关闭]
使用close(ch)关闭一个通道
只有发送方才能关闭通道，一旦通道关闭，发送者不能再往其中发送数据，否则panic
通道关闭作用:告诉接收者再无新数据可以到达了
通道关闭
	t，ok:= <-ch 或t:= <-ch 从通道中读取数据
	正在阻塞等待通道中的数据的接收者，由于通道被关闭，接收者将不再阻塞，获取数据失败ok为false，返回零值
	接收者依然可以访问关闭的通道而不阻塞。
		如果通道内还有剩余数据，ok为true，接收数据
		如果通道内剩余的数据被拿完了，继续接收不阻塞，ok为false，返回零值
已经关闭的通道，若再次关闭则panic，因此不要重复关闭

EG：
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(ch chan<- int) { //定义只写channel 生产 只写，只要该通道具有写能力就行
	for i := 0; i < 3; i++ {
		ch <- rand.Intn(10)
		fmt.Println("---------------")
	}
}

func consume(ch <-chan int, wg *sync.WaitGroup) { //定义只读channel 消费，只读，只要该通道具有读能力就行
	defer wg.Done()
	time.Sleep(time.Second * 10)
	for {
		t, ok := <-ch
		if ok {
			fmt.Println("消费者，从通道拿到了", t)
		} else {
			fmt.Println("Channel关闭了", ok, t)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch = make(chan int, 8) //无缓冲通道，被迫同步了

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
// 消费者，从通道拿到了 2
// 消费者，从通道拿到了 5
// 消费者，从通道拿到了 7
// Channel关闭了 false 0

[通道遍历]
1、nil通道
	发送、接收、遍历都阻塞
2、缓冲的、未关闭的通道
	相当于一个无限元素的通道，迭代不完，阻塞在等下一个元素到达

package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5) //缓冲，未关闭通道
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)
	c1 <- 111
	c1 <- 222
	c1 <- 333
	fmt.Println(<-c1, "@@@")

	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //看不到这句
}

3、缓冲的、关闭的通道
	关闭后，通道不能在进入新的元素，那么相当于遍历有限个元素容器，遍历完就结束了。
EG：
package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5) //缓冲，未关闭通道
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)
	c1 <- 111
	c1 <- 222
	c1 <- 333
	fmt.Println(<-c1, "@@@")

	close(c1) //关闭通道，不许再进数据
	for v := range c1 {
		fmt.Println(v, "###")
	}
	fmt.Println("========================") //看不到这句
}

// cl:0,5,0xc000150000n111 @@@
// 222 ###
// 333 ###
// ========================

4、非缓冲、未关闭通道
	 相当于一个无限元素的通道，迭代不完，阻塞在等下一个元素到达。
EG:
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 5) //缓冲，未关闭通道
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
	fmt.Println("========================") //看不到这句
}
	package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 5) //缓冲，未关闭通道
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
	fmt.Println("========================") //看不到这句
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

5.非缓冲、关闭通道
	关闭后，通道不能在进入新的元素，那么相当于遍历有限个元素容器，遍历完就结束了。
EG:
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) //非缓冲，未关闭通道
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
	fmt.Println("========================") //看不到这句
}

除nil通道外
	未关闭通道，如同一个无限的容器，将一直迭代通道内元素，没有元素就阻塞
	已关闭通道，将不能加入新的元素，迭代完当前通道内的元素，哪怕是0个元素，然后结束迭代	

[定时器]
EG：time.NewTicker(2 * time.Second)
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)
	for {
		fmt.Println(<-t.C) //通道每阻塞2秒就接受一次
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

EG：time.NewTimer(2 * time.Second)

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(2 * time.Second)
	for {
		fmt.Println(<-t.C) //通道阻塞2秒后只能接受一次
	}
}

	
[通道死锁]
	channel满了，就阻塞写:channel空了，就阻塞读。容量为0的通道可以理解为0个元素就满了阻塞了当前协程之后会交出CPU，去执行其他协程，
		希望其他协程帮助自己解除阻塞。main函数结束了，整个进程结束了。
	如果在main协程中，执行语句阻塞时，环顾四周，如果已经没有其他子协程可以执行，只剩主协程自己，解锁无望了，就自己把自己杀掉，报一个fatalerror deadlock
EG：
package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int) //非缓冲，未关闭通道
	fmt.Printf("cl:%d,%d,%vn", len(c1), cap(c1), c1)
	c1 <- 111 //当前协程阻塞，无人能解，死锁
}

**
	1.如果通道阻塞不在main协程中发生，而是发生在子协程中，子协程会继续阻塞着，也可能发生死锁。但是由于至少main协程是一个值得等待的希望，
		编译器不能帮你识别出死锁。如果真的无任何协程帮助该协程解除阻塞状态，那么事实上该子协程解锁无望，已经死锁了。
	2.死锁的危害可能会导致进程活着，但实际上某些协程未真正工作而阻塞，应该有良好的编码习惯，来减少死锁的出现。
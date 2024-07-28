Viedo Date: 2023/11/4
Made By:	BIRKHOFF
Date:	2024-06-28

【12-1-Python安装和基础语法】
=====================================================================


【12-2-并发原理】
=====================================================================
并发
	高并发，高的并发，高低相对的，针对当前可利用的资源(CPU、内存、带宽)
	并发的概念:
		currencyCon并和发，一起发生一段时间(你观察的时间段)内，有很多事情陆续或者同时发生
	并行的概念:
		parrallelism
		强调同一时刻，多个事情同时在发生
	晚上0点，上高速免费，0点之前，2万车辆都堵在闸口等待上高速。高并发，一段时间内你有很多事情要处理
		只开一个闸口，没有并行的可能，串行。请问，高并发能不能用串行解决?能，但是要注意每件事处理的速度，吞吐量
			硬盘技术发展:并行口IDE、sata串行口
			串行快还是并行快?
		开5个闸口，并行
			同一时刻，最多同时通过5辆车请问，高并发能不能用并行解决?可以，但是需要更多的资源，增加了成本
		排队就是队列，FIFO争抢，有的时候会产生饥饿
		Cache 缓存
			map来理解网站的静态内容可以cache
			CDN，缓存 就近
		
		并行、串行、争抢、队列、缓存都是解决并发的方案，综合运用多种手段来解决高并发问题，但注意成本

【12-3-进程线程和状态】
=====================================================================
进程
	多个进程
	进程间相互隔离，国家与国家，国与国可以通信，序列化和反序列化进程认为内存都是他家的，端口都是他家的，CPU是他家的
	
	线程
		多个线程
		每个线程拥绑指令(源代码编译后)
		CPU运行指令，CPu运行最小调度单元线程
		
		线程间通信
			省与省
				每个省都有自己的资源，不共享的，线程有自己的栈(函数压栈)
			共用同一个进程的内部资源
			没有序列化反序列化

优先级
	只能说，对进程、线程来说，一段时间内观察，分配的cPu时间占比更大
	不能这么说，优先级越高，CPU只运行我这个进程或线程

线程状态
	Ready，runable 可以被CPU运行的线程
	Running，正在CPU上运行，当前某个cPu核心上正在运行这个线程的指令
	Blocked，阻寒，线程被运行的过程中，卡住了
	Terminated,终止


【12-4-多线程开发和分析】
=====================================================================
线程启动
EG:
import threading
#最简单的线程程序
def worker():
    print("I'm working")
    print('Done')

t = threading.Thread(target=worker,name='worker') #创建线程对象
# target=worker方式为关键字传参，按名称对应
# Python中还有按照位置对应传参，按照顺序依次对应
t.start()#启动真正的线程

#I'm working
#Done

[线程无退出]
EG:
import threading
#最简单的线程程序
def worker():
    print("I'm working")
    print('Done')

t = threading.Thread(target=worker,name='worker') #创建线程对象
# target=worker方式为关键字传参，按名称对应
# Python中还有按照位置对应传参，按照顺序依次对应
t.start()#启动真正的线程

print("Process ID: ",t.ident) //查看进程ID
# I'm working
# Process ID:  34204
# Done

[线程退出]
Python没有提供线程退出的方法，线程在下面情况时退出
	1、线程函数内语句执行完毕 
	2、线程函数中抛出未处理的异常

#线程传参和函数传参没什么区别，本质上就是函数传参。

EG：
import threading
import time
def worker(x,y):
    s = "{} + {} = {}".format(x,y,x+y)
    print(s,threading.currentThread().ident)

t1 = threading.Thread(target=worker,name='worker',args=(4,5))
t1.start()
time.sleep(2)

[多线程编程]
顾名思义，多个线程，一个进程中如果有多个线程运行，就是多线程，实现并发。
想想下面有几个线程运行?
import threading
import time
import string

#最简单的线程程序
def count():
    c=1
    while True:
        time.sleep(5)
        print("count = ",c)
        c+=1
        print("count()中",threading.current_thread())

def char():
    s = string.ascii_lowercase
    for c in s:
        time.sleep(2)
        print("char = ",c)
        print("char()中",threading.current_thread())

print("count()中",threading.current_thread())
t1 = threading.Thread(target=count,name="count")
t2 = threading.Thread(target=char,name="char")
print("{},{}".format(t1.ident,threading.current_thread()))
t1.start()
t2.start()
print("{},{}".format(t1.ident,threading.current_thread()))

#3个线程，count、char、主线程。
#调整time.sleep(2)为20或更大
# 请问这个函数所在的执行线程怎么了?请问谁卡住(阻塞)了?谁又在运行?这和线程状态有什么关系?

#注:Python中有一个GIL全局解释器锁，大家初学可以忽略它，它对阻塞性10其实影响不大。
#重点:大家要从例子中找到并发执行的感觉，这对理解并发包括Goroutine都大有益处。


【12-5-socket开发】
=====================================================================
import socket

#TCP服务端编程
server = socket.socket() #创建socket对象
laddr = ('0.0.0.0',9999) #地址和端口的元组
server.bind(laddr) #绑定
server.listen(1024) #监听

#等待建立连接的客户端
#server.accept() 阻塞
conn,raddr = server.accept()  #三次握手建立好的连接，返回一个新的socket和元组(对方的地址)
print(conn,type(conn)) #负责客户端连接的socket对象
print(raddr,type(conn)) #对方IP地址和端口
print(conn.getpeername(),conn.getsockname()) #通过socket获取对端地址或本地地址

data = conn.recv(4096) #接受客户端信息
print(type(data),data)
conn.send(b"Hello BIRKHOFF")

conn.close()
server.close()
print('-' * 30)
Viedo Date: 2023/11/11
Made By:	BIRKHOFF
Date:	2024-06-29


【13-1-多线程加BIO】
=====================================================================
swiddler：https://swiddler.com/
#打开使用Swiddler进行模拟 发送收发消息
EG:
import socket

#TCP服务端编程
server = socket.socket() #创建socket对象
laddr = ('0.0.0.0',9999) #地址和端口的元组
server.bind(laddr) #绑定
server.listen(1024) #监听

#等待建立连接的客户端
#server.accept() 阻塞
conn,raddr = server.accept()  #三次握手建立好的连接，返回一个新的socket和元组(对方的地址)
print(1,conn,type(conn)) #负责客户端连接的socket对象
print(2,raddr,type(raddr)) #对方IP地址和端口
print(conn.getpeername(),conn.getsockname()) #通过socket获取对端地址或本地地址

data = conn.recv(4096) #接受客户端信息
print(type(data),data)
conn.send(b"Hello BIRKHOFF")

conn.close()
server.close()

print('-' * 30)

# 1 <socket.socket fd=424, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 9999), raddr=('127.0.0.1', 50667)> <class 'socket.socket'>
# 2 ('127.0.0.1', 50667) <class 'tuple'>
# ('127.0.0.1', 50667) ('127.0.0.1', 9999)
# <class 'bytes'> b'HEYE'

[多线程] Thread Per connect
EG：
import socket
import threading
import time


# 多线程 + 阻塞ＩＯ　Ｍｕｌｔｉｔｈｒｅａｄ＋　ＢｌｏｃｋＩＯ

def recv(conn: socket.socket, raddr):
    try:
        for i in range(3):
            data = conn.recv(1024)  # 阻塞 等待客户端消息来 Read
            print(3, data, type(data), threading.current_thread().name)
            if not data:
                print("Bye~~~", raddr)
                return
            conn.send(data)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recover抓住了异常并处理掉了，没有了异常
    finally:
        conn.close()


def accept(server):
    i = 1
    while True:
        # 分配一个新的服务员
        conn, raddr = server.accept()  # 三次握手建立好的连接，返回一个新的socket和元组(对方的地址)
        print(1, conn, type(conn))  # 负责客户端连接的socket对象
        print(2, raddr, type(raddr))  # 对方IP地址和端口
        # "T-线程ID"
        threading.Thread(target=recv, name="T-{}".format(i), args=(conn,raddr)).start()
        i += 1
        print('+' * 30)


if __name__ == '__main__':
    server = socket.socket()  # 创建socket对象
    laddr = ('0.0.0.0', 9999)  # 地址和端口的元组
    server.bind(laddr)  # 绑定
    server.listen(1024)  # 监听

    # accept() 也创建一个线程 这样下面print可以执行
    threading.Thread(target=accept, name="ac", args=(server,)).start()

    print('~' * 30)
    while True:
        time.sleep(5)
        print(threading.enumerate()) #每5秒查看进程还有多少

报文：连接上了4个客户端 发送1-1 ~ 4-3  查看是否发送3个消息自动关闭
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
1 <socket.socket fd=428, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 9999), raddr=('127.0.0.1', 52022)> <class 'socket.socket'>
2 ('127.0.0.1', 52022) <class 'tuple'>
++++++++++++++++++++++++++++++
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>]
1 <socket.socket fd=448, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 9999), raddr=('127.0.0.1', 52023)> <class 'socket.socket'>
2 ('127.0.0.1', 52023) <class 'tuple'>
++++++++++++++++++++++++++++++
1 <socket.socket fd=468, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 9999), raddr=('127.0.0.1', 52025)> <class 'socket.socket'>
2 ('127.0.0.1', 52025) <class 'tuple'>
++++++++++++++++++++++++++++++
1 <socket.socket fd=484, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 9999), raddr=('127.0.0.1', 52026)> <class 'socket.socket'>
2 ('127.0.0.1', 52026) <class 'tuple'>
++++++++++++++++++++++++++++++
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'1-1\r\n' <class 'bytes'> T-1
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'1-2' <class 'bytes'> T-1
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'2-1' <class 'bytes'> T-2
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'2-2' <class 'bytes'> T-2
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'3-1' <class 'bytes'> T-3
3 b'3-2' <class 'bytes'> T-3
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'4-1' <class 'bytes'> T-4
3 b'4-2' <class 'bytes'> T-4
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>, <Thread(T-4, started 36748)>]
3 b'4-3' <class 'bytes'> T-4
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>]
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>, <Thread(T-3, started 36340)>]
3 b'3-3' <class 'bytes'> T-3
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>]
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>, <Thread(T-2, started 33092)>]
3 b'2-3' <class 'bytes'> T-2
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>, <Thread(T-1, started 25032)>]
3 b'1-3' <class 'bytes'> T-1
[<_MainThread(MainThread, started 4324)>, <Thread(ac, started 30036)>] 

【13-2-WebServer原理和实现】
=====================================================================
Web Server 后端
	httpd，nginx
	HTTP协议
		基于文本
		基于TCP协议之上的应用层协议
		应用程序使用的
	HTML
		基于文本的
		描述图像、文字、格式
		文本文件，可以放在磁盘上
		图片文件，静态文件
	Js源代码
		文本文件，放在磁盘上
		封装函数、封装库第三方包，包管理
		框架
			React、Vue2 3
					router 前端路由浏览器端
						浏览器中url变化（没有发起对后台server的http请求），js拦截这个变化，js对应的回调函数做处理，导致页面变化
						后端通信，判断，异步请求HTTP Ajax() 
	动态生成
		HTML、JS文件能不能由代码动态生成这些源代码文本内容动态生成
		svg是文本，浏览器帮你绘制
		image，验证码
		
Browser浏览器(客户端、前端)
	HTTP协议https ssl tsl
		URL 全球 统一资源定位符，server端定位资源后，把资源的内容通过HTTp协议返回到浏览器端
		前端不管你后端是直接从文件读取(静态文件)还是后端通过代码动态生成内容(动态)
		和server通信，两个进程c/s通信，B/s本质就是c/s
		v1
			浏览器建立连接到服务端，发起一个请求，服务端响应，连接断开
			TCP连接断开，重建，很慢也很耗资源
		v1.1 keepalive 保持TCP连接一段时间,一段时间后空闲连接断开
		
		v2
	HTML解析(设计师)
		形成DOM
		Javascript操作DOM，实现动效
		渲染render
	CSS 控制HTML渲染
	JavaScript引整
		解释执行Js源代码
		Ecmascript 标准
		ES3、ES5、ES6+
		开发

Browser --> http请求报文(同步、异步Ajax) --> HTTP WEB Server(TCP listen) --> url method(router定位资源 路由，后端资源定位)
	--> 404 Not Found; 403
	--> 1 定位文件系统中的文件open read 内容
	--> 2 动态生成内容动态网页技术ASP JSPPHP
			(Go/Python/PHP/Java/c/C#/C++/Ruby/Javascript......)
			Database中查询内容回来
			凑出来文本(内容)

	-->响应response报文(含着内容) --> Browserl (HTML源码、Js源码、image) 
	html解析它，发现里面还有很多资源，要为每一个资源都要建立一个连接用完就断

OpenAI开发者大会，你我可能没有低级工作了

浏览器 https://www.bing.com/ 首页(静态、动态) -> Server 80 - > url -> response报文 HTML
	v1.1
HTML解析
	img src="/a.jpg"图片 会再次发起http请求
	脚本 src ="/js/t1.js" 会再次发起http请求
	n个资源都要发起请求，连接怎么办? 开三个TCP连接 keepalive

HTTP协议
	请求报文
	响应报文
		每一行\r\n结束
		第一行
			HTTP/1.1 200 OK 协议及版本 状态码 说明 
			状态码
				1xx客户端错误;
				2xx通信正常
					200 OK
					201 CREATED 资源创建成功 Restful
					204 No Content无内容返回 但正常，删除资源成功
				3xx跳转location
					301
					302
				4xx权限、资源错误
					401
					402
					403
					404
					405
				5xx 服务器错误
					500
					502nginx反向代理 上游服务器错误，不是ningx出错
			第二行开始
				xxx: yyy key-value对
				Connection：keep-alive
				Content-Length:1292 告诉浏览器接受response body的长度 字节
				Content-Type:image/png MIME类型
				content-type:text/html;charset=utf-8

EG:基于WebServer 实现了多线程加阻塞IO版本
import socket
import threading
import time

html_response_body = """\
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>BIRKHOFF</title>
</head>
<body>
    <h1>BIRKHOFF HTML TEST PAGE -- {}</h1>
</body>
</html>""".format("MultiThread + BIO").encode() #UTF-8 Bytes

response = """\
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: {}
X-Server: BIRKHOFF
content-type: text/html; charset=utf-8

""".format(len(html_response_body)).encode() + html_response_body

def recv(conn: socket.socket, raddr):
    try:
        data = conn.recv(1024)  # 阻塞 等待客户端消息来 Read
        # data就是请求报文。按道理这里要做请求报文的解析，解析库
        # 就应该按照method Path等信息，路由 找到静态文件读取内容返回，或动态生成内容
        print(3, data, type(data), threading.current_thread().name)
        if not data:
            print("Bye~~~", raddr)
            return
        # 省略几万行代码，响应报文封装
        conn.send(response)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recover抓住了异常并处理掉了，没有了异常
    finally:
        conn.close()

def accept(server):
    i = 1
    while True:
        # 分配一个新的服务员
        conn, raddr = server.accept()  # 三次握手建立好的连接，返回一个新的socket和元组(对方的地址)
        print(1, conn, type(conn))  # 负责客户端连接的socket对象
        print(2, raddr, type(raddr))  # 对方IP地址和端口
        # "T-线程ID"
        threading.Thread(target=recv, name="T-{}".format(i), args=(conn,raddr)).start()
        i += 1
        print('+' * 30)

if __name__ == '__main__':
    server = socket.socket()  # 创建socket对象
    laddr = ('0.0.0.0', 9999)  # 地址和端口的元组
    server.bind(laddr)  # 绑定
    server.listen(1024)  # 监听

    # accept() 也创建一个线程 这样下面print可以执行
    threading.Thread(target=accept, name="ac", args=(server,)).start()

    print('~' * 30)
    while True:
        time.sleep(5)
        print(threading.enumerate()) #每5秒查看进程还有多少

多线程+阻塞O，Multithread +Blocked Io
	thread per connection
	如果高并发请求，线程太多了，每一个线程都要创建(耗时)，分配线程栈内存MB
		静态资源，IO太多了，进入阻塞态
		动态资源，DB( 硬盘 网络+硬盘IO )，内存中内容，CPU计算，性能更差
		
线程复用
	池Poo1
		资源有限，所以复用
		泳道有限个，最大个数 8000
		第一个来防水，水放好后，下水，懒，太慢了
		第一个人来之前，提前把水放好，预热，最少空闲线程
		起水后，泳道放水吗?不放，复用		
		
		
EG:实现WEB服务器--线程池版 - 有问题没调式成功
import socket
import threading
import time

html_response_body = """\
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>BIRKHOFF</title>
</head>
<body>
    <h1>BIRKHOFF HTML TEST PAGE -- {}</h1>
</body>
</html>""".format("MultiThread + BIO").encode() #UTF-8 Bytes

response = """\
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: {}
X-Server: BIRKHOFF
content-type: text/html; charset=utf-8

""".format(len(html_response_body)).encode() + html_response_body

from concurrent.futures import ThreadPoolExecutor

def recv(conn: socket.socket, raddr):
    try:
        data = conn.recv(1024)  # 阻塞 等待客户端消息来 Read
        # data就是请求报文。按道理这里要做请求报文的解析，解析库
        # 就应该按照method Path等信息，路由 找到静态文件读取内容返回，或动态生成内容
        print(3, data, type(data), threading.current_thread().name)
        if not data:
            print("Bye~~~", raddr)
            return
        # 省略几万行代码，响应报文封装
        conn.send(response)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recover抓住了异常并处理掉了，没有了异常
    finally:
        conn.close()

def accept(server):
    i = 1
    while True:
        # 分配一个新的服务员
        conn, raddr = server.accept()  # 三次握手建立好的连接，返回一个新的socket和元组(对方的地址)
        print(1, conn, type(conn))  # 负责客户端连接的socket对象
        print(2, raddr, type(raddr))  # 对方IP地址和端口
        # "T-线程ID"
        # threading.Thread(target=recv, name="T-{}".format(i), args=(conn,raddr)).start()
        executor.submit(recv, conn, raddr)
        i += 1
        print('+' * 30)

# executor = concurrent.futures.ThreadPoolExecutor(1000)
executor = ThreadPoolExecutor(10)

if __name__ == '__main__':
    server = socket.socket()  # 创建socket对象
    laddr = ('0.0.0.0', 9999)  # 地址和端口的元组
    server.bind(laddr)  # 绑定
    server.listen(1024)  # 监听

    #daemon  守护线程，后台线程
    #main函数，默认是nondeamon线程
    # threading.Thread(target=accept, name="ac", args=(server,),daemon=True).start()
    executor.submit(accept, server) #帮助创建线程 并且复用

    print('~' * 30)
    while True:
        time.sleep(5)
        print(threading.enumerate()) #每5秒查看进程还有多少

【13-3-IO模型】
=====================================================================
CPU运行指令
	ring环
		ring0 kernel os 特权指令 操作硬件
		ring3 用户指令，用内存 map
	进程都会分配内存
		2块
			kernel用
			用户用

内核态
	运行内核指令(特权指令) 可以操作进程的kerne1的内存

用户态
	用户指令使用用户的内存

read一个磁盘上的文件
	f = open("/a.txt"，"r")
	data=f.read() // data 就是用户空间内存的数据结构;
	read函数，系统调用，
	OS提供访问硬件，陷入到内核态(特权指令)，从硬件把数据复制到进程的内核空间内存缓冲区buffer，copy数据到用户空间的内存buffer
	离开内核态，进入用户态，可以用户空间内存中data
	
同步、异步、阻塞、非阻塞完全不同的概念，不要混淆

同步、异步
	说的是结果，最终结果
	同步，不给最终结果，只要最终结果，哲不罢休
	异步，不给你最终结果，中间值(号码)，最终目的还是吃饭
	
阻塞、非阻塞
	卡不卡
	阻塞，条件不满足，石化了，调用者被阻塞，不能指向下一条指令
	非阻塞，条件不满足，调用者不卡住，就可能走到下一条指令，极有可能出错，要有特殊处理

同步同步阻塞:调用者只接受最终结果，不给最终结果就卡着
同步同步非阻塞:调用者只接受最终结果，不给最终结果不卡着

IO多路复用 Server
	管多路，交给选择器select 监管
	如果被监管某一路或者几路就绪，select 不阻塞了，返回就绪的keys和掩码
	for 遍历keys 拿到每一个key，通过key,data 关联数据做不同的操作
	while True死循环又让选择器.select()I

	没有了多线程，内核采用不同多路选择器技术完成,帮助我们从内核大大提高对多路控制，同时减少n个线程
	
【13-4-WebServer实战之IO多路复用】
=====================================================================
EG：
import socket
import threading
import time
import selectors
from selectors import EVENT_READ

html_response_body = """\
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>BIRKHOFF</title>
</head>
<body>
    <h1>BIRKHOFF HTML TEST PAGE -- {}</h1>
</body>
</html>""".format("Multiplexing").encode() #UTF-8 Bytes

response = """\
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: {}
X-Server: BIRKHOFF
content-type: text/html; charset=utf-8
Server: birkhoffxia.com

""".format(len(html_response_body)).encode() + html_response_body

selector = selectors.DefaultSelector() #初始化一个选择器对象
print("1",selector) #Linux Epoll

def recv(conn: socket.socket):
    try:
        data = conn.recv(1024)  # 阻塞 等待客户端消息来 Read
        if not data:
            print(conn.getpeername(), "Bye~~~")
        conn.send(response)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recover抓住了异常并处理掉了，没有了异常
    finally:
        selector.unregister(conn) #一定要反注册
        conn.close()

def accept(server):
    conn, raddr = server.accept()
    conn.setblocking(False) #要非阻塞
    selector.register(conn,EVENT_READ,recv)

if __name__ == '__main__':
    server = socket.socket()  # 创建socket对象
    server.setblocking(False) #要非阻塞
    laddr = ('0.0.0.0', 9999)  # 地址和端口的元组
    server.bind(laddr)  # 绑定
    server.listen(1024)  # 监听

    k = selector.register(server, EVENT_READ,accept)

    print('~' * 30)
    while True: #阻塞，第一阶段，阻塞到当前selector可读为止
        for key, event in selector.select(): #阻塞到有事件
            key.data(key.fileobj) #accpet(server)
            print("2",type(key),key == k)
            print("3",type(event),event)
            print("4",key.fileobj,key.fd,key.data,key.data)
            print("5",key.fileobj == server)

# 1 <selectors.SelectSelector object at 0x0000019263810278>
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# 2 <class 'selectors.SelectorKey'> True
# 3 <class 'int'> 1
# 4 <socket.socket fd=396, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('0.0.0.0', 9999)> 396 <function accept at 0x00000192638942F0> <function accept at 0x00000192638942F0>
# 5 True
# 2 <class 'selectors.SelectorKey'> False
# 3 <class 'int'> 1
# 4 <socket.socket [closed] fd=-1, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0> 424 <function recv at 0x00000192634B2E18> <function recv at 0x00000192634B2E18>
# 5 False



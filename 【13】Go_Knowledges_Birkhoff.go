Viedo Date: 2023/11/11
Made By:	BIRKHOFF
Date:	2024-06-29


��13-1-���̼߳�BIO��
=====================================================================
swiddler��https://swiddler.com/
#��ʹ��Swiddler����ģ�� �����շ���Ϣ
EG:
import socket

#TCP����˱��
server = socket.socket() #����socket����
laddr = ('0.0.0.0',9999) #��ַ�Ͷ˿ڵ�Ԫ��
server.bind(laddr) #��
server.listen(1024) #����

#�ȴ��������ӵĿͻ���
#server.accept() ����
conn,raddr = server.accept()  #�������ֽ����õ����ӣ�����һ���µ�socket��Ԫ��(�Է��ĵ�ַ)
print(1,conn,type(conn)) #����ͻ������ӵ�socket����
print(2,raddr,type(raddr)) #�Է�IP��ַ�Ͷ˿�
print(conn.getpeername(),conn.getsockname()) #ͨ��socket��ȡ�Զ˵�ַ�򱾵ص�ַ

data = conn.recv(4096) #���ܿͻ�����Ϣ
print(type(data),data)
conn.send(b"Hello BIRKHOFF")

conn.close()
server.close()

print('-' * 30)

# 1 <socket.socket fd=424, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=('127.0.0.1', 9999), raddr=('127.0.0.1', 50667)> <class 'socket.socket'>
# 2 ('127.0.0.1', 50667) <class 'tuple'>
# ('127.0.0.1', 50667) ('127.0.0.1', 9999)
# <class 'bytes'> b'HEYE'

[���߳�] Thread Per connect
EG��
import socket
import threading
import time


# ���߳� + �����ɣϡ��ͣ������������䣫���£����ɣ�

def recv(conn: socket.socket, raddr):
    try:
        for i in range(3):
            data = conn.recv(1024)  # ���� �ȴ��ͻ�����Ϣ�� Read
            print(3, data, type(data), threading.current_thread().name)
            if not data:
                print("Bye~~~", raddr)
                return
            conn.send(data)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recoverץס���쳣��������ˣ�û�����쳣
    finally:
        conn.close()


def accept(server):
    i = 1
    while True:
        # ����һ���µķ���Ա
        conn, raddr = server.accept()  # �������ֽ����õ����ӣ�����һ���µ�socket��Ԫ��(�Է��ĵ�ַ)
        print(1, conn, type(conn))  # ����ͻ������ӵ�socket����
        print(2, raddr, type(raddr))  # �Է�IP��ַ�Ͷ˿�
        # "T-�߳�ID"
        threading.Thread(target=recv, name="T-{}".format(i), args=(conn,raddr)).start()
        i += 1
        print('+' * 30)


if __name__ == '__main__':
    server = socket.socket()  # ����socket����
    laddr = ('0.0.0.0', 9999)  # ��ַ�Ͷ˿ڵ�Ԫ��
    server.bind(laddr)  # ��
    server.listen(1024)  # ����

    # accept() Ҳ����һ���߳� ��������print����ִ��
    threading.Thread(target=accept, name="ac", args=(server,)).start()

    print('~' * 30)
    while True:
        time.sleep(5)
        print(threading.enumerate()) #ÿ5��鿴���̻��ж���

���ģ���������4���ͻ��� ����1-1 ~ 4-3  �鿴�Ƿ���3����Ϣ�Զ��ر�
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

��13-2-WebServerԭ���ʵ�֡�
=====================================================================
Web Server ���
	httpd��nginx
	HTTPЭ��
		�����ı�
		����TCPЭ��֮�ϵ�Ӧ�ò�Э��
		Ӧ�ó���ʹ�õ�
	HTML
		�����ı���
		����ͼ�����֡���ʽ
		�ı��ļ������Է��ڴ�����
		ͼƬ�ļ�����̬�ļ�
	JsԴ����
		�ı��ļ������ڴ�����
		��װ��������װ�����������������
		���
			React��Vue2 3
					router ǰ��·���������
						�������url�仯��û�з���Ժ�̨server��http���󣩣�js��������仯��js��Ӧ�Ļص���������������ҳ��仯
						���ͨ�ţ��жϣ��첽����HTTP Ajax() 
	��̬����
		HTML��JS�ļ��ܲ����ɴ��붯̬������ЩԴ�����ı����ݶ�̬����
		svg���ı���������������
		image����֤��
		
Browser�����(�ͻ��ˡ�ǰ��)
	HTTPЭ��https ssl tsl
		URL ȫ�� ͳһ��Դ��λ����server�˶�λ��Դ�󣬰���Դ������ͨ��HTTpЭ�鷵�ص��������
		ǰ�˲���������ֱ�Ӵ��ļ���ȡ(��̬�ļ�)���Ǻ��ͨ�����붯̬��������(��̬)
		��serverͨ�ţ���������c/sͨ�ţ�B/s���ʾ���c/s
		v1
			������������ӵ�����ˣ�����һ�����󣬷������Ӧ�����ӶϿ�
			TCP���ӶϿ����ؽ�������Ҳ�ܺ���Դ
		v1.1 keepalive ����TCP����һ��ʱ��,һ��ʱ���������ӶϿ�
		
		v2
	HTML����(���ʦ)
		�γ�DOM
		Javascript����DOM��ʵ�ֶ�Ч
		��Ⱦrender
	CSS ����HTML��Ⱦ
	JavaScript����
		����ִ��JsԴ����
		Ecmascript ��׼
		ES3��ES5��ES6+
		����

Browser --> http������(ͬ�����첽Ajax) --> HTTP WEB Server(TCP listen) --> url method(router��λ��Դ ·�ɣ������Դ��λ)
	--> 404 Not Found; 403
	--> 1 ��λ�ļ�ϵͳ�е��ļ�open read ����
	--> 2 ��̬�������ݶ�̬��ҳ����ASP JSPPHP
			(Go/Python/PHP/Java/c/C#/C++/Ruby/Javascript......)
			Database�в�ѯ���ݻ���
			�ճ����ı�(����)

	-->��Ӧresponse����(��������) --> Browserl (HTMLԴ�롢JsԴ�롢image) 
	html���������������滹�кܶ���Դ��ҪΪÿһ����Դ��Ҫ����һ����������Ͷ�

OpenAI�����ߴ�ᣬ���ҿ���û�еͼ�������

����� https://www.bing.com/ ��ҳ(��̬����̬) -> Server 80 - > url -> response���� HTML
	v1.1
HTML����
	img src="/a.jpg"ͼƬ ���ٴη���http����
	�ű� src ="/js/t1.js" ���ٴη���http����
	n����Դ��Ҫ��������������ô��? ������TCP���� keepalive

HTTPЭ��
	������
	��Ӧ����
		ÿһ��\r\n����
		��һ��
			HTTP/1.1 200 OK Э�鼰�汾 ״̬�� ˵�� 
			״̬��
				1xx�ͻ��˴���;
				2xxͨ������
					200 OK
					201 CREATED ��Դ�����ɹ� Restful
					204 No Content�����ݷ��� ��������ɾ����Դ�ɹ�
				3xx��תlocation
					301
					302
				4xxȨ�ޡ���Դ����
					401
					402
					403
					404
					405
				5xx ����������
					500
					502nginx������� ���η��������󣬲���ningx����
			�ڶ��п�ʼ
				xxx: yyy key-value��
				Connection��keep-alive
				Content-Length:1292 �������������response body�ĳ��� �ֽ�
				Content-Type:image/png MIME����
				content-type:text/html;charset=utf-8

EG:����WebServer ʵ���˶��̼߳�����IO�汾
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
        data = conn.recv(1024)  # ���� �ȴ��ͻ�����Ϣ�� Read
        # data���������ġ�����������Ҫ�������ĵĽ�����������
        # ��Ӧ�ð���method Path����Ϣ��·�� �ҵ���̬�ļ���ȡ���ݷ��أ���̬��������
        print(3, data, type(data), threading.current_thread().name)
        if not data:
            print("Bye~~~", raddr)
            return
        # ʡ�Լ����д��룬��Ӧ���ķ�װ
        conn.send(response)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recoverץס���쳣��������ˣ�û�����쳣
    finally:
        conn.close()

def accept(server):
    i = 1
    while True:
        # ����һ���µķ���Ա
        conn, raddr = server.accept()  # �������ֽ����õ����ӣ�����һ���µ�socket��Ԫ��(�Է��ĵ�ַ)
        print(1, conn, type(conn))  # ����ͻ������ӵ�socket����
        print(2, raddr, type(raddr))  # �Է�IP��ַ�Ͷ˿�
        # "T-�߳�ID"
        threading.Thread(target=recv, name="T-{}".format(i), args=(conn,raddr)).start()
        i += 1
        print('+' * 30)

if __name__ == '__main__':
    server = socket.socket()  # ����socket����
    laddr = ('0.0.0.0', 9999)  # ��ַ�Ͷ˿ڵ�Ԫ��
    server.bind(laddr)  # ��
    server.listen(1024)  # ����

    # accept() Ҳ����һ���߳� ��������print����ִ��
    threading.Thread(target=accept, name="ac", args=(server,)).start()

    print('~' * 30)
    while True:
        time.sleep(5)
        print(threading.enumerate()) #ÿ5��鿴���̻��ж���

���߳�+����O��Multithread +Blocked Io
	thread per connection
	����߲��������߳�̫���ˣ�ÿһ���̶߳�Ҫ����(��ʱ)�������߳�ջ�ڴ�MB
		��̬��Դ��IO̫���ˣ���������̬
		��̬��Դ��DB( Ӳ�� ����+Ӳ��IO )���ڴ������ݣ�CPU���㣬���ܸ���
		
�̸߳���
	��Poo1
		��Դ���ޣ����Ը���
		Ӿ�����޸��������� 8000
		��һ������ˮ��ˮ�źú���ˮ������̫����
		��һ������֮ǰ����ǰ��ˮ�źã�Ԥ�ȣ����ٿ����߳�
		��ˮ��Ӿ����ˮ��?���ţ�����		
		
		
EG:ʵ��WEB������--�̳߳ذ� - ������û��ʽ�ɹ�
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
        data = conn.recv(1024)  # ���� �ȴ��ͻ�����Ϣ�� Read
        # data���������ġ�����������Ҫ�������ĵĽ�����������
        # ��Ӧ�ð���method Path����Ϣ��·�� �ҵ���̬�ļ���ȡ���ݷ��أ���̬��������
        print(3, data, type(data), threading.current_thread().name)
        if not data:
            print("Bye~~~", raddr)
            return
        # ʡ�Լ����д��룬��Ӧ���ķ�װ
        conn.send(response)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recoverץס���쳣��������ˣ�û�����쳣
    finally:
        conn.close()

def accept(server):
    i = 1
    while True:
        # ����һ���µķ���Ա
        conn, raddr = server.accept()  # �������ֽ����õ����ӣ�����һ���µ�socket��Ԫ��(�Է��ĵ�ַ)
        print(1, conn, type(conn))  # ����ͻ������ӵ�socket����
        print(2, raddr, type(raddr))  # �Է�IP��ַ�Ͷ˿�
        # "T-�߳�ID"
        # threading.Thread(target=recv, name="T-{}".format(i), args=(conn,raddr)).start()
        executor.submit(recv, conn, raddr)
        i += 1
        print('+' * 30)

# executor = concurrent.futures.ThreadPoolExecutor(1000)
executor = ThreadPoolExecutor(10)

if __name__ == '__main__':
    server = socket.socket()  # ����socket����
    laddr = ('0.0.0.0', 9999)  # ��ַ�Ͷ˿ڵ�Ԫ��
    server.bind(laddr)  # ��
    server.listen(1024)  # ����

    #daemon  �ػ��̣߳���̨�߳�
    #main������Ĭ����nondeamon�߳�
    # threading.Thread(target=accept, name="ac", args=(server,),daemon=True).start()
    executor.submit(accept, server) #���������߳� ���Ҹ���

    print('~' * 30)
    while True:
        time.sleep(5)
        print(threading.enumerate()) #ÿ5��鿴���̻��ж���

��13-3-IOģ�͡�
=====================================================================
CPU����ָ��
	ring��
		ring0 kernel os ��Ȩָ�� ����Ӳ��
		ring3 �û�ָ����ڴ� map
	���̶�������ڴ�
		2��
			kernel��
			�û���

�ں�̬
	�����ں�ָ��(��Ȩָ��) ���Բ������̵�kerne1���ڴ�

�û�̬
	�û�ָ��ʹ���û����ڴ�

readһ�������ϵ��ļ�
	f = open("/a.txt"��"r")
	data=f.read() // data �����û��ռ��ڴ�����ݽṹ;
	read������ϵͳ���ã�
	OS�ṩ����Ӳ�������뵽�ں�̬(��Ȩָ��)����Ӳ�������ݸ��Ƶ����̵��ں˿ռ��ڴ滺����buffer��copy���ݵ��û��ռ���ڴ�buffer
	�뿪�ں�̬�������û�̬�������û��ռ��ڴ���data
	
ͬ�����첽����������������ȫ��ͬ�ĸ����Ҫ����

ͬ�����첽
	˵���ǽ�������ս��
	ͬ�����������ս����ֻҪ���ս�����ܲ�����
	�첽�����������ս�����м�ֵ(����)������Ŀ�Ļ��ǳԷ�
	
������������
	������
	���������������㣬ʯ���ˣ������߱�����������ָ����һ��ָ��
	�����������������㣬�����߲���ס���Ϳ����ߵ���һ��ָ����п��ܳ���Ҫ�����⴦��

ͬ��ͬ������:������ֻ�������ս�����������ս���Ϳ���
ͬ��ͬ��������:������ֻ�������ս�����������ս��������

IO��·���� Server
	�ܶ�·������ѡ����select ���
	��������ĳһ·���߼�·������select �������ˣ����ؾ�����keys������
	for ����keys �õ�ÿһ��key��ͨ��key,data ������������ͬ�Ĳ���
	while True��ѭ������ѡ����.select()I

	û���˶��̣߳��ں˲��ò�ͬ��·ѡ�����������,�������Ǵ��ں˴����߶Զ�·���ƣ�ͬʱ����n���߳�
	
��13-4-WebServerʵս֮IO��·���á�
=====================================================================
EG��
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

selector = selectors.DefaultSelector() #��ʼ��һ��ѡ��������
print("1",selector) #Linux Epoll

def recv(conn: socket.socket):
    try:
        data = conn.recv(1024)  # ���� �ȴ��ͻ�����Ϣ�� Read
        if not data:
            print(conn.getpeername(), "Bye~~~")
        conn.send(response)  # echo Server
    except Exception as e:
        print(e, "----------------")  # recoverץס���쳣��������ˣ�û�����쳣
    finally:
        selector.unregister(conn) #һ��Ҫ��ע��
        conn.close()

def accept(server):
    conn, raddr = server.accept()
    conn.setblocking(False) #Ҫ������
    selector.register(conn,EVENT_READ,recv)

if __name__ == '__main__':
    server = socket.socket()  # ����socket����
    server.setblocking(False) #Ҫ������
    laddr = ('0.0.0.0', 9999)  # ��ַ�Ͷ˿ڵ�Ԫ��
    server.bind(laddr)  # ��
    server.listen(1024)  # ����

    k = selector.register(server, EVENT_READ,accept)

    print('~' * 30)
    while True: #��������һ�׶Σ���������ǰselector�ɶ�Ϊֹ
        for key, event in selector.select(): #���������¼�
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



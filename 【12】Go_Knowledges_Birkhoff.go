Viedo Date: 2023/11/4
Made By:	BIRKHOFF
Date:	2024-06-28

��12-1-Python��װ�ͻ����﷨��
=====================================================================


��12-2-����ԭ��
=====================================================================
����
	�߲������ߵĲ������ߵ���Եģ���Ե�ǰ�����õ���Դ(CPU���ڴ桢����)
	�����ĸ���:
		currencyCon���ͷ���һ����һ��ʱ��(��۲��ʱ���)�ڣ��кܶ�����½������ͬʱ����
	���еĸ���:
		parrallelism
		ǿ��ͬһʱ�̣��������ͬʱ�ڷ���
	����0�㣬�ϸ�����ѣ�0��֮ǰ��2����������բ�ڵȴ��ϸ��١��߲�����һ��ʱ�������кܶ�����Ҫ����
		ֻ��һ��բ�ڣ�û�в��еĿ��ܣ����С����ʣ��߲����ܲ����ô��н��?�ܣ�����Ҫע��ÿ���´�����ٶȣ�������
			Ӳ�̼�����չ:���п�IDE��sata���п�
			���п컹�ǲ��п�?
		��5��բ�ڣ�����
			ͬһʱ�̣����ͬʱͨ��5�������ʣ��߲����ܲ����ò��н��?���ԣ�������Ҫ�������Դ�������˳ɱ�
		�ŶӾ��Ƕ��У�FIFO�������е�ʱ����������
		Cache ����
			map�������վ�ľ�̬���ݿ���cache
			CDN������ �ͽ�
		
		���С����С����������С����涼�ǽ�������ķ������ۺ����ö����ֶ�������߲������⣬��ע��ɱ�

��12-3-�����̺߳�״̬��
=====================================================================
����
	�������
	���̼��໥���룬��������ң����������ͨ�ţ����л��ͷ����л�������Ϊ�ڴ涼�����ҵģ��˿ڶ������ҵģ�CPU�����ҵ�
	
	�߳�
		����߳�
		ÿ���߳�ӵ��ָ��(Դ��������)
		CPU����ָ�CPu������С���ȵ�Ԫ�߳�
		
		�̼߳�ͨ��
			ʡ��ʡ
				ÿ��ʡ�����Լ�����Դ��������ģ��߳����Լ���ջ(����ѹջ)
			����ͬһ�����̵��ڲ���Դ
			û�����л������л�

���ȼ�
	ֻ��˵���Խ��̡��߳���˵��һ��ʱ���ڹ۲죬�����cPuʱ��ռ�ȸ���
	������ô˵�����ȼ�Խ�ߣ�CPUֻ������������̻��߳�

�߳�״̬
	Ready��runable ���Ա�CPU���е��߳�
	Running������CPU�����У���ǰĳ��cPu������������������̵߳�ָ��
	Blocked���躮���̱߳����еĹ����У���ס��
	Terminated,��ֹ


��12-4-���߳̿����ͷ�����
=====================================================================
�߳�����
EG:
import threading
#��򵥵��̳߳���
def worker():
    print("I'm working")
    print('Done')

t = threading.Thread(target=worker,name='worker') #�����̶߳���
# target=worker��ʽΪ�ؼ��ִ��Σ������ƶ�Ӧ
# Python�л��а���λ�ö�Ӧ���Σ�����˳�����ζ�Ӧ
t.start()#�����������߳�

#I'm working
#Done

[�߳����˳�]
EG:
import threading
#��򵥵��̳߳���
def worker():
    print("I'm working")
    print('Done')

t = threading.Thread(target=worker,name='worker') #�����̶߳���
# target=worker��ʽΪ�ؼ��ִ��Σ������ƶ�Ӧ
# Python�л��а���λ�ö�Ӧ���Σ�����˳�����ζ�Ӧ
t.start()#�����������߳�

print("Process ID: ",t.ident) //�鿴����ID
# I'm working
# Process ID:  34204
# Done

[�߳��˳�]
Pythonû���ṩ�߳��˳��ķ������߳����������ʱ�˳�
	1���̺߳��������ִ����� 
	2���̺߳������׳�δ������쳣

#�̴߳��κͺ�������ûʲô���𣬱����Ͼ��Ǻ������Ρ�

EG��
import threading
import time
def worker(x,y):
    s = "{} + {} = {}".format(x,y,x+y)
    print(s,threading.currentThread().ident)

t1 = threading.Thread(target=worker,name='worker',args=(4,5))
t1.start()
time.sleep(2)

[���̱߳��]
����˼�壬����̣߳�һ������������ж���߳����У����Ƕ��̣߳�ʵ�ֲ�����
���������м����߳�����?
import threading
import time
import string

#��򵥵��̳߳���
def count():
    c=1
    while True:
        time.sleep(5)
        print("count = ",c)
        c+=1
        print("count()��",threading.current_thread())

def char():
    s = string.ascii_lowercase
    for c in s:
        time.sleep(2)
        print("char = ",c)
        print("char()��",threading.current_thread())

print("count()��",threading.current_thread())
t1 = threading.Thread(target=count,name="count")
t2 = threading.Thread(target=char,name="char")
print("{},{}".format(t1.ident,threading.current_thread()))
t1.start()
t2.start()
print("{},{}".format(t1.ident,threading.current_thread()))

#3���̣߳�count��char�����̡߳�
#����time.sleep(2)Ϊ20�����
# ��������������ڵ�ִ���߳���ô��?����˭��ס(����)��?˭��������?����߳�״̬��ʲô��ϵ?

#ע:Python����һ��GILȫ�ֽ�����������ҳ�ѧ���Ժ�����������������10��ʵӰ�첻��
#�ص�:���Ҫ���������ҵ�����ִ�еĸо��������Ⲣ������Goroutine�������洦��


��12-5-socket������
=====================================================================
import socket

#TCP����˱��
server = socket.socket() #����socket����
laddr = ('0.0.0.0',9999) #��ַ�Ͷ˿ڵ�Ԫ��
server.bind(laddr) #��
server.listen(1024) #����

#�ȴ��������ӵĿͻ���
#server.accept() ����
conn,raddr = server.accept()  #�������ֽ����õ����ӣ�����һ���µ�socket��Ԫ��(�Է��ĵ�ַ)
print(conn,type(conn)) #����ͻ������ӵ�socket����
print(raddr,type(conn)) #�Է�IP��ַ�Ͷ˿�
print(conn.getpeername(),conn.getsockname()) #ͨ��socket��ȡ�Զ˵�ַ�򱾵ص�ַ

data = conn.recv(4096) #���ܿͻ�����Ϣ
print(type(data),data)
conn.send(b"Hello BIRKHOFF")

conn.close()
server.close()
print('-' * 30)
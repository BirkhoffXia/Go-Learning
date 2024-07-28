Viedo Date: 2023/10/28
Made By:	BIRKHOFF
Date:	2024-06-26

��11-1-ʱ�䡿
=====================================================================
[ʱ��]
ʱ���Ƿǳ���Ҫ�ģ��뿪��ʱ�䣬����û���ĸ��������������ܹ������塣
��Go�����У�ʱ�䶨��ΪTime�ṹ�塣

	ʱ���ַ�������Parse
		��־��������־�����ı��ļ�
		�ַ���=>Timeʵ��
		����ʱ��Ҫ���ʽ���ַ��������ʱ���ַ���ƥ�䣬����err
		����ʱ
			��С��֮�⣬Ҫ���ʽһ��
			С��
				0����ȫһ��
				9������ƥ������λС��

	type Month int,�����Ϳ�����չ����
		String()string,Stringer�ӿ�
		int =>string,map

	timestamp
		Unixʱ���(��):1970��1��1��0���������ڵ�����(ʱ����0)
		Javascript ϰ��ʱ�����ms
		t��Timeʵ�������������ֳɵ�ʱ��ʵ��
			t.Unix()��t.UnixMilli()��t.UnixMicro(),t.UnixNano() ʵ��.����()			
			Unix()�Ƿ�������ʾʱ��ʵ�����ø�������չ�����ķ����õ�ʱ�������
	����ʱ��
		ʱ���ַ��� parse ʱ�����Time���淽��Time format => string
		time.Now() => Timeʵ��
		ʱ���int => Timeʵ��,��������û��Timeʵ��?
			time.Unix(sec int64,nsec int64)
 ʱ������
		time + time û������
		time-time =>ʱ������Ͷ���ʱ������-�����ã�goû���ṩ ���������
		time t delta  => time
I

EG:
package main

import (
	"fmt"
	"time"
)

func main() {
	var t = time.Now()
	fmt.Printf("%T\n", t) //time.Time
	fmt.Printf("%v: %+[1]v\n", t)
	// 2024-06-26 11:16:27.4683718 +0800 CST m=+0.004229301: 2024-06-26 11:16:27.4683718 +0800 CST m=+0.004229301
	//time.Time ʵ����GoString�ӿ�
	fmt.Printf("%#v\n", t)             //time.Date(2024, time.June, 26, 11, 15, 7, 912760200, time.Local)
	fmt.Printf("%T, %[1]v\n", t.UTC()) //time.Time, 2024-06-26 03:15:07.9127602 +0000 UTC

	//��ʱ���ʽ����
	//���ڸ�ʽ���� %y %Y %m %M %s Go������
	//1��2������3ʱ4��5��6��
	fmt.Println(t.Format("01*02*03*04*05*06 -0700"))    //06*26*11*21*13*24 +0800
	fmt.Println(t.Format("2006/01/02/ 15:04:05 -0700")) //2024/06/26/ 11:24:49 +0800
	fmt.Println(t.Format("2006/01/02/ 15:04:05"))       //2024/06/26/ 11:26:08
	//С��
	fmt.Println(t.Format("2006/01/02/ 15:04:05.000000000 -0700")) //2024/06/26/ 11:31:00.632596400 +0800
	fmt.Println(t.Format("2006/01/02/ 15:04:05.999999999 -0700")) //2024/06/26/ 11:31:00.6325964 +0800
	fmt.Println(t.UTC().Format("0102 030405 06 pm"))              //0626 032819 24 am

	//��ʱ�����������
	// s := "2019/06/28 06:27:00 +0800"
	s := "2019/06/28 06:27:00.1234567 +0800"

	// if t, err := time.Parse("2006/01/02 15:04:05 -0700", s); err == nil {
	if t, err := time.Parse("2006/01/02 15:04:05.0000000 -0700", s); err == nil { //����С��Ҫ����
		fmt.Println(t) //2019-06-28 06:27:00 +0800 CST
	} else {
		fmt.Println(err)
	}

	//��ʱ�����ԡ�
	//ȡ������ һ��ڼ����ȥ��
	fmt.Println(t.Year(), t.Month(), t.Day(), int(t.Month()), t.Month().String(), t.YearDay())
	//2024 June 26 6 June 178
	fmt.Println(t.Hour(), t.Minute(), t.Second()) //11 47 55
	fmt.Println(t.Nanosecond())                   //208191100
	fmt.Println(t.Weekday(), int(t.Weekday()))    //Wednesday 3
	fmt.Println(t.ISOWeek())                      //2024 26

	//��ʱ����� ʹ��t.UnixMilli()����ʱ���
	fmt.Println(t.Unix(), t.UnixMilli(), t.UnixMicro(), t.UnixNano()) //1719373890 1719373890743 1719373890743447 1719373890743447300

	//����
	ts := time.Unix(1698462248, 0)     //unix ʱ���
	fmt.Println("unix timestamp:", ts) //unix timestamp: 2023-10-28 11:04:08 +0800 CST
	//Unix(sec int64, nsec int64) time.Time

	//��ʱ����
	tt := "2019/06/28 06:27:00"
	if tt, err := time.Parse("2006/01/02 15:04:05", tt); err == nil {
		fmt.Println(tt.UTC())   //2019-06-27 22:27:00 +0000 UTC
		fmt.Println(tt.Local()) //2019-06-28 14:27:00 +0800 CST
		fmt.Println(tt)         //2019-06-28 06:27:00 +0000 UTC Ĭ����ʾ0ʱ��
	} else {
		fmt.Println(err)
	}
	/*����ָ��location����ʱ��*/
	tz, _ := time.LoadLocation("Asia/Shanghai")
	if t, err := time.ParseInLocation(
		"2006/01/02 15:04:05",
		"2019/06/28 06:27:00",
		tz,
	); err == nil {
		fmt.Println(t.UTC())   //2019-06-27 22:27:00 +0000 UTC
		fmt.Println(t.Local()) //2019-06-28 06:27:00 +0800 CST
		fmt.Println(t)         //2019-06-28 06:27:00 +0800 CST

	}

}

// ���ʱ��� "2019-06-28 06:27:00 +0800 CST" ���԰����·�ʽ��⣺

// "2019-06-28" ��ʾ���ڲ��֣������-�·�-���ڡ�
// "06:27:00" ��ʾʱ�䲿�֣���ʱ-��-�롣
// "+0800" ��ʾʱ��ƫ������������� UTC ʱ���ƫ����������������£�ƫ����Ϊ +08:00����ʾ����� UTC ʱ����ƫ�� 8 Сʱ��
// "CST" ��ʾʱ������д�������й���׼ʱ�䣨China Standard Time�������� UTC+8 ��ʱ����
// �ۺ���������ʱ�����ʾ�����й���׼ʱ�䣨UTC+8���µ� 2019 �� 6 �� 28 �� 06:27:00 ʱ�̡�

// ���ʱ��� "2019-06-27 22:27:00 +0000 UTC" ���԰����·�ʽ��⣺

// "2019-06-27" ��ʾ���ڲ��֣������-�·�-���ڡ�
// "22:27:00" ��ʾʱ�䲿�֣���ʱ-��-�롣
// "+0000" ��ʾʱ��ƫ������������� UTC ʱ���ƫ����������������£�ƫ����Ϊ +00:00����ʾ�� UTC ʱ�䱣��һ�£�û��ʱ��ƫ�ơ�
// "UTC" ��ʾЭ������ʱ��Coordinated Universal Time��������һ����׼��ȫ��ʱ���׼���൱������ͨ����˵�ĸ������α�׼ʱ�䣨GMT����
// �ۺ���������ʱ�����ʾ����Э������ʱ��UTC���µ� 2019 �� 6 �� 27 �� 22:27:00 ʱ�̣�����ʱ��ƫ����Ϊ +00:00����������α�׼ʱ�䣨GMT������һ�¡�


// A Month specifies a month of the year (January = 1, ...).
type Month int //�����Ϳ�����չ����

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

EG:
package main

import (
	"fmt"
	"time"
)

func main() {
	//��ʱ�����㡿
	tz, _ := time.LoadLocation("Asia/Shanghai")
	s1 := "2024/06/26 09:00:00"
	s2 := "2024/06/26 10:00:00"
	layout := "2006/01/02 15:04:05"
	t1, _ := time.ParseInLocation(layout, s1, tz)
	t2, _ := time.ParseInLocation(layout, s2, tz)
	fmt.Println(t1) //2024-06-26 09:00:00 +0800 CST

	fmt.Println(t2) //2024-06-26 10:00:00 +0800 CST

	//ʱ���
	delta := t2.Sub(t1)                           //t2-t1
	fmt.Printf("delta: %v,%[1]T\n", delta)        //delta: 1h0m0s,time.Duration
	fmt.Println(delta.Minutes(), delta.Seconds()) //60 3600	�������Сʱ����

	//����Duration
	ns3 := time.Duration(3)              //3����
	s3 := time.Duration(3 * time.Second) //3��
	h3 := time.Duration(3 * time.Hour)   //3Сʱ
	fmt.Println(ns3, s3, h3)             //3ns 3s 3h0m0s

	//ʱ��ƫ��
	t3 := t2.Add(h3)
	fmt.Println("t3:", t3) //t3: 2024-06-26 13:00:00 +0800 CST
	t4 := t2.Add(-h3)
	fmt.Println("t4:", t4)                     //t4: 2024-06-26 07:00:00 +0800 CST
	fmt.Println("t3.After:", t3.After(t4))     //t3.After: true
	fmt.Println("time.Since:", time.Since(t2)) //time.Since: 3h44m47.1194582s

	//ʱ���
	t := time.Unix(1, 0)
	fmt.Printf("%T %+[1]v\n", t) //time.Time 1970-01-01 08:00:01 +0800 CST
	fmt.Println(t.UTC())         //1970-01-01 00:00:01 +0000 UTC
}


��11-2-������
=====================================================================
[ģ�黯]
������
	ʹ��Ŀ¼��֯����Ŀ¼���ǰ�
		�����棬��������.go�ļ���.go�ļ���дpackage����
	����Сд�����ϱ�ʶ������Ҫ��
	��main���⣬�����������Ŀ¼�������Բ�һ������һ���Ļ����������鷳
	ͬһ��Ŀ¼����ͬһ����
		Ҫ������ .go�ļ�����ʹ��ͬһ������
		��xxx_test��������
	
	main������
		main����������main����
		��ִ���ļ�(windows .exe)�����и�������ֶΣ�-> main������ַ
	
	��Ŀ����
		go modulesģʽ
			go mod init test : test ���ذ�
			go mod init magedu.com/wayne/test Զ�ֿ̲���Ŀ��ַ�������
			go get github.com/magedu/tools git
			import "github.com/magedu/tools"
			
			go.mod�ļ� ���а���������������
				module ����
				go �汾
		 		Windows�汾 C:\Users\BIRKHOFF ALW\go\pkg\mod
		go mod vendor ʹ��vendoring���ƣ���Ŀ��Ŀ¼�������vendorĿ¼�������������Ŀ¼
		
		
		1.13�汾֮��Ĭ��on set GO111MODULE=

GOPROXY������������ָ�������ؾ���(�����ַ��ʱ��仯������չٷ������ĵ�)
GOPROXY=https://goproxy.cn,direct
GOPROXY=https://mirrors.aliyun.com/goproxy
GOPROXY=https://mirrors.cloud.tencent.com/go/
GOPROXY=https://repo.huaweicloud.com/repository/goproxy/

	F5������ԭ������ʱ���� Go build Go run ��Ϊ���ǰ���Go��������(��׼��GOROOT/src �� Go����GOROOT/bin) 
		��д��Դ������ϵ�����˵�Դ��(Դ����ϲ�����Ĵ�����) һ�����
	���뷽ʽ
			"github.com/vmihailenco/msgpack/v5" //���Ե���
			 m "test/calc" m��ʾ�������� ��������
			 . "test/calc/minus" //.��ǰminus���е�����ȫ�ֱ������뵽��ǰ��
			 _ "test/calc/minus" //_û������ minus���е�����ȫ�ֱ������뵽��ǰ�ڶ���Ҳ�����е����޷�ʹ�ã�Ҫ�����	
			 										// �������룬����Ϊ��init()û��ʹ����Դ,

	init()����
		����д���κ�.go�ļ���
		�޲Ρ��޷���ֵ
		һ��һ��go�ļ���ֻдһ��������д���
		��ͬ��go�ļ��ж�����д
		
		���Ե��롢�������붼����ָ��init
		���붼����ִ��init(),������������
		
		Ӧ�ã���������������������Լ�ʹ���������룬init()��ִ��
		
		init�������޲��޷���ֵ�����ܱ�������������
		���е�init��������main����֮ǰ�Զ�ִ��
		ÿ������init���������ж�����ҿ���λ�ڲ�ͬ���ļ���
		ͬһ���ļ��п����ж��init��������һ��һ���͹��ˣ�������д���
		ͬһ�����е�init����û����ȷ��ִ��˳�򣬲���Ԥ��
		��ͬ����init������ִ��˳���ɵ���˳�����
		
		init������Ҫ����һЩ��ʼ��������init��main������һ����ͬһ���ļ��С�
		import _ "xxx"������ʲô?ִֻ�иð���init�������޷�ʹ�ð�����Դ��
		import "xxx"������ʲô?Ҳ��ִ�иð���init������Ҳ����ʹ�ð�����Դ.��
		
		
[�Ӱ�]
��main.go��
package main

import (
	"fmt"
	"test/calc"
	// "github.com/vmihailenco/msgpack/v5" //���Ե���
	// m "test/calc" m��ʾ�������� ��������
	// . "test/calc/minus" .��ǰminus���е�����ȫ�ֱ������뵽��ǰ��
	// _ "test/calc/minus" _û������ minus���е�����ȫ�ֱ������뵽��ǰ�ڶ���Ҳ�����е����޷�ʹ�ã�Ҫ�����
)

func main() {
	// fmt.Printf(msgpack.Marshal())
	fmt.Printf("calc.Add(3,7): %v\n", calc.Add(3, 7))
	// fmt.Printf("minus.Minus(15,5) : %v\n", minus.Minus(15, 5))
}

// calc/calc.go Add:3 7
// calc/add.go fn:3 7
// calc.Add(3,7): 10
// calc/minus/minus.go Minus:15 5
// minus.Minus(15,5) : 10

// add.go [1] init
// add.go [2] init
// calc.go [1] init
// calc.go [2] init

��calc/add.go��
package calc

import "fmt"

func fn(x, y int) int {
	fmt.Printf("calc/add.go fn:%d %d\n", x, y)
	return x + y
}

func init() {
	fmt.Println("add.go [1] init")
}

func init() {
	fmt.Println("add.go [2] init")
}

��calc/calc.go��
package calc

import "fmt"

func Add(x, y int) int {
	fmt.Printf("calc/calc.go Add:%d %d\n", x, y)
	return fn(x, y)
}

func init() {
	fmt.Println("calc.go [1] init")
}

func init() {
	fmt.Println("calc.go [2] init")
}

[replace]
���ڽ�һ��ģ��汾�滻Ϊ����һ��ģ��汾


��11-3-log����
=====================================================================
��׼��1og
	��־��֮�����ܹ������ԭ���ڲ����ṹ��һ��Logger��־��¼��
	ȱʡ��־��¼��
		var std = New(os.stderr,""��LstdFlags) => *Logger
			1���outputSetoutput
			2 prefix
			3 flags���ñ�ǣ����ֱ���ṩһЩ����
	func Default()*Logger { return std }// log.Default()->std *Logger ֻ��getter
	������ݷ���
		log.Print * 0
		log.Fatal* log.Print* + os.Exit(1) ֮����벻��ִ��
		log.Panic* log.Print* + panic 2

	Logger����
		Print*��Fatal*��Panic*
	��־�ļ����
		ֻдwriteonly
		os.0penFile(name string, flag int, perm os.FileMode)
			name �ļ�·��
			flag������δ�ֻ����ֻд���ļ����ڴ������ļ��ڱ���׷��д�롢�ļ����������
			��λ������flag���������ʹ��
			os.O_CREATE�ļ��������½������ڲ���
			oS.O_WRONLY �ļ���дwritable��Ͻӿڷ���write����
			oS.O_APPEND׷��
	��־����
		ֻ��
		os.Open(name string)
	zerolog �Ƚϼ�
		һ�㶼��Ҫ����
		


[��׼��]
// ʹ��ȱʡLogger
log.Print("abcde\n")
log.Printf("%s\n","abcd")
log.Printin("abc")
log.Fatal("xyz")// �ȼ��� 1og.Print(��xyz");os.Exit(1)
logg.Panicln("Failed") // �ȼ��� log.Println("Failed");panic()


package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		err := recover()
		fmt.Println("err:", err)
	}() // reutnr or panic
	log.Println("abc") //��¼��־ �˳�״̬��Ϊ0 status 0
	fmt.Println("abc~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	log.Fatalf("Fatal Code: %v \n", "aabbcc") //2024/06/28 10:46:04 Fatal Code: aabbcc  �˳�״̬��Ϊ1 ֮����벻��ִ��
	//Fatal ֱ�� os.Exit(1)
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	log.Panicf("Panic: %v \n", "paninc information") //panic ����״̬��Ϊstatus 2
	fmt.Println("~~~~~~~~~~~~~~~~~~")

	log.Default() //func Default() *Logger { return std } => std *Logger

}

	log.Default().SetOutput(os.Stdout) //�ڲ���std ���κ�Logger�����Ը�
	log.Println("###abc###")           //2024/06/28 14:10:34 ###abc###
	
[�Զ���Logger]
�������ȱʡLogger std�����⣬����New����һ���Զ���Logger��ָ��ǰ׺��Flags��

package main

import (
	"log"
	"os"
)

func main() {
	/*�Զ���Logger*/

	// log.New()
	// func log.New(out io.Writer, prefix string, flag int) *log.Logger
	// out: out, prefix: prefix, flag: flag
	// type Writer interface {
	// 	Write(p []byte) (n int, err error)
	// }
	infologger := log.New(os.Stdout, "Info: ", log.LstdFlags)
	infologger.Println("����һ����ͨ��Ϣ1") //Info: 2024/06/28 14:18:43 ����һ����ͨ��Ϣ1

	////���Զ���ķź���
	infologger = log.New(os.Stdout, "Info: ", log.LstdFlags|log.Lmsgprefix)
	infologger.Println("����һ����ͨ��Ϣ2") //2024/06/28 14:20:00 Info: ����һ����ͨ��Ϣ
}

[д��־�ļ�]
New����ǩ�� New(out io.writer��prefix string��flag int)*Logger ��out�����ṩWriter�ӿڼ��ɣ���ô�Ϳ����ṩһ����д�ļ�����

package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(
		"D:/my_nginx.log",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, //ֻд���ļ������ڴ�����׷��
		os.ModePerm,                         //Unix permission bits��0o777
	) //��д ����һ���ļ�ָ��(���)
	if err != nil {
		log.Panicln(err)
	}

	defer f.Close() //��Ҫ�ùر�
	l := log.New(f, "BIRKHOFF Logger", log.LstdFlags)
	l.Println("����һ����д�ļ�����Ϣ Logger Title")
}


��11-4-zerolog��
=====================================================================
go get -u github.com/rs/zerolog/log
go: downloading github.com/rs/zerolog v1.33.0
go: downloading github.com/mattn/go-isatty v0.0.19
go: downloading golang.org/x/sys v0.12.0
go: downloading golang.org/x/sys v0.21.0
go: added github.com/mattn/go-colorable v0.1.13
go: added github.com/mattn/go-isatty v0.0.20
go: added github.com/rs/zerolog v1.33.0     
go: added golang.org/x/sys v0.21.0


����
	��Ϣ�ļ���
		{"level":"warn","time":"2024-06-28T16:27:24+08:00","message":"Warn CPU HIGH"}
		{"level":"error","time":"2024-06-28T16:27:24+08:00","message":"Error run too many times"}


const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota //0
	// InfoLevel defines info log level.
	InfoLevel //1
	// WarnLevel defines warn log level.
	WarnLevel //2
	// ErrorLevel defines error log level.
	ErrorLevel //3
	// FatalLevel defines fatal log level.
	FatalLevel //4
	// PanicLevel defines panic log level.
	PanicLevel //5
	// NoLevel defines an absent log level.
	NoLevel //6
	// Disabled disables the logger.
	Disabled //7


zerolog
	����
		��Ϣ�ļ���
			logger.Warn()��logger.Debug()
		logger�ļ���New���ɵ�LoggerĬ�ϼ�����trace -1
			logger.Level(?) => new child logger
			log.New() => default level trace -l
		gLevel -l trace
			SetGlobalLevel(?)

		��Ϣ�������:��Ϣ���� >= Max(gLevel��logger.Level)
	
	zerolog/log 1og.Xxx()��ݷ���
		log.Warn().Msg("warn string")
		һ����¼ level":"warn","time":"2023-10-28T18:07:17+08:00","message":"warn string"
		Warn()=>"level":"warn"��Ϣ�ļ���
		Msg("warn string") => "message":"warn string"
		time�ֶ�һ��Ҫ��
	
	var Logger = zerolog.New(os.stderr).with().Timestamp().Logger()������ã�ȫ��ȱʡ
		log.Logger
		logger.Warn()
		logger.Debug()������Ϣ����
		logger.Level(zerolog.XxxxLevel) ����Logger����

package main

import (
	// "log"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	/*ȫ��Level trace*/
	fmt.Println(zerolog.GlobalLevel(), "####") //trace ####
	//����ȫ��Level
	// zerolog.SetGlobalLevel(zerolog.WarnLevel) //���������warning����ĵ�����������޷���ʾ
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	fmt.Println(zerolog.GlobalLevel(), "####") //Info ####

	/*log.Logger default logger*/
	//zerolog�����ʽdebug����
	log.Print("Kramer Lamer")
	// {"level":"debug","time":"2024-06-28T16:27:24+08:00","message":"Kramer Lamer"}
	log.Debug().Msg("Debug Kramer Lamer")
	//{"level":"debug","time":"2024-06-28T16:40:51+08:00","message":"Debug Kramer Lamer"}
	log.Info().Msg("Info string")
	//{"level":"info","time":"2024-06-28T16:35:45+08:00","message":"Info string"}
	log.Warn().Msg("Warn CPU HIGH")
	// {"level":"warn","time":"2024-06-28T16:27:24+08:00","message":"Warn CPU HIGH"}
	log.Error().Msg("Error run too many times")
	// {"level":"error","time":"2024-06-28T16:27:24+08:00","message":"Error run too many times"}
	fmt.Println(log.Logger.GetLevel()) //Ĭ���ǣ�trace -1
	fmt.Println("------------------------------")

	// mylog := zerolog.New(os.Stdout)
	// fmt.Println(mylog.GetLevel()) //trace
	// mylog.Info().Msg("mylog info")
	// mylog.Warn().Msg("mylog warning")
	// // // {"level":"info","message":"mylog info"}
	// // // {"level":"warn","message":"mylog warning"}
	// fmt.Println("------------------------------")

	mylog := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.WarnLevel)
	mylog.Info().Msg("mylog info")
	mylog.Warn().Msg("mylog warning")
	//��Ϊ�����˼���ΪWarning ���Ա����ͼ���Ĳ������
	// {"level":"warn","time":"2024-06-28T16:43:24+08:00","message":"mylog warning"}
	fmt.Println(mylog.GetLevel()) //Ĭ���ǣ�warn
	fmt.Println("------------------------------")

	/*����Level*/
	log1 := log.Level(zerolog.ErrorLevel) //��Ϊ������Error���� �������涼�����
	fmt.Println(log1.GetLevel())          //error
	log1.Info().Msg("log1 Info msg")
	log1.Warn().Msg("log1 Warn msg")
	log1.Error().Msg("log1 Error msg")
}

		
	�����ֶ�
		����Loggerʱ��With =>Context .Logger() => Logger  	
		ʹ��Loggerʱ��mylog.Warn().Msg("warn infomation")
		
package main

import (
	"fmt"
	"errors"
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	/*����Ĭ�ϸ�ʽ*/
	// zerolog.TimeFieldFormat =

	zerolog.TimeFieldFormat = "2006"

	mylog := zerolog.New(os.Stdout).With().Timestamp().Bool("Success", true).Str("Name", "Nginx").Logger().Level(zerolog.WarnLevel)
	fmt.Println(mylog.GetLevel())

	mylog.Info().Msg("mylog info")
	mylog.Warn().Msg("mylog warning")
	mylog.Warn().Msg("mylog error")
	// warn
	// {"level":"warn","Success":true,"Name":"Nginx","time":"2024-06-28T17:19:30+08:00","message":"mylog warning"}
	// {"level":"warn","Success":true,"Name":"Nginx","time":"2024-06-28T17:19:30+08:00","message":"mylog error"}

	// log1 := log.Level(zerolog.ErrorLevel)
	// log1.Error().Bool("test", false).Str("name", "Tom").Floats32("scores", []float32{60.6, 2, 4}).Msg("log1 Error msg")
	//{"level":"error","test":false,"name":"Tom","scores":[60.6,2,4],"time":"2024-06-28T18:28:48+08:00","message":"log1 Error msg"}
	log1 := log.Level(zerolog.ErrorLevel).With().Caller().Logger() //Caller���ӡ ��������
	// {"level":"error","error":"�Զ������","time":"2024","caller":"e:/goprojects/main.go:35"}
	// {"level":"fatal","error":"�Զ������","time":"2024","caller":"e:/goprojects/main.go:37"}

	//log1.Error() ��Ϣ����
	log1.Error().Err(errors.New("�Զ������")).Msg("")
	//{"level":"error","error":"�Զ������","time":"2024-06-28T18:37:04+08:00"}
	log1.Fatal().Err(errors.New("�Զ������")).Send()
	//{"level":"fatal","error":"�Զ������","time":"2024-06-28T18:38:21+08:00"}
}
			
		
[д�ļ�]
package main

import (
	// "log"

	// "internal/types/errors"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	f, err := os.OpenFile("D:/my.log", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Panic().Err(err).Send() //�ڲ�����panic
	}
	defer f.Close()

	multi := zerolog.MultiLevelWriter(f, os.Stdout) //���֧д �ļ��Ϳ���̨
	//Timesta()���ȫ�µ�lOgger����ʱ������
	logger := zerolog.New(multi).With().Timestamp().Logger()
	logger.Info().Msg("Write to Control Screen and File") //{"level":"info","time":1719571970,"message":"Write to Control Screen and File"}

}
		
[������־]	��׼��û���ṩ 

go get gopkg.in/natefinch/lumberjack.v2

	������־�ļ�������
	����:����һ����ͬ���ļ�����ԭ�����ļ�������
		ʱ��
		��С

package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	f, err := os.OpenFile("D:/my.log", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Panic().Err(err).Send() //�ڲ�����panic
	}

	l := &lumberjack.Logger{
		Filename:   "d:/routine.log",
		MaxBackups: 2,     //����ǰ����д�����־�ļ��⣬��ʷ��־��ౣ��2��
		Compress:   false, //ȱʡ��ѹ��
		MaxAge:     1,     //1��
		MaxSize:    1,     //ȱʡֵ100 ��100M ����1M��������
	} //ʵ����Writer�ӿ� �õ�һ���Զ������Logger
	defer l.Close()

	multi := zerolog.MultiLevelWriter(f, os.Stdout, l) //���֧д �ļ��Ϳ���̨
	logger := zerolog.New(multi).With().Timestamp().Logger()
	for {
		time.Sleep(1 * time.Microsecond)
		t := time.Now().Format("[06-01-02 15:04:05 -0700]")
		logger.Info().Msg(t)
	}
}

[��ӡ����ջ]
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func add(x, y int) int {
	return x + y
}
func division(x, y int) int {
	return x / y
}
func calc(x, y int, fn func(int, int) int) int {
	return fn(x, y)
}
func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	defer func() {
		err := recover()
		switch v := err.(type) {
		case nil:
			log.Debug().Msg("û�д���")
		case runtime.Error:
			fmt.Println(string(debug.Stack()))
			log.Error().Caller(3).Err(v).Str("stack", string(debug.Stack())).Send()
		default:
			log.Debug().Msg(fmt.Sprintf("��������", v))
		}
	}()

	fmt.Println(calc(10, 0, add))
	fmt.Println(calc(10, 0, division))
}

/*
goroutine 1 [running]:
runtime/debug.Stack()
	D:/InstalledApp/Go/src/runtime/debug/stack.go:24 +0x7a
main.main.func1()
	e:/goprojects/main.go:33 +0xb7
panic({0x1a32a0, 0x25d860})
	D:/InstalledApp/Go/src/runtime/panic.go:890 +0x262
main.division(0xa, 0x0)
	e:/goprojects/main.go:16 +0x4c
main.calc(0xa, 0x0, 0x1c43f0)
	e:/goprojects/main.go:19 +0x34
main.main()
	e:/goprojects/main.go:41 +0x109

{"level":"error","caller":"e:/goprojects/main.go:16","error":"runtime error: integer divide by zero","stack":"goroutine 1 [running]:\nruntime/debug.Stack()\n\tD:/InstalledApp/Go/src/runtime/debug/stack.go:24 +0x7a\nmain.main.func1()\n\te:/goprojects/main.go:34 +0x249\npanic({0x1a32a0, 0x25d860})\n\tD:/InstalledApp/Go/src/runtime/panic.go:890 +0x262\nmain.division(0xa, 0x0)\n\te:/goprojects/main.go:16 +0x4c\nmain.calc(0xa, 0x0, 0x1c43f0)\n\te:/goprojects/main.go:19 +0x34\nmain.main()\n\te:/goprojects/main.go:41 +0x109\n","time":1719573745}\
*/

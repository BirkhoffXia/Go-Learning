Viedo Date: 2023/10/28
Made By:	BIRKHOFF
Date:	2024-06-26

【11-1-时间】
=====================================================================
[时间]
时间是非常重要的，离开了时间，几乎没有哪个生产环境数据能够有意义。
在Go语言中，时间定义为Time结构体。

	时间字符串解析Parse
		日志分析，日志就是文本文件
		字符串=>Time实例
		解析时，要求格式化字符串必须和时间字符串匹配，否则err
		解析时
			除小数之外，要求格式一致
			小数
				0，完全一致
				9，可以匹配任意位小数

	type Month int,新类型可以扩展方法
		String()string,Stringer接口
		int =>string,map

	timestamp
		Unix时间戳(秒):1970年1月1日0点整到现在的秒数(时区，0)
		Javascript 习惯时间戳是ms
		t是Time实例，我手里有现成的时间实例
			t.Unix()、t.UnixMilli()，t.UnixMicro(),t.UnixNano() 实例.方法()			
			Unix()是方法，表示时间实例调用该类型扩展出来的方法得到时间戳整数
	构建时间
		时间字符串 parse 时间对象Time，逆方向Time format => string
		time.Now() => Time实例
		时间戳int => Time实例,我手里有没有Time实例?
			time.Unix(sec int64,nsec int64)
 时间运算
		time + time 没有意义
		time-time =>时间差类型对象，时间增量-不能用，go没有提供 运算符重载
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
	//time.Time 实现了GoString接口
	fmt.Printf("%#v\n", t)             //time.Date(2024, time.June, 26, 11, 15, 7, 912760200, time.Local)
	fmt.Printf("%T, %[1]v\n", t.UTC()) //time.Time, 2024-06-26 03:15:07.9127602 +0000 UTC

	//【时间格式化】
	//日期格式化符 %y %Y %m %M %s Go都不用
	//1月2日下午3时4分5秒6年
	fmt.Println(t.Format("01*02*03*04*05*06 -0700"))    //06*26*11*21*13*24 +0800
	fmt.Println(t.Format("2006/01/02/ 15:04:05 -0700")) //2024/06/26/ 11:24:49 +0800
	fmt.Println(t.Format("2006/01/02/ 15:04:05"))       //2024/06/26/ 11:26:08
	//小数
	fmt.Println(t.Format("2006/01/02/ 15:04:05.000000000 -0700")) //2024/06/26/ 11:31:00.632596400 +0800
	fmt.Println(t.Format("2006/01/02/ 15:04:05.999999999 -0700")) //2024/06/26/ 11:31:00.6325964 +0800
	fmt.Println(t.UTC().Format("0102 030405 06 pm"))              //0626 032819 24 am

	//【时间解析分析】
	// s := "2019/06/28 06:27:00 +0800"
	s := "2019/06/28 06:27:00.1234567 +0800"

	// if t, err := time.Parse("2006/01/02 15:04:05 -0700", s); err == nil {
	if t, err := time.Parse("2006/01/02 15:04:05.0000000 -0700", s); err == nil { //解析小数要对齐
		fmt.Println(t) //2019-06-28 06:27:00 +0800 CST
	} else {
		fmt.Println(err)
	}

	//【时间属性】
	//取年月日 一年第几天过去了
	fmt.Println(t.Year(), t.Month(), t.Day(), int(t.Month()), t.Month().String(), t.YearDay())
	//2024 June 26 6 June 178
	fmt.Println(t.Hour(), t.Minute(), t.Second()) //11 47 55
	fmt.Println(t.Nanosecond())                   //208191100
	fmt.Println(t.Weekday(), int(t.Weekday()))    //Wednesday 3
	fmt.Println(t.ISOWeek())                      //2024 26

	//【时间戳】 使用t.UnixMilli()来做时间戳
	fmt.Println(t.Unix(), t.UnixMilli(), t.UnixMicro(), t.UnixNano()) //1719373890 1719373890743 1719373890743447 1719373890743447300

	//【】
	ts := time.Unix(1698462248, 0)     //unix 时间戳
	fmt.Println("unix timestamp:", ts) //unix timestamp: 2023-10-28 11:04:08 +0800 CST
	//Unix(sec int64, nsec int64) time.Time

	//【时区】
	tt := "2019/06/28 06:27:00"
	if tt, err := time.Parse("2006/01/02 15:04:05", tt); err == nil {
		fmt.Println(tt.UTC())   //2019-06-27 22:27:00 +0000 UTC
		fmt.Println(tt.Local()) //2019-06-28 14:27:00 +0800 CST
		fmt.Println(tt)         //2019-06-28 06:27:00 +0000 UTC 默认显示0时区
	} else {
		fmt.Println(err)
	}
	/*所以指定location定义时间*/
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

// 这个时间戳 "2019-06-28 06:27:00 +0800 CST" 可以按如下方式理解：

// "2019-06-28" 表示日期部分，即年份-月份-日期。
// "06:27:00" 表示时间部分，即时-分-秒。
// "+0800" 表示时区偏移量，即相对于 UTC 时间的偏移量。在这种情况下，偏移量为 +08:00，表示相对于 UTC 时间向东偏移 8 小时。
// "CST" 表示时区的缩写，代表中国标准时间（China Standard Time），它是 UTC+8 的时区。
// 综合起来，该时间戳表示的是中国标准时间（UTC+8）下的 2019 年 6 月 28 日 06:27:00 时刻。

// 这个时间戳 "2019-06-27 22:27:00 +0000 UTC" 可以按如下方式理解：

// "2019-06-27" 表示日期部分，即年份-月份-日期。
// "22:27:00" 表示时间部分，即时-分-秒。
// "+0000" 表示时区偏移量，即相对于 UTC 时间的偏移量。在这种情况下，偏移量为 +00:00，表示与 UTC 时间保持一致，没有时区偏移。
// "UTC" 表示协调世界时（Coordinated Universal Time），它是一个标准的全球时间标准，相当于我们通常所说的格林威治标准时间（GMT）。
// 综合起来，该时间戳表示的是协调世界时（UTC）下的 2019 年 6 月 27 日 22:27:00 时刻，由于时区偏移量为 +00:00，与格林威治标准时间（GMT）保持一致。


// A Month specifies a month of the year (January = 1, ...).
type Month int //新类型可以扩展方法

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
	//【时间运算】
	tz, _ := time.LoadLocation("Asia/Shanghai")
	s1 := "2024/06/26 09:00:00"
	s2 := "2024/06/26 10:00:00"
	layout := "2006/01/02 15:04:05"
	t1, _ := time.ParseInLocation(layout, s1, tz)
	t2, _ := time.ParseInLocation(layout, s2, tz)
	fmt.Println(t1) //2024-06-26 09:00:00 +0800 CST

	fmt.Println(t2) //2024-06-26 10:00:00 +0800 CST

	//时间差
	delta := t2.Sub(t1)                           //t2-t1
	fmt.Printf("delta: %v,%[1]T\n", delta)        //delta: 1h0m0s,time.Duration
	fmt.Println(delta.Minutes(), delta.Seconds()) //60 3600	共差多少小时和秒

	//构造Duration
	ns3 := time.Duration(3)              //3纳秒
	s3 := time.Duration(3 * time.Second) //3秒
	h3 := time.Duration(3 * time.Hour)   //3小时
	fmt.Println(ns3, s3, h3)             //3ns 3s 3h0m0s

	//时间偏移
	t3 := t2.Add(h3)
	fmt.Println("t3:", t3) //t3: 2024-06-26 13:00:00 +0800 CST
	t4 := t2.Add(-h3)
	fmt.Println("t4:", t4)                     //t4: 2024-06-26 07:00:00 +0800 CST
	fmt.Println("t3.After:", t3.After(t4))     //t3.After: true
	fmt.Println("time.Since:", time.Since(t2)) //time.Since: 3h44m47.1194582s

	//时间戳
	t := time.Unix(1, 0)
	fmt.Printf("%T %+[1]v\n", t) //time.Time 1970-01-01 08:00:01 +0800 CST
	fmt.Println(t.UTC())         //1970-01-01 00:00:01 +0000 UTC
}


【11-2-包管理】
=====================================================================
[模块化]
包管理
	使用目录组织包，目录就是包
		包里面，可以若干.go文件，.go文件中写package包名
	包名小写，符合标识符定义要求
	除main以外，建议包名就是目录名，可以不一样。不一样的话，用起来麻烦
	同一个目录就是同一个包
		要求所以 .go文件必须使用同一个包名
		但xxx_test包名除外
	
	main包特殊
		main函数必须在main包里
		可执行文件(windows .exe)里面有个表，入口字段，-> main函数地址
	
	项目管理
		go modules模式
			go mod init test : test 本地包
			go mod init magedu.com/wayne/test 远程仓库项目地址，网络包
			go get github.com/magedu/tools git
			import "github.com/magedu/tools"
			
			go.mod文件 所有包依赖管理在里面
				module 名称
				go 版本
		 		Windows版本 C:\Users\BIRKHOFF ALW\go\pkg\mod
		go mod vendor 使用vendoring机制，项目根目录下如果有vendor目录，优先搜索这个目录
		
		
		1.13版本之后默认on set GO111MODULE=

GOPROXY环境变量可以指定包下载镜像(镜像地址有时会变化，请参照官方最新文档)
GOPROXY=https://goproxy.cn,direct
GOPROXY=https://mirrors.aliyun.com/goproxy
GOPROXY=https://mirrors.cloud.tencent.com/go/
GOPROXY=https://repo.huaweicloud.com/repository/goproxy/

	F5工作有原理？背后时编译 Go build Go run 因为我们按照Go开发环境(标准库GOROOT/src 、 Go命令GOROOT/bin) 
		你写的源代码加上导入别人的源码(源代码合并到你的代码中) 一起编译
	导入方式
			"github.com/vmihailenco/msgpack/v5" //绝对导入
			 m "test/calc" m表示别名引用 别名导入
			 . "test/calc/minus" //.当前minus包中导出的全局变量导入到当前包
			 _ "test/calc/minus" //_没有名字 minus包中导出的全局变量导入到当前黑洞，也是所有导出无法使用，要它干嘛？	
			 										// 匿名导入，就是为了init()没法使用资源,

	init()函数
		可以写在任何.go文件中
		无参、无返回值
		一般一个go文件中只写一个，可以写多个
		不同的go文件中都可以写
		
		绝对导入、别名导入都可以指向init
		导入都可以执行init(),包括匿名导入
		
		应用：驱动程序，驱动程序包你自己使用匿名导入，init()被执行
		
		init函数，无参无返回值，不能被其他函数调用
		包中的init函数将在main函数之前自动执行
		每个包中init函数可以有多个，且可以位于不同的文件中
		同一个文件中可以有多个init函数，但一般一个就够了，不建议写多个
		同一个包中的init函数没有明确的执行顺序，不可预期
		不同包的init函数的执行顺序由导入顺序决定
		
		init函数主要是做一些初始化工作。init和main函数不一定在同一个文件中。
		import _ "xxx"作用是什么?只执行该包的init函数，无法使用包内资源。
		import "xxx"作用是什么?也会执行该包的init函数，也可以使用包内资源.。
		
		
[子包]
《main.go》
package main

import (
	"fmt"
	"test/calc"
	// "github.com/vmihailenco/msgpack/v5" //绝对导入
	// m "test/calc" m表示别名引用 别名导入
	// . "test/calc/minus" .当前minus包中导出的全局变量导入到当前包
	// _ "test/calc/minus" _没有名字 minus包中导出的全局变量导入到当前黑洞，也是所有导出无法使用，要它干嘛？
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

《calc/add.go》
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

《calc/calc.go》
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
用于将一个模块版本替换为另外一个模块版本


【11-3-log包】
=====================================================================
标准库1og
	日志库之所以能够输出，原因内部都会构建一个Logger日志记录器
	缺省日志记录器
		var std = New(os.stderr,""，LstdFlags) => *Logger
			1输出outputSetoutput
			2 prefix
			3 flags配置标记，各种标记提供一些能力
	func Default()*Logger { return std }// log.Default()->std *Logger 只读getter
	三个快捷方法
		log.Print * 0
		log.Fatal* log.Print* + os.Exit(1) 之后代码不再执行
		log.Panic* log.Print* + panic 2

	Logger类型
		Print*、Fatal*、Panic*
	日志文件输出
		只写writeonly
		os.0penFile(name string, flag int, perm os.FileMode)
			name 文件路径
			flag控制如何打开只读、只写、文件不在创建、文件在报错、追加写入、文件有内容清空
			移位创建的flag，可以组合使用
			os.O_CREATE文件不存在新建，存在不管
			oS.O_WRONLY 文件可写writable配合接口方法write方法
			oS.O_APPEND追加
	日志分析
		只读
		os.Open(name string)
	zerolog 比较简单
		一般都需要级别
		


[标准库]
// 使用缺省Logger
log.Print("abcde\n")
log.Printf("%s\n","abcd")
log.Printin("abc")
log.Fatal("xyz")// 等价于 1og.Print(“xyz");os.Exit(1)
logg.Panicln("Failed") // 等价于 log.Println("Failed");panic()


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
	log.Println("abc") //记录日志 退出状态码为0 status 0
	fmt.Println("abc~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	log.Fatalf("Fatal Code: %v \n", "aabbcc") //2024/06/28 10:46:04 Fatal Code: aabbcc  退出状态码为1 之后代码不再执行
	//Fatal 直接 os.Exit(1)
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	log.Panicf("Panic: %v \n", "paninc information") //panic 返回状态码为status 2
	fmt.Println("~~~~~~~~~~~~~~~~~~")

	log.Default() //func Default() *Logger { return std } => std *Logger

}

	log.Default().SetOutput(os.Stdout) //内部的std 、任何Logger都可以改
	log.Println("###abc###")           //2024/06/28 14:10:34 ###abc###
	
[自定义Logger]
如果觉得缺省Logger std不满意，可以New构建一个自定义Logger并指定前缀、Flags。

package main

import (
	"log"
	"os"
)

func main() {
	/*自定义Logger*/

	// log.New()
	// func log.New(out io.Writer, prefix string, flag int) *log.Logger
	// out: out, prefix: prefix, flag: flag
	// type Writer interface {
	// 	Write(p []byte) (n int, err error)
	// }
	infologger := log.New(os.Stdout, "Info: ", log.LstdFlags)
	infologger.Println("这是一个普通消息1") //Info: 2024/06/28 14:18:43 这是一个普通消息1

	////把自定义的放后面
	infologger = log.New(os.Stdout, "Info: ", log.LstdFlags|log.Lmsgprefix)
	infologger.Println("这是一个普通消息2") //2024/06/28 14:20:00 Info: 这是一个普通消息
}

[写日志文件]
New方法签名 New(out io.writer，prefix string，flag int)*Logger 中out参数提供Writer接口即可，那么就可以提供一个可写文件对象。

package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(
		"D:/my_nginx.log",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, //只写、文件不存在创建、追加
		os.ModePerm,                         //Unix permission bits，0o777
	) //读写 返回一个文件指针(句柄)
	if err != nil {
		log.Panicln(err)
	}

	defer f.Close() //不要用关闭
	l := log.New(f, "BIRKHOFF Logger", log.LstdFlags)
	l.Println("这是一个读写文件的消息 Logger Title")
}


【11-4-zerolog】
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


级别
	消息的级别
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
	级别
		消息的级别
			logger.Warn()、logger.Debug()
		logger的级别，New生成的Logger默认级别是trace -1
			logger.Level(?) => new child logger
			log.New() => default level trace -l
		gLevel -l trace
			SetGlobalLevel(?)

		消息输出条件:消息级别 >= Max(gLevel，logger.Level)
	
	zerolog/log 1og.Xxx()快捷方法
		log.Warn().Msg("warn string")
		一条记录 level":"warn","time":"2023-10-28T18:07:17+08:00","message":"warn string"
		Warn()=>"level":"warn"消息的级别
		Msg("warn string") => "message":"warn string"
		time字段一定要有
	
	var Logger = zerolog.New(os.stderr).with().Timestamp().Logger()包外可用，全局缺省
		log.Logger
		logger.Warn()
		logger.Debug()产生消息级别
		logger.Level(zerolog.XxxxLevel) 设置Logger级别

package main

import (
	// "log"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	/*全局Level trace*/
	fmt.Println(zerolog.GlobalLevel(), "####") //trace ####
	//设置全局Level
	// zerolog.SetGlobalLevel(zerolog.WarnLevel) //如果设置了warning下面的低于这个警告无法显示
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	fmt.Println(zerolog.GlobalLevel(), "####") //Info ####

	/*log.Logger default logger*/
	//zerolog输出格式debug级别
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
	fmt.Println(log.Logger.GetLevel()) //默认是：trace -1
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
	//因为设置了级别为Warning 所以比它低级别的不输出了
	// {"level":"warn","time":"2024-06-28T16:43:24+08:00","message":"mylog warning"}
	fmt.Println(mylog.GetLevel()) //默认是：warn
	fmt.Println("------------------------------")

	/*设置Level*/
	log1 := log.Level(zerolog.ErrorLevel) //因为设置了Error级别 所以下面都不输出
	fmt.Println(log1.GetLevel())          //error
	log1.Info().Msg("log1 Info msg")
	log1.Warn().Msg("log1 Warn msg")
	log1.Error().Msg("log1 Error msg")
}

		
	增加字段
		定义Logger时，With =>Context .Logger() => Logger  	
		使用Logger时，mylog.Warn().Msg("warn infomation")
		
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
	/*设置默认格式*/
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
	log1 := log.Level(zerolog.ErrorLevel).With().Caller().Logger() //Caller多打印 错误行数
	// {"level":"error","error":"自定义错误","time":"2024","caller":"e:/goprojects/main.go:35"}
	// {"level":"fatal","error":"自定义错误","time":"2024","caller":"e:/goprojects/main.go:37"}

	//log1.Error() 消息级别
	log1.Error().Err(errors.New("自定义错误")).Msg("")
	//{"level":"error","error":"自定义错误","time":"2024-06-28T18:37:04+08:00"}
	log1.Fatal().Err(errors.New("自定义错误")).Send()
	//{"level":"fatal","error":"自定义错误","time":"2024-06-28T18:38:21+08:00"}
}
			
		
[写文件]
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
		log.Panic().Err(err).Send() //内部调用panic
	}
	defer f.Close()

	multi := zerolog.MultiLevelWriter(f, os.Stdout) //多分支写 文件和控制台
	//Timesta()这个全新的lOgger增加时间戳输出
	logger := zerolog.New(multi).With().Timestamp().Logger()
	logger.Info().Msg("Write to Control Screen and File") //{"level":"info","time":1719571970,"message":"Write to Control Screen and File"}

}
		
[滚动日志]	标准库没有提供 

go get gopkg.in/natefinch/lumberjack.v2

	单个日志文件会膨胀
	滚动:生成一个新同名文件，把原来的文件重命名
		时间
		大小

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
		log.Panic().Err(err).Send() //内部调用panic
	}

	l := &lumberjack.Logger{
		Filename:   "d:/routine.log",
		MaxBackups: 2,     //出当前正在写入的日志文件外，历史日志最多保留2个
		Compress:   false, //缺省不压缩
		MaxAge:     1,     //1天
		MaxSize:    1,     //缺省值100 ，100M 超过1M立即滚动
	} //实现了Writer接口 得到一个自定义滚动Logger
	defer l.Close()

	multi := zerolog.MultiLevelWriter(f, os.Stdout, l) //多分支写 文件和控制台
	logger := zerolog.New(multi).With().Timestamp().Logger()
	for {
		time.Sleep(1 * time.Microsecond)
		t := time.Now().Format("[06-01-02 15:04:05 -0700]")
		logger.Info().Msg(t)
	}
}

[打印错误栈]
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
			log.Debug().Msg("没有错误")
		case runtime.Error:
			fmt.Println(string(debug.Stack()))
			log.Error().Caller(3).Err(v).Str("stack", string(debug.Stack())).Send()
		default:
			log.Debug().Msg(fmt.Sprintf("其他错误", v))
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

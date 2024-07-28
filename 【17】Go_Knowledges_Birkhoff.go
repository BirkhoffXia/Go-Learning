Viedo Date: 2023/12/9
Made By:	BIRKHOFF
Date:	2024-07-03

【17-1-ORM概念和数据库连接】
=====================================================================
对象关系映射(Object Relational Mapping，ORM)。指的是对象和关系之间的映射，使用面向对象的方式操作数据库。
关系模型和Go对象之间的映射
table  => struct    表映射为结构体 
row    => object    行映射为实例
column => property  字段映射为属性

* 可以认为ORM是一种高级抽象，操作的结构会被封装成对象。
	对象的操作最终还是会转换成对应关系数据库操作的SQL语句，数据库

[GORM]
GORM是一个友好的、功能全面的、性能不错的基于Go语言实现的ORM库

go get -u github.com/go-sql-driver/mysql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

EG:
package main

import (
	//标准库提供的统一编程接口
	"fmt"
	"log"

	"gorm.io/driver/mysql" //驱动 针对不同数据库做适配
	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)

/*
	以上代码难道不需要使用驱动吗?
	在"gorm.io/driver/mysql/mysql.go"中
	import了"github.com/go-sql-driver/mysql",也就是说驱动也导入了Dialector的Initialize方法中使用了sql.Open
*/

var db *gorm.DB

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

	// db: &{0xc00013c510 <nil> 0 0xc0001f8000 1}
	// ==============
}

func main() {
	fmt.Println(db) //&{0xc000170510 <nil> 0 0xc000240000 1}
}


【17-2-模型定义和迁移】
=====================================================================
官网：https://gorm.io/zh_CN/docs/models.html

GORM 倾向于约定优于配置
	约定使用名为ID的属性会作为主键
	约定使用snake_cases作为表名，结构体名使用驼峰命名
		结构体命名为employee，那么数据库表名就是employees
	约定使用snake_case作为字段名，字段采用首字母大写的大驼峰命名
		属性名为FirstName，默认对应数据库表的字段名为first_name

约定
	约定大于配置，如果你遵守约定可以少些很多代码
	构建模型
		最好按照约定来，否则要手动配置
		建立Model类、结构体(注意首字母大小写)，定义好字段(首字母大写)
	配置

EG：
package main

import (
	//标准库提供的统一编程接口
	"fmt"
	"log"

	"gorm.io/driver/mysql" //驱动 针对不同数据库做适配
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

/*
	以上代码难道不需要使用驱动吗?
	在"gorm.io/driver/mysql/mysql.go"中
	import了"github.com/go-sql-driver/mysql",也就是说驱动也导入了Dialector的Initialize方法中使用了sql.Open
*/

var db *gorm.DB

// 可以看出下列属性名顺序无所谓，属性名首字母要大写
type Emp struct {
	//使用 gorm:"primaryKey"来指定字段为主键，默认使用名为ID的属性作为主键。primaryKey是tag名大小写不敏感，但建议小驼峰。
	EmpNo     int    `gorm:"primaryKey"` //不是ID为主键
	FirstName string //首字母大写，对应字段first_name
	LastName  string
	Gender    byte
	// BirthDate string
	// BirthDate string `gorm:"column:birth_date"`
	//如果未按照约定定义字段，需要定义结构体属性时指定数据库字段名称时什么
	Xyz string `gorm:"column:birth_date"`
}

func (Emp) TableName() string {
	return "employees"
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //日志级别，默认为Slient 即打印慢SQL和错误
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

	// db: &{0xc00013c510 <nil> 0 0xc0001f8000 1}
	// ==============
}

func main() {
	fmt.Println(db) //&{0xc000170510 <nil> 0 0xc000240000 1}
	var e Emp
	result := db.Take(&e) //等价于Limit 1 ，取1条
	fmt.Println("result : ", result)
	fmt.Println("result.Error : ", result.Error)
	fmt.Println("result.RowsAffected : ", result.RowsAffected)
	fmt.Println("Value : ", e)
}

// db: &{0xc00013c6c0 <nil> 0 0xc0001f8000 1}
// ==============
// &{0xc00013c6c0 <nil> 0 0xc0001f8000 1}

// 2024/07/03 13:21:44 e:/goprojects/main.go:54
// [1.017ms] [rows:1] SELECT * FROM `employees` LIMIT 1
// result :  &{0xc00013c6c0 <nil> 1 0xc0001f81c0 0}
// result.Error :  <nil>
// result.RowsAffected :  1
// Value :  {10001 Georgi Facello 2 1953-09-02}


[迁移]
迁移Migration
	代码中Model生成数据库中的表
		Model名称->表名
		Model中定义的属性->表中的字段
EG：
package main

import (
	//标准库提供的统一编程接口
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //驱动 针对不同数据库做适配
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int       //缺省逐渐bigint AUTO_INCREMENT
	Name     string    `gorm:"not null;type:varchar(48);comment:姓名"`
	Age      byte      //byte => tinyint unsigned
	Birthday time.Time //datetime 也可以用时间戳int64
	Gender   byte      `gorm:"type:tinyint"`
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //日志级别，默认为Slient 即打印慢SQL和错误
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

func main() {

	db.Migrator().CreateTable(
		&Student{},
	)
	/*
		//执行CreateTable
					CREATE TABLE `students` (
								`id` bigint AUTO_INCREMENT,
								`name` varchar(48) NOT NULL COMMENT '姓名',
								`age` tinyint unsigned,
								`birthday` datetime(3) NULL,
								`gender` tinyint,
								PRIMARY KEY (`id`))
				//DDL
							CREATE TABLE `students` (
						  `id` bigint(20) NOT NULL AUTO_INCREMENT,
						  `name` varchar(48) NOT NULL COMMENT '姓名',
						  `age` tinyint(3) unsigned DEFAULT NULL,
						  `birthday` datetime(3) DEFAULT NULL,
						  `gender` tinyint(4) DEFAULT NULL,
						  PRIMARY KEY (`id`)
						) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	*/
}

【17-3-新增和时间时区处理】
=====================================================================
[新增]
EG：
package main

import (
	//标准库提供的统一编程接口
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //驱动 针对不同数据库做适配
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int        //缺省逐渐bigint AUTO_INCREMENT
	Name     string     `gorm:"not null;type:varchar(48);comment:姓名"`
	Age      byte       //byte => tinyint unsigned
	Birthday *time.Time //datetime 也可以用时间戳int64
	Gender   byte       `gorm:"type:tinyint"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d : %s %d", s.ID, s.Name, s.Age)
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //日志级别，默认为Slient 即打印慢SQL和错误
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

func main() {
	n := time.Now()
	s := Student{Name: "Tom", Age: 20, Birthday: &n} //构建实例
	fmt.Println(s)
	// result := db.Create(&s) //新增一条
	result := db.Create([]*Student{&s, &s, &s}) //新增多条
	fmt.Println(s)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

// db: &{0xc00013c5a0 <nil> 0 0xc0001f8000 1}
// ==============
// {0 Tom 20 2024-07-03 15:17:05.3666566 +0800 CST m=+0.007288001 0}

// 2024/07/03 15:17:05 e:/goprojects/main.go:47
// [3.819ms] [rows:1] INSERT INTO `students` (`name`,`age`,`birthday`,`gender`) VALUES ('Tom',20,'2024-07-03 07:17:05.366',0)
// {1 Tom 20 2024-07-03 15:17:05.3666566 +0800 CST m=+0.007288001 0}
// <nil>
// 1

[查询一条]
package main

import (
	//标准库提供的统一编程接口
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //驱动 针对不同数据库做适配
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int        //缺省逐渐bigint AUTO_INCREMENT
	Name     string     `gorm:"not null;type:varchar(48);comment:姓名"`
	Age      byte       //byte => tinyint unsigned
	Birthday *time.Time //datetime 也可以用时间戳int64
	Gender   byte       `gorm:"type:tinyint"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d : %s %d", s.ID, s.Name, s.Age)
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //日志级别，默认为Slient 即打印慢SQL和错误
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

func main() {
	var s Student
	fmt.Println(s) //零值
	// r := db.Take(&s) //LIMIT 1
	// r := db.First(&s) //SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	// r := db.Last(&s) //SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	s = Student{ID: 2}
	r := db.Take(&s) //加入条件

	fmt.Println("Fetch : ", s)
	fmt.Println("Local:", s.Birthday.Format("2006-01-02 15:04:05 -0700")) //Local: 2024-07-03 07:05:05 +0000
	fmt.Println("Process After time: ", s.Birthday.Local().Format("2006-01-02 15:04:05 -0700"))
	//[加loc=Local]
	// Local:               2024-07-03 07:05:05 +0800
	// Process After time:  2024-07-03 07:05:05 +0800
	//[不加loc=Local]
	// Local:               2024-07-03 07:17:05 +0000
	// Process After time:  2024-07-03 15:17:05 +0800
	fmt.Println(r)
	//时间错误 加入charset=utf8mb4&parseTime=true
	fmt.Println("Error : ", r.Error)
	//Error :  sql: Scan error on column index 3, name "birthday": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	fmt.Println(r.RowsAffected)
}

【17-4-查询、分组、聚合、Join】&&【17-5-更新、删除】
=====================================================================
EG:
package main

import (
	//标准库提供的统一编程接口
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //驱动 针对不同数据库做适配
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int        //缺省逐渐bigint AUTO_INCREMENT
	Name     string     `gorm:"not null;type:varchar(48);comment:姓名"`
	Age      byte       //byte => tinyint unsigned
	Birthday *time.Time //datetime 也可以用时间戳int64
	Gender   byte       `gorm:"type:tinyint"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d : %s %d", s.ID, s.Name, s.Age)
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //日志级别，默认为Slient 即打印慢SQL和错误
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

// 【查询多行】
func FindMany() {
	var students []*Student
	r := db.Find(&students)
	fmt.Println(r.Error)
	fmt.Println(students) //[1 : Tom 20 2 : Tom 20 3 : Tom 20 4 : Tom 20 5 : Tom 20]
}

// 【去重 Distinct】
func distinct() {
	var students []*Student
	r := db.Distinct("id").Find(&students)
	fmt.Println(r.Error)
	fmt.Println(students)
}

// 【投影】
func Projection() {
	var students []*Student
	// db.Select("id", "name", "age").Offset(2).Limit(2).Find(&students) //[0.751ms] [rows:2] SELECT `id`,`name`,`age` FROM `students` LIMIT 2 OFFSET 2
	db.Select([]string{"id", "name", "age"}).Offset(2).Limit(2).Find(&students)
	fmt.Println(students)
}

// 【条件 Where】
func Where() {
	//and、or、not、like、between and 、in
	var students []*Student
	// db.Find(&students, 1, 3) //[]int{1,3}
	// db.Where("id = 1 or id = 3").Find(&students)
	// db.Where("id = ? or id = ?", 3, 4).Find(&students) //SELECT * FROM `students` WHERE id = 3 or id = 4
	// db.Where("id in (?,?)", 1, 3).Find(&students) //SELECT * FROM `students` WHERE id in (1,3)
	// db.Where([]int{1,3}).Find(&students)
	// db.Find(&students) //all
	// db.Where(&Student{}).Find(&students)
	// db.Where(&Student{ID: 3}).Find(&students)
	// db.Where(&Student{ID: 0, Name: "Tom"}).Find(&students) //零值忽略
	// db.Where(map[string]any{"id": 1, "name": "Tom"}).Find(&students)
	// db.Not("id in (?,?)", 1, 3).Find(&students)
	// db.Where("name = ?", "Tom").Or("name = ?", "Jerry").Find(&students) //Or
	db.Where("name = ?", "Tom").Or(&Student{Name: "Jerry"}).Find(&students) //Or
	fmt.Println(students)
}

// 【排序】
func OrderBy() {
	var students []*Student
	// db.Order("name desc,id desc").Find(&students)
	db.Order("name").Order("id desc").Find(&students)
	fmt.Println(students)
}

// 【分组】
// 自定义个结构体 为了装name,count(id) as c
type Result struct {
	Name string
	C    int `gorm:"colume:c"`
}

func (s Result) String() string {
	return fmt.Sprintf("%s %d", s.Name, s.C)
}

func GroupBy() {
	// var students []*Student
	// db.Group("id").Find(&students)
	// db.Group("name").Find(&students) //SELECT * FROM `students` GROUP BY `name`
	var Results []*Result
	// db.Table("students").Select("name,count(id) as c").Group("name").Find(&Results)
	// rows, _ := db.Table("students").Select("name,count(id) as c").Group("name").Rows()
	rows, _ := db.Table("students").Select("name,count(id) as c").Group("name").Having("c > ?", 1).Rows()

	for rows.Next() {
		var r Result
		rows.Scan(&r.Name, &r.C)
		fmt.Println(r)
	}
	fmt.Println(Results)

}

// 【Join】
type ResultJoin struct {
	EmpNo  int
	Name   string
	Salary int
}

func (s ResultJoin) String() string {
	return fmt.Sprintf("%d - %s - %d\n", s.EmpNo, s.Name, s.Salary)
}

func Join() {
	var Results []*ResultJoin
	r := db.Table("employees as e").Select("e.emp_no,concat(first_name,' ',last_name) as name,salary").
		Joins("join salaries as s on e.emp_no = s.emp_no").Find(&Results)
	if r.Error != nil {
		panic(r.Error)
	}
	fmt.Println(Results)
}

// 【更新-Update】
func Update1() {
	//持久态，数据中有该条记录
	// var s = Student{ID: 1} //没有和数据库数据做任何关联，新建对象
	// s.Name = "XIAKAISHENG"
	var s Student
	fmt.Println(s.ID, s)
	db.Take(&s, 1)
	fmt.Println(s.ID, s) //查有此数据id=1 ,持久化的
	s.Name = "XIAKAISHENG"
	s.Age += 1
	db.Save(&s) //有些ORM库报错的，有些不报错
	//[2.046ms] [rows:1] UPDATE `students` SET `name`='XIAKAISHENG',`age`=89,`birthday`='2024-07-03 07:17:05.366',`gender`=1 WHERE `id` = 1
}
func Update2() {
	//更新一个字段
	r := db.Model(&Student{}).Where(1).Update("name", "Sammy") //UPDATE `students` SET `name`='Sammy' WHERE `students`.`id` = 1
	// r := db.Model(&Student{ID: 1}).Update("name", "Sam")
	fmt.Println(r.RowsAffected)
}
func Update3() {
	//多个字段
	// r := db.Model(&Student{ID: 4}).Updates(map[string]any{"name": "john", "age": 22}) //[rows:1] UPDATE `students` SET `age`=22,`name`='john' WHERE `id` = 4
	r := db.Model(&Student{ID: 4}).Updates(Student{Name: "JUICE", Age: 18})
	fmt.Println(r.RowsAffected)
}

// 【删除 Delete】
func Del1() {
	// db.Where("id>0").Delete(&Student{})
	r := db.Delete(&Student{})
	fmt.Println(r.Error, "#####")
}
func Del2() {
	// db.Delete(&Student{}, []int{3, 4})
	db.Where([]int{3, 4}).Delete(&Student{}) //[3.120ms] [rows:2] DELETE FROM `students` WHERE `students`.`id` IN (3,4)
}

func main() {
	// FindMany()
	// distinct() //SELECT DISTINCT `id` FROM `students`
	// Projection()
	// Where()
	// OrderBy()
	// GroupBy()
	// Join()
	// Update1()
	// Update2()
	// Update3()
	// Del2()

	var s Student
	fmt.Println(s) //零值
	// r := db.Take(&s) //LIMIT 1
	// r := db.First(&s) //SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	// r := db.Last(&s) //SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	s = Student{ID: 2}
	r := db.Take(&s) //加入条件

	fmt.Println("Fetch : ", s)
	fmt.Println("Local:", s.Birthday.Format("2006-01-02 15:04:05 -0700")) //Local: 2024-07-03 07:05:05 +0000
	fmt.Println("Process After time: ", s.Birthday.Local().Format("2006-01-02 15:04:05 -0700"))
	//[加loc=Local]
	// Local:               2024-07-03 07:05:05 +0800
	// Process After time:  2024-07-03 07:05:05 +0800
	//[不加loc=Local]
	// Local:               2024-07-03 07:17:05 +0000
	// Process After time:  2024-07-03 15:17:05 +0800
	fmt.Println(r)
	//时间错误 加入charset=utf8mb4&parseTime=true
	fmt.Println("Error : ", r.Error)
	//Error :  sql: Scan error on column index 3, name "birthday": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	fmt.Println(r.RowsAffected)
}

【17-6-MongoDB及数据库连接】
=====================================================================
MongoDB
	NOSOL文档型 
	Bson格式
	分布式集群
	对比 关系型数据库
		关系型数据库字段时固定的。会提前定义预留字段
		MongoDB 字段随意

Windows安装-启动

EG:
package main

import (
	//标准库提供的统一编程接口
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	dsn := "mongodb://localhost:27017/" //客户端可以读取全局环境变量或者配置文件
	opts := options.Client()
	opts.ApplyURI(dsn).SetConnectTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), opts) //localhost:27017 配置服务端指定用户名和密码
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
}
func main() {
	fmt.Println("abc")
}















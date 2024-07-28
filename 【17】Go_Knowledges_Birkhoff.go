Viedo Date: 2023/12/9
Made By:	BIRKHOFF
Date:	2024-07-03

��17-1-ORM��������ݿ����ӡ�
=====================================================================
�����ϵӳ��(Object Relational Mapping��ORM)��ָ���Ƕ���͹�ϵ֮���ӳ�䣬ʹ���������ķ�ʽ�������ݿ⡣
��ϵģ�ͺ�Go����֮���ӳ��
table  => struct    ��ӳ��Ϊ�ṹ�� 
row    => object    ��ӳ��Ϊʵ��
column => property  �ֶ�ӳ��Ϊ����

* ������ΪORM��һ�ָ߼����󣬲����Ľṹ�ᱻ��װ�ɶ���
	����Ĳ������ջ��ǻ�ת���ɶ�Ӧ��ϵ���ݿ������SQL��䣬���ݿ�

[GORM]
GORM��һ���Ѻõġ�����ȫ��ġ����ܲ���Ļ���Go����ʵ�ֵ�ORM��

go get -u github.com/go-sql-driver/mysql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

EG:
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"

	"gorm.io/driver/mysql" //���� ��Բ�ͬ���ݿ�������
	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)

/*
	���ϴ����ѵ�����Ҫʹ��������?
	��"gorm.io/driver/mysql/mysql.go"��
	import��"github.com/go-sql-driver/mysql",Ҳ����˵����Ҳ������Dialector��Initialize������ʹ����sql.Open
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


��17-2-ģ�Ͷ����Ǩ�ơ�
=====================================================================
������https://gorm.io/zh_CN/docs/models.html

GORM ������Լ����������
	Լ��ʹ����ΪID�����Ի���Ϊ����
	Լ��ʹ��snake_cases��Ϊ�������ṹ����ʹ���շ�����
		�ṹ������Ϊemployee����ô���ݿ��������employees
	Լ��ʹ��snake_case��Ϊ�ֶ������ֶβ�������ĸ��д�Ĵ��շ�����
		������ΪFirstName��Ĭ�϶�Ӧ���ݿ����ֶ���Ϊfirst_name

Լ��
	Լ���������ã����������Լ��������Щ�ܶ����
	����ģ��
		��ð���Լ����������Ҫ�ֶ�����
		����Model�ࡢ�ṹ��(ע������ĸ��Сд)��������ֶ�(����ĸ��д)
	����

EG��
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"

	"gorm.io/driver/mysql" //���� ��Բ�ͬ���ݿ�������
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

/*
	���ϴ����ѵ�����Ҫʹ��������?
	��"gorm.io/driver/mysql/mysql.go"��
	import��"github.com/go-sql-driver/mysql",Ҳ����˵����Ҳ������Dialector��Initialize������ʹ����sql.Open
*/

var db *gorm.DB

// ���Կ�������������˳������ν������������ĸҪ��д
type Emp struct {
	//ʹ�� gorm:"primaryKey"��ָ���ֶ�Ϊ������Ĭ��ʹ����ΪID��������Ϊ������primaryKey��tag����Сд�����У�������С�շ塣
	EmpNo     int    `gorm:"primaryKey"` //����IDΪ����
	FirstName string //����ĸ��д����Ӧ�ֶ�first_name
	LastName  string
	Gender    byte
	// BirthDate string
	// BirthDate string `gorm:"column:birth_date"`
	//���δ����Լ�������ֶΣ���Ҫ����ṹ������ʱָ�����ݿ��ֶ�����ʱʲô
	Xyz string `gorm:"column:birth_date"`
}

func (Emp) TableName() string {
	return "employees"
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //��־����Ĭ��ΪSlient ����ӡ��SQL�ʹ���
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
	result := db.Take(&e) //�ȼ���Limit 1 ��ȡ1��
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


[Ǩ��]
Ǩ��Migration
	������Model�������ݿ��еı�
		Model����->����
		Model�ж��������->���е��ֶ�
EG��
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //���� ��Բ�ͬ���ݿ�������
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int       //ȱʡ��bigint AUTO_INCREMENT
	Name     string    `gorm:"not null;type:varchar(48);comment:����"`
	Age      byte      //byte => tinyint unsigned
	Birthday time.Time //datetime Ҳ������ʱ���int64
	Gender   byte      `gorm:"type:tinyint"`
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //��־����Ĭ��ΪSlient ����ӡ��SQL�ʹ���
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
		//ִ��CreateTable
					CREATE TABLE `students` (
								`id` bigint AUTO_INCREMENT,
								`name` varchar(48) NOT NULL COMMENT '����',
								`age` tinyint unsigned,
								`birthday` datetime(3) NULL,
								`gender` tinyint,
								PRIMARY KEY (`id`))
				//DDL
							CREATE TABLE `students` (
						  `id` bigint(20) NOT NULL AUTO_INCREMENT,
						  `name` varchar(48) NOT NULL COMMENT '����',
						  `age` tinyint(3) unsigned DEFAULT NULL,
						  `birthday` datetime(3) DEFAULT NULL,
						  `gender` tinyint(4) DEFAULT NULL,
						  PRIMARY KEY (`id`)
						) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	*/
}

��17-3-������ʱ��ʱ������
=====================================================================
[����]
EG��
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //���� ��Բ�ͬ���ݿ�������
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int        //ȱʡ��bigint AUTO_INCREMENT
	Name     string     `gorm:"not null;type:varchar(48);comment:����"`
	Age      byte       //byte => tinyint unsigned
	Birthday *time.Time //datetime Ҳ������ʱ���int64
	Gender   byte       `gorm:"type:tinyint"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d : %s %d", s.ID, s.Name, s.Age)
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //��־����Ĭ��ΪSlient ����ӡ��SQL�ʹ���
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

func main() {
	n := time.Now()
	s := Student{Name: "Tom", Age: 20, Birthday: &n} //����ʵ��
	fmt.Println(s)
	// result := db.Create(&s) //����һ��
	result := db.Create([]*Student{&s, &s, &s}) //��������
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

[��ѯһ��]
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //���� ��Բ�ͬ���ݿ�������
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int        //ȱʡ��bigint AUTO_INCREMENT
	Name     string     `gorm:"not null;type:varchar(48);comment:����"`
	Age      byte       //byte => tinyint unsigned
	Birthday *time.Time //datetime Ҳ������ʱ���int64
	Gender   byte       `gorm:"type:tinyint"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d : %s %d", s.ID, s.Name, s.Age)
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //��־����Ĭ��ΪSlient ����ӡ��SQL�ʹ���
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

func main() {
	var s Student
	fmt.Println(s) //��ֵ
	// r := db.Take(&s) //LIMIT 1
	// r := db.First(&s) //SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	// r := db.Last(&s) //SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	s = Student{ID: 2}
	r := db.Take(&s) //��������

	fmt.Println("Fetch : ", s)
	fmt.Println("Local:", s.Birthday.Format("2006-01-02 15:04:05 -0700")) //Local: 2024-07-03 07:05:05 +0000
	fmt.Println("Process After time: ", s.Birthday.Local().Format("2006-01-02 15:04:05 -0700"))
	//[��loc=Local]
	// Local:               2024-07-03 07:05:05 +0800
	// Process After time:  2024-07-03 07:05:05 +0800
	//[����loc=Local]
	// Local:               2024-07-03 07:17:05 +0000
	// Process After time:  2024-07-03 15:17:05 +0800
	fmt.Println(r)
	//ʱ����� ����charset=utf8mb4&parseTime=true
	fmt.Println("Error : ", r.Error)
	//Error :  sql: Scan error on column index 3, name "birthday": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	fmt.Println(r.RowsAffected)
}

��17-4-��ѯ�����顢�ۺϡ�Join��&&��17-5-���¡�ɾ����
=====================================================================
EG:
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" //���� ��Բ�ͬ���ݿ�������
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Student struct {
	ID       int        //ȱʡ��bigint AUTO_INCREMENT
	Name     string     `gorm:"not null;type:varchar(48);comment:����"`
	Age      byte       //byte => tinyint unsigned
	Birthday *time.Time //datetime Ҳ������ʱ���int64
	Gender   byte       `gorm:"type:tinyint"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d : %s %d", s.ID, s.Name, s.Age)
}

func init() {
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //��־����Ĭ��ΪSlient ����ӡ��SQL�ʹ���
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("db: %v\n", db)
	fmt.Println("==============")

}

// ����ѯ���С�
func FindMany() {
	var students []*Student
	r := db.Find(&students)
	fmt.Println(r.Error)
	fmt.Println(students) //[1 : Tom 20 2 : Tom 20 3 : Tom 20 4 : Tom 20 5 : Tom 20]
}

// ��ȥ�� Distinct��
func distinct() {
	var students []*Student
	r := db.Distinct("id").Find(&students)
	fmt.Println(r.Error)
	fmt.Println(students)
}

// ��ͶӰ��
func Projection() {
	var students []*Student
	// db.Select("id", "name", "age").Offset(2).Limit(2).Find(&students) //[0.751ms] [rows:2] SELECT `id`,`name`,`age` FROM `students` LIMIT 2 OFFSET 2
	db.Select([]string{"id", "name", "age"}).Offset(2).Limit(2).Find(&students)
	fmt.Println(students)
}

// ������ Where��
func Where() {
	//and��or��not��like��between and ��in
	var students []*Student
	// db.Find(&students, 1, 3) //[]int{1,3}
	// db.Where("id = 1 or id = 3").Find(&students)
	// db.Where("id = ? or id = ?", 3, 4).Find(&students) //SELECT * FROM `students` WHERE id = 3 or id = 4
	// db.Where("id in (?,?)", 1, 3).Find(&students) //SELECT * FROM `students` WHERE id in (1,3)
	// db.Where([]int{1,3}).Find(&students)
	// db.Find(&students) //all
	// db.Where(&Student{}).Find(&students)
	// db.Where(&Student{ID: 3}).Find(&students)
	// db.Where(&Student{ID: 0, Name: "Tom"}).Find(&students) //��ֵ����
	// db.Where(map[string]any{"id": 1, "name": "Tom"}).Find(&students)
	// db.Not("id in (?,?)", 1, 3).Find(&students)
	// db.Where("name = ?", "Tom").Or("name = ?", "Jerry").Find(&students) //Or
	db.Where("name = ?", "Tom").Or(&Student{Name: "Jerry"}).Find(&students) //Or
	fmt.Println(students)
}

// ������
func OrderBy() {
	var students []*Student
	// db.Order("name desc,id desc").Find(&students)
	db.Order("name").Order("id desc").Find(&students)
	fmt.Println(students)
}

// �����顿
// �Զ�����ṹ�� Ϊ��װname,count(id) as c
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

// ��Join��
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

// ������-Update��
func Update1() {
	//�־�̬���������и�����¼
	// var s = Student{ID: 1} //û�к����ݿ��������κι������½�����
	// s.Name = "XIAKAISHENG"
	var s Student
	fmt.Println(s.ID, s)
	db.Take(&s, 1)
	fmt.Println(s.ID, s) //���д�����id=1 ,�־û���
	s.Name = "XIAKAISHENG"
	s.Age += 1
	db.Save(&s) //��ЩORM�ⱨ��ģ���Щ������
	//[2.046ms] [rows:1] UPDATE `students` SET `name`='XIAKAISHENG',`age`=89,`birthday`='2024-07-03 07:17:05.366',`gender`=1 WHERE `id` = 1
}
func Update2() {
	//����һ���ֶ�
	r := db.Model(&Student{}).Where(1).Update("name", "Sammy") //UPDATE `students` SET `name`='Sammy' WHERE `students`.`id` = 1
	// r := db.Model(&Student{ID: 1}).Update("name", "Sam")
	fmt.Println(r.RowsAffected)
}
func Update3() {
	//����ֶ�
	// r := db.Model(&Student{ID: 4}).Updates(map[string]any{"name": "john", "age": 22}) //[rows:1] UPDATE `students` SET `age`=22,`name`='john' WHERE `id` = 4
	r := db.Model(&Student{ID: 4}).Updates(Student{Name: "JUICE", Age: 18})
	fmt.Println(r.RowsAffected)
}

// ��ɾ�� Delete��
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
	fmt.Println(s) //��ֵ
	// r := db.Take(&s) //LIMIT 1
	// r := db.First(&s) //SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	// r := db.Last(&s) //SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	s = Student{ID: 2}
	r := db.Take(&s) //��������

	fmt.Println("Fetch : ", s)
	fmt.Println("Local:", s.Birthday.Format("2006-01-02 15:04:05 -0700")) //Local: 2024-07-03 07:05:05 +0000
	fmt.Println("Process After time: ", s.Birthday.Local().Format("2006-01-02 15:04:05 -0700"))
	//[��loc=Local]
	// Local:               2024-07-03 07:05:05 +0800
	// Process After time:  2024-07-03 07:05:05 +0800
	//[����loc=Local]
	// Local:               2024-07-03 07:17:05 +0000
	// Process After time:  2024-07-03 15:17:05 +0800
	fmt.Println(r)
	//ʱ����� ����charset=utf8mb4&parseTime=true
	fmt.Println("Error : ", r.Error)
	//Error :  sql: Scan error on column index 3, name "birthday": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	fmt.Println(r.RowsAffected)
}

��17-6-MongoDB�����ݿ����ӡ�
=====================================================================
MongoDB
	NOSOL�ĵ��� 
	Bson��ʽ
	�ֲ�ʽ��Ⱥ
	�Ա� ��ϵ�����ݿ�
		��ϵ�����ݿ��ֶ�ʱ�̶��ġ�����ǰ����Ԥ���ֶ�
		MongoDB �ֶ�����

Windows��װ-����

EG:
package main

import (
	//��׼���ṩ��ͳһ��̽ӿ�
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	dsn := "mongodb://localhost:27017/" //�ͻ��˿��Զ�ȡȫ�ֻ����������������ļ�
	opts := options.Client()
	opts.ApplyURI(dsn).SetConnectTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), opts) //localhost:27017 ���÷����ָ���û���������
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















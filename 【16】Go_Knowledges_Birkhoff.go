Viedo Date: 2023/11/18
Made By:	BIRKHOFF
Date:	2024-07-02


【16-1-SQL之增删改】
=====================================================================
【16-2-SQL之分页、条件】
=====================================================================
【16-3-SQL之排序、分组、聚合】
=====================================================================
【16-4-SQL之子查询和Join】
=====================================================================
【16-5-数据库标准库开发】
=====================================================================
ALTER USER 'root'@'localhost' IDENTIFIED BY 'sheca';

安装mysql的Go驱动
go get -u github.com/go-sql-driver/mysql
go: downloading github.com/go-sql-driver/mysql v1.8.1
go: downloading filippo.io/edwards25519 v1.1.0
go: added filippo.io/edwards25519 v1.1.0
go: added github.com/go-sql-driver/mysql v1.8.1

导入

连接
package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //1 驱动安装和导入
)

/*db类型是*sq1.DB，是一个操作数据库的句柄，底层是一个多协程安全的连接池。*/
var db *sql.DB

func init() {
	//2 连接数据库和配置
	var err error
	db, err = sql.Open("mysql", "root:sheca@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Second * 30) //超时事件
	db.SetMaxOpenConns(0)                   //设置最大连接数，默认为0表示不限制
	db.SetMaxIdleConns(10)                  //设置空闲连接数
}

[操作]- 增删改查 
EG：
package main

import (
	"database/sql" //标准库提供的统一编程接口
	"fmt"
	"log"
	"time"

	//[1].驱动安装和导入
	_ "github.com/go-sql-driver/mysql"
	//第三方驱动，TCP连接和mysql协议。只执行该包的init函数，资源包括结构体、接口、全局变量常量都用不了
	//init函数中一定用了map，往map中注册一个名字和驱动对象的映射，名字就叫mysql。map是标准库sql包drivers映射
)

/* db类型是*sq1.DB，是一个操作数据库的句柄，底层是一个多协程安全的连接池。*/
var db *sql.DB

func init() {
	//[2].连接数据库和配置
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = sql.Open("mysql", dsn) //这里的db使用的是全局变量 不能用:=
	// db, err = sql.Open("mysql", "root:sheca@(192.168.40.103:3306)/test") //可省略tcp

	if err != nil {
		log.Fatal(err)
	}
	//
	db.SetConnMaxLifetime(time.Second * 30) //超时事件
	db.SetMaxOpenConns(0)                   //设置最大连接数，默认为0表示不限制
	db.SetMaxIdleConns(10)                  //设置空闲连接数
}

// [3].定义结构体
type Emp struct { //和字段对应的变量或结构体定义，最好和数据库字段顺序对应
	emp_no     int
	birth_date string
	first_name string
	last_name  string
	gender     int16
	hire_date  string
}

func main() {
	//单行查询
	emp := Emp{}                                                          //定义实例
	row := db.QueryRow("select * from employees where emp_no = ?", 10010) //查询行 只能查一行 Limit-1
	if row.Err() != nil {
		log.Fatal(row.Err().Error())
	}
	err := row.Scan(&emp.emp_no, &emp.birth_date, &emp.first_name,
		&emp.last_name, &emp.gender, &emp.hire_date) //字段顺序要手动对应一致
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emp) //
	fmt.Println("================================")

	// {10010 1963-06-01 Duangkaew Piveteau 2 1989-08-24}
	// ================================

	//批量查询，预编译
	//有字符串的拼接使用 1.Prepare 2.Query
	//SQL注入攻击 select * from employees where emp_no = 10010 or 1=1
	stmt, err := db.Prepare("select * from employees where emp_no > ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//Query使用给定的参数执行准备好的查询语句，并以*Rows的形式返回查询结果。
	//查询使用上下文。内部背景；要指定上下文，请使用QueryContext。
	rows, err := stmt.Query(10018) //这里直接给参数
	// query := "select * from employees"
	// rows, err := db.Query(query,1,2) //这里可以传入sql和参数

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() { //遍历，每一趟rows内部指向当前行 相当于游标
		//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
		err := rows.Scan(&emp.emp_no, &emp.birth_date, &emp.first_name,
			&emp.last_name, &emp.gender, &emp.hire_date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(emp)
		fmt.Printf("Transfer Before [%T] [%[1]v]; ", emp.birth_date)
		t, err := time.Parse("2006-01-02", emp.birth_date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Transfer After [%T] [%[1]v]\n", t)
	}

	//更新数据
	result, err := db.Exec("UPDATE employees set gender=? where emp_no=?", 2, 10001)
	if err != nil {
		fmt.Printf("Insert failed, err: %v\n", err)
		return
	}
	successrowupdate, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Println("success row update:", successrowupdate) //success row update: 1

	//插入
	result1, err := db.Exec("Insert into employees(emp_no,birth_date,first_name,last_name,gender,hire_date) values(?,?,?,?,?,?)", 10021, "1994-03-19", "KaiSheng", "Xia", 1, "2015-12-01")
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result1.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	successrowupdate1, err := result1.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", successrowupdate1) //RowsAffected: 1

	//删除
	result2, err := db.Exec("delete from employees where emp_no=?", 10021)
	if err != nil {
		fmt.Printf("Delete failed,err:%v", err)
		return
	}
	successrowupdate2, err := result2.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", successrowupdate2) //RowsAffected: 1

}

// {10010 1963-06-01 Duangkaew Piveteau 2 1989-08-24}
// ================================
// {10019 1953-01-23 Lillian Haddadi 1 1999-04-30}
// Transfer Before [string] [1953-01-23]; Transfer After [time.Time] [1953-01-23 00:00:00 +0000 UTC]
// {10020 1952-12-24 Mayuko Warwick 1 1991-01-26}
// Transfer Before [string] [1952-12-24]; Transfer After [time.Time] [1952-12-24 00:00:00 +0000 UTC]
// success row update: 0
// LastInsertID: 0
// RowsAffected: 1
// RowsAffected: 1

/*
【总结】
	驱动安装和导入，例如import _ "github.com/go-sq1-driver/mysq1
	连接数据库并返回数据库操作句柄，例如 sq1.open("root:sheca@tcp(192.168.40.103:3306)/test")
	使用db提供的接口函数
	使用db.Prepare预编译并使用参数化查询
		对预编译的SQL语句进行缓存，省去了每次解析优化该SQL语句的过程
		防止注入攻击
		使用返回的sql.Stmt操作数据库
*/


【16-6-SQLBuilder】
=====================================================================
"github.com/huandu/go-sqlbuilder"

EG:
	//sqlbuilder 帮助写SQL
	query := sqlbuilder.
		Select("emp_no", "first_name", "last_name", "gender", "birth_date").
		From("employees").
		Where("emp_no > 10015").
		Offset(2).Limit(2).
		OrderBy("emp_no").Desc().
		String()
	fmt.Printf("SQLBUILDER : %s\n", query)
	//SQLBUILDER : SELECT emp_no, first_name, last_name, gender, birth_date FROM employees WHERE emp_no > 10015 ORDER BY emp_no DESC LIMIT 2 OFFSET 2

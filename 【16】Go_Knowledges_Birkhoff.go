Viedo Date: 2023/11/18
Made By:	BIRKHOFF
Date:	2024-07-02


��16-1-SQL֮��ɾ�ġ�
=====================================================================
��16-2-SQL֮��ҳ��������
=====================================================================
��16-3-SQL֮���򡢷��顢�ۺϡ�
=====================================================================
��16-4-SQL֮�Ӳ�ѯ��Join��
=====================================================================
��16-5-���ݿ��׼�⿪����
=====================================================================
ALTER USER 'root'@'localhost' IDENTIFIED BY 'sheca';

��װmysql��Go����
go get -u github.com/go-sql-driver/mysql
go: downloading github.com/go-sql-driver/mysql v1.8.1
go: downloading filippo.io/edwards25519 v1.1.0
go: added filippo.io/edwards25519 v1.1.0
go: added github.com/go-sql-driver/mysql v1.8.1

����

����
package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //1 ������װ�͵���
)

/*db������*sq1.DB����һ���������ݿ�ľ�����ײ���һ����Э�̰�ȫ�����ӳء�*/
var db *sql.DB

func init() {
	//2 �������ݿ������
	var err error
	db, err = sql.Open("mysql", "root:sheca@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Second * 30) //��ʱ�¼�
	db.SetMaxOpenConns(0)                   //���������������Ĭ��Ϊ0��ʾ������
	db.SetMaxIdleConns(10)                  //���ÿ���������
}

[����]- ��ɾ�Ĳ� 
EG��
package main

import (
	"database/sql" //��׼���ṩ��ͳһ��̽ӿ�
	"fmt"
	"log"
	"time"

	//[1].������װ�͵���
	_ "github.com/go-sql-driver/mysql"
	//������������TCP���Ӻ�mysqlЭ�顣ִֻ�иð���init��������Դ�����ṹ�塢�ӿڡ�ȫ�ֱ����������ò���
	//init������һ������map����map��ע��һ�����ֺ����������ӳ�䣬���־ͽ�mysql��map�Ǳ�׼��sql��driversӳ��
)

/* db������*sq1.DB����һ���������ݿ�ľ�����ײ���һ����Э�̰�ȫ�����ӳء�*/
var db *sql.DB

func init() {
	//[2].�������ݿ������
	var err error
	dsn := "root:sheca@tcp(192.168.40.103:3306)/test"
	db, err = sql.Open("mysql", dsn) //�����dbʹ�õ���ȫ�ֱ��� ������:=
	// db, err = sql.Open("mysql", "root:sheca@(192.168.40.103:3306)/test") //��ʡ��tcp

	if err != nil {
		log.Fatal(err)
	}
	//
	db.SetConnMaxLifetime(time.Second * 30) //��ʱ�¼�
	db.SetMaxOpenConns(0)                   //���������������Ĭ��Ϊ0��ʾ������
	db.SetMaxIdleConns(10)                  //���ÿ���������
}

// [3].����ṹ��
type Emp struct { //���ֶζ�Ӧ�ı�����ṹ�嶨�壬��ú����ݿ��ֶ�˳���Ӧ
	emp_no     int
	birth_date string
	first_name string
	last_name  string
	gender     int16
	hire_date  string
}

func main() {
	//���в�ѯ
	emp := Emp{}                                                          //����ʵ��
	row := db.QueryRow("select * from employees where emp_no = ?", 10010) //��ѯ�� ֻ�ܲ�һ�� Limit-1
	if row.Err() != nil {
		log.Fatal(row.Err().Error())
	}
	err := row.Scan(&emp.emp_no, &emp.birth_date, &emp.first_name,
		&emp.last_name, &emp.gender, &emp.hire_date) //�ֶ�˳��Ҫ�ֶ���Ӧһ��
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emp) //
	fmt.Println("================================")

	// {10010 1963-06-01 Duangkaew Piveteau 2 1989-08-24}
	// ================================

	//������ѯ��Ԥ����
	//���ַ�����ƴ��ʹ�� 1.Prepare 2.Query
	//SQLע�빥�� select * from employees where emp_no = 10010 or 1=1
	stmt, err := db.Prepare("select * from employees where emp_no > ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//Queryʹ�ø����Ĳ���ִ��׼���õĲ�ѯ��䣬����*Rows����ʽ���ز�ѯ�����
	//��ѯʹ�������ġ��ڲ�������Ҫָ�������ģ���ʹ��QueryContext��
	rows, err := stmt.Query(10018) //����ֱ�Ӹ�����
	// query := "select * from employees"
	// rows, err := db.Query(query,1,2) //������Դ���sql�Ͳ���

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() { //������ÿһ��rows�ڲ�ָ��ǰ�� �൱���α�
		//row.scan�е��ֶα����ǰ������ݿ�����ֶε�˳�򣬷��򱨴�
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

	//��������
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

	//����
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

	//ɾ��
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
���ܽ᡿
	������װ�͵��룬����import _ "github.com/go-sq1-driver/mysq1
	�������ݿⲢ�������ݿ������������� sq1.open("root:sheca@tcp(192.168.40.103:3306)/test")
	ʹ��db�ṩ�Ľӿں���
	ʹ��db.PrepareԤ���벢ʹ�ò�������ѯ
		��Ԥ�����SQL�����л��棬ʡȥ��ÿ�ν����Ż���SQL���Ĺ���
		��ֹע�빥��
		ʹ�÷��ص�sql.Stmt�������ݿ�
*/


��16-6-SQLBuilder��
=====================================================================
"github.com/huandu/go-sqlbuilder"

EG:
	//sqlbuilder ����дSQL
	query := sqlbuilder.
		Select("emp_no", "first_name", "last_name", "gender", "birth_date").
		From("employees").
		Where("emp_no > 10015").
		Offset(2).Limit(2).
		OrderBy("emp_no").Desc().
		String()
	fmt.Printf("SQLBUILDER : %s\n", query)
	//SQLBUILDER : SELECT emp_no, first_name, last_name, gender, birth_date FROM employees WHERE emp_no > 10015 ORDER BY emp_no DESC LIMIT 2 OFFSET 2

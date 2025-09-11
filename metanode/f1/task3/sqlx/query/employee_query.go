package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Employee 结构体对应 employees 表
type Employee struct {
	ID         int          `db:"id"`
	Name       string       `db:"name"`
	Department string       `db:"department"`
	Salary     int          `db:"salary"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
}

// 查询工资最高员工
func queryHighestPaidEmployee(db *sqlx.DB) {
	var topEmployee Employee
	err := db.Get(&topEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n工资最高员工: ID: %d, 姓名: %s, 部门: %s, 薪资: %d\n",
		topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary)
}

// Book 结构体对应 books 表
type Book struct {
	ID        int          `db:"id"`
	Title     string       `db:"title"`
	Author    string       `db:"author"`
	Price     float64      `db:"price"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func queryExpensiveBooks(db *sqlx.DB) {
	// 查询价格大于50元的书籍
	var expensiveBooks []Book
	err := db.Select(&expensiveBooks, "SELECT * FROM books WHERE price > ?", 50.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n价格大于50元的书籍共 %d 本:\n", len(expensiveBooks))
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("mysql", "root:guo0909@tcp(localhost:3306)/knowledge?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询技术部员工
	var techEmployees []Employee
	err = db.Select(&techEmployees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("技术部员工共 %d 人:\n", len(techEmployees))
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 薪资: %d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询 employees 表中工资最高的员工信息
	queryHighestPaidEmployee(db)

	// 查询价格大于 50 元的书籍
	queryExpensiveBooks(db)
}

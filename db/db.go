package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	_ "github.com/go-sql-driver/mysql"
// )

// DB 结构体 数据库
type Database struct {
	*sql.DB
}

// NewDatabase 实例化Datebase
func NewDatabase(db *sql.DB) *Database {
	return &Database{
		db,
	}
}

// CreateDB 创建 名为 name 的数据库
func (d *Database) CreateDB(name string) error {
	// SQL语句
	sqlStr := "create database if not exists %s"

	// 执行SQL语句
	_, err := d.Exec(fmt.Sprintf(sqlStr, name))
	if err != nil {
		return fmt.Errorf("create database failed, err:%w", err)
	}
	fmt.Println("create database success")
	return nil
}

// CreateTable 创建数据表
func (d *Database) CreateTableSongs() error {
	// SQL语句
	sqlStr := `CREATE TABLE IF NOT EXISTS songs (
		id INT UNSIGNED AUTO_INCREMENT,
		name VARCHAR(255) DEFAULT NULL,
		singer VARCHAR(255) DEFAULT NULL,
		duration VARCHAR(255) DEFAULT NULL,
		link VARCHAR(255) DEFAULT NULL,
		cover VARCHAR(255) DEFAULT NULL,
		tag VARCHAR(64) DEFAULT NULL,
		platform INT UNSIGNED DEFAULT NULL,
		PRIMARY KEY ( id )
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	// 执行SQL语句
	_, err := d.Exec(sqlStr)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("create table success")
	return nil
}

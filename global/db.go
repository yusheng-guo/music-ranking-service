package global

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yushengguo557/music-ranking/db"
)

// var DB *sql.DB
var DB *db.Database

// 初始化数据库
// - 实例化Database
// - 数据库连接
// - 创建数据表
func InitDB() error {
	// 实例化 Database
	mysqldb, err := sql.Open("mysql", "root:123456@tcp(119.91.204.226:3306)/")
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	DB = db.NewDatabase(mysqldb)

	// 测试连接是否成功
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("Datebase init successfully!")
	// 创建 colly 数据库
	if err = DB.CreateDB("colly"); err != nil {
		return fmt.Errorf("create db, err: %w", err)
	}
	// 使用 colly 数据库
	_, err = DB.Exec("USE colly")
	if err != nil {
		panic(err.Error())
	}
	// 创建 songs 数据表
	if err = DB.CreateTableSongs(); err != nil {
		return fmt.Errorf("create table songs, err: %w", err)
	}
	return nil
}

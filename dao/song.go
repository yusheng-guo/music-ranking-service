package dao

import (
	"database/sql"
	"fmt"

	"github.com/yushengguo557/music-ranking/model"
)

// Insert 插入数据
func (d *Dao) InsertSong(s *model.Song) error {
	// SQL语句
	sqlStr := "insert into songs (name, singer, duration, link, cover) values (?, ?, ?, ?, ?)"
	// 执行SQL语句
	ret, err := d.DB.Exec(sqlStr, s.Name, s.Singer, s.Duration, s.Link, s.Cover)
	if err != nil {
		return fmt.Errorf("insert failed, err: %w", err)
	}
	// 获取插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		return fmt.Errorf("get id failed, err:%w", err)
	}
	fmt.Println("insert success, id:", id)
	return nil
}

// QueryOne 查询单条数据
func (d *Dao) QueryOneSong(db *sql.DB) error {
	//SQL语句
	sqlStr := "select name, singer, duration, link, cover from songs where id=?"
	var s model.Song
	err := d.DB.QueryRow(sqlStr, 1).Scan(&s.Name, &s.Singer, &s.Duration, &s.Link, &s.Cover)
	if err != nil {
		return fmt.Errorf("query failed, err:%w", err)
	}
	fmt.Println("query success:", s)
	return nil
}

// 查询多条数据
func (d *Dao) QueryMultiSongs() error {
	// SQL语句
	sqlStr := "select name, singer, duration, link, cover from songs where id > ?"
	rows, err := d.DB.Query(sqlStr, 0)
	if err != nil {
		return fmt.Errorf("query failed, err:%w", err)
	}

	//关闭Rows对象，释放连接资源
	defer rows.Close()

	var songs []model.Song
	//循环遍历Rows对象，获取每一行的数据
	for rows.Next() {
		var s model.Song
		err := rows.Scan(&s.Name, &s.Singer, &s.Duration, &s.Link, &s.Cover)
		if err != nil {
			return fmt.Errorf("scan failed,err:%w", err)
		}
		songs = append(songs, s)
	}
	fmt.Println("query success:", songs)
	return nil
}
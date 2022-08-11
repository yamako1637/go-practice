package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dbName := "./test.db"
	os.Remove(dbName)

	// 接続
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブル作成
	sqlStmt := `
		CREATE TABLE test (
			id	INTEGER NOT NULL,
			name	TEXT,
			PRIMARY KEY(id)
		);
		DELETE FROM test;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO test(id,name)
		VALUES(?,?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("HELLO %03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	// トランザクション終了
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	timeStart := time.Now()
	// クエリ
	rows, err := db.Query(`
		SELECT id,name FROM test
	`)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(id, name)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("処理時間: %vms\n", time.Since(timeStart).Milliseconds())
	// os.Remove(dbName)

}

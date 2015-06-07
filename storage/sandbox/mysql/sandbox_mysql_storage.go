package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"github.com/screwyprof/gosandbox/storage/sandbox"
	//"bitbucket.org/lazadaweb/squirrel"
	//	"fmt"
)

type mysqlStorage struct {
	db *sql.DB
}

func NewInstance(db *sql.DB) storage.ISandboxStorage {
	return &mysqlStorage{db: db}
}

func (s mysqlStorage) LoadUserNameById(id int64) string {
	//	db, err := sql.Open("mysql", "root:gjgjrfntgtnkm@tcp(127.0.0.1:3306)/tender")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer db.Close()

	var name string

	err := s.db.QueryRow("select username from core__user where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	return name

	/*
		rows, err := db.Query("select username from core__user where id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/
}

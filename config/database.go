package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Setting struct {
	Username string
	Password string
	DB       string
}

func (c Setting) DatabaseConnect() *sql.DB {

	db, err := sql.Open("mysql", c.Username+":"+c.Password+"@tcp(127.0.0.1:3306)/"+c.DB)

	if err != nil {
		log.Print(err.Error())
	}

	return db

}

type Tag struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
}

func test() {

	db := Setting{Username: "sina", Password: "sina", DB: "test"}.DatabaseConnect()

	defer db.Close()

	// results, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// for results.Next() {
	// 	var tag Tag

	// 	err = results.Scan(&tag.ID, &tag.Name, &tag.Family, &tag.Email)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Printf("Name (%s), Fmaily (%s), Email (%s) and ID:%d\n", tag.Name, tag.Family, tag.Email, tag.ID)
	// }

}

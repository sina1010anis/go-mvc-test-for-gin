package model

import (
	"database/sql"
	"web/config"
)

var Userdb = config.Setting{Username: "sina", Password: "sina", DB: "test"}.DatabaseConnect()

type UserFilable struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
}

func UserGet() any {

	// defer db.Close()
	var res []map[string]string
	results, err := Userdb.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag UserFilable

		err = results.Scan(&tag.ID, &tag.Name, &tag.Family, &tag.Email)
		if err != nil {
			panic(err.Error())
		}

		res = append(res, map[string]string{
			"id":     tag.ID,
			"name":   tag.Name,
			"family": tag.Family,
			"email":  tag.Email,
		})

	}

	return res

}

func UserStore(name, family, email string) (sql.Result, error) {

	res, err := Userdb.Exec("INSERT INTO users (name, family, email) VALUES (?, ?, ?) ;", name, family, email)

	return res, err

}

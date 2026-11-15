package model

import (
	"web/config"
)

var db = config.Setting{Username: "sina", Password: "sina", DB: "test"}.DatabaseConnect()

type Filable struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
}

func UserGet() any {

	// defer db.Close()
	var res []map[string]string
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Filable

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

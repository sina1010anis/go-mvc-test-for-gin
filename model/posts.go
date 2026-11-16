package model

import (
	"web/config"
)

var Postdb = config.Setting{Username: "sina", Password: "sina", DB: "test"}.DatabaseConnect()

type PostFilable struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
	Price  string `json:"price"`
}

func PostGet() any {

	// defer db.Close()
	var res []map[string]string
	results, err := Postdb.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag PostFilable

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

package main

import "database/sql"

type User struct {
	Id           string         `db:"id"`
	FirstName    string         `db:"first_name"`
	Surname      string         `db:"surname"`
	Email        sql.NullString `db:"email"`
	DepartmentId string         `db:"department_id"`
	Position     string         `db:"position"`
}

type users struct {
	Map map[string]User
}

var Users *users = &users{make(map[string]User)}

func (t *users) Load() {
	//res, err := DB.Query("SELECT id, first_name person FROM person")
	//
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//for res.Next() {
	//	var id string
	//	var name string
	//	res.Scan(&id, &name)
	//	log.Println(id, name)
	//}

	users := []User{}

	err := DB.Select(&users, `
		SELECT
			id,
			first_name,
			surname,
			email,
			department_id,
			position
		FROM person
		WHERE is_active = 'y'
	`)

	if err != nil {
		panic(err)
	}

	for _, user := range users {
		t.Map[user.Id] = user
	}

	ml.Printf("Users [%d] loaded\n", len(t.Map))
}

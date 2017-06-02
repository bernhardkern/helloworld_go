package data

import (
	"database/sql"
	"fmt"
)

type mySql struct {
	db *sql.DB
}

func (m *mySql) Create(person Person) error {
	fmt.Println("we create a person", person.FirstName)

	stmt, err := m.db.Prepare("INSERT INTO person (first_name, last_name) VAlUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(person.FirstName, person.LastName)

	return err
}

func (m *mySql) ReadAll() ([]Person, error) {

	result := make([]Person, 0)
	rows, err := m.db.Query("SELECT * from person")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		person := Person{}
		err = rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}
		result = append(result, person)
	}
	return result, err
}

func NewMySQL(user string, password string, schema string) (PersonRepository, error) {
	db, err := sql.Open("mysql", user+":"+password+"@/"+schema)

	if err != nil {
		return nil, err
	}

	return &mySql{db: db}, nil
}

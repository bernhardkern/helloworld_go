package data

import (
	"database/sql"
	"log"
)

type mySql struct {
	db *sql.DB
}

func NewMySQL(user string, password string, schema string) (PersonRepository, error) {
	db, err := sql.Open("mysql", user+":"+password+"@/"+schema)
	if err != nil {
		return nil, err
	}

	mySQL := &mySql{db: db}
	mySQL.initTables()
	return mySQL, nil
}

func (m *mySql) Create(person Person) error {
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

// hacky schema creation for demo purposes
func (m *mySql) initTables() {
	_, err := m.db.Exec("CREATE TABLE person(id BIGINT(20) NOT NULL AUTO_INCREMENT,first_name VARCHAR(50) NULL DEFAULT NULL, last_name VARCHAR(50) NULL DEFAULT NULL, PRIMARY KEY (id), INDEX id (id))")
	if err != nil {
		return
	}
	log.Println("Created table(s)")
}

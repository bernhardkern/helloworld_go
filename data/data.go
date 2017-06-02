package data

type PersonRepository interface {
	Create(Person) error
}

type Person struct {
	Id        int64
	FirstName string
	LastName  string
}

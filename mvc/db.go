package mvc

import "fmt"

type DBI interface {
	Create(value interface{})
	First(out interface{}, where ...interface{})
	Save(value interface{})
	Delete(value interface{})
	Error() error
}

type DB struct {}

func (db *DB) Create(value interface{}) *DB {
	fmt.Println("\tDB Create called")

	return db
}

func (db *DB) First(out interface{}, where ...interface{}) *DB {
	fmt.Println("\tDB First called")

	return db
}

func (db *DB) Save(value interface{}) *DB {
	fmt.Println("\tDB Save called")

	return db
}

func (db *DB) Delete(value interface{}) *DB {
	fmt.Println("\tDB Delete called")

	return db
}

func (db *DB) Error(err error) *DB {
	return db
}

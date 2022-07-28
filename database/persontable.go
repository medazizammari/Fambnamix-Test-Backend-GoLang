package database

import (
	. "crudapp/models"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
)

func GetAllPerson() []Person {
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/crud", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error())
	}

	var personList []Person
	for results.Next() {
		var person Person

		err := results.Scan(&person.ID, &person.FirstName, &person.LastName, &person.DateOfBirth, &person.Address, &person.PhotoId)
		if err != nil {
			panic(err.Error())
		}
		dob, _ := time.Parse("2006-01-02", person.DateOfBirth)
		fmt.Println(dob)
		age := int((time.Since(dob).Hours()) / 8760)
		person.Age = age
		personList = append(personList, person)
	}
	return personList
}

func GetPerson(id string) []Person {
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/crud", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM person WHERE id=%s", id)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var personList []Person
	for results.Next() {
		var person Person

		err := results.Scan(&person.ID, &person.FirstName, &person.LastName, &person.DateOfBirth, &person.Address, &person.PhotoId)
		if err != nil {
			panic(err.Error())
		}
		dob, _ := time.Parse("2006-01-02", person.DateOfBirth)
		age := int((time.Since(dob).Hours()) / 8760)
		person.Age = age
		personList = append(personList, person)
	}
	return personList
}

func AddPerson(person Person) bool {
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/crud", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf(
		`INSERT INTO person(firstname, lastname, date_of_birth, address, photo_id)
		VALUES("%s", "%s", "%s", "%s", "%s")`,
		person.FirstName, person.LastName, person.DateOfBirth, person.Address, person.PhotoId)
	fmt.Println(query)
	results, err := db.Query(query)

	fmt.Println(results)
	return err == nil
}

func DeletePerson(id string) bool {
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/crud", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf(`DELETE FROM person WHERE id=%s`, id)
	fmt.Println(query)
	results, err := db.Query(query)
	fmt.Println(err)
	fmt.Println(results)
	return err == nil
}

func UpdatePerson(id string, person Person) bool {
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/crud", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf(
		`UPDATE person SET firstname="%s", lastname="%s", date_of_birth="%s", address="%s" WHERE id="%s"`,
		person.FirstName, person.LastName, person.DateOfBirth, person.Address, id)
	fmt.Println(query)
	results, err := db.Query(query)

	fmt.Println(results)
	return err == nil
}

func UpdateProfilePicture(id string, photo_url string) bool {
	connString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/crud", os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := fmt.Sprintf(
		`UPDATE person SET photo_id="%s" WHERE id="%s"`,
		photo_url, id)
	fmt.Println(query)
	results, err := db.Query(query)

	fmt.Println(results)
	return err == nil
}

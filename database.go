package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)


func getConnection (path string) *sql.DB{
	database, err := sql.Open("sqlite3", path )
	if err != nil {
		log.Fatal("Database Connection Error !!")
		os.Exit(1)
	}
	return database
}

func createTableIfNotCreated(db *sql.DB)  {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS contact (id INTEGER PRIMARY KEY, name TEXT NOT NULL, contact TEXT UNIQUE, email TEXT NOT NULL UNIQUE) ;")
	if err != nil {
		log.Fatal("Database Write Error !!")
	}
	statement.Exec()
}

func insertContact(db *sql.DB ,contact Contact) error {
	statement, err := db.Prepare("INSERT INTO contact (name, contact, email) VALUES (? ,? ,?) ;")
	if err != nil {
		fmt.Println("Error from Insertion Statement Creation")
		fmt.Println(err.Error())
		return err
	}
	result , err := statement.Exec(contact.Name, contact.Contact, contact.Email)
	if err != nil {
		fmt.Println("Error In Statement Execution")
	}
	fmt.Println(result)
	return err
}

func getContactsFromDb(db *sql.DB) ([]Contact, error) {
	rows, err := db.Query("SELECT id, name, contact, email FROM contact ;")
	var contacts []Contact
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var name string
		var contact string 
		var email string 
		rows.Scan(&id, &name, &contact, &email)
		contacts = append(contacts, Contact{id, name, contact, email})
	}
	return contacts, nil
}

func getSingleContactFromDb(db *sql.DB, id int) ([]Contact, error) {
	statement, err := db.Prepare("SELECT id, name, contact, email FROM contact WHERE id = ? ;")
	if err != nil {
		fmt.Println("Error In Getting Single Record")
		fmt.Println(err.Error())
		return []Contact{}, err
	}
	var contacts []Contact
	rows, err := statement.Query(id)
	if err != nil {
		fmt.Println(err.Error())
		return []Contact{}, err
	}
	for rows.Next() {
		var id int
		var name string
		var contact string 
		var email string 
		rows.Scan(&id, &name, &contact, &email)
		contacts = append(contacts, Contact{id, name, contact, email})
	}
	return contacts, nil
}

func updateContactInDB(db *sql.DB, id int, contact Contact) error {
	statement, err := db.Prepare("UPDATE contact SET name = ? , contact = ? , email = ? WHERE id = ?;")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = statement.Exec(contact.Name, contact.Contact, contact.Email, id)
	if err != nil {
		return nil
	}
	return err
}

func deleteSingleContact(db *sql.DB, id int) error{
	statement, err := db.Prepare("DELETE FROM contact WHERE id = ?;")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return nil
	}
	return err
}
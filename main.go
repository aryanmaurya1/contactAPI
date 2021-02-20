package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Contact : Data Model
type Contact struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Contact string `json:"contact"`
	Email string `json:"email"`
}

// Error : Error object
type Error struct {
	Err string `json:"err"`
}

// Success : sucess object
type Success struct {
	Msg string `json:"msg"`
}

var dbPath string = "./contacts.db"

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := getConnection(dbPath)
	defer db.Close()
	contacts, err := getContactsFromDb(db)
	if err != nil {
		var e Error
		e.Err = err.Error()
		json.NewEncoder(w).Encode(contacts)
		return
	}
	json.NewEncoder(w).Encode(contacts)
	fmt.Println("GET ALL Sucess")
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var e Error
	db := getConnection(dbPath)
	defer db.Close()
	
	var id string = mux.Vars(r)["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}
	contacts, err := getSingleContactFromDb(db, intID)
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}

	if len(contacts) == 0 || contacts == nil {
		e.Err = "No result"
		json.NewEncoder(w).Encode(e)
		return
	}
	json.NewEncoder(w).Encode(contacts) 
	fmt.Println("GET Single Sucess")
}

func addContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)
	fmt.Println(contact)
	contact.ID = int(time.Now().UnixNano())
	db := getConnection(dbPath)
	defer db.Close()

	err :=	insertContact(db, contact)
	if err != nil {
		var e Error
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return 
	}

	var s Success
	s.Msg = "Sucess"
	json.NewEncoder(w).Encode(s)
	fmt.Println("ADD Sucess")
}


func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := getConnection(dbPath)
	var e Error
	defer db.Close()
	var contact Contact
	var ID, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}

	json.NewDecoder(r.Body).Decode(&contact)
	err = updateContactInDB(db, ID, contact)
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}
	contacts, err := getSingleContactFromDb(db, ID)
	json.NewEncoder(w).Encode(contacts)
	fmt.Println("UPDATE ONE Sucess")
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := getConnection(dbPath)
	var e Error
	defer db.Close()
	var contact Contact
	var ID, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}
	contacts, err := getSingleContactFromDb(db, ID)
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}
	contact = contacts[0]
	err = deleteSingleContact(db, ID)
	if err != nil {
		fmt.Println(err.Error())
		e.Err = err.Error()
		json.NewEncoder(w).Encode(e)
		return
	}
	json.NewEncoder(w).Encode(contact)
}




func main() {
	db := getConnection(dbPath)
	createTableIfNotCreated(db)
	db.Close()


	r := mux.NewRouter()
	r.HandleFunc("/api/contacts", getContacts).Methods("GET")
	r.HandleFunc("/api/contacts/{id}", getContact).Methods("GET")	
	r.HandleFunc("/api/contacts", addContact).Methods("POST")	
	r.HandleFunc("/api/contacts/{id}", updateContact).Methods("PUT")	
	r.HandleFunc("/api/contacts/{id}", deleteContact).Methods("DELETE")	


	http.ListenAndServe(":8000", r)
}
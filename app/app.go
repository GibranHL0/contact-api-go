// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/asaskevich/govalidator"
// 	_ "github.com/lib/pq"
// )

// type Contact struct {
// 	Email    string `json:"email" valid:"required,email"`
// 	Name     string `json:"name" valid:"required,alpha"`
// 	LastName string `json:"last_name" valid:"required,alpha"`
// 	Company  string `json:"company" valid:"optional"`
// 	Phone    string `json:"phone" valid:"numeric,optional"`
// }

// func validateContact(c Contact) (Message, bool) {
// 	var msg Message

// 	if !govalidator.IsEmail(c.Email) && govalidator.IsNotNull(c.Email) {
// 		msg.Msg = "Email is not valid"

// 		return msg, false
// 	}

// 	if !govalidator.IsAlpha(c.Name) && govalidator.IsNotNull(c.Name) {
// 		msg.Msg = "Name is not valid"

// 		return msg, false
// 	}

// 	if !govalidator.IsAlpha(c.LastName) && govalidator.IsNotNull(c.LastName) {
// 		msg.Msg = "Last Name is not valid"

// 		return msg, false
// 	}

// 	valid, _ := govalidator.ValidateStruct(c)

// 	if !valid {
// 		msg.Msg = "Information is not well formatted"

// 		return msg, false
// 	}

// 	return msg, true
// }

// type Message struct {
// 	Msg string `json:"msg"`
// }

// func getDB() *sql.DB {
// 	url := os.Getenv("DATABASE")

// 	db, err := sql.Open("postgres", url)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }

// func contactsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	switch r.Method {
// 	case "GET":
// 		contacts, err := getContacts()

// 		if err != nil {
// 			log.Print(err)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(contacts)

// 	case "POST":
// 		body, err := ioutil.ReadAll(r.Body)

// 		if err != nil {
// 			log.Print(err)
// 			return
// 		}

// 		msg, status := postContact(body)

// 		w.WriteHeader(status)
// 		json.NewEncoder(w).Encode(msg)

// 	case "DELETE":
// 		email := r.URL.Query().Get("email")
// 		msg, status := deleteContact(string(email))

// 		w.WriteHeader(status)
// 		json.NewEncoder(w).Encode(msg)

// 	case "PUT":
// 		body, err := ioutil.ReadAll(r.Body)

// 		if err != nil {
// 			log.Print(err)
// 			return
// 		}

// 		msg, status := putContact(body)

// 		w.WriteHeader(status)
// 		json.NewEncoder(w).Encode(msg)

// 	default:
// 		http.Error(w, "Action not supported", http.StatusBadRequest)
// 	}
// }

// func getContacts() ([]Contact, error) {
// 	contacts := make([]Contact, 0)

// 	db := getDB()
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM contacts")

// 	if err != nil {
// 		return contacts, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var email string
// 		var name string
// 		var last_name string
// 		var company string
// 		var phone string

// 		err = rows.Scan(&email, &name, &last_name, &company, &phone)

// 		if err != nil {
// 			panic(err)
// 		}

// 		contacts = append(contacts, Contact{
// 			Email:    email,
// 			Name:     name,
// 			LastName: last_name,
// 			Company:  company,
// 			Phone:    phone,
// 		})
// 	}

// 	return contacts, nil
// }

// func postContact(body []byte) (Message, int) {
// 	var contact Contact
// 	var msg Message

// 	json.Unmarshal(body, &contact)

// 	msg, valid := validateContact(contact)

// 	if !valid {
// 		return msg, http.StatusPartialContent
// 	}

// 	db := getDB()

// 	found, err := findEmail(contact.Email, db)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	if found {
// 		msg.Msg = "Email already exists"

// 		return msg, http.StatusPartialContent
// 	}

// 	sqlStatement := `
// 	INSERT INTO contacts (email, name, last_name, company, phone)
// 	VALUES ($1, $2, $3, $4, $5)`

// 	_, err = db.Exec(sqlStatement, contact.Email, contact.Name, contact.LastName,
// 		contact.Company, contact.Phone)

// 	if err != nil {
// 		log.Print(err)
// 		msg.Msg = "Something went wrong"

// 		return msg, http.StatusInternalServerError
// 	} else {
// 		msg.Msg = "Created"
// 	}

// 	return msg, http.StatusCreated
// }

// func deleteContact(email string) (Message, int) {
// 	var msg Message

// 	if !govalidator.IsEmail(email) && govalidator.IsNotNull(email) {
// 		msg.Msg = "Invalid Email"

// 		return msg, http.StatusBadRequest
// 	}

// 	sqlStatement := `DELETE FROM contacts WHERE email = $1;`
// 	db := getDB()

// 	found, err := findEmail(email, db)

// 	if err != nil {
// 		msg.Msg = "Something went wrong"

// 		return msg, http.StatusInternalServerError
// 	}

// 	if !found {
// 		msg.Msg = "Email does not exist"

// 		return msg, http.StatusPartialContent
// 	}

// 	_, err = db.Exec(sqlStatement, email)

// 	if err != nil {
// 		log.Print(err)
// 		return msg, http.StatusInternalServerError
// 	} else {
// 		msg.Msg = fmt.Sprintf("%s deleted", email)
// 	}

// 	return msg, http.StatusOK
// }

// func putContact(body []byte) (Message, int) {
// 	var contact Contact
// 	var msg Message

// 	json.Unmarshal(body, &contact)

// 	msg, valid := validateContact(contact)

// 	if !valid {
// 		return msg, http.StatusPartialContent
// 	}

// 	db := getDB()

// 	sqlStatement := `
// 	UPDATE contacts
// 	SET name = $2, last_name = $3, company = $4, phone = $5
// 	WHERE email = $1
// 	`

// 	_, err := db.Exec(sqlStatement, contact.Email, contact.Name, contact.LastName,
// 		contact.Company, contact.Phone)

// 	if err != nil {
// 		log.Print(err)
// 		msg.Msg = "Something went wrong"

// 		return msg, http.StatusInternalServerError
// 	} else {
// 		msg.Msg = fmt.Sprintf("%s modified", contact.Email)
// 	}

// 	return msg, http.StatusOK
// }

// func findEmail(email string, db *sql.DB) (bool, error) {
// 	sqlStatement := `SELECT email from contacts WHERE email = $1`

// 	err := db.QueryRow(sqlStatement, email).Scan(&email)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return false, nil
// 		}
// 		return true, err
// 	}

// 	return true, nil
// }

// func main() {
// 	mux := http.NewServeMux()
// 	port := os.Getenv("PORT")

// 	mux.HandleFunc("/contacts", contactsHandler)

// 	log.Fatal(http.ListenAndServe(":"+port, mux))
// }

package app

import (
	"log"
	"net/http"

	"github.com/GibranHL0/contact-api-go/app/database"
	"github.com/GibranHL0/contact-api-go/app/router"
)

type App struct {
	Server *http.ServeMux
	Db     database.Database
}

func New() *App {
	app := &App{
		Server: http.NewServeMux(),
		Db:     &database.Postgres{},
	}

	check(app.Db.Open())

	app.initRoutes()

	return app
}

func (a *App) initRoutes() {
	a.Server.HandleFunc("/contacts", router.ContactsRouter(a.Db))
}

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

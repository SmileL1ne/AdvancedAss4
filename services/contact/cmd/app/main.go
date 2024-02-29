package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain/contact"
	pgRepo "architecture_go/services/contact/internal/repository/storage/postgres"
	ucContact "architecture_go/services/contact/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	dbInfo := &postgres.Database{
		Host:     "localhost",
		Port:     "7777",
		User:     "postgres",
		Password: "postgres",
		Name:     "ap_ass3",
	}

	db, err := postgres.OpenDB(dbInfo)
	if err != nil {
		panic("nah" + err.Error())
	}

	contactRepo := pgRepo.New(db)
	contactUC := ucContact.New(contactRepo)

	http.HandleFunc("/contact/create", func(w http.ResponseWriter, req *http.Request) {
		id, err := contactUC.Create(contact.Contact{})
		if err != nil {
			panic("cannot create contact")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("new contact with id %d created", id)))
	})

	err = http.ListenAndServe("127.0.0.1:7000", nil)
	log.Fatal(err)
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
	Success bool
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/Message.gtpl"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
		Success: true,
	}

	// do something with details
	_ = details
	log.Println("Email:" + details.Email)
	log.Println("Subject:" + details.Subject)
	log.Println("Message:" + details.Message)
	tmpl.Execute(w, struct{Data ContactDetails}{ details})
}
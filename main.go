package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/scorredoira/email"
)

var temp *template.Template = template.Must(template.ParseFiles("../Project/Login.html"))
var temp0 *template.Template = template.Must(template.ParseFiles("../Project/UserPortal.html"))
var temp1 *template.Template = template.Must(template.ParseFiles("../Project/Management.html"))

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/Login", Login).Methods("GET")
	r.HandleFunc("/UserPortal", UserPortal).Methods("GET")
	r.HandleFunc("/Management", Management).Methods("GET")
	http.ListenAndServe(":3030", r)

}

func Login(w http.ResponseWriter, r *http.Request) {
	temp.Execute(w, nil)
}
func UserPortal(w http.ResponseWriter, r *http.Request) {
	temp0.Execute(w, nil)
}
func Management(w http.ResponseWriter, r *http.Request) {
	temp1.Execute(w, nil)
	SendEmail()
}

//------------send mail--------------

func SendEmail() {
	m := email.NewMessage("Customer management ", "All details are available in this pdf")
	m.From = mail.Address{Name: "cratan588@gmail.com", Address: "Ayantika@12"}
	m.To = []string{"beginnerprogrammer58@gmail.com"}

	err := m.Attach("./InvoicesDetails.pdf")
	if err != nil {
		log.Println(err)
	} else {
		email.Send("smtp.gmail.com:587", smtp.PlainAuth("", "cratan588@gmail.com", "Ayantika@12", "smtp.gmail.com"), m)
		fmt.Println("PDF sended via email successfully!")
	}
}

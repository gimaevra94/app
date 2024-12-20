package main

import (
	"database/sql"
	"net/http"

	"github.com/gimaevra94/app/app/database"
	"github.com/gimaevra94/app/app/mailsendler"
)

var mail string

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	http.HandleFunc("/email_send_button", func(w http.ResponseWriter, r *http.Request) {
		mail = r.FormValue("email")
		mailsendler.MailSendler(mail)
		http.ServeFile(w, r, "web/codesendform.html")
	})

	http.HandleFunc("/code_send_button", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		if code != mailsendler.Authcode_str {
			http.ServeFile(w, r, "web/wrongcode.html")
		} else {
			db.SqlConn()
			db.EmailAdd(db.UsersEmail{Email: mail}, database)
			http.ServeFile(w, r, "web/home.html")
		}
	})

	http.HandleFunc("/back_to_input_button", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/codesendform.html")
	})

	http.HandleFunc("/code_not_arrived_button", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/codesendform.html")
		mailsendler.MailSendler(mail)
	})

	http.ListenAndServe(":8000", nil)

}
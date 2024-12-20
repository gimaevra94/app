package main

import (
	"fmt"
	"net/http"

	"github.com/gimaevra94/auth/email_auth/app/database"
	"github.com/gimaevra94/auth/email_auth/app/mailsendler"
)

func main() {
	var email string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/web/index.html")
	})

	http.HandleFunc("/email_send_button", func(w http.ResponseWriter, r *http.Request) {
		email = r.FormValue("email")
		mailsendler.MailSendler(email)
		http.ServeFile(w, r, "app/codesendform.html")
	})

	http.HandleFunc("/code_send_button", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		if code != mailsendler.Authcode_str {
			http.ServeFile(w, r, "app/web/wrongcode.html")
		} else {
			db, err := database.SqlConn()
			if err != nil {
				fmt.Printf("SqlConn: %v", err)
				return
			}

			database.EmailAdd(database.Users{Email: email}, db)
			http.ServeFile(w, r, "app/web/home.html")
		}
	})

	http.HandleFunc("/back_to_input_button", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/web/codesendform.html")
	})

	http.HandleFunc("/code_not_arrived_button", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/web/codesendform.html")
		mailsendler.MailSendler(email)
	})

	http.ListenAndServe(":8000", nil)
}

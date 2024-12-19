package main

import (
	"net/http"

	"github.com/gimaevra94/app/app/mailsendler"
)

var mail string

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/web/index.html")
	})

	http.HandleFunc("/user_input", func(w http.ResponseWriter, r *http.Request) {
		mail = r.FormValue("mail")
		mailsendler.MailSendler(mail)
		http.ServeFile(w, r, "app/web/codesendform.html")
	})

	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		if code != mailsendler.Authcode_str {
			http.ServeFile(w, r, "app/web/wrongcode.html")
		} else {
			http.ServeFile(w, r, "app/web/home.html")
		}
	})

	http.HandleFunc("/backtoinput", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/web/codesendform.html")
	})

	http.HandleFunc("/codenotarrived", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/web/codesendform.html")
		mailsendler.MailSendler(mail)
	})

	http.ListenAndServe(":8000", nil)

}
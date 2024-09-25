package auth

import (
	DB "SCTI/database"
	"SCTI/fileserver"
	HTMX "SCTI/htmx"
	"fmt"
	"net/http"
	"net/url"
	"os"

	gomail "gopkg.in/mail.v2"
)

func GetSenha(w http.ResponseWriter, r *http.Request) {
	var t = fileserver.Execute("template/senha.gohtml")
	t.Execute(w, nil)
}

func PostSenha(w http.ResponseWriter, r *http.Request) {
	from := os.Getenv("GMAIL_SENDER")
	pass := os.Getenv("GMAIL_PASS")

	code, err := DB.GetCodeByEmail(r.FormValue("Email"))
	if err != nil {
		HTMX.Failure(w, "Nenhum usuário encontrado com este email: ", err)
		return
	}

	encodedEmail := url.QueryEscape(r.FormValue("Email"))
	verificationLink := fmt.Sprintf("%s/trocar?email=%s&code=%s", os.Getenv("URL"), encodedEmail, code)

	htmlBody := `
  <!DOCTYPE html>
  <html>
    <head>
      <style>
        a {
          display: inline-block;
          padding: 10px 20px;
          font-size: 16px;
          cursor: pointer;
          text-align: center;
          text-decoration: none;
          outline: none;
          color: #ffffff !important;
          background-color: #FF00C7 !important;
          border: none;
          border-radius: 15px;
          box-shadow: 0 9px #999;
        }
        a:hover { filter: brightness(1.2) !important;}
        a:active {
          background-color: #FF00C7;
          box-shadow: 0 5px #666;
          transform: translateY(4px);
        }
      </style>
    </head>
    <body>
      <p>Clique no botão abaixo para trocar sua senha:</p>
      <a href="` + verificationLink + `">Trocar Senha</a>
    </body>
  </html>
  `

	plainBody := "Clique no botão abaixo para trocar sua senha:\n" + verificationLink

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", r.FormValue("Email"))
	msg.SetHeader("Subject", "Trocar Senha SCTI")
	msg.SetBody("text/plain", plainBody)
	msg.AddAlternative("text/html", htmlBody)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

	if err := dialer.DialAndSend(msg); err != nil {
		HTMX.Failure(w, "Falha ao enviar email: ", err)
		return
	}
	HTMX.Success(w, "Email enviado com sucesso")
}

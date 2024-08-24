package auth

import (
  "SCTI/fileserver"
  DB "SCTI/database"
  "fmt"
  "os"
  "net/url"
  "net/http"
  gomail "gopkg.in/mail.v2"
)

func GetSenha(w http.ResponseWriter, r *http.Request) {
  var t = fileserver.Execute("template/senha.gohtml")
  t.Execute(w, nil)
}

func PostSenha(w http.ResponseWriter, r *http.Request) {
  from := os.Getenv("GMAIL_SENDER")
  pass := os.Getenv("GMAIL_PASS")

  _, err := DB.GetCodeByEmail(r.FormValue("Email"))
  if err != nil {
    w.Header().Set("Content-Type", "text/html")
      w.Write([]byte(`
        <div class="red">
          Nenhum usuário encontrado com este email
        </div>
      `))
    return
  }

  encodedEmail := url.QueryEscape(r.FormValue("Email"))
  verificationLink := fmt.Sprintf("http://localhost:8080/trocar?email=%s", encodedEmail)

  htmlBody := `
    <!DOCTYPE html>
    <html>
    <head>
        <style>
            .button {
                display: inline-block;
                padding: 10px 20px;
                font-size: 16px;
                cursor: pointer;
                text-align: center;
                text-decoration: none;
                outline: none;
                color: #ffffff;
                background-color: #4CAF50;
                border: none;
                border-radius: 15px;
                box-shadow: 0 9px #999;
            }
            .button:hover {background-color: #3e8e41}
            .button:active {
                background-color: #3e8e41;
                box-shadow: 0 5px #666;
                transform: translateY(4px);
            }
        </style>
    </head>
    <body>
        <p>Clique no botão abaixo para trocar sua senha:</p>
        <a href="` + verificationLink + `" class="button">Trocar Senha</a>
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
    fmt.Printf("smtp error: %s\n", err)
    w.Header().Set("Content-Type", "text/html")
      w.Write([]byte(`
        <div class="red">
          Falha ao enviar email
        </div>
      `))
    return
  }

  w.Header().Set("Content-Type", "text/html")
  w.Write([]byte(`
    <div class="green">
      Email enviado com sucesso!!
    </div>
  `))
}

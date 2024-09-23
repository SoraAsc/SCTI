package dashboard

import (
	DB "SCTI/database"
	Erros "SCTI/erros"
	HTMX "SCTI/htmx"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	qrcode "github.com/skip2/go-qrcode"
	gomail "gopkg.in/mail.v2"
)

func UserSentQR(w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
		HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
		return
	}
	email := r.FormValue("Email")
	code, err := DB.GetCodeByEmail(email)
	if err != nil {
		HTMX.Failure(w, "Falha ao validar código: ", err)
		return
	}
	user := DB.User{
		Code:  code,
		Email: email,
	}
	sendQRToUser(user)
	HTMX.Success(w, "QR Code Enviado!")
}

func AllUsersSentQR(w http.ResponseWriter, r *http.Request) {
	if !CheckAdmin(w, r) {
		HTMX.Failure(w, "Endpoint exclusivo de admins", fmt.Errorf("Acesso proibido a usuários não admin"))
		return
	}
	users, err := DB.GetAllUsers()
	if err != nil {
		HTMX.Failure(w, "Falha ao obter lista de usuários: ", err)
		return
	}
	sentCount := 0
	failedCount := 0
	for _, user := range users {
		qrSent, err := DB.IsUserQR(user.Email)
		if err != nil {
			Erros.LogError("dashboard/qrcode.go", fmt.Errorf("Erro ao verificar status do QR para usuário %s: %v\n", user.Email, err))
			failedCount++
			continue
		}
		if !qrSent {
			err := sendQRToUser(user)
			if err != nil {
				Erros.LogError("dashboard/qrcode.go", fmt.Errorf("Falha ao enviar QR para %s: %v\n", user.Email, err))
				failedCount++
			} else {
				sentCount++
			}
		}
	}
	HTMX.Success(w, fmt.Sprintf("Processo concluído. QR codes enviados: %d, Falhas: %d", sentCount, failedCount))
}

func sendQRToUser(user DB.User) error {
	code, err := DB.GetCodeByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("falha ao obter código para %s: %v", user.Email, err)
	}
	encodedEmail := url.QueryEscape(user.Email)
	qrContent := fmt.Sprintf("%s/presenca?email=%v&code=%v", os.Getenv("URL"), encodedEmail, code)
	qr, err := qrcode.Encode(qrContent, qrcode.Medium, 256)
	if err != nil {
		return fmt.Errorf("falha ao gerar QR para %s: %v", user.Email, err)
	}
	qrBase64 := base64.StdEncoding.EncodeToString(qr)
	err = sendEmail(user.Email, qrBase64)
	if err != nil {
		return fmt.Errorf("falha ao enviar e-mail para %s: %v", user.Email, err)
	}
	err = DB.SetSentQR(user.Email)
	if err != nil {
		return fmt.Errorf("falha ao atualizar status de envio para %s: %v", user.Email, err)
	}
	return nil
}

func sendEmail(email, qrBase64 string) error {
	from := os.Getenv("GMAIL_SENDER")
	pass := os.Getenv("GMAIL_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verificação de email SCTI")

	// Corpo HTML com referência à imagem anexada
	htmlBody := `
  <!DOCTYPE html>
  <html>
  <body>
  <p>Este e-mail contém um QR code para verificação.</p>
  <img src="cid:qrcode.png" alt="QR code">
  </body>
  </html>
  `
	m.SetBody("text/html", htmlBody)

	// Decodifica a string base64 para bytes
	qrBytes, err := base64.StdEncoding.DecodeString(qrBase64)
	if err != nil {
		return fmt.Errorf("falha ao decodificar QR code: %v", err)
	}

	// Anexa o QR code como um arquivo
	m.Embed("qrcode.png", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(qrBytes)
		return err
	}))

	d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)
	return d.DialAndSend(m)
}

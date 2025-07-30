package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"mime/quotedprintable"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

type EmailService struct {
	cfg EmailConfig
}

func NewEmailService(cfg EmailConfig) *EmailService {
	return &EmailService{cfg: cfg}
}

func (s *EmailService) EnviarEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.cfg.Username, s.cfg.Password, s.cfg.SMTPHost)

	msg := []byte("From: " + s.cfg.From + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)
	return smtp.SendMail(addr, auth, s.cfg.From, []string{to}, msg)
}

func (s *EmailService) EnviarEmailConAdjunto(to, subject, htmlBody, filePath string) error {
	auth := smtp.PlainAuth("", s.cfg.Username, s.cfg.Password, s.cfg.SMTPHost)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	boundary := writer.Boundary()

	// Encabezado del mensaje
	headers := make(map[string]string)
	headers["From"] = s.cfg.From
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "multipart/mixed; boundary=" + boundary

	var message strings.Builder
	for k, v := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	message.WriteString("\r\n")

	bodyPart := fmt.Sprintf("--%s\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n"+
		"Content-Transfer-Encoding: quoted-printable\r\n\r\n%s\r\n", boundary, quotedprintableBody(htmlBody))
	message.WriteString(bodyPart)

	if filePath != "" {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("no se pudo leer el archivo: %w", err)
		}

		encoded := base64.StdEncoding.EncodeToString(fileContent)
		filename := filepath.Base(filePath)

		attachment := fmt.Sprintf("--%s\r\nContent-Type: application/octet-stream\r\n"+
			"Content-Disposition: attachment; filename=\"%s\"\r\n"+
			"Content-Transfer-Encoding: base64\r\n\r\n%s\r\n", boundary, filename, splitBase64(encoded))

		message.WriteString(attachment)
	}

	message.WriteString("--" + boundary + "--")

	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)
	return smtp.SendMail(addr, auth, s.cfg.From, []string{to}, []byte(message.String()))
}

func quotedprintableBody(body string) string {
	var buf bytes.Buffer
	qp := quotedprintable.NewWriter(&buf)
	qp.Write([]byte(body))
	qp.Close()
	return buf.String()
}

func splitBase64(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i += 76 {
		end := i + 76
		if end > len(s) {
			end = len(s)
		}
		result.WriteString(s[i:end] + "\r\n")
	}
	return result.String()
}

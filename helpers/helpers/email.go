package helpers

import (
	"fmt"
	"net/smtp"
)

func EnviarCorreoVerificacion(destinatario string, codigo string) error {
	// Configuración de tu servidor SMTP (Ejemplo usando Gmail)
	// Nota: Si usas Gmail, debes generar una "Contraseña de aplicación" en tu cuenta de Google.
	from := "tu_correo@gmail.com"
	password := "tu_contraseña_de_aplicacion"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Cuerpo del mensaje
	subject := "Subject: Código de Verificación - GESTOCK\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; padding: 20px; color: #333;">
			<h2 style="color: #2c3e50;">GESTOCK - Verificación de Identidad</h2>
			<p>Has solicitado un código de verificación. Tu código de acceso es:</p>
			<h1 style="background: #f4f4f4; padding: 10px; display: inline-block; letter-spacing: 5px; color: #e74c3c;">%s</h1>
			<p>Este código es válido por unos minutos. Si no solicitaste este cambio, ignora este mensaje.</p>
		</div>
	`, codigo)

	message := []byte(subject + mime + body)

	// Autenticación SMTP
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Envío del correo
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{destinatario}, message)
	return err
}
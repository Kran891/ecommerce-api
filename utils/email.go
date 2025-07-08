// utils/email.go
package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendPasswordResetEmail(email, token string) error {
	godotenv.Load()
	from := mail.NewEmail("eCommerce Support", os.Getenv("EMAIL_SENDER"))
	to := mail.NewEmail("User", email)
	resetLink := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_RESET_URL"), token)
	subject := "Hello Form Demo"
	content := fmt.Sprintf("Demo of : %s", resetLink)
	message := mail.NewSingleEmail(from, subject, to, content, content)
	client := sendgrid.NewSendClient(os.Getenv("MY_API"))
	_, err := client.Send(message)
	fmt.Print(err)
	return err
}

// func main() {
// 	fmt.Println(uuid.New())
// }

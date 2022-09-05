package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {

	//If you want to compile this application, this code starts the command line interface (for windows).
	cmd := exec.Command("sth")
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	cmd.Run()

	//Establishing a secure connection to the mail server.
	c, err := smtp.DialTLS("smtp.server.ru:port", nil) // For example: "smtp.mail.ru:465" / "smtp.gmail.com:465"
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Соеденение с сервером установлено ( ͡❛ ͜ʖ ͡❛)")
	}

	//Set current month.
	today := time.Now()
	mounyh := today.Month().String()
	switch mounyh {
	case "January":
		mounyh = "Январь"
	case "February":
		mounyh = "Февраль"
	case "March":
		mounyh = "Март"
	case "April":
		mounyh = "Апрель"
	case "May":
		mounyh = "Май"
	case "June":
		mounyh = "Июнь"
	case "July":
		mounyh = "Июль"
	case "August":
		mounyh = "Август"
	case "September":
		mounyh = "Сентябрь"
	case "October":
		mounyh = "Октябрь"
	case "November":
		mounyh = "Ноябрь"
	case "December":
		mounyh = "Декабрь"
	}

	//Getting user input.
	var water, electricity string
	fmt.Printf("Показания счетчиков за: %s.\n", mounyh)
	fmt.Println("Вода:")
	fmt.Scan(&water)
	fmt.Println("Электричество:")
	fmt.Scan(&electricity)
	fmt.Println("Отправить показания?\nВведите: Y/N")
	var answer string
	for {
		fmt.Scan(&answer)
		if answer == "Y" || answer == "y" {
			break
		} else if answer == "N" || answer == "n" {
			os.Exit(0)
			break
		} else {
			fmt.Println("Введите: Y/N")
		}
	}

	// Setup authentication information.
	auth := sasl.NewPlainClient("", "sender@mail.ru", "password")
	err = c.Auth(auth)
	if err != nil {
		log.Fatal(err)
	}

	//Set the recipient/recipients.
	to := []string{"recipient@mail.ru"}

	//Email Configuration.
	msg := strings.NewReader("To: recipient@mail.ru\r\n" +
		"Subject: Показания\r\n" +
		"\r\n" +
		"This is the email body \r\n\n" + mounyh + "\n" + "Электричесвто: " + electricity + "\n" + "Вода: " + water + "\n")

	//Sending an email.
	err = c.SendMail("sender@mail.ru", to, msg)
	if err != nil {
		log.Fatal(err)
	}

	//Exit from the application.
	os.Exit(0)
}

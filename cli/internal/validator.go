package internal

import (
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

func IsValidEmail(email string) bool {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`) //regex format
	return regex.MatchString(email)
}

func IsDomainValid(domain string) bool {
	_, err := net.LookupMX(domain) // DNS MX record
	if err != nil {
		fmt.Println("Error resolving domain:", err)
		return false
	}
	return true
}

func IsEmailReachable(email string) bool {
	parts := strings.Split(email, "@")
	domain := parts[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("No MX records found:", err)
		return false
	}

	
	smtpServer := mxRecords[0].Host

	conn, err := net.DialTimeout("tcp", smtpServer+":25", 10*time.Second) //tcp connection with port :25 for smtp
	if err != nil {
		fmt.Println("SMTP connection failed:", err)
		return false
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		fmt.Println("Failed to create SMTP client:", err)
		return false
	}
	defer client.Close()

	client.Hello("localhost") //msg 
	err = client.Mail("sender@mail.com") // sender
	if err != nil {
		fmt.Println("MAIL FROM failed:", err)
		return false
	}

	err = client.Rcpt(email) //recipient
	if err != nil {
		fmt.Println("RCPT TO failed, email likely invalid: ", err)
		return false
	}

	return true
}
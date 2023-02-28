package main

import (
	"log"
	"net/http"
	"net/smtp"
	"time"
)

const (
	domain   = "https://tayfunguler.org"
	from     = "info@tayfunguler.org"
	to       = "gulertayfun@outlook.com"
	host     = "smtp.tayfunguler.org"
	port     = "587"
	password = "password"
	msg      = "From: " + from + "\n" + "To: " + to + "\n" + "Subject: Down!! \n\n" + "Project down!!"
	period   = 5
)

func checkStatus(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false, err
	}

	return true, nil
}

func sendMail(msg string) {
	toMail := []string{to}
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toMail, []byte(msg))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	for {
		status, _ := checkStatus(domain)

		if !status {
			sendMail(msg)
		}

		time.Sleep(time.Minute * period)
	}
}

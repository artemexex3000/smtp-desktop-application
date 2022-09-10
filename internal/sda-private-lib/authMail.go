package sda_private_lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Identity string `json:"identity"`
	Host     string `json:"host"`
}

func Auth() {
	users := jsonReader()

	pass, resPass := os.LookupEnv("SECRET_KEY")
	if !resPass {
		fmt.Println("There is no password")
		os.Exit(1)
	}

	email, resEmail := os.LookupEnv("MAIL")
	if !resEmail {
		fmt.Println("There is no email")
	}

	auth := smtp.PlainAuth(
		users.Users[0].Identity,
		email,
		pass,
		users.Users[0].Host,
	)

	fmt.Println("Successfully connected to host")

	to := []string{email}

	msg := []byte("To: " + email + "\r\n" +
		"Subject: Sory za spam!\r\n" +
		"\r\n" +
		"duze vazlivo.\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Check your email!")
}

func jsonReader() *Users {
	var users *Users

	jsonAuth, err := os.Open("./api/authMailConf.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened json file")

	defer func(jsonAuth *os.File) {
		err := jsonAuth.Close()
		if err != nil {

		}
	}(jsonAuth)

	byteValue, _ := ioutil.ReadAll(jsonAuth)

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		log.Fatal(err)
	}

	return users
}

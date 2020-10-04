package main

import (
	"fmt"

	"gitlab.com/mahdiidarabi/exg/authexg"
	"gitlab.com/mahdiidarabi/exg/dbexg"
)

func main() {
	fmt.Println("this is main func")
	authexg.Auth()

	_, err := dbexg.SetConnection()
	if err != nil {
		fmt.Println("got error in setting connection to database")
	}
}

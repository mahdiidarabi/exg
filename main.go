package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/authexg"
	"gitlab.com/mahdiidarabi/exg/dbexg"
)

func main() {
	fmt.Println("this is main func")

	_, err := dbexg.SetConnection()
	if err != nil {
		fmt.Println("got error in setting connection to database")
	}

	r := gin.Default()

	public := r.Group("/")
	public.POST("/createtable", dbexg.CreateTable)
	public.POST("/register", authexg.Register)

	r.Run("localhost:8080")
}

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/authexg"
	"gitlab.com/mahdiidarabi/exg/dbexg"
)

func main() {

	fmt.Println("you know, this is main function")

	_, err := dbexg.SetConnection()
	if err != nil {
		fmt.Println("got error in setting connection to database")
	}

	r := gin.Default()

	public := r.Group("/")
	public.POST("/register", authexg.Register)
	public.POST("/login", authexg.Login)
	public.GET("/", dbexg.GetAllUsers)

	private := r.Group("/user")
	private.Use(jwt.Auth(os.Getenv("ACCESS_SECRET")))

	private.POST("/createtable", dbexg.CreateTable)
	private.GET("/:email", dbexg.GetUser)
	private.GET("/", dbexg.GetAllUsers)
	private.DELETE("/:email", dbexg.DeleteUser)
	private.PATCH("/:email", dbexg.UpdateUser)
	//private.PUT("/:email", dbexg.UpdateUser)

	r.Run("localhost:8080")
}

package authexg

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/dbexg"
	"gitlab.com/mahdiidarabi/exg/model"
)

func Register(c *gin.Context) {
	fmt.Println("this is register in authexg package")
	dbexg.AddUser(c)
}

func Login(c *gin.Context) {

	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided in login api")
		panic(err)
	}

	output := dbexg.GetUserbyEmail(input.Email)

	// TODO
	//compare the user from the request, with the one we defined:
	if input.Email != output.Email || input.Password != output.Password {
		c.JSON(http.StatusUnauthorized, "Invalid username or password")
	} else {
		CreateToken(c)
	}
}

func CreateToken(c *gin.Context) {

	secret := os.Getenv("ACCESS_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = "email"
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("got err in creating token")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

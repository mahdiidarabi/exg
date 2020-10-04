package authexg

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/dbexg"
)

func Register(c *gin.Context) {
	fmt.Println("this is register in authexg package")
	dbexg.AddUser(c)
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
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

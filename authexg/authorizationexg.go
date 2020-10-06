package authexg

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
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
	atClaims["isAdmin"] = false
	atClaims["authorized"] = true
	atClaims["user_email"] = c.Request.Header["Email"]
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("got err in creating token")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ExtractToken(c *gin.Context) {

	fmt.Println(c.Request.Header["Email"])

	bearerToken := c.Request.Header.Get("Authorization")
	//
	fmt.Println(bearerToken)
	fmt.Println(reflect.TypeOf(bearerToken))

	bearerTokenFinal := bearerToken[7:]
	fmt.Println(bearerTokenFinal)
	//
	//// claims := bearerToken.Claims(jwt.MapClaims{})

	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDE5NzY1NTMsImlzQWRtaW4iOmZhbHNlLCJ1c2VyX2VtYWlsIjoiZW1haWwifQ.RciKSL3niB_vccLAbYEGN5q66Thfl2FEqydrfXngsd4"

	tokenString := bearerTokenFinal

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_email"], claims["exp"])
	} else {
		fmt.Println(err)
	}

}

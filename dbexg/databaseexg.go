package dbexg

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/mahdiidarabi/exg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetConnection() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("got error in loading .env")
	}
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Println(dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("got error in setConnection")
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return DB, nil
}

func AddUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("got error in addUser in dbexg package")
	}

	result := DB.Create(&user)
	if result.Error != nil {
		fmt.Println("got error in addUser, when creating DB row")
	}

	fmt.Println("Successfully user added")

}

func CreateTable(c *gin.Context) {

	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("got error in createTable")
		panic(err)
	}
	fmt.Println("SUccessfully migrated ( table created)")
}

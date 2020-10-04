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

func CreateTable(c *gin.Context) {

	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("got error in createTable")
		panic(err)
	}
	fmt.Println("SUccessfully migrated ( table created)")
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

func DeleteUser(c *gin.Context) {

	var input model.User
	var output model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided in deleteUser function in dbexg package")
	}

	if err := DB.First(&output, "Email = ?", input.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		fmt.Println("got error in deleteUser in dbexg package")
	}

	DB.Delete(&output)

	c.JSON(http.StatusOK, gin.H{"data": output})
}

func UpdateUser(c *gin.Context) {

	var input model.User
	var output model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided in updateUser function in dbexg package")
	}

	if err := DB.First(&output, "Email = ?", input.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		fmt.Println("got error in updateUser in dbexg package")
	}

	DB.Model(&output).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": output})

}

func GetAllUsers(c *gin.Context) {
	var users []model.User

	DB.Find(&users)

	fmt.Println("getAllUsers function in dbexg package")

	c.JSON(http.StatusOK, gin.H{"data": users})

}

func GetUser(c *gin.Context) {

	// var input model.User
	var output model.User

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, "Invalid json provided in getUser function in dbexg package")
	// }

	if err := DB.First(&output, "Email = ?", c.Request.Header["Email"]).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		fmt.Printf("got error in getUser in dbexg package. email")

	}

	fmt.Println(c.Request.URL)

	c.JSON(http.StatusOK, gin.H{"data": output})
}

func GetUserbyEmail(email string) *model.User {

	var user model.User

	if err := DB.First(&user, "Email = ?", email).Error; err != nil {
		fmt.Println("got error in getUserbyEmail in dbexg package")
		return nil
	}

	return &user
}

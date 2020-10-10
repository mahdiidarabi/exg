package dbexg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/mahdiidarabi/exg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var DB *gorm.DB

func SetConnection() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("got error in loading .env")
	}
	var dsn = fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	fmt.Println(dsn)

	//var dsn = fmt.Sprintf("user=mahdi password=mahdipass dbname=test port=5432 sslmode=disable")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("got error in setConnection")
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return DB, nil
}

func CreateTable(c *gin.Context) {

	if DB.Migrator().HasTable(&model.User{}) {

		var s =fmt.Sprintf("We already have a table, so dont create a new one")
		c.JSON(http.StatusOK, gin.H{"data": s})

	} else {

		err := DB.AutoMigrate(&model.User{})
		if err != nil {
			var s =fmt.Sprintf("got error in createTable")
			c.JSON(http.StatusBadRequest, gin.H{"data": s})
		}
		var s = fmt.Sprintf("Successfully migrated ( table created)")
		c.JSON(http.StatusOK, gin.H{"data": s})

	}
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
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

	fmt.Println("Successfully user added")
	c.JSON(http.StatusOK, gin.H{"Data": user})

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

	result := DB.Delete(&output)
	if result.Error != nil {
		fmt.Println("got error in deleteUser, when deleting DB row")
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

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

	result := DB.Model(&output).Updates(input)
	if result.Error != nil {
		fmt.Println("got error in updateUser, when updating DB row")
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

	c.JSON(http.StatusOK, gin.H{"data": output})

}

func GetAllUsers(c *gin.Context) {
	var users []model.User

	result := DB.Find(&users)

	if result.Error != nil {
		fmt.Println("got error in getAllUsers, when fetching DB rows")
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/model"
	"gitlab.com/mahdiidarabi/exg/dbexg"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// var DB *gorm.DB


func TestSetConnection(t *testing.T) {
	DB, err := dbexg.SetConnection()

	if err != nil {
		t.Errorf("got error in testing SetConnection")
	}

	if DB.Migrator().HasTable(&model.User{}) {

		fmt.Println("We already have a table, so dont create a new one")

	} else {
		t.Errorf("can not creating databse in SetConnection")
	}
	fmt.Println("Connection is set in the test")
}

func TestCreateTable(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	//if err != nil {
	//	t.Errorf("can not make a gin context with CreateTestContext")
	//}

	requstBody := model.User{}

	bufVar := new(bytes.Buffer)

	json.NewEncoder(bufVar).Encode(requstBody)

	request := httptest.NewRequest("POST", "localhost:8080/register", bufVar)

	c.Request = request

	dbexg.CreateTable(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	fmt.Println("Table is created in the test")


}


func TestAddUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	//if err != nil {
	//	t.Errorf("can not make a gin context with CreateTestContext")
	//}

	// requstBody := user10
	bufVar := new(bytes.Buffer)

	json.NewEncoder(bufVar).Encode(dbexg.User00)

	request := httptest.NewRequest("POST", "localhost:8080/register", bufVar)

	c.Request = request

	dbexg.AddUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email != dbexg.User00.Email || result.Data.Password != dbexg.User00.Password{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	fmt.Println("a user is added in the test")

}

func TestUpdateUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	//if err != nil {
	//	t.Errorf("can not make a gin context with CreateTestContext")
	//}

	// requstBody := user11

	bufVar := new(bytes.Buffer)

	json.NewEncoder(bufVar).Encode(dbexg.User01)

	request := httptest.NewRequest("PATCH", "localhost:8080/user/", bufVar)

	c.Request = request

	dbexg.UpdateUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email != dbexg.User01.Email || result.Data.Phone == dbexg.User00.Phone{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	fmt.Println("the user is updated in the test")

}


func TestDeleteUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	//if err != nil {
	//	t.Errorf("can not make a gin context with CreateTestContext")
	//}

	// requstBody := user11

	bufVar := new(bytes.Buffer)

	json.NewEncoder(bufVar).Encode(dbexg.User01)

	request := httptest.NewRequest("PATCH", "localhost:8080/user/", bufVar)

	c.Request = request

	dbexg.DeleteUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email != dbexg.User01.Email || result.Data.Password != dbexg.User01.Password{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	fmt.Println("the user is deleted in the test")

}

func TestGetUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	//if err != nil {
	//	t.Errorf("can not make a gin context with CreateTestContext")
	//}

	// requstBody := user11

	bufVar := new(bytes.Buffer)

	json.NewEncoder(bufVar).Encode(dbexg.User01)

	request := httptest.NewRequest("PATCH", "localhost:8080/user/", bufVar)

	c.Request = request

	dbexg.GetUser(c)

	// should not find the user , since it was deleted in last step
	if w.Code != 400 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email == dbexg.User01.Email || result.Data.Password == dbexg.User01.Password{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	fmt.Println("the user is not found, since it was deleted in the last step")

}



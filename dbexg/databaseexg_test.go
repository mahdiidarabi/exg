package dbexg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// var DB *gorm.DB


func TestSetConnection(t *testing.T) {
	DB, err := SetConnection()

	if err != nil {
		t.Errorf("got error in testing SetConnection")
	}

	if DB.Migrator().HasTable(&model.User{}) {

		fmt.Println("We already have a table, so dont create a new one")

	} else {
		t.Errorf("can not creating databse in SetConnection")
	}
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

	CreateTable(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}


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

	json.NewEncoder(bufVar).Encode(User00)

	request := httptest.NewRequest("POST", "localhost:8080/register", bufVar)

	c.Request = request

	AddUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email != User00.Email || result.Data.Password != User00.Password{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

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

	json.NewEncoder(bufVar).Encode(User01)

	request := httptest.NewRequest("PATCH", "localhost:8080/user/", bufVar)

	c.Request = request

	UpdateUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email != User01.Email || result.Data.Phone == User00.Phone{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

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

	json.NewEncoder(bufVar).Encode(User01)

	request := httptest.NewRequest("PATCH", "localhost:8080/user/", bufVar)

	c.Request = request

	DeleteUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result
	// var user model.User

	json.Unmarshal(w.Body.Bytes(), &result)
	// json.Unmarshal(bufVar.Bytes(), &user)


	// fmt.Println(result.Data)

	if result.Data.Email != User01.Email || result.Data.Password != User01.Password{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

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

	json.NewEncoder(bufVar).Encode(User01)

	request := httptest.NewRequest("PATCH", "localhost:8080/user/", bufVar)

	c.Request = request

	GetUser(c)

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

	if result.Data.Email == User01.Email || result.Data.Password == User01.Password{
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

}



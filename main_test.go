package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/mahdiidarabi/exg/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"gitlab.com/mahdiidarabi/exg/dbexg"
)



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


}


func TestAddUser(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	//if err != nil {
	//	t.Errorf("can not make a gin context with CreateTestContext")
	//}

	requstBody := model.User{
		Username:    "mahdi56",
		Email:       "mahdi@gmail.com56",
		Phone:       "phone56",
		Password:    "mahdipass56",
		BtcBalance:  0,
		EthBalance:  0,
		DashBalance: 0,
		TethBalance: 0,
		XrpBalance:  0,
		BinBalance:  0,
		EosBalance:  0,
	}

	bufVar := new(bytes.Buffer)

	json.NewEncoder(bufVar).Encode(requstBody)

	request := httptest.NewRequest("POST", "localhost:8080/register", bufVar)

	c.Request = request

	dbexg.AddUser(c)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}

	var result model.Result

	json.Unmarshal(w.Body.Bytes(), &result)


	fmt.Println(result.Data.Email)

}


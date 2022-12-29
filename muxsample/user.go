package muxsample

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

const DNS = "root:123456@tcp(127.0.0.1:3306)/gomuxsample?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// set return header
	w.Header().Set("Content-Type", "application/json")

	// get users from db
	var users []User
	DB.Find(&users)

	// return users
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// set return header
	w.Header().Set("Content-Type", "application/json")

	// map of params
	params := mux.Vars(r)

	// get by id
	var user User
	DB.First(&user, params["id"])

	// return added user
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// set return header
	w.Header().Set("Content-Type", "application/json")

	// read user from body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// add user to db
	DB.Create(&user)

	// return added user
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// set return header
	w.Header().Set("Content-Type", "application/json")

	// map of params
	params := mux.Vars(r)

	// get by id
	var user User
	DB.First(&user, params["id"])

	// update user
	json.NewDecoder(r.Body).Decode(&user)

	DB.Save(&user)

	// return user
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// set return header
	w.Header().Set("Content-Type", "application/json")

	// map of params
	params := mux.Vars(r)

	// get by id
	var user User
	DB.Delete(&user, params["id"])

	// return added user
	json.NewEncoder(w).Encode("User deleted!")
}

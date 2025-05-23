package user

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterRunners = []rune("qwertyuiopasdfghjklzxcvbnm1234567890")

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt time.Time `json:"updateAt"`
}


func (acc *User) Ouput() {
  color.Cyan(acc.Name)
  color.Cyan(acc.Password)
  color.Cyan(acc.Url)
}


func (acc *User) generetPassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunners[rand.IntN(len(letterRunners))]
	}
	acc.Password = string(result)
}

func NewAccount(name,password,urlString string) (*User, error) {
	if name == ""{
	return nil,errors.New("INVALID_LOGIN")
   } 
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &User{
	CreatedAt:  time.Now(),
	UpdateAt: time.Now(),
	Name: name,
	Password: password,
	Url: urlString,
	}
	if password == ""{
       newAcc.generetPassword(8)
	}
	return newAcc,nil
}

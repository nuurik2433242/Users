package main

import (
	// "demo/app-5/cloud"
	"demo/app-5/encripter"
	"demo/app-5/failes"
	"demo/app-5/ouput"
	"demo/app-5/user"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*user.VaultWithDb){
  "1": createdAccount,
  "2": findAccountByUrl,
  "3": findAccountByName,
  "4": deleteAccount,
}

var menuVariants = []string{
  "1.Создать аккаунт",
    "2.Найти аккаунт по URL",
    "3.Найти аккаунт по логину",
    "4.Удалить аккаунт",
    "5.Выход",
    "Выберите вариант",
  }

func main() {
  fmt.Println("__Менеджер паролей__")
  err := godotenv.Load()
  if err != nil {
    ouput.PrintError("Не удалось найти env файл")
  }
  vault := user.NewVault(failes.NewJsonDb("data.vault"), *encripter.NewEncripter())
  // vault := user.NewVault(cloud.NewCloudDb("https://google.ru"))
  Menu:
  for {  
  menuAcc := addUser(menuVariants...)
  menuFunc := menu[menuAcc]
  if menuFunc == nil {
    break Menu
  }
  menuFunc(vault)
  }
  
}

func createdAccount(vault *user.VaultWithDb)  {
  name := addUser("Вывидите логин")
  password := addUser("Вывидите пароль")
  url := addUser("Вывидите url")
  myAccount,err := user.NewAccount(name,password,url)
  if err != nil {
  ouput.PrintError("Неверный логин или неверный формат URL")
  return
  }
  vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *user.VaultWithDb) { 
  url := addUser("Вывидите url для пойска") 
  users := vault.SearchAccount(url,func(acc user.User,str string)bool{
    return strings.Contains(acc.Url, str)
  })
  ouputResult(&users)
}

func findAccountByName(vault *user.VaultWithDb) { 
  url := addUser("Вывидите Login для пойска") 
  users := vault.SearchAccount(url,func(acc user.User,str string)bool{
    return strings.Contains(acc.Name, str)
  })
 ouputResult(&users)
}

func ouputResult(users *[]user.User) {
  if len(*users) == 0{
    ouput.PrintError("Аккаунтов не найдено")
  }
  for _, user := range *users {
   user.Ouput()
  }
}

func deleteAccount(vault *user.VaultWithDb) {
  url := addUser("Вывидите url для удаления")
  isDelete := vault.DeleteAccount(url)
  if isDelete{
    color.Green("Удаленно")
  }else{
   ouput.PrintError("Не найденно")
  }
}

func addUser(add ...string) string {
	for i,value := range add{
    if i == len(add) - 1{
      fmt.Printf("%v:",value)
    }else{
      fmt.Println(value)
    }
  }
  var res string
	fmt.Scanln(&res)
  return res
}













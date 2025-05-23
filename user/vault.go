package user

import (
	"demo/app-5/encripter"
	"demo/app-5/ouput"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface{
	Read() ([]byte, error)
}

type ByteWriter interface{
    Write([]byte)
}

type Db interface{
	ByteReader
	ByteWriter
}

type Vault struct {
	Users  []User  `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

type VaultWithDb struct{
	Vault
	db Db
	enc encripter.Encripter
}

func NewVault(db Db, enc encripter.Encripter) *VaultWithDb {
	file,err := db.Read()
	if err != nil {
		return &VaultWithDb{
		Vault: Vault{
		  Users: []User{},
		  UpdateAt:  time.Now(),
		},
		db: db,
		enc: enc,
	}
	}
	data := enc.Decryter(file)
	var vault Vault
	err = json.Unmarshal(data,&vault)
	color.Cyan("Найдено аккаунтов %d", len(vault.Users))
	if err != nil {
		ouput.PrintError("Не удалось разобрать файл")
		return &VaultWithDb{
		Vault: Vault{
			Users: []User{},
		    UpdateAt:  time.Now(),
		},
		    db: db,
			enc: enc,
	}
	}
	return &VaultWithDb{
		Vault: vault,
		db: db,
		enc: enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc User) {
	vault.Users = append(vault.Users, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte,error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil,err
	}
	return file,nil
}

func (vault *VaultWithDb) SearchAccount(str string,checker func(User,string)bool) []User {
  var users []User
  for _,account := range vault.Users {
	isMatch := checker(account,str)
	if isMatch {
		users = append(users, account)
	}
  }
  return users
}

func (vault *VaultWithDb) DeleteAccount(url string) bool {
  var users []User
  usersBool := false
  for _,account := range vault.Users {
	isMatch := strings.Contains(account.Url,url)
	if !isMatch {
		users = append(users, account)
		continue
	} 
	usersBool = true
  }
	vault.Users = users
	vault.save()
    return usersBool
}

func (vault *VaultWithDb) save() {
   vault.UpdateAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encripter(data)
	if err != nil {
		ouput.PrintError("Не удалось разобрать файл")
	}
	vault.db.Write(encData)
}
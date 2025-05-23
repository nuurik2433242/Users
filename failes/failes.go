package failes

import (
	"demo/app-5/ouput"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb (name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error){
   date,err := os.ReadFile(db.filename)
   if err != nil {
    return nil, err
   }
   return date,nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		ouput.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil{
		ouput.PrintError(err)
		return 
	}
	color.Green("Запись успешно")
}
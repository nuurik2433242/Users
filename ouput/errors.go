package ouput

import "github.com/fatih/color"

func PrintError(value any) {
    intValue,ok := value.(int)
	if ok {
		color.Red("Неверный формат ошибки: %d", intValue)
		return
	}
	strValue,ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Неизвестный формат")
	
	// switch t := value.(type) {
	// case string:
	// 	color.Red(t)
	// case int:
	// 	color.Red("Неверный формат ошибки: %d", t)
	// case error: 
	// color.Red(t.Error())
	// default: color.Red("Неизвестный формат")
	// }
}
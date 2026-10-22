package actioninfo

import (
	"fmt"
	"log"
)

// Интерфейс, в котором объявлены сигнатуры методов Pasrt() и ActionInfo(),
// реализованных в пакетах trainings и daysteps
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Метод принимает набор строк с сырыми данными об активности и объект интерфейса DataParser,
// парсит каждую строку, выводит отчеты на экран и логирует возможные ошибки
func Info(dataset []string, dp DataParser) {
	for _, item := range dataset {
		err := dp.Parse(item)
		if err != nil {
			log.Println("Возникла ошибка при парсинге", err)
			continue
		}

		infoStr, err := dp.ActionInfo()
		if err != nil {
			log.Println("Возникла ошибка при формировании информации об активности", err)
			continue
		}
		fmt.Print(infoStr)
	}
}

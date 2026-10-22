package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

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

package debug

import (
	"guest/settings"
	"log"
)

func PrintInterface(data ...interface{}) {
	for _, v := range data {
		body, _ := settings.Stringify(v, settings.FormatJson)

		log.Println(string(body))
	}
}

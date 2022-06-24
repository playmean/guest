package debug

import (
	"log"

	"github.com/playmean/guest/settings"
)

func PrintInterface(data ...interface{}) {
	for _, v := range data {
		body, _ := settings.Stringify(v, settings.FormatJson)

		log.Println(string(body))
	}
}

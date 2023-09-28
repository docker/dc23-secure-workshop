package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func main() {
	fmt.Println("Starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		e := message.NewPrinter(language.English)
		enMessage := e.Sprintf("%v bottles of beer on the wall.", number.Decimal(1234))

		f := message.NewPrinter(language.French)
		frMessage := f.Sprintf("%v bouteilles de bière sur le mur", number.Decimal(1234))

		_, _ = fmt.Fprintf(w, "<html><body><ul><li>%s</li><li>%s</li></ul></body></html>", enMessage, frMessage)
	})

	log.Fatal(http.ListenAndServe(":9090", nil))
}

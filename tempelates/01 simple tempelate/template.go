package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//tempelate.Parefiles can be used to store multiple files into these tempelate containers
	// tpl, err := template.ParseFiles("tpl.html")
	tpl, err := template.ParseGlob("./*.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdin, "tpl.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"html/template"
	"os"
)

func main() {
	tmpl, _ := template.New("inline").Parse("From: {{.Sender}}\nTo:   {{.Recipient}}\n")
	data := struct {
		Sender    string
		Recipient string
	}{"Luke", "Krista"}
	_ = tmpl.Execute(os.Stdout, data)
}

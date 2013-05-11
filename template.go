package main

import (
	"os"
	"text/template"
)

func main() {
	t, _ := template.New("foo").Parse("Hello, {{.}}!")
	t.Execute(os.Stdout, "devcamp 2013")
}

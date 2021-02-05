package utils

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"html/template"
)

func BuildEmail(userdata interface{}) (string, error) {

	html, err := template.ParseFiles("email.min.html")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var tpl bytes.Buffer

	if err := html.Execute(&tpl, userdata); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

package tools

import (
	"strings"
	"text/template"
)

func TextTemplate(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	return b.String()
}

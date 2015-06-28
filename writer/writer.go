package writer

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
	"text/template"

	"github.com/jordwest/xexpressions/compiler"
)

var wordRegexp = regexp.MustCompile("[0-9A-Za-z]+")

type Pipeline struct {
	RegularExpressions []compiler.Regexp
	TemplateFilename   string
}

func WriteRegexps(re []compiler.Regexp, templateFilename string, outputFile io.Writer) error {
	data, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		return err
	}
	tmpl, err := template.New("output").Funcs(template.FuncMap{
		"Line":           Underline,
		"UpperCamelCase": UpperCamelCase,
		"LowerCamelCase": LowerCamelCase,
		"UpperCase":      UpperCase,
		"LowerCase":      LowerCase,
	}).Parse(string(data))
	if err != nil {
		return err
	}

	pipeline := Pipeline{
		RegularExpressions: re,
		TemplateFilename:   templateFilename,
	}

	err = tmpl.Execute(outputFile, pipeline)
	if err != nil {
		return err
	}

	return nil
}

// Gives a line string, eg "-----" of the same length as the text argument
func Underline(text string) string {
	line := ""
	for len(line) < len(text) {
		line += "-"
	}
	return fmt.Sprintf("%s", line)
}

func UpperCamelCase(text string) string {
	chunks := wordRegexp.FindAllString(text, -1)
	for idx, val := range chunks {
		chunks[idx] = strings.Title(val)
	}
	return strings.Join(chunks, "")
}

func LowerCamelCase(text string) string {
	chunks := wordRegexp.FindAllString(text, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = strings.Title(val)
		} else {
			chunks[idx] = strings.ToLower(val)
		}
	}
	return strings.Join(chunks, "")
}

func LowerCase(text string) string {
	chunks := wordRegexp.FindAllString(text, -1)
	for idx, val := range chunks {
		chunks[idx] = strings.ToLower(val)
	}
	return strings.Join(chunks, "_")
}

func UpperCase(text string) string {
	chunks := wordRegexp.FindAllString(text, -1)
	for idx, val := range chunks {
		chunks[idx] = strings.ToUpper(val)
	}
	return strings.Join(chunks, "_")
}

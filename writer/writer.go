package writer

import (
	"io"
	"io/ioutil"
	"text/template"

	"github.com/jordwest/xexpressions/compiler"
)

type Pipeline struct {
	RegularExpressions []compiler.Regexp
}

func WriteRegexps(re []compiler.Regexp, templateFilename string, outputFile io.Writer) error {
	data, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		return err
	}
	tmpl, err := template.New("output").Parse(string(data))
	if err != nil {
		return err
	}

	pipeline := Pipeline{
		RegularExpressions: re,
	}

	tmpl.Execute(outputFile, pipeline)

	return nil
}

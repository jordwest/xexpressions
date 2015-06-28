package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jordwest/xexpressions/compiler"
	"github.com/jordwest/xexpressions/lexer"
	"github.com/jordwest/xexpressions/writer"
)

var templateFile string

func init() {
	flag.StringVar(&templateFile, "template", "", "Specify the template used to output the regular expressions")
	flag.Parse()
}

func main() {
	if len(flag.Args()) < 1 {
		fmt.Printf("One or more X-Expression files must be specified for compilation\n")
		os.Exit(1)
	}

	if templateFile == "" {
		fmt.Printf("Please specify the template file to use. Eg: -template=\"templates/javascript.js\"\n")
		os.Exit(1)
	}

	ast := lexer.NewASTNode(nil, lexer.Command{})

	// Parse each of the files provided on the command line
	for _, file := range flag.Args() {
		fileAst, err := lexer.ParseFile(file)
		if err != nil {
			fmt.Printf("Error parsing file %s\n\t%s\n", file, err)
			os.Exit(2)
		}

		ast.Include(fileAst)
	}

	regularExpressions, _, err := compiler.CompileRoot(*ast)
	if err != nil {
		fmt.Printf("Compilation error\n\t%s\n", err)
		os.Exit(2)
	}

	// Execute the template and send output to Stdout
	err = writer.WriteRegexps(regularExpressions, templateFile, os.Stdout)
	if err != nil {
		fmt.Printf("Output error\n\t%s\n", err)
		os.Exit(2)
	}
}

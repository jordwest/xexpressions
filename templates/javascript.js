// The following Regular Expressions were generated from XExpressions
// https://github.com/jordwest/xexpressions

var regular_expressions = {};

{{range .RegularExpressions}}
/**
  * {{.TextName}}
  * ---------
  * {{.Description}}
  * Definition: {{.Source}}
  */
regular_expressions['{{.TextName}}'] = /{{.RegexpText}}/;
{{end}}

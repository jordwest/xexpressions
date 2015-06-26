// The following Regular Expressions were generated from XExpressions
// https://github.com/jordwest/xexpressions

var regular_expressions = {};

{{range .RegularExpressions}}
/** {{.TextName}}
  * ---------
  *
  * {{.Description}}
  */
regular_expressions['{{.TextName}}'] = /{{.RegexpText}}/;

{{end}}

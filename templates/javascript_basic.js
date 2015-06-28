// The following Regular Expressions were generated from XExpressions
// https://github.com/jordwest/xexpressions
// Template: {{.TemplateFilename}}

var REGULAR_EXPRESSIONS = (function(){
  var expressions = {};
  {{range .RegularExpressions}}
  // {{.Description}}
  expressions.{{UpperCase .TextName}} = /{{.RegexpText}}/
  {{end}}
  return expressions;
})();

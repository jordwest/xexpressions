// The following Regular Expressions were generated from XExpressions
// https://github.com/jordwest/xexpressions
// Template: {{.TemplateFilename}}

window.XEXPRESSIONS = (function(){

  var XExpression = function(regexp) {
    this.regexp = regexp;
    this.groups = {'full': 0 };
  }

  XExpression.prototype.exec = function(text) {
    var result = this.regexp.exec(text);
    if(result === null || result.length < 1) {
      return null;
    }
    var match_data = {};
    for(group_name in this.groups) {
      match_data[group_name] = result[this.groups[group_name]];
    }
    return match_data;
  }
  XExpression.prototype.test = function(text) {
    return this.regexp.test(text);
  }
  XExpression.prototype.add_capture_group = function(index, name) {
    this.groups[name] = index;
    return this;
  }

  var xExpressions = {};

  {{range .RegularExpressions}}
  /**
    * {{.TextName}}
    * {{Line .TextName}}
    * {{.Description}}
    * Definition: {{.Source}}
    */
  xExpressions.{{UpperCase .TextName}} = new XExpression(/{{.RegexpText}}/);
  xExpressions.{{UpperCase .TextName}}{{range .CaptureGroups}}
    .add_capture_group({{.Index}}, '{{LowerCase .Name}}'){{end}};
  {{end}}

  return xExpressions;
})();

Xtreme-Expressions
=============

##### Because Regular Expressions are too regular

Xtreme-Expression (or X-Expression) is an extremely verbose, self-documenting regular expression language designed for maintainability in projects that use complex regular expressions.

Write your regular expressions in one place (an `.xexp` file) then compile them into a library for use in your chosen language.

A few features:

 - Built-in unit testing
 - Named groups
 - Self documenting
 - Reuse common patterns by defining custom aliases
 - Write your regular expressions in one place, compile to multiple languages

Here's an example of a regexp converted to an X-Expression.

#### Regular Expression

	^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$

Can you guess what the above regular expression is used for? How about in the following:

#### X-Expression

	Alias: Separator
		'[- /.]'

	XExpression: Date
		Description: Matches a valid date between year 1900 and 2099
		Example:
			Match: 2015-02-21
			Match: 2015/02/21
			Match: 2015 02 21
			Non Match: 3056-02-16
			Non Match: 2015-2-8
			Non Match: 2015-05-34
			Non Match: 2015-13-05

		Group[Capture]: Year
			'(19|20)': Must be in range 1900-2099
			Digit
			Digit
		Separator
		Group[Capture]: Month
			Select:
				Case: Jan to Sept
					'0[1-9]'
				Case: Oct to Dec
					'1[012]'
		Separator
		Group[Capture]: Day
			Select:
				Case: 1st to 9th
					'0[1-9]'
				Case: 10th to 29th
					'[12][0-9]'
				Case: 30th to 31st
					'3[01]'

Which one is easier to understand? (Original example from [Regular-Expressions.info](http://www.regular-expressions.info/examples.html))

Getting Started
---------------

### Installing from source

Once you have golang installed, install this package with `go get`:

	go get github.com/jordwest/xexpressions

Copy the `example` file into your project, then:

  cd example
	xexpressions -template="templates/javascript.js" common.xexpr demos.xexpr > demos.js

Syntax Overview
---------------

 - Each 'command' is in the format `CommandName[Parameters]: Comment`. Usually the comment is
	 ignored. In some commands the comment is used, for example to define the name of
	 the X-Expression or the name of a capture group, or to specify a test case.
 - Insert literal regular expression text by enclosed it in 'single quotes', eg `'[A-z0-9]'`
 - Indentation must be tabs. The compiler will ignore spaces.
 - An `Alias` can do almost anything that an XExpression can do, except that the
	 alias won't be output on it's own. It must be called from inside an XExpression.
 - To call an alias, simply type out it's name.
 - Names of aliases or X-Expressions may contain spaces.

Project Status
--------

This is still in a prototype state - I wouldn't use it for any mission critical applications, but it might still be useful for creating and testing your regexps.
It's the first time I've attempted to write
any kind of compiler, and the syntax is open for debate so the language and compiler may
change significantly. If you have a suggestion for an improvement to the syntax
or idea for an additional feature, please post an issue.

Many things need to be improved for this to be a real compiler, for example:

 - [ ] Lots of tidying up and refactoring
 - [ ] Read and parse files as a stream, instead of loading the whole thing into memory
 - [ ] Cleaner interface for defining commands/tokens
 - [ ] Lexer should not assume 1 line == 1 command

Made in Japan @ [Open Source Cafe](http://www.osscafe.net/en/)

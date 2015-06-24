Xtreme-Expressions
=============

##### Because Regular Expressions are too regular

Xtreme-Expression (or X-Expression) is an extremely verbose, self-documenting regular expression language designed for maintainability in projects that use complex regular expressions.

Write your regular expressions in one place (an `.xexp` file) then compile them into a library for use in your chosen language.

A few features:

 - Built-in unit testing
 - Named groups
 - Self documenting
 - Easy to reason about
 - Reuse common patterns by defining custom aliases
 - Embed regular expressions where you need them

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

Made in Japan @ [Open Source Cafe](http://www.osscafe.net/en/)

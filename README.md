Xtreme-Expressions
=============

##### Because Regular Expressions are too regular

Xtreme-Expressions (or X-Expressions for short) are verbose, self-documenting regular expressions designed for maintainability in projects that use complex regular expressions.

Write your regular expressions in one place -- an `.xexp` file -- then compile them into a library for your chosen language.

Here's an example of a date validation regexp converted to an X-Expression

#### Regular Expression

	^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$

#### X-Expression

	Alias: Separator
		'[- /.]'

	XExpression: Date
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

Which one do you think is easier to understand?

Made in Japan @ [Open Source Cafe](http://www.osscafe.net/en/)

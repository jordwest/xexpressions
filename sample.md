Match IP address
=================

Regular Expression
------------------

	\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
		(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
		(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.
		(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b

X-Expression
------------

	Alias: Byte
		Select:
			Case: 250 or more
				'25[0-5]'
			Case: 200 to 249
				'2[0-4][0-9]'
			Case: Under 200
				'[01]?[0-9][0-9]'

	XExpression: IP Address
		Word Boundary
		Group[Capture]: Byte 1
			Byte
		'.'
		Group[Capture]: Byte 2
			Byte
		'.'
		Group[Capture]: Byte 3
			Byte
		'.'
		Group[Capture]: Byte 4
			Byte
		Word Boundary

Matching a date
===============

Regular Expression
------------------

	^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$

X-Expression
------------

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

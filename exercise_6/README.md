
# Exercise 6 

Design a BNF grammar to parse a time specification. All of the
following examples should be accepted:

`4pm, 7:38pm, 23:42, 3:16, 3:16am`

# Solution

```
<time> ::= <12-hour-time> | <24-hour-time>

<12-hour-time> ::= <12-hour-hour><period> | <12-hour-hour>:<minutes><period>

<24-hour-time> ::= <24-hour-hour>:<minutes>

<period> ::= pm | am

<integer> ::= <digit> | <digit><digit>

<12-hour-hour> ::= <positive-digit> | 10 | 11 | 12

<24-hour-value> ::= <12-hour-hour> | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 00 

<minutes> ::= <left-minute><digit> 

<left-minute> ::= 0 | 1 | 2 | 3 | 4 | 5 

<digit> ::= 0 | <positive-digit>

<positive-digit> ::= 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
```

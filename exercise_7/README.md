# Exercise 7 

Implement a parser for the BNF grammar in the previous
exercise using a PEG parser generator in the language of your
choice. The output should be an integer containing the number
of minutes past midnight.

### What does this mean?

Parsers usually consume source code to produce something like an AST 
that can be more easily handled by a compiler.

Parser code tends to grow, thats why this exercise suggests using a PEG parser generator.
This way you can produce the code to parse a grammar from the grammar itself.

## Step 1

I need to translate my BNF grammar into PEG notation so the PEG parser generator
can consume it.

BNF:
```
<time> ::= <12-hour-time> | <24-hour-time>

<12-hour-time> ::= <12-hour-hour><period> | <12-hour-hour>:<minutes><period>

<24-hour-time> ::= <24-hour-hour>:<minutes>

<period> ::= pm | am

<12-hour-hour> ::= <positive-digit> | 10 | 11 | 12

<24-hour-value> ::= 00 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23

<minutes> ::= <left-minute><digit> 

<left-minute> ::= 0 | 1 | 2 | 3 | 4 | 5 

<digit> ::= 0 | <positive-digit>

<positive-digit> ::= 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
``` 

PEG:
```

```
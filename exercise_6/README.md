
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

<12-hour-hour> ::= <positive-digit> | 10 | 11 | 12

<24-hour-value> ::= 00 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23

<minutes> ::= <left-minute><digit> 

<left-minute> ::= 0 | 1 | 2 | 3 | 4 | 5 

<digit> ::= 0 | <positive-digit>

<positive-digit> ::= 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9

```

## Testing

### ¿Is 4pm a valid time?

4pm

4 is a <positive-digit> and therefore  a <12-hour-hour>

pm is a <period>

4pm is a <12-hour-hour><period> therefore it is <time>



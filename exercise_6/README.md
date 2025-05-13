
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

### Is 4pm a valid time?

```
4pm

4 is a <positive-digit> and therefore  a <12-hour-hour>

pm is a <period>

4pm is a <12-hour-hour><period> therefore it is <time>

### ¿Is 4:38pm a valid time?
```

### Is 7:38pm a valid time?
```
7:38pm

7 is a <positive-digit> and therefore a <12-hour-hour>

: is a symbol

3 is a <left-minute>

8 is a <positive-digit> and therefore a <digit> 

38 is a <left-minute><digit> and therefore a <minutes>

pm is a <period>

7:38pm is a <12-hour-hour>:<minutes><period> and therefore a <12-hour-time> and therefore a <time>
```

### Is 23:42 a valid time?

```
23:42

23 is a <24-hour-value>

: is a symbol

4 is a <left-minute>

2 is a <positive-digit> and therefore a <digit>

42 is a <left-minute><digit> and therefore a <minutes>

23:34 is a <24-hour-value>:<minutes> and therefore a <24-hour-time> and therefore a <time>
```






















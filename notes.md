# Notes

This will contain notes from the book.

## Lexer

Code starts as **source code**. The **lexer** turs the code into **tokens**, then a **parser** creates an **abstract syntax tree** that can then be evaluated.

As an example, the source code is:

```
let x = 5 + 5
```

The lexer creates:

```
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]
```

**NOTE:** A professional program would attach extra data to each token, such as the line and column number for better error messages.

A lexer does not care if syntax is valid or not - only that it can recognize the characters correctly to tokenize them.

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

## Parser

A **Parser** takes input data and builds a data structure - often some form of tree - and checks for correct syntax in the process.

Generally things are parsed into an **Abstract Syntax Tree** (AST), which discards several components of the language that are only used to guide the parser. 

Parser Generators do exist, which can take **context-free grammer** (CFG) which describes the language - typically in **Backus-Naur Form** or **Extended Backus-Naur Form** and creates a parser.

### Strategies

Two strategies exist for parsing: top-down or bottom-up. Examples of top-down parsers include "recursive descent", "early", and "predictive". Monkey uses recursive descent, or specifically a Pratt Parser.

Terminology:

- **Statements** do not produce a value
- **Expressions** do produce a value
- **prefix operator**: is "in front of" its operand (`--5`)
- **postfix operator**: is after is operand (`foobar++`)
- **infix operator**: is between its operands (`5 * 8`)
- **operator precendence**: order of operations; which operators take priority
- 
# InterpreterInGo

Working through the book *Writing an Interpreter In Go* by Thorsten Ball. In addition to practing Go and learning about writing an interpreter, I also want to practice using some other tools in Github.

## Monkey Programming Language

Monkey has the following features:

- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hashh data structure

Binding names:

```
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
let myArray = [1, 2, 3, 4, 5];
let thorsten = {"name": "Thorsten", "age": 28};
let add = fn(a, b) { return a + b; };
```

Accessing elements:

```
myArray[0] // => 1
thorsten["name"] // => "Thorsten"
```

Calling a function:

```
add(1, 2);
```
Monkey only supports ASCII characters, it has no Unicode support.

## Lexer

The lexer needs a few things:

1. Tokens
2. A way to create a new lexer
3. A way to get the next token

That is basically all that is needed from the lexer.
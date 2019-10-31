`go build` is used to build a binary. Without any params, it just outputs binary to the file with the same name.

`go run` is interpreter mode. Need to test interpreter vs compiler performance.

`main` package declared in each file present inside this directory, so be sure to pass a specific file to compiler/interpreter.

As it turns out, Golang has a strict scope inheritance for almost any language construct that works with a code block (for, if, etc.). All variables from parent scope are implicitely derived in child scope, but also can be redeclared. _Variables and consts cannot be redeclared in the same scope!_

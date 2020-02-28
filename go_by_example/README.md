`go build` is used to build a binary. Without any params, it just outputs binary to the file with the same name.

`go run` is interpreter mode. Need to test interpreter vs compiler performance. Flag `-race` can be used to ask
interpreter to check for any race conditions in code while accessing internal data structures.

`main` package declared in each file present inside this directory, so be sure to pass a specific file to
compiler/interpreter.

As it turns out, Golang has a strict scope inheritance for almost any language construct that works with a code block
(for, if, etc.). All variables from parent scope are implicitely derived in child scope, but also can be redeclared.
_Variables and consts cannot be redeclared in the same scope!_

One of the most detailed Golang documentations is GitHub wiki one. It has descriptive articles on language pitfalls like
[this one](https://github.com/golang/go/wiki/SliceTricks)

Arrays in Golang have a strict structure and capacity, which is different from the C lang. But, slices are something
similar to C arrays - it's just a pointer to underlying array beginning, also length and capacity.   
To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice
into it. This technique is how dynamic array implementations from other languages work behind the scenes.
Great article to understand slices: https://blog.golang.org/go-slices-usage-and-internals

Golang support declaring exported/unexported fields or functions by making the first name letter uppercase (exported) or
lowercase (unexported). It works similar to public/private methods and fields in OOP languages, but here it's just
a convention, which is used by many libraries, like JSON (e.g. JSON lib will not marshal unexported fields in structs).
package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var p = fmt.Println

func main() {
	p("Contains:    ", s.Contains("test", "es"))
	p("Count:       ", s.Count("test", "t"))
	p("HasPrefix:   ", s.HasPrefix("test", "te"))
	p("HasSuffix:   ", s.HasSuffix("test", "st"))
	p("Index:       ", s.Index("test", "e"))
	p("Join:        ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:      ", s.Repeat("a", 5))
	p("Replace:     ", s.Replace("foo", "o", "0", -1))
	p("Replace:     ", s.Replace("foo", "o", "0", 1))
	p("Split:       ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:     ", s.ToLower("TEST"))
	p("ToUpper:     ", s.ToUpper("test"))
	p("ContainsAny :", s.ContainsAny("test", "e"))
	p("ContainsRune:", s.ContainsRune("test", 's'))
	p("FieldsFunc  :", s.FieldsFunc("1foo09923bar9893", func(r rune) bool {
		return unicode.IsNumber(r)
	}))

	p("Len:  ", len("hello"))
	p("Char: ", fmt.Sprintf("%T", "hello"[1]))
	p("Rune: ", fmt.Sprintf("%T", 'e'))
}
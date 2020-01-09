package main

import s "strings" //note that you can rename an import for ease of use
import "fmt"

var out = fmt.Println //please don't do this

func main() {

	//(http://golang.org/pkg/strings/)

	out("Contains:  ", s.Contains("test", "es"))
	out("Count:     ", s.Count("test", "t"))
	out("HasPrefix: ", s.HasPrefix("test", "te"))
	out("HasSuffix: ", s.HasSuffix("test", "st"))
	out("Index:     ", s.Index("test", "e"))
	out("Join:      ", s.Join([]string{"a", "b"}, "-"))
	out("Repeat:    ", s.Repeat("a", 5))
	out("Replace:   ", s.Replace("foo", "o", "0", -1))
	out("Replace:   ", s.Replace("foo", "o", "0", 1))
	out("Split:     ", s.Split("a-b-c-d-e", "-"))
	out("ToLower:   ", s.ToLower("TEST"))
	out("ToUpper:   ", s.ToUpper("test"))
	out("Title:     ", s.Title("harry potter and the goblet of fire"))
	out()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	out("Len: ", len("hello"))
	out("Char:", "hello"[3:])
}

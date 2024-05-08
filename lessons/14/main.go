package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.\n", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s\n", b.Title, b.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Learning to use Interfaces in Go",
		Author: "Ivonei Ramos",
	}
	Print(a)

	b := Book{
		Title:  "GoLang chronicles",
		Author: "Ivonei Gomes",
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Printf(s.String())
}

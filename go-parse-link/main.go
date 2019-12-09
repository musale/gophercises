package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/musale/gophercises/go-parse-link/parser"
)

func main() {
	htmlString := `
	<a href="/dog">
		<span>Something in a span</span>
		Text not in a span
		<b>Bold text!</b>
	</a>
	`
	r := strings.NewReader(htmlString)
	links, err := parser.Parse(r)
	if err != nil {
		log.Printf("An error occured parsing HTML: %v", err)
	}
	fmt.Println(links)
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	lang := flag.String("lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(*lang))
	fmt.Println(greeting)
}


type language string

 var phrasebook = map[language]string{
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"el": "Χαίρετε Κόσμε",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

func greet(l language) string {

	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
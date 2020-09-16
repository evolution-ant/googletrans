package main

import (
	"fmt"

	"github.com/mind1949/googletrans"
	// "golang.org/x/text/language"
)

func main() {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			params := googletrans.TranslateParams{
				Src:  "auto",
				Dest: "zh",
				Text: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. ",
			}
			translated, err := googletrans.Translate(params)
			if err != nil {
				translated, err = googletrans.Translate(params)
			}
			if err != nil {
				panic(err)
			}
			// fmt.Printf("text: %q \npronunciation: %q", translated.Text, translated.Pronunciation)
			fmt.Printf("text: %q \n index: %d", translated.Text, i)
		}(i)
	}
	for {

	}
}

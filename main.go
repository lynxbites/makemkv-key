package main

import (
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {

	file, err := os.Create("makemkv-key.txt")
	if err != nil {
		panic(err)
	}

	collector := colly.NewCollector()
	key := "null"
	collector.OnHTML("code", func(e *colly.HTMLElement) {
		key = e.Text
	})

	err = collector.Visit("https://forum.makemkv.com/forum/viewtopic.php?f=5&t=1053")
	if err != nil {
		panic(err)
	}
	file.WriteString(key)
}

package main

import (
	"fmt"
	"os"

	colly "github.com/gocolly/colly"
)

func main() {

	new_file, err := os.Create("Mined.txt")
	if err != nil {
		fmt.Printf("An error occured %s", err)
	}

	defer new_file.Close()

	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("There was an error %s", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)
	})

	c.OnHTML("div", func(k *colly.HTMLElement) {
		new_file.WriteString(k.Text)
		// fmt.Println(k.Text)
	})

	c.Visit("https://news.google.com/topics/CAAqJggKIiBDQkFTRWdvSUwyMHZNRFp1ZEdvU0FtVnVHZ0pWVXlnQVAB?hl=en-US&gl=US&ceid=US%3Aen")
}

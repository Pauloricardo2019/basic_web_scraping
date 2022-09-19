package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func WriteFile(data, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(data)
}

func main() {
	url := "https://techcrunch.com/"

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	river, err := doc.Find("div.river").Html()
	if err != nil {
		fmt.Println(err)
	}

	WriteFile(river, "index.html")
	fmt.Println("Arquivo gerado com sucesso")

}

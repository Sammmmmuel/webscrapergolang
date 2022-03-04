package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	// "os"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID int `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement){
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("couldnt get id")
		}

		factDesc := element.Text

		fact := Fact{
			ID: factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/area-51-facts")

	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data,""," ")
	if err != nil {
		log.Println("cant create json file")
		return
	}

	_ = ioutil.WriteFile("Area51facts.json", file, 0644)
}
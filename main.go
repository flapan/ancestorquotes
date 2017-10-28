package main

import (
	"fmt"
	"math/rand"
	"os"
	"github.com/urfave/cli"
	"encoding/json"
	"time"
	"githubcom/bruno-chavez/ancestorquotes_test/quotes"
)

type Quotes struct {
	Quote string `json:"quote"`
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	quoteSlice := make([]Quotes, 0)

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons"
	app.Version = "1.0"
	app.Action = func(c *cli.Context) error {

		json.Unmarshal(quotes.Q(), &quoteSlice)
		selectedQuote := quoteSlice[rand.Intn(len(quoteSlice))]
		fmt.Printf("%v", selectedQuote.Quote)
		return nil
	}
	app.Run(os.Args)
}
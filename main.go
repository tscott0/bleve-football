package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

const indexPath = "football.bleve"

func main() {
	index, err := bleve.Open(indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		// open a new index
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
		if err != nil {
			fmt.Println(err)
			return
		}

		season := readSeason()

		for _, r := range season.Rounds {
			for _, m := range r.Matches {
				id := fmt.Sprintf("%s vs. %s %s",
					m.Team1.Code, m.Team2.Code, m.Date)

				// index some data
				index.Index(id, m)
			}
		}
	}
	// search for some text
	query := bleve.NewMatchQuery("LEI ARS 2016-08")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}

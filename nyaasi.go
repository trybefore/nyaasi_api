package nyaasi

import (
	"github.com/gocolly/colly"
)

// API is where you access the nyaasi api from
type API struct {
	collector *colly.Collector

	torrents <-chan Torrent
}

// Search search nyaa.si for your torrent(s).
// category is the result of using CategoryByName("your category").Format(), or typing it manually.
// query are the keywords you'd like to use.
// filter is the filter you'd like to use.
// torrents will be sent on the channel specified in NewAPI.
func (a *API) Search(query, category, filter string) error {
	return a.SearchPage(query, category, filter, 1)
}

// SearchPage identical to Search, but includes pagination
func (a *API) SearchPage(query, category, filter string, page int) error {
	url, err := buildURLWithPage(query, category, filter, page)
	if err != nil {
		return err
	}

	return a.collector.Visit(url.String())
}

func setupCollector(collector *colly.Collector) <-chan Torrent {

	torrents := make(chan Torrent)

	collector.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("tr", func(i int, row *colly.HTMLElement) bool {

			torrent := Torrent{}

			torrent.Type = parseType(row.Attr("class"))

			row.ForEachWithBreak("td", func(i int, cell *colly.HTMLElement) bool {

				parseCell(cell, &torrent)

				return true
			})

			torrents <- torrent
			return true
		})

	})

	return torrents
}

// NewAPI creates a new API.
// The channel returned is where all torrents will be sent
func NewAPI() (*API, <-chan Torrent) {
	collector := colly.NewCollector(colly.AllowedDomains("nyaa.si"))

	torrents := setupCollector(collector)

	a := &API{
		collector: collector,
		torrents:  torrents,
	}

	return a, torrents
}

package nyaasi

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const (
	baseURL = "https://nyaa.si"
)

func buildURL(query, category, filter string) (*url.URL, error) {
	return buildURLWithPage(query, category, filter, 1)
}

func buildURLWithPage(query, category, filter string, page int) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s/?f=%s&c=%s&q=%s&p=%d", baseURL, filter, category, url.QueryEscape(query), page))
}

func parseType(class string) string {
	switch class {
	case "success":
		return "trusted"
	case "danger":
		return "remake"
	default:
		fallthrough
	case "default":
		return "default"
	}
}

func parseCell(cell *colly.HTMLElement, torrent *Torrent) {
	switch cell.Index {
	case 0: // Category
		torrent.Category = ParseCategory(cell.ChildAttr("a[href]", "title"))
	case 1: // Title
		torrent.Title = cell.ChildText("a[href]:not(.comments)")
	case 2: // Torrent, Magnet link
		attrs := cell.ChildAttrs("a", "href")

		for _, a := range attrs {
			if strings.HasPrefix(a, "/download") {
				torrent.TorrentLink = baseURL + a
			} else if strings.HasPrefix(a, "magnet:") {
				torrent.MagnetLink = a
			}
		}
	case 3:
		torrent.Size = cell.Text
	case 4:
		// 2019-11-02 19:57
		timestamp, err := time.Parse("2006-01-02 15:04", cell.Text)

		if err != nil {
			timestamp = time.Now()
		}

		torrent.Timestamp = timestamp
	case 5:
		i, _ := strconv.Atoi(cell.Text)
		torrent.Seeders = i
	case 6:
		i, _ := strconv.Atoi(cell.Text)
		torrent.Leechers = i
	case 7:
		i, _ := strconv.Atoi(cell.Text)
		torrent.Completed = i
	}
}

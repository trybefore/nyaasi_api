package nyaasi

import (
	"testing"
	"time"
)

func TestAPISearch(t *testing.T) {

	nyaa, torrents := NewAPI()

	go func(t *testing.T) {
		if err := nyaa.Search("Attack on Titan", GetCategoryByName("Anime").FormatWithSubCategoryName("English-translated"), GetFilterByName("Trusted only").ID); err != nil {
			t.Fatal(err)
		}

	}(t)

	ticker := time.NewTicker(time.Second * 1)

outer:
	for {
		select {
		case <-ticker.C:
			break outer
		case torrent, ok := <-torrents:
			if !ok {
				break outer
			}
			t.Logf("[%v] %s", torrent.Timestamp, torrent.Title)

		}
	}

	ticker.Stop()

}

func TestAPISearchPage(t *testing.T) {

	nyaa, torrents := NewAPI()

	go func(t *testing.T) {
		for i := 0; i < 5; i++ {
			if err := nyaa.SearchPage(
				"Attack on Titan",
				GetCategoryByName("Anime").FormatWithSubCategoryName("English-translated"),
				GetFilterByName("Trusted only").ID,
				i,
			); err != nil {
				t.Fatal(err)
			}
		}
	}(t)

	ticker := time.NewTicker(time.Second * 1)

outer:
	for {
		select {
		case <-ticker.C:
			break outer
		case torrent, ok := <-torrents:
			if !ok {
				break outer
			}
			t.Logf("[%v] %s", torrent.Timestamp, torrent.Title)

		}
	}

	ticker.Stop()

}

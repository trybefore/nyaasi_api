package nyaasi

import "testing"

func TestAPISearch(t *testing.T) {

	nyaa, torrents := NewAPI()

	go func(t *testing.T) {
		if err := nyaa.Search("Attack on Titan", GetCategoryByName("Anime").FormatWithSubCategoryName("English-translated"), GetFilterByName("Trusted only").ID); err != nil {
			t.Fatal(err)
		}

	}(t)

outer:
	for {
		select {
		case torrent, ok := <-torrents:
			if !ok {
				break outer
			}
			t.Logf("[%v] %s", torrent.Timestamp, torrent.Title)

		}
	}

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

	for {
		select {
		case torrent := <-torrents:
			t.Logf("[%v] %s", torrent.Timestamp, torrent.Title)
		}
	}

}

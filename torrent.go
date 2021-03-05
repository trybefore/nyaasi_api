package nyaasi

import (
	"encoding/json"
	"time"
)

// Torrent represents a torrent entry on the nyaa.si website
type Torrent struct {
	// TODO

	Type     string    // default, trusted, remake
	Category *Category // main category
	Title    string
	Size     string

	MagnetLink  string // both MagnetLink and TorrentLink CAN be empty
	TorrentLink string

	Timestamp time.Time

	Seeders   int
	Leechers  int
	Completed int // Completed downloads
}

func (t Torrent) String() string {
	bs, _ := json.MarshalIndent(t, " ", "  ")

	return string(bs)
}

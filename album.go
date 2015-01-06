package spotify

// SimpleAlbum contains basic data about an album.
type SimpleAlbum struct {
	// The name of the album.
	Name string `json:"name"`
	// The type of the album: one of "album",
	// "single", or "compilation".
	AlbumType string `json:"album_type"`
	// The SpotifyID for the album.
	ID ID `json:"id"`
	// The SpotifyURI for the album.
	URI URI `json:"uri"`
	// The markets in which the album is available,
	// identified using ISO 3166-1 alpha-2 country
	// codes.  Note that al album is considered
	// available in a market when at least 1 of its
	// tracks is available in that market.
	AvailableMarkets []string `json:"available_markets"`
	// A link to the Web API enpoint providing full
	// details of the album.
	Endpoint string `json:"href"`
	// The cover art for the album in various sizes,
	// widest first.
	Images []Image `json:"images"`
	// Known external URLs for this album.
	ExternalURLs ExternalURL `json:"external_urls"`
}

// Copyright contains the copyright statement
// associated with an album.
type Copyright struct {
	// The copyright text for the album.
	Text string `json:"text"`
	// The type of copyright.
	Type string `json:"type"`
}

// FullAlbum provides extra album data in addition
// to the data provided by SimpleAlbum.
type FullAlbum struct {
	SimpleAlbum
	// The artists of the album.
	Artists []SimpleArtist `json:"artists"`
	// The copyright statements of the album.
	Copyrights []Copyright `json:"copyrights"`
	//A list of genres used to classify the album.
	// For example, "Prog Rock" or "Post-Grunge".
	// If not yet classified, the slice is empty.
	Genres []string `json:"genres"`
	// The popularity of the album.  The value will
	// be between 0 and 100, with 100 being the most
	// popular.  Popularity of an album is calculated
	// from the popularify of the album's individual
	// tracks.
	Popularity int `json:"popularity"`
	// The date the album was first released.  For
	// example, "1981-12-15".  Depending on the
	// ReleaseDatePrecision, it might be shown as
	// "1981" or "1981-12".
	ReleaseDate string `json:"release_date"` // TODO change to Timestamp
	// The precision with which ReleaseDate value
	// is known: "year", "month", or "day"
	ReleaseDatePrecision string `json:"release_date_precision"`
	// The tracks of the album.
	Tracks []SimpleTrack `json:"tracks"`
	// Known external IDs for the album.
	ExternalIDs ExternalID `json:"external_ids"`
}

func (a *SimpleAlbum) String() string {
	return "SimpleAlbum: " + a.Name
}

func (a *FullAlbum) String() string {
	return "FullAlbum: " + a.Name
}

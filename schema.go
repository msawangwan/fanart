package fanart

// ImageResponse represents a set of common fields used in movie response payloads.
type ImageResponse struct {
	ID       string `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	Lang     string `json:"lang,omitempty"`
	Likes    string `json:"likes,omitempty"`
	Disc     string `json:"disc,omitempty"`
	DiscType string `json:"disc_type,omitempty"`
}

// MovieFields is the set of fields common to movie response payloads.
type MovieFields struct {
	Name   string `json:"name,omitempty"`
	TMDbID string `json:"tmdb_id,omitempty"`
	IMDbID string `json:"imdb_id,omitempty"`
}

// MovieResponse is the response returned when submitting a movie query.
type MovieResponse struct {
	*MovieFields

	HDMovieLogo     []ImageResponse `json:"hdmovielogo,omitempty"`
	MovieDisc       []ImageResponse `json:"moviedisc,omitempty"`
	MovieLogo       []ImageResponse `json:"movielogo,omitempty"`
	MoviePoster     []ImageResponse `json:"movieposter,omitempty"`
	HDMovieClearArt []ImageResponse `json:"hdmovieclearart,omitempty"`
	MovieArt        []ImageResponse `json:"movieart,omitempty"`
	MovieBackground []ImageResponse `json:"moviebackground,omitempty"`
	MovieBanner     []ImageResponse `json:"moviebanner,omitempty"`
	MovieThumb      []ImageResponse `json:"moviethumb,omitempty"`
}

// NotFoundResponse represents a resource not found error.
type NotFoundResponse struct {
	ErrorMessage *string
	Status       *string
}

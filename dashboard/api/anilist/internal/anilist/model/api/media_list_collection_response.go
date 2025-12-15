package api

type List struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Media    Media  `json:"media"`
	Progress *int   `json:"progress"`
	Status   string `json:"status"` // Personal Status
}

type Media struct {
	ID    int `json:"id"`
	Title struct {
		Romaji  string `json:"romaji"`
		English string `json:"english"`
	} `json:"title"`
	Episodes          *int `json:"episodes"`
	NextAiringEpisode *struct {
		ID       int    `json:"id"`
		Episode  int    `json:"episode"`
		AiringAt *int64 `json:"airingAt"`
	} `json:"nextAiringEpisode"`
	Status string `json:"status"` // Anime Status
}

type MediaListCollectionResponse struct {
	Data struct {
		MediaListCollection struct {
			Lists []List `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}

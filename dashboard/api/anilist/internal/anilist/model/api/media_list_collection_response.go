package api

type MediaListCollectionResponse struct {
	Data struct {
		MediaListCollection struct {
			Lists []struct {
				Entries []struct {
					Media struct {
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
					} `json:"media"`
					Progress *int   `json:"progress"`
					Status   string `json:"status"` // Personal Status
				} `json:"entries"`
			} `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}

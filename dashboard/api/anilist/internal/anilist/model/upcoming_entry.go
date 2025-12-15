package model

import "fmt"

type UpcomingEntry struct {
	Title         string  `json:"title"`
	NextEpisode   int     `json:"nextEpisode"`
	TotalEpisodes *int    `json:"totalEpisodes"`
	URL           string  `json:"url"`
	WatchURL      string  `json:"watchUrl"`
	Watched       int     `json:"watched"`
	AiringAt      *string `json:"airingAt"`
}

func UpcomingEntryQuery(username string) string {
	return fmt.Sprintf(`
	query {
	  MediaListCollection(
		userName: "%s",
		type: ANIME,
		status_in: [CURRENT, PLANNING]
	  ) {
		lists {
		  entries {
			media {
			  id
			  title { romaji english }
			  nextAiringEpisode { id episode airingAt }
			}
			progress
		  }
		}
	  }
	}`, username)
}

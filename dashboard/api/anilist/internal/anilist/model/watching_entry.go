package model

import "fmt"

type WatchingEntry struct {
	Title         string `json:"title"`
	NextEpisode   int    `json:"nextEpisode"`
	TotalEpisodes *int   `json:"totalEpisodes"`
	URL           string `json:"url"`
	WatchURL      string `json:"watchUrl"`
	Watched       int    `json:"watched"`
}

func WatchingEntryQuery(username string) string {
	return fmt.Sprintf(`
	query {
	  MediaListCollection(
		userName: "%s",
		type: ANIME,
		status: CURRENT
	  ) {
		lists {
		  entries {
			media {
			  id
			  title { romaji english }
			  episodes
			}
			progress
		  }
		}
	  }
	}`, username)
}

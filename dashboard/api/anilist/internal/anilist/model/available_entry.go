package model

import "fmt"

type AvailableEntry struct {
	Title         string `json:"title"`
	NextEpisode   int    `json:"nextEpisode"`
	TotalEpisodes *int   `json:"totalEpisodes"`
	URL           string `json:"url"`
	WatchURL      string `json:"watchUrl"`
	Watched       int    `json:"watched"`
}

func AvailableEntryQuery(username string) string {
	return fmt.Sprintf(`
	query {
	  MediaListCollection(
		userName: "%s",
		type: ANIME,
		status: PLANNING
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

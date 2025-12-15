package model

import (
	"fmt"
	"tomtruyen/anilist/internal/anilist/model/api"
)

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

func FormatAvailableEntry(media api.Media, progress int, anilistAnimeUrl, anikaiBrowseUrl string) *AvailableEntry {
	title := media.Title.English
	if title == "" {
		title = media.Title.Romaji
	}

	nextEpisode := progress + 1

	// Skip if we already finished the Anime and the Anime is Finished
	if media.Status == "FINISHED" && progress >= *media.Episodes {
		return nil
	}

	// Skip if Anime is not yet released
	if media.Status == "NOT_YET_RELEASED" {
		return nil
	}

	return &AvailableEntry{
		Title:         title,
		NextEpisode:   nextEpisode,
		TotalEpisodes: media.Episodes,
		URL:           fmt.Sprintf(anilistAnimeUrl, media.ID),
		WatchURL:      fmt.Sprintf(anikaiBrowseUrl, title),
		Watched:       progress,
	}
}

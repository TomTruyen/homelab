package model

import (
	"fmt"
	"tomtruyen/anilist/internal/anilist/model/api"
)

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

func FormatWatchingEntry(media api.Media, progress int, anilistAnimeUrl, anikaiBrowseUrl string) *WatchingEntry {
	title := media.Title.English
	if title == "" {
		title = media.Title.Romaji
	}

	nextEpisode := progress + 1

	return &WatchingEntry{
		Title:         title,
		NextEpisode:   nextEpisode,
		TotalEpisodes: media.Episodes,
		URL:           fmt.Sprintf(anilistAnimeUrl, media.ID),
		WatchURL:      fmt.Sprintf(anikaiBrowseUrl, title),
		Watched:       progress,
	}
}

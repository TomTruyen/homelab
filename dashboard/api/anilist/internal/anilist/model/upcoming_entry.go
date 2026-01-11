package model

import (
	"fmt"
	"tomtruyen/anilist/internal/anilist/model/api"
	"tomtruyen/anilist/internal/anilist/util"
)

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
			  status
			}
			progress
		  }
		}
	  }
	}`, username)
}

func FormatUpcomingEntry(media api.Media, progress int, anilistAnimeUrl, anikaiBrowseUrl string) *UpcomingEntry {
	title := media.Title.English
	if title == "" {
		title = media.Title.Romaji
	}

	var nextEpisode int
	var airingAt *string
	if media.NextAiringEpisode == nil {
		if media.Status != "NOT_YET_RELEASED" {
			return nil
		}
		nextEpisode = 1 // Not yet released, first episode will be episode 1
	} else {
		nextEpisode = media.NextAiringEpisode.Episode
		formatted := util.FormatAiringAt(media.NextAiringEpisode.AiringAt)
		if formatted != nil {
			airingAt = formatted
		}
	}

	if airingAt == nil {
		n := "N/A"
		airingAt = &n
	}

	return &UpcomingEntry{
		Title:         title,
		NextEpisode:   nextEpisode,
		TotalEpisodes: media.Episodes,
		URL:           fmt.Sprintf(anilistAnimeUrl, media.ID),
		WatchURL:      fmt.Sprintf(anikaiBrowseUrl, title),
		Watched:       progress,
		AiringAt:      airingAt,
	}
}

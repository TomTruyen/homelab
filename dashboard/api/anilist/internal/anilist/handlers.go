package anilist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"tomtruyen/anilist/internal/anilist/model"
	"tomtruyen/anilist/internal/anilist/model/api"
	"tomtruyen/anilist/internal/anilist/util"
)

const anilistAnimeUrl = "https://anilist.co/anime/%d"
const anikaiBrowseUrl = "https://anikai.to/browser?keyword=%s"

func (s *Service) GetUpcomingAnimes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Prepare Query
	query := model.UpcomingEntryQuery(s.username)

	// Perform GraphQL Request
	var resp api.MediaListCollectionResponse
	if err := s.doGraphQL(ctx, query, &resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Flatten and Modify Entries
	upcoming := make([]model.UpcomingEntry, 0) // Initialize empty array
	for _, list := range resp.Data.MediaListCollection.Lists {
		for _, entry := range list.Entries {
			media := entry.Media

			progress := 0
			if entry.Progress != nil {
				progress = *entry.Progress
			}

			title := media.Title.English
			if title == "" {
				title = media.Title.Romaji
			}

			nextEpisode := progress + 1

			var airingAt *string

			if media.NextAiringEpisode != nil {
				formatted := util.FormatAiringAt(media.NextAiringEpisode.AiringAt)
				if formatted != nil {
					airingAt = formatted
				}
			}

			if airingAt == nil {
				n := "N/A"
				airingAt = &n
			}

			upcoming = append(upcoming, model.UpcomingEntry{
				Title:         title,
				NextEpisode:   nextEpisode,
				TotalEpisodes: media.Episodes,
				URL:           fmt.Sprintf(anilistAnimeUrl, media.ID),
				WatchURL:      fmt.Sprintf(anikaiBrowseUrl, title),
				Watched:       progress,
				AiringAt:      airingAt,
			})
		}
	}

	sort.SliceStable(upcoming, func(i, j int) bool {
		a := upcoming[i]
		b := upcoming[j]

		if a.Watched != b.Watched {
			return a.Watched > b.Watched
		}

		aUnknown := a.AiringAt == nil
		bUnknown := b.AiringAt == nil

		switch {
		case aUnknown && bUnknown:
			return false
		case aUnknown:
			return false
		case bUnknown:
			return true
		default:
			return *a.AiringAt < *b.AiringAt
		}
	})

	json.NewEncoder(w).Encode(upcoming)
}

// Available Animes
func (s *Service) FetchAvailable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Prepare Query
	query := model.AvailableEntryQuery(s.username)

	// Perform GraphQL Request
	var resp api.MediaListCollectionResponse
	if err := s.doGraphQL(ctx, query, &resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Flatten and Modify Entries
	available := make([]model.AvailableEntry, 0) // Initialize empty array
	for _, list := range resp.Data.MediaListCollection.Lists {
		for _, entry := range list.Entries {
			media := entry.Media

			progress := 0
			if entry.Progress != nil {
				progress = *entry.Progress
			}

			title := media.Title.English
			if title == "" {
				title = media.Title.Romaji
			}

			nextEpisode := progress + 1

			// Skip if we already finished the Anime and the Anime is Finished
			if media.Status == "FINISHED" && *entry.Progress >= *media.Episodes {
				continue
			}

			// Skip if Anime is not yet released
			if media.Status == "NOT_YET_RELEASED" {
				continue
			}

			available = append(available, model.AvailableEntry{
				Title:         title,
				NextEpisode:   nextEpisode,
				TotalEpisodes: media.Episodes,
				URL:           fmt.Sprintf(anilistAnimeUrl, media.ID),
				WatchURL:      fmt.Sprintf(anikaiBrowseUrl, title),
				Watched:       progress,
			})
		}
	}

	sort.SliceStable(available, func(i, j int) bool {
		return available[i].Watched > available[j].Watched
	})

	json.NewEncoder(w).Encode(available)
}

// Currently Watching Animes
func (s *Service) FetchWatching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Prepare Query
	query := model.WatchingEntryQuery(s.username)

	// Perform GraphQL Request
	var resp api.MediaListCollectionResponse
	if err := s.doGraphQL(ctx, query, &resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Flatten and Modify Entries
	watching := make([]model.WatchingEntry, 0) // Initialize empty array
	for _, list := range resp.Data.MediaListCollection.Lists {
		for _, entry := range list.Entries {
			media := entry.Media

			progress := 0
			if entry.Progress != nil {
				progress = *entry.Progress
			}

			title := media.Title.English
			if title == "" {
				title = media.Title.Romaji
			}

			nextEpisode := progress + 1

			watching = append(watching, model.WatchingEntry{
				Title:         title,
				NextEpisode:   nextEpisode,
				TotalEpisodes: media.Episodes,
				URL:           fmt.Sprintf(anilistAnimeUrl, media.ID),
				WatchURL:      fmt.Sprintf(anikaiBrowseUrl, title),
				Watched:       progress,
			})
		}
	}

	sort.SliceStable(watching, func(i, j int) bool {
		return watching[i].Watched > watching[j].Watched
	})

	json.NewEncoder(w).Encode(watching)
}

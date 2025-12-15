package anilist

import (
	"encoding/json"
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
	upcoming := util.FlattenMediaEntries(resp.Data.MediaListCollection.Lists, func(media api.Media, progress int) *model.UpcomingEntry {
		return model.FormatUpcomingEntry(media, progress, anilistAnimeUrl, anikaiBrowseUrl)
	})

	// Sort
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
	available := util.FlattenMediaEntries(resp.Data.MediaListCollection.Lists, func(media api.Media, progress int) *model.AvailableEntry {
		return model.FormatAvailableEntry(media, progress, anilistAnimeUrl, anikaiBrowseUrl)
	})

	// Sort
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
	watching := util.FlattenMediaEntries(resp.Data.MediaListCollection.Lists, func(media api.Media, progress int) *model.WatchingEntry {
		return model.FormatWatchingEntry(media, progress, anilistAnimeUrl, anikaiBrowseUrl)
	})

	sort.SliceStable(watching, func(i, j int) bool {
		return watching[i].Watched > watching[j].Watched
	})

	// Sort
	json.NewEncoder(w).Encode(watching)
}

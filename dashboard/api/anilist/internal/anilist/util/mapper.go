package util

import "tomtruyen/anilist/internal/anilist/model/api"

func FlattenMediaEntries[T any](lists []api.List, format func(media api.Media, progress int) *T) []T {
	result := make([]T, 0)

	for _, list := range lists {
		for _, entry := range list.Entries {
			progress := 0
			if entry.Progress != nil {
				progress = *entry.Progress
			}

			formatted := format(entry.Media, progress)
			if formatted != nil {
				result = append(result, *formatted) // ok because formatted is *T
			}
		}
	}

	return result
}

package util

import "time"

func FormatAiringAt(airingAt *int64) *string {
	if airingAt == nil {
		return nil
	}

	t := time.Unix(*airingAt, 0)
	formatted := t.Format("Jan 2 03:04 PM")

	return &formatted
}

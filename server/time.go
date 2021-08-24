package server

import "time"

// return Tomorrow 00:00:00.000.
func Tomorrow() time.Time {
	t := time.Now().AddDate(0, 0, 1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

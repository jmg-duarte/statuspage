package internal

import "time"

type History map[time.Duration]map[string]string

func (h *History) AddEntry(time time.Duration, statuses map[string]string) {
	(*h)[time] = statuses
}

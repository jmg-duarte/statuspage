package internal

import (
	"time"
)

type History map[time.Time]map[string]string

func (h *History) AddEntry(time time.Time, statuses map[string]string) {
	(*h)[time] = statuses
}

type ServiceHistory map[string]History

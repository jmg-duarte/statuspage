package internal

import (
	"time"
)

type History map[time.Time]map[string]string

func (h *History) AddEntry(time time.Time, statuses map[string]string) {
	(*h)[time] = statuses
}

func (h History) CSV() [][]string {
	var currentRecord []string
	var fields []string
	records := make([][]string, len(h)+1)
	fieldsIdx := make(map[string]int)
	for _, statuses := range h {
		nFields := 1 + len(statuses)
		fields = make([]string, nFields)
		currentRecord = make([]string, nFields)
		fields[0] = "time_utc"
		i := 1
		for component := range statuses {
			fields[i] = component
			fieldsIdx[component] = i
			i++
		}
		records[0] = make([]string, nFields)
		for i, id := range fields {
			records[0][i] = id //strings.Replace(id, ",", "", -1)
		}
		break
	}

	currentRecordIdx := 1
	for t, statuses := range h {
		//currentRecord = make([]string, 1+len(statuses))
		currentRecord[0] = t.UTC().String()
		// fields = make([]string, 0, 1 + len(statuses))
		for service, status := range statuses {
			currentRecord[fieldsIdx[service]] = status
		}
		records[currentRecordIdx] = currentRecord
		currentRecordIdx++
	}
	return records
}

type ServiceHistory map[string]History

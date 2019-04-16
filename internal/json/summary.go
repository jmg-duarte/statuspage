package json

import "strings"

const SummaryJson = "summary.json"

type Summary struct {
	Page                  `json:"page"`
	Components            `json:"components"`
	Incidents             []string `json:"incidents"`
	ScheduledMaintenances []string `json:"scheduled_maintenances"`
	Status                `json:"status"`
}

func (s Summary) BriefStatus() string {
	return s.Status.Description
}

func (s Summary) FullStatus() string {
	var status strings.Builder
	for _, component := range s.Components {
		status.Write([]byte("\t"))
		status.Write([]byte(component.Name))
		status.Write([]byte(": "))
		status.Write([]byte(component.Status))
		status.Write([]byte("\n"))
	}
	return status.String()
}

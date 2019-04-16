package json

type Component struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Position           int    `json:"position"`
	Description        string `json:"description"`
	Showcase           bool   `json:"showcase"`
	GroupID            string `json:"group_id"`
	PageID             string `json:"page_id"`
	Group              bool   `json:"group"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded"`
}

type Components []Component

package dto

import "device-manager/model"

type SortingRequest struct {
	Order  string `json:"order"`
	Column string `json:"column"`
}

type RobotRequest struct {
	Sorting    SortingRequest         `json:"sorting"`
	Filters    map[string]interface{} `json:"filters"`
	PageNumber int                    `json:"page_number"`
	Limit      int                    `json:"limit"`
}

type RobotResponse struct {
	Robots      []model.Robot `json:"robots"`
	TotalCount  int64         `json:"total_count"`
	TotalPages  int           `json:"total_pages"`
	CurrentPage int           `json:"current_page"`
}

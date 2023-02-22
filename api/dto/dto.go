package dto

type PageData struct {
	Page      int         `json:"page"`
	Size      int         `json:"size"`
	Total     int64       `json:"total"`
	TotalPage float64     `json:"total_page"`
	Data      interface{} `json:"data"`
}

package dto

type NewsRequest struct{
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Author string`json:"author"`
}
type NewsResponse struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Author string`json:"author"`
	CreatedAt string `json:"created-at"`
	UpdatedAt string `json:"updated-at"`
}
type NewsListResponse struct{
	TotalCount int `json:"total-count"`
	TotalPages int 	`json:"total-pages"`
	Page int `json:"page"`
	Size int `json:"size"`
	HasMore bool `json:"has-more"`
	News []NewsResponse `json:"news"`
}

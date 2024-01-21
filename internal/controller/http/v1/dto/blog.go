package dto

type BlogRequest struct{
	Title string `json:"title"`
	Content string `json:"content"`
	Author string`json:"author"`
}
type BlogResponse struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Author string`json:"author"`
	CreatedAt string `json:"created-at"`
	UpdatedAt string `json:"updated-at"`
}
type BlogListResponse struct{
	TotalCount int `json:"total-count"`
	TotalPages int 	`json:"total-pages"`
	Page int `json:"page"`
	Size int `json:"size"`
	HasMore bool `json:"has-more"`
	Blog []BlogResponse `json:"news"`
}

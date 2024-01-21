package dto

type CommentRequest struct{
	BlogID string `json:"blog_id"`
	Body string `json:"body"`
}
type CommentResponse struct{
	ID string `json:"id"`
	BlogID string `json:"blog_id"`
	Body string `json:"body"`
}
type CommentListResponse struct{
	TotalCount int `json:"total-count"`
	TotalPages int 	`json:"total-pages"`
	Page int `json:"page"`
	Size int `json:"size"`
	HasMore bool `json:"has-more"`
	Comment []CommentResponse `json:"comment"`
}
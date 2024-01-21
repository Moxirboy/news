package models

type Comment struct{
	ID string
	BlogId string
	Body string
}
type CommentList struct {
	TotalCount int
	TotalPages int
	Page int
	Size int
	HasMore bool
	Comment []*Comment
}
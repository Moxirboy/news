package models

type Blog struct{
	ID string
	Title string
	Content string
	By
	At
}
type BlogList struct {
	TotalCount int
	TotalPages int
	Page int
	Size int
	HasMore bool
	Blog []*Blog
}
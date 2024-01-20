package models

type News struct{
	ID string
	Title string
	Content string
	Category string
	By
	At
}
type NewsList struct {
	TotalCount int
	TotalPages int
	Page int
	Size int
	HasMore bool
	News []*News
}
package dto
type CreateNewsRequest struct{
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Author string`json:"author"`
}
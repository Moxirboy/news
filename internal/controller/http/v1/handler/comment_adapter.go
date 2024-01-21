package http

import (
	"news/internal/controller/http/v1/dto"
	"news/internal/models"

)


func fromCreateCommentRequest(
	req *dto.CommentRequest,
	) *models.Comment {
	return &models.Comment{
		Body: req.Body,
		BlogId: req.BlogID,
		}}
	func ToCommentResponse(
		comment *models.Comment,
		) dto.CommentResponse {
		return dto.CommentResponse{
			ID: comment.ID,
			BlogID: comment.BlogId,
			Body: comment.Body,
			}
		}
		func ToCommentResponseList(
			comment *models.CommentList,
			) dto.CommentListResponse{
			lenth:=len(comment.Comment)
			Comment:=make([]dto.CommentResponse,lenth)
			for _,instance:=range comment.Comment{
				Comment=append(Comment,ToCommentResponse(instance))
			}
			return dto.CommentListResponse{
				TotalCount: 	comment.TotalCount,
				TotalPages: 	comment.TotalPages,
				Page: 	comment.Page,
				Size: 	comment.Size,
				HasMore: 	comment.HasMore,
				Comment: 	Comment,
				}
			}
				
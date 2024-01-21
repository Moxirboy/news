package http

import (
	"news/internal/controller/http/v1/dto"
	"news/internal/models"
	"time"
)


func fromCreateBlogRequest(
	req *dto.BlogRequest,
	) *models.Blog {
	return &models.Blog{
		Title:    req.Title,
		Content:  req.Content,
		By: models.By{
			CreatedBy: req.Author,
			},
			At: models.At{
			CreatedAt: time.Now(),
			},
			}
	}
	func ToBlogResponse(
		blog *models.Blog,
		) dto.BlogResponse {
		return dto.BlogResponse{
			ID:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			Author:    blog.CreatedBy,
			CreatedAt: blog.CreatedAt.Format("2006-01-02"),
			UpdatedAt: blog.UpdatedAt.Format("2006-01-02"),
			}
		}
		func ToBlogResponseList(
			blog *models.BlogList,
			) dto.BlogListResponse{
			lenth:=len(blog.Blog)
			Blog:=make([]dto.BlogResponse,lenth)
			for _,instance:=range blog.Blog{
				Blog=append(Blog,ToBlogResponse(instance))
			}
			return dto.BlogListResponse{
				TotalCount: 	blog.TotalCount,
				TotalPages: 	blog.TotalPages,
				Page: 	blog.Page,
				Size: 	blog.Size,
				HasMore: 	blog.HasMore,
				Blog: 	Blog,
				}
			}
			
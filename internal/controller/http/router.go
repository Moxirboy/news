package http

import (
	"github.com/gin-gonic/gin"
	v1 "news/internal/controller/http/v1/handler"
	"news/internal/service/usecase"

)

func SetUp(
	g *gin.Engine,
	uc usecase.IUseCase,

) {
	SetUpHandlerV1(
		g.Group("/api/v1"),
		uc,
		)

}
func SetUpHandlerV1(
	group *gin.RouterGroup,
	uc usecase.IUseCase,
) {
	v1.NewNewsHandler(
		group,
		uc.NewsUseCase(),
		)
	v1.NewBlogHandler(
		group,
		uc.BlogUseCase(),
		)
	v1.NewCommentHandler(
		group,
		uc.CommentUseCase(),
		)
}

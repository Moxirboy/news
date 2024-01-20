package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/internal/controller/http/v1/dto"
	"news/internal/service/usecase"
	"news/pkg/utils"
)
type newsHandler struct{
	NewsUseCase usecase.INewsUsecase
}
func NewNewsHandler(
	g *gin.RouterGroup,
	NewsUseCase usecase.INewsUsecase,
	){
	handler:=newsHandler{
		NewsUseCase: NewsUseCase,
	}
	news:=g.Group("/news")
	news.POST("",handler.create)
	
}
// @Router /v1/news [post]
func (h *newsHandler) create(c *gin.Context){
	req:=dto.CreateNewsRequest{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	params:=utils.Validate(req)
	if params!=nil{
		c.JSON(http.StatusInternalServerError,params)
	}
	
}
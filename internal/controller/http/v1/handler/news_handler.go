package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/internal/controller/http/v1/dto"
	"news/internal/service/usecase"
	"news/pkg/errors"
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
	news.POST("/create",handler.create)
	news.GET("/get/:id",handler.GetOneById)
	news.GET("/getall/:id",handler.GetAll)
	news.PUT("/update/:id",handler.update)
	news.DELETE("/delete/:id",handler.delete)
	
}
// @Router /v1/news [post]
func (h *newsHandler) create(c *gin.Context){
	req:=dto.NewsRequest{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	params:=utils.Validate(req)
	if params!=nil{
		c.JSON(http.StatusInternalServerError,params)
	}
	news:=fromCreateNewsRequest(&req)

	newsID,err:=h.NewsUseCase.Create(c,news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"newsID": newsID})
}
func (h *newsHandler) GetOneById(c *gin.Context){
	id := c.Param("id")
	news,err:=h.NewsUseCase.GetOneByID(c,id)
	if err!=nil{
		c.JSON(errors.ErrorResponse(err))
	}
	res:=ToNewsResponse(news)
	c.JSON(http.StatusOK, res)
}
func (h *newsHandler) GetAll(c *gin.Context){
	pq,err:=utils.GetPaginationFromCtx(c)
	if err!=nil{
		c.JSON(errors.ErrorResponse(err))
	}
	NewsList,err:=h.NewsUseCase.GetAll(c,*pq)
	newsList:=ToNewsResponseList(NewsList)
	c.JSON(http.StatusOK,newsList)
}
func (h *newsHandler) update(c *gin.Context){
	id := c.Param("id")
	req:=dto.NewsRequest{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	news:=fromCreateNewsRequest(&req)
	news.ID=id
	err:=h.NewsUseCase.Update(c,news)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
	})
}
func (h *newsHandler) delete(c *gin.Context){
	id:=c.Param("id")
	err:=h.NewsUseCase.Delete(c,id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		})
}
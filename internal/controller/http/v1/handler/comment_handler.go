package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/internal/controller/http/v1/dto"
	"news/internal/service/usecase"
	"news/pkg/errors"
	"news/pkg/utils"
)
type commentHandler struct{
	CommentUseCase usecase.ICommentUseCase
}
func NewCommentHandler(
	g *gin.RouterGroup,
	CommentUseCase usecase.ICommentUseCase,
	){
	handler:=commentHandler{
		CommentUseCase: CommentUseCase,
		}
		news:=g.Group("/comment")
		news.POST("/create",handler.create)
		news.GET("/getall/:id",handler.GetAll)
		news.PUT("/update/:id",handler.update)
		news.DELETE("/delete/:id",handler.delete)

	}
	// @Router /v1/news [post]
	func (h *commentHandler) create(c *gin.Context){
		req:=dto.CommentRequest{}
		if err:=c.ShouldBindJSON(&req);err!=nil{
			c.JSON(http.StatusInternalServerError,err)
		}
		params:=utils.Validate(req)
		if params!=nil{
			c.JSON(http.StatusInternalServerError,params)
		}
		comment:=fromCreateCommentRequest(&req)

		err:=h.CommentUseCase.Create(c,comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message":"success"})
	}
func (h *commentHandler) GetAll(c *gin.Context){
	pq,err:=utils.GetPaginationFromCtx(c)
	if err!=nil{
		c.JSON(errors.ErrorResponse(err))
	}
	CommentList,err:=h.CommentUseCase.GetAll(c,*pq)
	commentList:=ToCommentResponseList(CommentList)
	c.JSON(http.StatusOK,commentList)
}
func (h *commentHandler) update(c *gin.Context){
	id := c.Param("id")
	req:=dto.CommentRequest{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	comment:=fromCreateCommentRequest(&req)
	comment.ID=id
	err:=h.CommentUseCase.Update(c,comment)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		})
}
func (h *commentHandler) delete(c *gin.Context){
	id:=c.Param("id")
	err:=h.CommentUseCase.Delete(c,id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		})
}

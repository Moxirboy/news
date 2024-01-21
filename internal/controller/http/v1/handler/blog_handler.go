package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/internal/controller/http/v1/dto"
	"news/internal/service/usecase"
	"news/pkg/errors"
	"news/pkg/utils"
)
type blogHandler struct{
	BlogUseCase usecase.IBlogUsecase
}
func NewBlogHandler(
	g *gin.RouterGroup,
	BlogUseCase usecase.IBlogUsecase,
	){
	handler:=blogHandler{
		BlogUseCase: BlogUseCase,
		}
		news:=g.Group("/blog")
		news.POST("/create",handler.create)
		news.GET("/get/:id",handler.GetOneById)
		news.GET("/getall/:id",handler.GetAll)
		news.PUT("/update/:id",handler.update)
		news.DELETE("/delete/:id",handler.delete)

	}
// @Router /v1/news [post]
func (h *blogHandler) create(c *gin.Context){
	req:=dto.BlogRequest{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	params:=utils.Validate(req)
	if params!=nil{
		c.JSON(http.StatusInternalServerError,params)
	}
	blog:=fromCreateBlogRequest(&req)

	blogID,err:=h.BlogUseCase.Create(c,blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogID": blogID})
}
func (h *blogHandler) GetOneById(c *gin.Context){
	id := c.Param("id")
	blog,err:=h.BlogUseCase.GetOneByID(c,id)
	if err!=nil{
		c.JSON(errors.ErrorResponse(err))
	}
	res:=ToBlogResponse(blog)
	c.JSON(http.StatusOK, res)
}
func (h *blogHandler) GetAll(c *gin.Context){
	pq,err:=utils.GetPaginationFromCtx(c)
	if err!=nil{
		c.JSON(errors.ErrorResponse(err))
	}
	BlogList,err:=h.BlogUseCase.GetAll(c,*pq)
	blogList:=ToBlogResponseList(BlogList)
	c.JSON(http.StatusOK,blogList)
}
func (h *blogHandler) update(c *gin.Context){
	id := c.Param("id")
	req:=dto.BlogRequest{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	blog:=fromCreateBlogRequest(&req)
	blog.ID=id
	err:=h.BlogUseCase.Update(c,blog)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		})
}
func (h *blogHandler) delete(c *gin.Context){
	id:=c.Param("id")
	err:=h.BlogUseCase.Delete(c,id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		})
}
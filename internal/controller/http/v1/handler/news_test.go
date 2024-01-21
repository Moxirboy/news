package http

import (
	"bytes"
	"errors"
	"news/internal/models"
	"news/internal/service/usecase"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"news/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUseCase struct {
}

func MockNewHandler(
	group gin.RouterGroup,
	uc usecase.INewsUsecase,
	)usecase.INewsUsecase{
	return &mockUseCase{}
}
func (m *mockUseCase) Create(
	ctx context.Context,
	news *models.News,
	) (newsID string, err error) {
			return "1", nil
}

func (m *mockUseCase) GetOneByID(
		ctx context.Context,
		id string) (
		*models.News, error,
		){
	if id == "existing_id" {
		return &models.News{ID: "existing_id", Title: "Test News", Content: "This is a test news"}, nil
	} else if id == "non_existing_id" {
		return nil, errors.New("news not found")
	} 
		return nil, errors.New("invalid news ID")
	
}

func (m *mockUseCase) GetAll(
		ctx context.Context,
		query utils.PaginationQuery,
		) (
		*models.NewsList,
		error,
		) {
	if 	query.Page == 1 {
		return &models.NewsList{
			2,
			1,
			1,
			1,
			false,
			[]*models.News{
				{ID: "news_id_1", Title: "News 1", Content: "Content 1"},
				{ID: "news_id_2", Title: "News 2", Content: "Content 2"}},
			}, nil
	} else if query.Page == 2 {
		return &models.NewsList{
			2,
			1,
			1,
			1,
			false,
			[]*models.News{
				{ID: "news_id_3", Title: "News 3", Content: "Content 3"},
				}}, nil
	
	}
		return &models.NewsList{}, errors.New("invalid pagination")

}

func (m *mockUseCase) Update(
	ctx context.Context,
	news *models.News,
	) error {
	if news.ID== "existing_id" {
		return nil
	} else if news.ID == "non_existing_id" {
		return errors.New("news not found")
	}
		return errors.New("invalid news ID")
	
}

func (m *mockUseCase) Delete(
	ctx context.Context,
	id string,
	) error {
	if id == "existing_id" {
		return nil
	} else if id == "non_existing_id" {
		return errors.New("news not found")
	} else {
		return errors.New("invalid news ID")
	}
}

func TestNewsHandler_Create(t *testing.T) {
	router := gin.Default()
	mock:=mockUseCase{}
		NewNewsHandler(router.Group("/v1"), &mock)

	t.Run("Create news successfully", func(t *testing.T) {
		reqBody := `{"title": "Test News", "content": "This is a test news"}`
		req, err := http.NewRequest("POST", "/v1/news/create", bytes.NewBufferString(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Create news with invalid JSON", func(t *testing.T) {
		reqBody := `{"title": "Test News", "content": "This is a test news"`
		req, err := http.NewRequest("POST", "/v1/news/create", bytes.NewBufferString(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

}

func TestNewsHandler_GetOneById(t *testing.T) {
	router := gin.Default()
	mock:=mockUseCase{}
		NewNewsHandler(router.Group("/v1"), &mock)
	t.Run("Get news by existing ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/news/get/existing_id", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get news by non-existing ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/news/get/non_existing_id", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("Get news with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/news/get/invalid_id", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestNewsHandler_GetAll(t *testing.T) {
	router := gin.Default()
	mock:=mockUseCase{}
		NewNewsHandler(router.Group("/v1"), &mock)
	t.Run("Get all news - Page 1", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/news/getall?page=1", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get all news - Page 2", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/news/getall?page=2", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Get all news with invalid page", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/v1/news/getall?page=invalid", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

}

func TestNewsHandler_Update(t *testing.T) {
	router := gin.Default()
	mock:=mockUseCase{}
	NewNewsHandler(router.Group("/v1"), &mock)

	t.Run("Update news successfully", func(t *testing.T) {
		reqBody := `{"title": "Updated News", "content": "This is an updated news"}`
		req, err := http.NewRequest("PUT", "/v1/news/update/existing_id", bytes.NewBufferString(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Update news with non-existing ID", func(t *testing.T) {
		reqBody := `{"title": "Updated News", "content": "This is an updated news"}`
		req, err := http.NewRequest("PUT", "/v1/news/update/non_existing_id", bytes.NewBufferString(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("Update news with invalid ID", func(t *testing.T) {
		reqBody := `{"title": "Updated News", "content": "This is an updated news"}`
		req, err := http.NewRequest("PUT", "/v1/news/update/invalid_id", bytes.NewBufferString(reqBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestNewsHandler_Delete(t *testing.T) {
	router := gin.Default()
	mock:=mockUseCase{}
	NewNewsHandler(router.Group("/v1"), &mock)

	t.Run("Delete news successfully", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/v1/news/delete/existing_id", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Delete news with non-existing ID", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/v1/news/delete/non_existing_id", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("Delete news with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/v1/news/delete/invalid_id", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

}
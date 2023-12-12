package handler

import (
	"fmt"
	"layanan-kependudukan-api/article"
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	articleService article.Service
	authService    auth.Service
}

func NewArticleHandler(articleService article.Service, authService auth.Service) *articleHandler {
	return &articleHandler{articleService, authService}
}

func (h *articleHandler) CreateArticle(c *gin.Context) {
	var input article.CreateArticleInput

	currentUser, _ := c.Get("currentUser")
	userObject := currentUser.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		fmt.Print(input)
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		fmt.Print(err.Error())
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	path := helper.FormatFileName(file.Filename)
	filePath := filepath.Join("images/articles", path)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	input.ImagePath = filePath
	input.Author = userObject.Name
	newarticle, err := h.articleService.CreateArticle(input)
	if err != nil {
		response := helper.APIResponse("Failed create article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := article.FormatArticle(newarticle)
	response := helper.APIResponse("Success create article", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) UpdateArticle(c *gin.Context) {
	var inputID article.GetArticleDetailInput
	var inputData article.CreateArticleInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newarticle, err := h.articleService.UpdateArticle(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := article.FormatArticle(newarticle)
	response := helper.APIResponse("Success Update article", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) DeleteArticle(c *gin.Context) {
	var inputID article.GetArticleDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.articleService.DeleteArticle(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete article", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetArticles(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.articleService.GetArticles(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get article", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	articles, _ := pagination.Data.([]article.Article)
	pagination.Data = article.FormatArticles(articles)

	response := helper.APIResponse("Success get article", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

func (h *articleHandler) GetArticle(c *gin.Context) {
	var inputID article.GetArticleDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get Article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newArticle, err := h.articleService.GetArticleByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get Article", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := article.FormatArticle(newArticle)
	response := helper.APIResponse("Success Get Article", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

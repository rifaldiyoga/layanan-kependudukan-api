package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/job"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jobHandler struct {
	jobService  job.Service
	authService auth.Service
}

func NewJobHandler(jobService job.Service, authService auth.Service) *jobHandler {
	return &jobHandler{jobService, authService}
}

func (h *jobHandler) CreateJob(c *gin.Context) {
	var input job.CreateJobInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create job", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newjob, err := h.jobService.CreateJob(input)
	if err != nil {
		response := helper.APIResponse("Failed create job", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := job.FormatJob(newjob)
	response := helper.APIResponse("Success create job", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jobHandler) UpdateJob(c *gin.Context) {
	var inputID job.GetJobDetailInput
	var inputData job.CreateJobInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update job", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update job", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newjob, err := h.jobService.UpdateJob(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update job", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := job.FormatJob(newjob)
	response := helper.APIResponse("Success Update job", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jobHandler) DeleteJob(c *gin.Context) {
	var inputID job.GetJobDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete job", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.jobService.DeleteJob(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete job", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete job", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *jobHandler) GetJobs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.jobService.GetJobs(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get job", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	jobs, _ := pagination.Data.([]job.Job)
	pagination.Data = job.FormatJobs(jobs)

	response := helper.APIResponse("Success get job", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

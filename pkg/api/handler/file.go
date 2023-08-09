package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/response"
	usecase "github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/usecase/interfaces"
)

type fileHandler struct {
	usecase usecase.FileUseCase
}

func NewFileHandler(usecase usecase.FileUseCase) interfaces.FileHandler {
	return &fileHandler{
		usecase: usecase,
	}
}

// UploadSingleFile godoc
// @summary api for upload single file
// @Accept multipart/form-data
// @tags File
// @id UploadSingleFile
// @Param     file   formData     file   true   "Upload one file"
// @Param name formData string true "File Name"
// @Param description formData string true "File Description"
// @Router /uploads/one [post]
// @Success 201 {object} response.Response{} "Successfully file uploaded"
// @Failure 400 {object} response.Response{} "Invalid inputs"
// @Failure 500 {object} response.Response{} "Failed to upload file"
func (c *fileHandler) UploadSingleFile(ctx *gin.Context) {

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "failed to get file from request", err, nil)
		return
	}

	name, err1 := request.GetFormValuesAsString(ctx, "name")
	description, err2 := request.GetFormValuesAsString(ctx, "description")
	err = errors.Join(err1, err2)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "failed to get form values", err, nil)
		return
	}

	singleFile := request.SingleFile{
		Name:        name,
		Description: description,
		FileHeader:  fileHeader,
	}

	err = c.usecase.UploadSingleFile(ctx, singleFile)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Failed to upload file", err, nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully file uploaded")
}

// GetAllFiles godoc
// @summary api for retrieve all files
// @tags File
// @id GetAllFiles
// @Param   page_number query int false "Page Number"
// @Param count query int false "Count of Files"
// @Router /all [get]
// @Success 201 {object} response.Response{data=[]response.SingleFile} "Successfully retrieved all files"
// @Failure 500 {object} response.Response{} "Failed to retrieve all files"
func (c *fileHandler) GetAllFiles(ctx *gin.Context) {

	pagination := request.GetPagination(ctx)

	files, err := c.usecase.GetAllFiles(ctx, pagination)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve all files", err, nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved all files", files)
}

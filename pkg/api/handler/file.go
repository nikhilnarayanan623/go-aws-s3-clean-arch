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
// @Accept       json
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
		response.ErrorResponse(ctx, http.StatusBadRequest, "failed to upload file", err, nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully file uploaded")
}

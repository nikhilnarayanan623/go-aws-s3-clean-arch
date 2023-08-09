package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Get query params as string from request url
func GetParamAsUint(ctx *gin.Context, key string) (string, error) {

	param := ctx.Param(key)

	if param == "" {
		return "", fmt.Errorf("failed to get %s from param as int", key)
	}

	return param, nil
}

// Get values from request form as string
func GetFormValuesAsString(ctx *gin.Context, name string) (value string, err error) {

	value = ctx.Request.PostFormValue(name)
	if value == "" {
		return "", fmt.Errorf("failed to get %s from request body", name)
	}

	return value, nil
}

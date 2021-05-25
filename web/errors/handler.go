package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func DefaultErrorHandler(ctx *gin.Context) bool {

	for _, err := range ctx.Errors {
		switch err.Err.(type) {
		case NotImplementedError:
			ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, &ApiError{
				Title:   http.StatusText(http.StatusMethodNotAllowed),
				Message: err.Err.Error(),
				Details: err.Err,
			})
			return true
		case ResourceNotFoundError:
			ctx.AbortWithStatusJSON(http.StatusNotFound, &ApiError{
				Title:   http.StatusText(http.StatusNotFound),
				Message: err.Err.Error(),
				Details: err.Err,
			})
			return true
		case DuplicatedEntryError:
			ctx.AbortWithStatusJSON(http.StatusConflict, &ApiError{
				Title:   http.StatusText(http.StatusConflict),
				Message: err.Err.Error(),
				Details: err.Err,
			})
			return true
		case ForbiddenError:
			ctx.AbortWithStatusJSON(http.StatusForbidden, &ApiError{
				Title:   http.StatusText(http.StatusForbidden),
				Message: err.Err.Error(),
				Details: err.Err,
			})
			return true
		case InvalidArgumentError:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ApiError{
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Err.Error(),
				Details: err.Err,
			})
			return true
		case *strconv.NumError:
			var value = err.Err.(*strconv.NumError).Num
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ApiError{
				Title:   http.StatusText(http.StatusBadRequest),
				Message: fmt.Sprintf("error parsing the value \"%s\" to a number", value),
				Details: struct {
					Value string `json:"value"`
				}{
					Value: value,
				},
			})
			return true
		case validator.ValidationErrors:
			vErrs, _ := err.Err.(validator.ValidationErrors)
			errs := make(map[string]string)
			for _, f := range vErrs {
				errs[f.Field()] = f.ActualTag()
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ApiError{
				Title:   http.StatusText(http.StatusBadRequest),
				Message: err.Err.Error(),
				Details: errs,
			})
			return true
		}
	}

	return false
}

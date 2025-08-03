package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the standard response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// SuccessResponse sends a standardized success response
func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends a standardized error response
func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   message,
	})
}

// ValidationErrorResponse sends a standardized validation error response
func ValidationErrorResponse(ctx *gin.Context, errors interface{}) {
	ctx.JSON(http.StatusUnprocessableEntity, Response{
		Success: false,
		Message: "Validation failed",
		Error:   errors,
	})
}

// PaginatedResponse sends a standardized paginated response
func PaginatedResponse(ctx *gin.Context, message string, data interface{}, total int64, page int, pageSize int) {
	ctx.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   message,
		"data":      data,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// UnauthorizedResponse sends a standardized unauthorized response
func UnauthorizedResponse(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: message,
		Error:   "Unauthorized",
	})
}

// NotFoundResponse sends a standardized not found response
func NotFoundResponse(ctx *gin.Context, resource string) {
	ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: resource + " not found",
		Error:   "Not Found",
	})
}

// InternalServerErrorResponse sends a standardized internal server error response
func InternalServerErrorResponse(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: message,
		Error:   "Internal Server Error",
	})
}

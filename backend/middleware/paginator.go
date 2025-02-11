package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	constants "github.com/tejiriaustin/lema/constants"
	"github.com/tejiriaustin/lema/response"
)

func ReadPaginationOptions() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if pageNumber := ctx.Query("pageNumber"); pageNumber != "" {
			pageNum, err := strconv.ParseInt(pageNumber, 10, 64)
			if err != nil {
				response.FormatResponse(ctx, http.StatusBadRequest, "page number must be a number", nil)
				ctx.Abort()
				return
			}
			ctx.Set(string(constants.ContextKeyPageNumber), pageNum)
		}

		if pageSize := ctx.Query("pageSize"); pageSize != "" {
			perPageNum, err := strconv.ParseInt(pageSize, 10, 64)
			if err != nil {
				response.FormatResponse(ctx, http.StatusBadRequest, "page size must be a number", nil)
				ctx.Abort()
				return
			}
			ctx.Set(string(constants.ContextKeyPageSize), perPageNum)
		}

		ctx.Next()
	}
}

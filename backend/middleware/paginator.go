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
		params := ctx.Params

		if pageNumber, _ := params.Get("pageNumber"); pageNumber != "" {
			pageNum, err := strconv.ParseInt(pageNumber, 10, 64)
			if err != nil {
				response.FormatResponse(ctx, http.StatusBadRequest, "page number must be a number", nil)
				return
			}

			ctx.Set(string(constants.ContextKeyPageNumber), pageNum)
		}

		if pageSize, _ := params.Get("pageSize"); pageSize != "" {
			perPageNum, err := strconv.ParseInt(pageSize, 10, 64)
			if err != nil {
				response.FormatResponse(ctx, http.StatusBadRequest, "per page number must be a number", nil)
				return
			}

			ctx.Set(string(constants.ContextKeyPageNumber), perPageNum)
		}

	}
}

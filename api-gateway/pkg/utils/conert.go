package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/api-gateway/pkg/utils/request"
)

const (
	defaultPageNumber = 1
	defaultPageCount  = 10
)

func GetPagination(ctx *gin.Context) request.Pagination {

	pagination := request.Pagination{
		PageNumber: defaultPageNumber,
		Count:      defaultPageCount,
	}

	pn, err := strconv.ParseUint(ctx.Query("page_number"), 10, 64)
	if err == nil && pn != 0 {
		pagination.PageNumber = pn
	}

	c, err := strconv.ParseUint(ctx.Query("count"), 10, 64)
	if err == nil && c != 0 {
		pagination.Count = c
	}

	return pagination
}

func StringToUint64(str string) (uint64, error) {

	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid string failed to convert \nerror:%w", err)
	}

	return num, nil
}

package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLimit(c *gin.Context) int {
	limitStr := c.Query("limit")
	limit := 100
	if len(strings.TrimSpace(limitStr)) > 0 {
		limit, err := strconv.Atoi(limitStr)
		if limit > 100 || err != nil || limit <= 0 {
			limit = 100
		}
	}
	return limit
}

func GetOffset(c *gin.Context) int {
	offsetStr := c.Query("offset")
	offset := 0
	if len(strings.TrimSpace(offsetStr)) > 0 {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			offset = 0
		}
	}
	return offset
}

package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, status, code int, msg string, data any) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	})
}

func ParseQueryToUint(ctx *gin.Context, paramName string) (uint, error) {
	paramInt, err := strconv.Atoi(ctx.Query(paramName))
	if err != nil {
		return 0 ,err
	} else {
		paramUint := uint(paramInt)
		return paramUint, err
	}
}


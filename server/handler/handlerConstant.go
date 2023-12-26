package handler

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, status, code int, msg string, data any) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg": msg,
		"data": data,
	})
}


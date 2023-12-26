package handler

import (
	"server/db"
	"server/model"
	"github.com/gin-gonic/gin"
	"fmt"
)

// user register

func RegisterHandler(ctx *gin.Context) {
	//1. 获取数据,创建结构体 通过注册界面创建的用户默认为 1 
	var user model.User
	if err := ctx.ShouldBindJSON(&user) ; err != nil {
		Response(ctx, 200, 1, "解析参数失败:" + err.Error(), "")
		return 
	}
	user.Email = "@163.com"
	user.IsAdmin = false
	user.RoleID = 1
	// 2. 获取数据库连接
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)
	// 3. 插入数据
	if err := conn.Create(&user).Error ; err != nil {
		Response(ctx, 200, 1, err.Error(), "")
	} else {
		Response(ctx, 200, 0, "注册成功,请登录", "")
	}
}

// user login

func LoginHandler(ctx *gin.Context) {
	// 1. 获取用户名和密码
	var user model.User
	if err := ctx.ShouldBindJSON(&user) ; err != nil {
		Response(ctx, 200, 1, "解析参数失败:" + err.Error(), "")
		return 
	} 
	// 2. 获取数据库连接
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)
	// 3. 查询用户是否存在
	err := db.GetUserByOption(conn, user.Username, user.Password, &user)
	if err != nil {
		Response(ctx, 200, 1, "未找到用户,请检查用户名和密码:" + err.Error(), "")
		return 
	}
	// 如果是管理员 则直接返回数据
	if user.IsAdmin {
		Response(ctx, 200, 0, "登录成功", user)
		return
	}
	// 如果不是管理员 则查询用户的所有对应界面的权限
	accessWithLevels, _ := db.GetAccessByRoleId(conn, user.RoleID)
	fmt.Println("accessWithLevels = ", accessWithLevels)
	Response(ctx, 200, 0, "登录成功", map[string]interface{}{"user": user, "access": accessWithLevels})
}
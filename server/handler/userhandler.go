package handler

import (
	"server/db"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"fmt"
)

// 获取所有的角色
func GetUser(ctx *gin.Context) {
	
	// 1.数据库相关操作
	
	// 获取数据库连接 
	conn := db.GetDatabaseConnection()
	// 关闭数据库连接
	defer db.CloseDatabaseConnection(conn)

	// 2.获取请求数据

	pageNum, err := strconv.Atoi(ctx.Query("pageNum")) 
	if err != nil {
		Response(ctx, 500, 1, "查询用户失败", nil)
	}
	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		Response(ctx, 500, 1, "查询用户失败", nil)
	}
	userName := ctx.Query("userName")
	roleName := ctx.Query("roleName")

	// 2. 查询user表的所有数据

	users, err := db.GetAllUser(conn, pageNum, pageSize, userName, roleName) 
	if err != nil {
		Response(ctx, 500, 1, "查询用户失败", nil)
	} else {         
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg": "查询用户成功",
			"data": users,
			"total": db.GetUserCount(conn),
		})
	}
}

func AddUser(ctx *gin.Context) {
	// 1.获取数据
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	rolename := ctx.PostForm("rolename")
	// 2. 获取数据库连接
	db := db.GetDatabaseConnection()
	// 2.查询该角色ID
	var providedRole model.Role
	if err := db.Where("name = ?", rolename).First(&providedRole).Error; err != nil {
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": err,
		})
		return 
	}
	// 3.将数据存入数据库
	user := model.NewUser(username, password, email, false, true, providedRole.ID)
	if err := db.Create(&user).Error ; err != nil {
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg": "添加用户成功",
		})
	}
}

func UpdateUser(ctx *gin.Context) {
	// 1. 获取数据
	id := ctx.PostForm("id")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	rolename := ctx.PostForm("rolename")

	// 2. 获取数据库连接
	db := db.GetDatabaseConnection()

	// 3. 根据角色名查询该角色是否存在
	var providedRole model.Role
	if err := db.Where("name = ?", rolename).First(&providedRole).Error; err != nil {
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": err,
		})
		return 
	}
	// 4. 更新数据
	if err := db.Model(&model.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"username": username,
		"password": password,
		"email": email,
		"role_id": providedRole.ID,
	}).Error ; err != nil {
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0, 
			"msg": "更新成功",
		})
	}
}

func DeleteUser(ctx *gin.Context) {

	fmt.Println("执行逻辑")
	// 获取传来的数据
	id := ctx.Query("id")
	// 获取数据库连接

	fmt.Println("id = ", id)

	conn := db.GetDatabaseConnection()

	defer db.CloseDatabaseConnection(conn)

	result := conn.Delete(&model.User{}, id)

	// 检查删除操作是否出现错误
	if result.Error != nil {
		fmt.Println(result.Error)
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": result.Error,
		})
		return 
	}

	// 检查是否成功删除
	if result.RowsAffected == 0 {
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": "删除失败,未找到记录",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg": "删除成功",
		})
	}
}
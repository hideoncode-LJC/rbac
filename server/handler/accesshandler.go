package handler

import (
	"fmt"
	"server/db"
	"server/model"
	"github.com/gin-gonic/gin"
)

/*
当点击每个角色时，根据每个角色获取每个角色的权限，进行显示
*/

func GetAccessByRoleIDHandler(ctx *gin.Context) {
	// 1. 获取数据库连接和数据
	conn := db.GetDatabaseConnection()
	// 2. 获取所有的权限
	accesses, err := db.GetAllAccess(conn)
	if err != nil {
		Response(ctx, 500, 1, "查询失败", "")
		return 
	}
	// 3. 根据角色ID和所有的权限获取权限等级

	// 定义返回数据
	var accessWithLevels []model.AccessWithLevel
	for i := 0 ; i < len(accesses) ; i ++ {
		var roleWithAccess model.RoleAccess
		var accessWithLevel model.AccessWithLevel
		if err := conn.Where("role_id = ? AND access_id = ?", ctx.Param("id"), accesses[i].ID).First(&roleWithAccess).Error; err == nil {
			// 查询到权限，进行赋值
			accessWithLevel.Access = accesses[i]
			accessWithLevel.Level = roleWithAccess.Level
		} else {
			accessWithLevel.Access = accesses[i]
			accessWithLevel.Level = 0
		}
		accessWithLevels = append(accessWithLevels, accessWithLevel)
	}
	// 4. 返回数据
	Response(ctx, 200, 0, "查询成功", accessWithLevels)
}


func AddAccess(ctx *gin.Context) {
	//1. 获取数据
	var addAccess model.AddAccessModel
	fmt.Println("access = ", addAccess)
	if err := ctx.ShouldBindJSON(&addAccess) ; err != nil {
		fmt.Println("err = ", err.Error())
		Response(ctx, 500, 1, "参数错误", "")
		return 
	}
	access := model.Access{
		Title: addAccess.Title,
		Content: addAccess.Content,
	}
	// 2. 获取数据库连接 以及 延迟关闭
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)

	// 3.  检查数据库中是否已存在相同名称的 Role
	err := db.HasAccessWithTitle(conn, access.Title)
	if err == nil {
		Response(ctx, 200, 1, "该页面的名字已经存在,请更换一下角色的名字", "")
		return 
	}
	// 如果不存在同名 Role，则创建新的 Role 对象并保存到数据库中
	// 在数据库中创建 Role 记录
	if err := db.AddAccess(conn, &access); err != nil  {
		Response(ctx, 500, 1, "创建页面失败", "")
	} else {
		Response(ctx, 200, 0, "创建页面成功", "")
	} 
}

func GetAccess(ctx *gin.Context) {
	// 1. 获取数据库连接
	conn := db.GetDatabaseConnection()
	// 2. 基于数据库连接进行关联查询
	accesses, err := db.GetAllAccess(conn)
	if err != nil {
		Response(ctx, 500, 1, "查询失败", "")
	} else {
		Response(ctx, 200, 0, "查询成功", accesses)
	}
}

func DeleteAccess(ctx *gin.Context) {
	// 获取要删除的页面ID
	accessID := ctx.Query("id")
	//获取数据库连接
	db := db.GetDatabaseConnection()
	// 先查询角色是否存在
    var access model.Access
    if err := db.First(&access, accessID).Error; err != nil {
		// 角色不存在
		ctx.JSON(500, gin.H{
			"code": 1,
			"msg": err,
		})
		return 
    }
    // 删除关联的权限
    if err := db.Where("access_id = ?", accessID).Delete(&model.RoleAccess{}).Error; err != nil {
        ctx.JSON(500, gin.H{
			"code": 1,
			"msg": err,
		})
		return 
    }

    // 删除角色
    if err := db.Delete(&access).Error; err != nil {
        ctx.JSON(500, gin.H{
			"code": 1,
			"msg": err,
		}) 
    } else {
		// 删除成功 返回响应
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg": "删除角色成功",
		})
	}
}

func UpdateAccess(ctx *gin.Context) {
	// 1. 获取数据
	id := ctx.PostForm("id")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	// 2. 获取数据库连接
	db := db.GetDatabaseConnection()

	// 3. 根据角色名查询该角色是否存在
	var providedAccess model.Access
	if err := db.Where("id = ?", id).First(&providedAccess).Error; err != nil {
		ctx.JSON(500, gin.H{
			"code": 1, 
			"error": err,
		})
		return 
	}
	// 4. 更新数据
	if err := db.Model(&model.Access{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title": title,
		"content": content,
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
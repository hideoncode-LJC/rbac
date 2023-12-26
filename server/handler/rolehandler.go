package handler

import (
	"server/db"
	"server/model"
	"github.com/gin-gonic/gin"
	"fmt"
)

// 获取所有角色
func GetRole(ctx *gin.Context) {
	// 1. 获取数据库连接
	conn := db.GetDatabaseConnection()
	// 关闭数据库连接
	defer db.CloseDatabaseConnection(conn)
	// 2. 基于数据库连接进行查询所有角色
	roles, err := db.GetAllRole(conn)
	if err != nil {
		Response(ctx, 500, 1, "查询失败", "")
	} else {
		Response(ctx, 200, 0, "查询成功", roles)
	}	


}

func AddRole(ctx *gin.Context) {
	//1. 获取角色的名字和相关权限
	var addRoleModel model.AddRoleModel
	if err := ctx.ShouldBindJSON(&addRoleModel); err != nil {
		fmt.Println("err = ", err)
		Response(ctx, 500, 1, "参数错误", "")
		return
	}
	fmt.Println("addRoleModel = ", addRoleModel)
	// 2. 获取数据库连接
	conn := db.GetDatabaseConnection()
	// 关闭数据库连接
	defer db.CloseDatabaseConnection(conn)
	/*
	检查数据库中是否已存在相同名称的 Role
	如果 err == nil, 则表示数据库中已经存在同名 Role，返回错误响应
	*/

	_, err := db.GetRoleByName(conn, addRoleModel.RoleName)
	if err == nil {
		Response(ctx, 500, 1, "该角色名已经存在", "")
		return 
	}
	// 如果不存在同名 Role，则创建新的 Role 对象并保存到数据库中
	
	role := model.NewRole(addRoleModel.RoleName)
	access := new(model.Access)
	access.ID = 1
	role.Accesses = append(role.Accesses, *access)
	// 在数据库中创建 Role 记录
	err = db.AddRole(conn, *role)
	if err != nil {
		Response(ctx, 1, 500, "创建角色失败", "")
	} else {
		Response(ctx, 200, 0, "创建角色成功", "")
	}
}

func UpdateRole(c *gin.Context) {

}

func DeleteRole(c *gin.Context) {
	// 获取要删除的角色ID
	roleID := c.Param("id")
	//获取数据库连接
	db := db.GetDatabaseConnection()
	// 先查询角色是否存在
    var role model.Role
    if err := db.First(&role, roleID).Error; err != nil {
		// 角色不存在
		c.JSON(500, gin.H{
			"code": 1,
			"msg": err,
		})
		return 
    }
    // 删除关联的权限
    if err := db.Where("role_id = ?", roleID).Delete(&model.RoleAccess{}).Error; err != nil {
        c.JSON(500, gin.H{
			"code": 1,
			"msg": err,
		})
		return 
    }

    // 删除角色
    if err := db.Delete(&role).Error; err != nil {
        c.JSON(500, gin.H{
			"code": 1,
			"msg": err,
		}) 
    } 

	// 删除成功 返回响应
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "删除角色成功",
	})

}
package handler

import (
	"server/db"
	"server/log"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取所有角色
func GetRole(ctx *gin.Context) { 
	conn := db.GetDatabaseConnection() // 获取数据库连接
	defer db.CloseDatabaseConnection(conn) // 关闭数据库连接
	roles, err := db.GetAllRole(conn) 	// 基于数据库连接进行查询所有角色
	if err != nil {
		Response(ctx, 500, 1, "查询失败", "")
	} else {
		Response(ctx, 200, 0, "查询成功", roles)
	}	
}

func GetChildRole(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		Response(ctx, 200, 1, "参数错误", "")
		return
	}
	roleId := uint(param)
	conn := db.GetDatabaseConnection() // 获取数据库连接
	defer db.CloseDatabaseConnection(conn) // 关闭数据库连接
	roles, err := db.GetChildRole(conn, roleId) // 基于数据库连接进行查询所有角色
	if err != nil {
		Response(ctx, 200, 1, "查询失败", "")
	} else {
		Response(ctx, 200, 0, "查询成功", roles)
	}	
}

func CheckChildLevel(ctx *gin.Context) {
	log.Info("得到请求")
	accessIDint, err := ParseQueryToUint(ctx, "accessid")
	if err != nil {
		Response(ctx, 200, 1, "参数错误" + err.Error(), "")
	}
	accessID := uint(accessIDint)
	roleIDint, err := ParseQueryToUint(ctx, "roleid")

	if err != nil {
		Response(ctx, 200, 1, "参数错误" + err.Error(), "")
	}
	roleID := uint(roleIDint)
	conn := db.GetDatabaseConnection() // 获取数据库连接
	defer db.CloseDatabaseConnection(conn) // 关闭数据库连接
	level, err := db.GetChildLevel(conn, roleID, accessID)
	if err != nil {
		Response(ctx, 200, 1, "查询失败" + err.Error(), "")
	} else {
		Response(ctx, 200, 0, "查询成功", level)
	}

}


func AddRole(ctx *gin.Context) {

	var addRoleModel model.AddRoleModel // 定义接收数据的类型
	if err := ctx.ShouldBindJSON(&addRoleModel); err != nil { // 绑定数据
		Response(ctx, 200, 1, "参数错误" + err.Error(), "")
		return
	}
	log.Info(addRoleModel)
	conn := db.GetDatabaseConnection() // 获取数据库连接
	defer db.CloseDatabaseConnection(conn) 	// 关闭数据库连接
	/*
	检查数据库中是否已存在相同名称的 Role
	如果 err == nil, 则表示数据库中已经存在同名 Role，返回错误响应
	*/
	_, err := db.GetRoleByName(conn, addRoleModel.RoleName)
	if err == nil {
		Response(ctx, 200, 1, "该角色名已经存在", "")
		return 
	}
	// 如果不存在同名 Role，则创建新的 Role 对象并保存到数据库中
	var role model.Role
	role.Name = addRoleModel.RoleName
	result := conn.Create(&role)
	if result.Error != nil {
		Response(ctx, 200, 1, "创建角色失败", "")
		return
	}
	for i := 0 ; i < len(addRoleModel.Accesses) ; i ++ {
		level := addRoleModel.Accesses[i].Level
		if level > 0 {
			var roleAccess model.RoleAccess
			roleAccess.RoleID = role.ID
			roleAccess.AccessID = addRoleModel.Accesses[i].ID
			roleAccess.Level = level
			if err := conn.Create(&roleAccess).Error; err != nil {
				Response(ctx, 200, 1, "创建角色失败", "")
				return
			}
		}
	}
	for i := 0 ; i < len(addRoleModel.ChildID) ; i ++ {
		childID := uint(addRoleModel.ChildID[i])
		db.AddChildRole(conn, role.ID, childID)
	}
	Response(ctx, 200, 0, "创建角色成功", "")
}

func UpdateRole(ctx *gin.Context) {
	// 1. 获取角色ID参数
	RoleID, err := ParseQueryToUint(ctx, "id")
	if err != nil {
		Response(ctx, 200, 1, "解析url参数错误" + err.Error(), "")
		return
	}
	log.Info("更新角色")
	log.Info("获取角色ID成功")
	log.Info(RoleID)
	var addRoleModel model.AddRoleModel // 定义接收数据的类型
	if err := ctx.ShouldBindJSON(&addRoleModel); err != nil { // 绑定数据
		Response(ctx, 200, 1, "解析JSON参数错误" + err.Error(), "")
		return
	}
	log.Info("获取角色信息成功")
	log.Info(addRoleModel)
	conn := db.GetDatabaseConnection() // 获取数据库连接
	defer db.CloseDatabaseConnection(conn) // 关闭数据库连接 
	
	if err := db.UpdateRole(conn, RoleID, &addRoleModel); err != nil {
		Response(ctx, 200, 1, "更新角色失败" + err.Error(), "")
		return
	}
	// 更新角色权限
	db.DeleteAllChild(conn, RoleID);
	for i := 0 ; i < len(addRoleModel.ChildID) ; i ++ {
		childID := uint(addRoleModel.ChildID[i])
		db.AddChildRole(conn, RoleID, childID)
	}

	Response(ctx, 200, 0, "更新角色成功", "")
}

func DeleteRole(ctx *gin.Context) {
	// 获取要删除的角色ID
	param, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		Response(ctx, 200, 1, "参数错误" + err.Error(), "")
	}
	roleID := uint(param)
	//获取数据库连接
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn) // 关闭数据库连接
	// 先查询角色是否存在
    var role model.Role
    if err := conn.First(&role, roleID).Error; err != nil {
		// 角色不存在
		Response(ctx, 200, 1, "角色不存在", "")
    }
    // 删除关联的权限
    if err := conn.Where("role_id = ?", roleID).Delete(&model.RoleAccess{}).Error; err != nil {
        Response(ctx, 200, 1, "删除角色失败", "")
		return 
    }

    // 删除角色
    if err := conn.Delete(&role).Error; err != nil {
		Response(ctx, 200, 1, "删除角色失败", "")
    } 
	// 将原来拥有该觉得用户ID进行修改
	db.UpdateUserRole(conn, roleID)

	

	// 删除成功 返回响应
	Response(ctx, 200, 0, "删除角色成功", "")

}
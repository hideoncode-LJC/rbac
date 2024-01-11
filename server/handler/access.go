package handler

import (
	"fmt"
	"server/db"
	"server/model"
	"strconv"
	"github.com/gin-gonic/gin" 
	"server/log"
)

/*
复用函数
传递来的参数 
roleID 代表角色ID 则获取该角色的所有权限的level
accessID 代表权限ID 如果权限ID是0 代表获取所有权限 不是0 代表获取单个权限的内容
*/

func GetAccess(ctx *gin.Context) {
	// 1. 解析从前端获取来的参数 roleid & accessid
	RoleID, err := ParseQueryToUint(ctx, "roleid")
	if err != nil {
		Response(ctx, 200, 1, "参数解析错误" + err.Error(), "")
	}
	AccessID, err := ParseQueryToUint(ctx, "accessid")
	if err != nil {
		Response(ctx, 200, 1, "参数解析错误" + err.Error(), "")
	}
	// 2. 获取数据库连接
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)
	// 3. 根据返回的参数进行查询
	data, err := db.GetAccess(conn, RoleID, AccessID)
	log.Info("这里是/access处理函数,返回的数据为")
	log.Info(data)
	if err != nil {
		Response(ctx, 200, 1, "查询失败" + err.Error(), "")
	} else {
		Response(ctx, 200, 0, "查询成功", data)
	}
	// accesses, err := db.GetAllAccess(conn) 
	// if id == 0 {
	// 	if err != nil {
	// 		Response(ctx, 200, 1, "查询所有权限失败", "")
	// 	} else {
	// 		Response(ctx, 200, 0, "查询所有权限成功", accesses)
	// 	}
	// 	return 	 
	// } 
	// // 根据角色ID和所有的权限获取权限等级
	// var accessWithLevels []model.AccessWithLevel
	// for i := 0 ; i < len(accesses) ; i ++ {
	// 	var roleWithAccess model.RoleAccess
	// 	var accessWithLevel model.AccessWithLevel
	// 	accessWithLevel.ID = accesses[i].ID
	// 	accessWithLevel.Title = accesses[i].Title
	// 	accessWithLevel.Content = accesses[i].Content
	// 	if err := conn.Where("role_id = ? AND access_id = ?", id, accesses[i].ID).First(&roleWithAccess).Error; err == nil {
	// 		// 查询到权限，进行赋值
	// 		accessWithLevel.Level = roleWithAccess.Level
	// 	} else {
	// 		accessWithLevel.Level = 0
	// 	}
	// 	accessWithLevels = append(accessWithLevels, accessWithLevel)
	// }
	// // 4. 返回数据
	// Response(ctx, 200, 0, "查询成功", accessWithLevels)
}



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
		accessWithLevel.ID = accesses[i].ID
		accessWithLevel.Title = accesses[i].Title
		accessWithLevel.Content = accesses[i].Content
		if err := conn.Where("role_id = ? AND access_id = ?", ctx.Param("id"), accesses[i].ID).First(&roleWithAccess).Error; err == nil {
			// 查询到权限，进行赋值
			accessWithLevel.Level = roleWithAccess.Level
		} else {
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

func DeleteAccess(ctx *gin.Context) {
	log.Info("开始删除权限")
	// 获取要删除的页面ID
	param, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		log.Warn("解析参数失败")
		Response(ctx, 200, 1, "解析参数失败", "")
	}
	accessID := uint(param)
	log.Info("要删除的权限ID为" + fmt.Sprint(accessID))
	//获取数据库连接
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)
	// 先查询角色是否存在
    var access model.Access
    if err := conn.
	Model(&model.Access{}).
	Where("id = ?", accessID).
	First(&access).Error; err != nil {
		log.Warn("查询权限是否存在出现错误" + err.Error())
		Response(ctx, 200, 1, "该权限不存在", "")
		return 
    }
    // 删除关联的权限
    if err := conn.
	Model(&model.RoleAccess{}).
	Delete("access_id = ?", accessID).Error; err != nil {
		log.Warn("查询关联表记录是否存在出现错误" + err.Error())
        Response(ctx, 200, 1, "查询关联表记录出现错误", "")
		return 
    }

    if result := conn.
	Model(&model.Access{}).
	Delete("id = ?", accessID); result.Error != nil {
        log.Warn("删除权限是否发生错误" + result.Error.Error())
		Response(ctx, 200, 1, "删除权限发生错误", "")
    } else {
		log.Info("删除权限影响的行数为" + fmt.Sprint(result.RowsAffected))
		if result.RowsAffected == 0 {
			log.Warn("删除权限失败")
			Response(ctx, 200, 0, "删除权限失败", "")
		}
		log.Info("删除权限成功")
		Response(ctx, 200, 0, "删除权限成功", "")
	}
}

type UpdateAccessModel struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func UpdateAccess(ctx *gin.Context) {
	accessID, err := ParseQueryToUint(ctx, "id")
	if err != nil {
		Response(ctx, 200, 1, "解析参数错误" + err.Error(), "")
	}
	var updateAccessModel UpdateAccessModel
	if err := ctx.ShouldBindJSON(&updateAccessModel); err != nil {
		Response(ctx, 200, 1, "参数错误" + err.Error(), "")
	}

	// 2. 获取数据库连接
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)
	// 4. 更新数据
	if err := conn.
	Model(&model.Access{}).
	Where("id = ?", accessID).
	Updates(map[string]interface{}{
		"title": updateAccessModel.Title,
		"content": updateAccessModel.Content,
	}).Error ; err != nil {
		ctx.JSON(200, gin.H{
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



func GetAccessById(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Query("accessid"))
	if err != nil {
		Response(ctx, 200, 1, "解析参数错误", "")
		return 
	}
	conn := db.GetDatabaseConnection()
	defer db.CloseDatabaseConnection(conn)
	accessID := uint(param)
	var access model.Access
	if err := conn.Where("id = ?", accessID).First(&access).Error; err != nil {
		Response(ctx, 200, 1, "查询失败", "")
	} else {
		Response(ctx, 200, 0, "查询成功", access)
	}

}
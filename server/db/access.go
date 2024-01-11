package db

import (
	"fmt"
	"server/log"
	"server/model"

	"gorm.io/gorm"
)

func GetAccess(db *gorm.DB, RoleID, AccessID uint) (any, error) {
	// 查询的是权限，不和角色ID存在关系
	if RoleID == 0 {
		if AccessID == 0 {
			// 查询所有权限
			return GetAllAccess(db)
		} 
		return GetAccessByID(db, AccessID)
	} else {
		// 查询角色的权限
		return GetAccessByRoleId(db, RoleID)
	}
}

func GetAccessByID(db *gorm.DB, id uint) (model.Access, error) {
	var access model.Access
	result := db.Where("id = ?", id).First(&access)
	log.Info("查询权限ID为" + fmt.Sprint(id) + "的权限查询到的记录个数为" + fmt.Sprint(result.RowsAffected))
	return access, result.Error
}

// 查找所有权限 不是关联查询

func GetAllAccess(db *gorm.DB) ([]model.Access, error) {
	var accesses []model.Access
	result := db.Model(model.Access{}).Find(&accesses)
	log.Info("查询所有权限查询到的记录个数为" + fmt.Sprint(result.RowsAffected))
	return accesses, result.Error
}

// 数据库中已经有该access了 返回true
func HasAccessWithTitle(db *gorm.DB, title string) (error) {
	var access model.Access
	err := db.Where("title = ?", title).First(&access).Error
	return err
}

func AddAccess(db *gorm.DB, access *model.Access) error {
	return db.Create(access).Error
}

func UpdateAccess(db *gorm.DB, access *model.Access) error {
	return db.Save(access).Error
}

func GetAccessByRoleId(db *gorm.DB, roleId uint, ) ([]model.AccessWithLevel, error) {
	// 首先获取权限列表
	accesses, err := GetAllAccess(db)
	if err != nil {
		return nil, err
	}
	// 根据角色ID和所有的权限获取权限等级
	var accessWithLevels []model.AccessWithLevel
	for i := 0 ; i < len(accesses) ; i ++ {
		var roleWithAccess model.RoleAccess
		var accessWithLevel model.AccessWithLevel
		accessWithLevel.ID = accesses[i].ID
		accessWithLevel.Title = accesses[i].Title
		accessWithLevel.Content = accesses[i].Content
		if err := db.Where("role_id = ? AND access_id = ?", roleId, accesses[i].ID).First(&roleWithAccess).Error; err == nil {
			// 查询到权限，进行赋值
			accessWithLevel.Level = roleWithAccess.Level
		} else {
			accessWithLevel.Level = 0
		}
		accessWithLevels = append(accessWithLevels, accessWithLevel)
	}
	return accessWithLevels, nil
}
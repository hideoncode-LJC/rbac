package db

import (
	"gorm.io/gorm"
	"server/model"
)

// 查找所有权限 不是关联查询

func GetAllAccess(db *gorm.DB) ([]model.Access, error) {
	var accesses []model.Access
	err := db.Find(&accesses).Error
	return accesses, err	
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
		if err := db.Where("role_id = ? AND access_id = ?", roleId, accesses[i].ID).First(&roleWithAccess).Error; err == nil {
			// 查询到权限，进行赋值
			accessWithLevel.Access = accesses[i]
			accessWithLevel.Level = roleWithAccess.Level
		} else {
			accessWithLevel.Access = accesses[i]
			accessWithLevel.Level = 0
		}
		accessWithLevels = append(accessWithLevels, accessWithLevel)
	}
	return accessWithLevels, nil
}
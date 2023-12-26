package db

import (
	"gorm.io/gorm"
	"server/model"
)

// 查找所有角色 不是关联查询

func GetAllRole(db *gorm.DB) ([]model.Role, error) {
	var roles []model.Role
	err := db.Find(&roles).Error
	return roles, err	
}

// 根据name查询角色
func GetRoleByName(db *gorm.DB, name string) (model.Role, error) {
	var role model.Role
	err := db.Where("name = ?", name).First(&role).Error
	return role, err
}

// 添加新的角色
func AddRole(db *gorm.DB, role model.Role) error {
	return db.Create(&role).Error
}
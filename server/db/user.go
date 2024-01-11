package db

import (
	"server/log"
	"server/model"

	"gorm.io/gorm"
)


func GetAllUser(db *gorm.DB, pageNum, pageSize int, userName, roleName string) ([]model.User, error) {
	offset := (pageNum - 1) * pageSize
	var users []model.User
	if userName == "" && roleName == "" {
		err := db.
		Limit(pageSize).
		Offset(offset).
		Preload("Role").
		Find(&users).
		Error
		return users, err
	} else if userName != "" && roleName == "" {
		err := db.
		Table("users").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("users.username = ?", userName).
		Preload("Role").
		Limit(pageSize).
		Offset(offset).
		Find(&users).
		Error
		return users, err
	} else if userName == "" && roleName != "" {
		err := db.
		Table("users").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("roles.name = ?", roleName).
		Preload("Role").
		Limit(pageSize).
		Offset(offset).
		Find(&users).
		Error
		return users, err
	} else {
		err := db.
		Table("users").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("users.username = ? AND roles.name = ?", userName, roleName).
		Preload("Role").
		Limit(pageSize).
		Offset(offset).
		Find(&users).Error
		return users, err
	}
}

func GetUserCount(db *gorm.DB) int {
	var users []model.User
	if err := db.Find(&users).Error ; err != nil {
		panic(err)
	}
	return len(users)
}

func GetUserByOption(db *gorm.DB, username, password string) (model.User, error) {
	var user model.User
	err := db.Preload("Role").Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err
}

func FindUser(db *gorm.DB, username, password string) (model.User, error) {
	var user model.User
	err := db.Preload("Role").Where("username = ?", username).First(&user).Error
	return user, err
}

// func AddUser(db *gorm.DB, user model.User) error {

// }

// func UpdateUser(db *gorm.DB, map[string]interface{}) error {

// }


func UpdateUserRole(db *gorm.DB, roleID uint) {
	// 先找到普通用户的roleID
	var role model.Role
	if err := db.Where("name = ?", "普通用户").First(&role).Error; err != nil {
		log.Info("根据角色名查询角色失败")
	}
	// 更新roleID
	result := db.
	Model(&model.User{}).
	Where("role_id = ?", roleID).
	Update("role_id", role.ID)

	log.Info("更新角色ID的影响行数为")
	log.Info(result.RowsAffected)
}

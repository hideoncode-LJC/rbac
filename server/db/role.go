package db

import (
	"server/log"
	"server/model"

	"gorm.io/gorm"
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

func UpdateRoleName(db *gorm.DB, roleID uint, rolename string) error {
	result := db.
	Model(&model.Role{}).
	Where("id = ?", roleID).
	Update("name", rolename)
	log.Info("更新角色名字, 收到影响的行数为")
	log.Info(result.RowsAffected)
	return result.Error
}


func hasAddRoleAccess(db *gorm.DB, roleID uint, accessID uint) bool {
	var roleAccess model.RoleAccess
	result := db.Where("role_id = ? and access_id = ?", roleID, accessID).First(&roleAccess)
	if result.Error != nil {
		log.Info("更新RoleAccess权限表发生错误")
		log.Info(result.Error)
		return false
	}
	return true
}

func deleteRoleAccess(db *gorm.DB, roleID uint, accessID uint) {
	result := db.Delete(&model.RoleAccess{}, "role_id = ? and access_id = ?", roleID, accessID)
	if result.RowsAffected == 0 {
		log.Warn("删除记录失败")
	}
	log.Info("删除记录成功")
}

func updateRoleAccess(db *gorm.DB, roleID uint, accessID uint, level int) {
	result := db.Model(&model.RoleAccess{}).
	Where("role_id = ? and access_id = ?", roleID, accessID).
	Update("level", level)
	if result.RowsAffected == 0 {
		log.Warn("更新记录失败")
	}
	log.Info("更新记录成功")
}

func createRoleAccess(db *gorm.DB, roleAccess model.RoleAccess) {
	result := db.Create(&roleAccess)
	if result.RowsAffected == 0 {
		log.Warn("创建记录失败")
	}
	log.Info("创建记录成功")
}

func UpdateRoleAccess(db *gorm.DB, roleID uint, addRoleModel *model.AddRoleModel) {
	for i := 0 ; i < len(addRoleModel.Accesses) ; i++ {
		// 判断该记录之前是否存在
		if hasAddRoleAccess(db, roleID, addRoleModel.Accesses[i].ID) {
			// 该记录已经存在
			// 如果level > 0 修改记录
			if addRoleModel.Accesses[i].Level > 0 {
				updateRoleAccess(db, roleID, addRoleModel.Accesses[i].ID, addRoleModel.Accesses[i].Level)
			} else { // 如果level == 0 删除记录
				deleteRoleAccess(db, roleID, addRoleModel.Accesses[i].ID)
			}
		} else {
			// 该记录不存在
			// 如果level > 0 创建记录
			createRoleAccess(db, model.RoleAccess{
				RoleID: roleID,
				AccessID: addRoleModel.Accesses[i].ID,
				Level: addRoleModel.Accesses[i].Level,
			})
		}
	}
}

func UpdateRole(db *gorm.DB, roleID uint, addRoleModel *model.AddRoleModel) error {
	err := UpdateRoleName(db, roleID, addRoleModel.RoleName)
	UpdateRoleAccess(db, roleID, addRoleModel)
	return err
}

func AddChildRole(db *gorm.DB, parentID uint, childID uint) error {
	roleInheritance := model.RoleInheritance{
		ParentID: parentID,
		ChildID: childID,
	}
	return db.Create(&roleInheritance).Error
}

func DeleteAllChild(db *gorm.DB, parentID uint) error {
	return db.
	Where("parent_id = ?", parentID).
	Delete(&model.RoleInheritance{}).Error
}

func GetChildRole(db *gorm.DB, parentID uint) ([]model.Role, error) {
	var roles []model.Role
	err := db.
	Where("id in (select child_id from role_inheritances where parent_id = ?)", parentID).
	Find(&roles).Error
	return roles, err
}



func GetChildLevel(db *gorm.DB, roleId, accessID uint) (int ,error) {
	roles, _ := GetChildRole(db, roleId)
	var level int = 0
	for i := 0 ; i < len(roles) ; i ++ {
		var roleAccess model.RoleAccess
		db.Where("role_id = ? and access_id = ?", roles[i].ID, accessID).First(&roleAccess)
		if roleAccess.Level > level {
			level = roleAccess.Level
		}
	}
	return level, nil
}